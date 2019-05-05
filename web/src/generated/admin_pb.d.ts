import * as jspb from "google-protobuf"

export class ScrapeRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ScrapeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ScrapeRequest): ScrapeRequest.AsObject;
  static serializeBinaryToWriter(message: ScrapeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ScrapeRequest;
  static deserializeBinaryFromReader(message: ScrapeRequest, reader: jspb.BinaryReader): ScrapeRequest;
}

export namespace ScrapeRequest {
  export type AsObject = {
  }
}

export class ScrapeResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ScrapeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ScrapeResponse): ScrapeResponse.AsObject;
  static serializeBinaryToWriter(message: ScrapeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ScrapeResponse;
  static deserializeBinaryFromReader(message: ScrapeResponse, reader: jspb.BinaryReader): ScrapeResponse;
}

export namespace ScrapeResponse {
  export type AsObject = {
  }
}

