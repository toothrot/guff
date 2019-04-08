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

# default target
all: proto

$(GO_PROTO_FILES): $(PROTO_FILES)
	$(PROTOC) -I proto/ $< --go_out=plugins=grpc:go/generated

$(JS_PROTO_FILES) $(TS_PROTO_FILES) $(JS_SERVICE_PROTO_FILES) $(TS_SERVICE_PROTO_FILES): $(PROTO_FILES)
	NODE_PATH="$(NODE_PATH)" $(PROTOC) \
		-I proto/ $< \
		--plugin=protoc-gen-ts=$(PROTOC_GEN_TS) \
		--js_out=import_style=commonjs,binary:$(JS_PROTO_PATH) \
		--ts_out=service=true:$(JS_PROTO_PATH)

.PHONY: proto
proto: $(GO_PROTO_FILES) $(JS_PROTO_FILES) $(TS_PROTO_FILES) $(JS_SERVICE_PROTO_FILES) $(TS_SERVICE_PROTO_FILES)

.PHONY: clean
clean:
	rm $(GO_PROTO_FILES) $(JS_PROTO_FILES) $(TS_PROTO_FILES) $(JS_SERVICE_PROTO_FILES) $(TS_SERVICE_PROTO_FILES)