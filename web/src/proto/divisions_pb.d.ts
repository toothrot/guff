// package: guff.proto
// file: divisions.proto

import * as jspb from "google-protobuf";

export class GetDivisionsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDivisionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDivisionsRequest): GetDivisionsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetDivisionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDivisionsRequest;
  static deserializeBinaryFromReader(message: GetDivisionsRequest, reader: jspb.BinaryReader): GetDivisionsRequest;
}

export namespace GetDivisionsRequest {
  export type AsObject = {
  }
}

export class GetDivisionsResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDivisionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDivisionsResponse): GetDivisionsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetDivisionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDivisionsResponse;
  static deserializeBinaryFromReader(message: GetDivisionsResponse, reader: jspb.BinaryReader): GetDivisionsResponse;
}

export namespace GetDivisionsResponse {
  export type AsObject = {
  }
}

