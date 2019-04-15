PWD := $(shell pwd)
PROTOC := $(shell which protoc)
NODE_PATH := $(PWD)/web
PROTOC_GEN_TS := $(PWD)/web/node_modules/.bin/protoc-gen-ts

PROTO_PATH := proto
GO_PROTO_PATH := go/generated
JS_PROTO_PATH := web/src/generated
PROTO_FILES := $(wildcard $(PROTO_PATH)/*.proto)
GO_PROTO_FILES := $(patsubst $(PROTO_PATH)/%.proto,$(GO_PROTO_PATH)/%.pb.go,$(PROTO_FILES))
TS_PROTO_FILES := $(patsubst $(PROTO_PATH)/%.proto,$(JS_PROTO_PATH)/%_pb.d.ts,$(PROTO_FILES))
JS_PROTO_FILES := $(patsubst $(PROTO_PATH)/%.proto,$(JS_PROTO_PATH)/%_pb.js,$(PROTO_FILES))
TS_SERVICE_PROTO_FILES := $(patsubst $(PROTO_PATH)/%.proto,$(JS_PROTO_PATH)/%_pb_service.d.ts,$(PROTO_FILES))
JS_SERVICE_PROTO_FILES := $(patsubst $(PROTO_PATH)/%.proto,$(JS_PROTO_PATH)/%_pb_service.js,$(PROTO_FILES))
ALL_PROTO_FILES := $(GO_PROTO_FILES) $(JS_PROTO_FILES) $(TS_PROTO_FILES) $(JS_SERVICE_PROTO_FILES) $(TS_SERVICE_PROTO_FILES)
WEB_SOURCES := $(shell git ls-files web/)

# default target
.PHONY: all
all: proto web-prod fmt

.PHONY: clean
clean:
	rm -f $(GO_PROTO_FILES) $(JS_PROTO_FILES) $(TS_PROTO_FILES) $(JS_SERVICE_PROTO_FILES) $(TS_SERVICE_PROTO_FILES)
	rm -rf web/dist/

#
# PROTO
#

$(GO_PROTO_FILES): $(PROTO_FILES)
	$(PROTOC) -I proto/ $^ --go_out=plugins=grpc:go/generated

$(JS_PROTO_FILES) $(TS_PROTO_FILES) $(JS_SERVICE_PROTO_FILES) $(TS_SERVICE_PROTO_FILES): $(PROTO_FILES)
	NODE_PATH="$(NODE_PATH)" $(PROTOC) \
		-I proto/ $^ \
		--plugin=protoc-gen-ts=$(PROTOC_GEN_TS) \
		--js_out=import_style=commonjs,binary:$(JS_PROTO_PATH) \
		--ts_out=service=true:$(JS_PROTO_PATH)

web/dist/prod/%: $(WEB_SOURCES) $(ALL_PROTO_FILES)
	cd web; npm run ng -- build --prod --output-path=./dist/prod

.PHONY: proto
proto: $(ALL_PROTO_FILES)

#
# BUILD
#

.PHONY: fmt
fmt:
	go fmt ./go/...

.PHONY: web-prod
web-prod: web/dist/prod/*

#
# DEV
#
#.secrets/oauth2-secret-dev:
.secrets/session-key-secret-dev:
	openssl rand -base64 -out .secrets/session-key-secret-dev 64

.secrets/oauth2-secret-dev.json:
	cat ./doc/oauth_dev_credentials.md

#
# DOCKER
#

.PHONY: docker-dev
docker-dev: proto
	docker-compose build

.PHONY: docker-dev-run
docker-dev-run: docker-dev
	$(MAKE) docker-dev-stop
	docker-compose up -d

.PHONY: docker-dev-stop
docker-dev-stop:
	docker-compose down || true

.PHONY: watch-docker-dev
watch-docker-dev:
	git ls-files | entr bash -c "time $(MAKE) docker-dev-run"
