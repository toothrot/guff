import * as jspb from "google-protobuf"

export class TestEchoRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestEchoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TestEchoRequest): TestEchoRequest.AsObject;
  static serializeBinaryToWriter(message: TestEchoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestEchoRequest;
  static deserializeBinaryFromReader(message: TestEchoRequest, reader: jspb.BinaryReader): TestEchoRequest;
}

export namespace TestEchoRequest {
  export type AsObject = {
    message: string,
  }
}

export class TestEchoResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TestEchoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TestEchoResponse): TestEchoResponse.AsObject;
  static serializeBinaryToWriter(message: TestEchoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TestEchoResponse;
  static deserializeBinaryFromReader(message: TestEchoResponse, reader: jspb.BinaryReader): TestEchoResponse;
}

export namespace TestEchoResponse {
  export type AsObject = {
    message: string,
  }
}

