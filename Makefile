PROTOC := $(shell which protoc)
PWD := $(shell pwd)

PROTO_PATH := $(PWD)/proto
GO_PROTO_PATH := $(PWD)/go/generated
PROTO_FILES := $(wildcard $(PROTO_PATH)/*.proto)
GO_PROTO_FILES := $(patsubst $(PROTO_PATH)/%.proto,$(GO_PROTO_PATH)/%.pb.go,$(PROTO_FILES))

$(GO_PROTO_FILES): $(PROTO_FILES)
	$(PROTOC) -I proto/ $< --go_out=plugins=grpc:go/generated

.PHONY: clean
clean:
	rm $(GO_PROTO_FILES)