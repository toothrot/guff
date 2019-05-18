PWD := $(shell pwd)
PROTOC := $(shell which protoc)
NODE_PATH := $(PWD)/web
PROTOC_GEN_TS := $(PWD)/web/node_modules/.bin/protoc-gen-ts

PROTO_PATH := proto
GO_PROTO_PATH := backend/generated
JS_PROTO_PATH := web/src/generated
PROTO_FILES := $(wildcard $(PROTO_PATH)/*.proto)
GO_PROTO_FILES := $(patsubst $(PROTO_PATH)/%.proto,$(GO_PROTO_PATH)/%.pb.go,$(PROTO_FILES))
TS_PROTO_FILES := $(patsubst $(PROTO_PATH)/%.proto,$(JS_PROTO_PATH)/%_pb.d.ts,$(PROTO_FILES))
JS_PROTO_FILES := $(patsubst $(PROTO_PATH)/%.proto,$(JS_PROTO_PATH)/%_pb.js,$(PROTO_FILES))
ALL_PROTO_FILES := $(GO_PROTO_FILES) $(JS_PROTO_FILES) $(TS_PROTO_FILES) $(JS_SERVICE_PROTO_FILES) $(TS_SERVICE_PROTO_FILES)
WEB_SOURCES := $(shell git ls-files web/)

# default target
.PHONY: all
all: fmt proto web-prod dist/guff

.PHONY: clean
clean:
	rm -f $(GO_PROTO_FILES) $(JS_PROTO_FILES) $(TS_PROTO_FILES) $(JS_SERVICE_PROTO_FILES) $(TS_SERVICE_PROTO_FILES)
	rm -rf web/dist/
	rm -rf dist/

#
# PROTO
#

$(GO_PROTO_FILES): $(PROTO_FILES)
	$(PROTOC) -I proto/ $^ --go_out=plugins=grpc:$(GO_PROTO_PATH)

$(JS_PROTO_FILES) $(TS_PROTO_FILES): $(PROTO_FILES)
	NODE_PATH="$(NODE_PATH)" $(PROTOC) \
		-I proto/ $^ \
		--js_out=import_style=commonjs:$(JS_PROTO_PATH) \
		--grpc_web_out=import_style=typescript,mode=grpcwebtext:$(JS_PROTO_PATH)

web/dist/prod/%: $(WEB_SOURCES) $(ALL_PROTO_FILES)
	cd web; npm run ng -- build --prod --output-path=./dist/prod

.PHONY: proto
proto: $(ALL_PROTO_FILES)

#
# TEST
#
.PHONY: test
test: proto go-test web-test

.PHONY: go-test
go-test:
	cd backend; go test ./...

.PHONY: watch-go-test
watch-go-test:
	git ls-files | entr bash -c "time $(MAKE) go-test"

.PHONY: web-test
web-test:
	cd web; npm run -- ng test --watch=false --browsers=ChromeHeadless

.PHONY: watch-web-test
watch-web-test:
	cd web; npm run -- ng test

#
# BUILD
#

.PHONY: fmt
fmt:
	go fmt ./backend/...

.PHONY: web-prod
web-prod: web/dist/prod/*

dist:
	mkdir -p dist

dist/guff: dist
	cd backend; go build -o guff . && mv guff ../dist

#
# DEV
#
.PHONY: secrets
secrets: .secrets/postgres-password .secrets/oauth2-secret-dev.json .secrets/postgres-guff-password .secrets/session-key-secret-dev

.secrets/session-key-secret-dev:
	openssl rand -base64 -out .secrets/session-key-secret-dev 64

.secrets/oauth2-secret-dev.json:
	cat ./doc/oauth_dev_credentials.md

.secrets/postgres-password:
	openssl rand -hex -out .secrets/postgres-password 32

.secrets/postgres-guff-password:
	openssl rand -hex -out .secrets/postgres-guff-password 32

#
# DOCKER
#
.PHONY: docker
docker:
	docker build --build-arg configuration=production -t guff:1 .

.PHONY: docker-push
docker-push: docker
	docker tag guff:1 gcr.io/shuffleboardclub/guff:1
	docker push gcr.io/shuffleboardclub/guff:1

.PHONY: docker-test
docker-test:
	docker-compose run --rm backend-test
	docker-compose run --rm web-test

.PHONY: docker-dev
docker-dev: proto secrets
	docker-compose build

.PHONY: docker-dev-run
docker-dev-run: docker-dev
	$(MAKE) docker-dev-stop
	docker-compose up -d web

.PHONY: docker-dev-stop
docker-dev-stop:
	docker-compose stop web || true

.PHONY: docker-dev-down
docker-dev-down:
	docker-compose down || true

.PHONY: docker-dev-clean
docker-dev-clean:
	docker-compose down --rmi local || true

.PHONY: watch-docker-dev
watch-docker-dev:
	while true; do git ls-files | entr bash -c "time $(MAKE) docker-dev-run"; done

.PHONY: watch-docker-test
watch-docker-test:
	while true; do git ls-files | entr bash -c "time $(MAKE) docker-test"; done
