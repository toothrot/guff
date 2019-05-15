import * as jspb from "google-protobuf"

export class GetDivisionsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDivisionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDivisionsRequest): GetDivisionsRequest.AsObject;
  static serializeBinaryToWriter(message: GetDivisionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDivisionsRequest;
  static deserializeBinaryFromReader(message: GetDivisionsRequest, reader: jspb.BinaryReader): GetDivisionsRequest;
}

export namespace GetDivisionsRequest {
  export type AsObject = {
  }
}

export class GetDivisionsResponse extends jspb.Message {
  getDivisionsList(): Array<Division>;
  setDivisionsList(value: Array<Division>): void;
  clearDivisionsList(): void;
  addDivisions(value?: Division, index?: number): Division;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDivisionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDivisionsResponse): GetDivisionsResponse.AsObject;
  static serializeBinaryToWriter(message: GetDivisionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDivisionsResponse;
  static deserializeBinaryFromReader(message: GetDivisionsResponse, reader: jspb.BinaryReader): GetDivisionsResponse;
}

export namespace GetDivisionsResponse {
  export type AsObject = {
    divisionsList: Array<Division.AsObject>,
  }
}

export class Division extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Division.AsObject;
  static toObject(includeInstance: boolean, msg: Division): Division.AsObject;
  static serializeBinaryToWriter(message: Division, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Division;
  static deserializeBinaryFromReader(message: Division, reader: jspb.BinaryReader): Division;
}

export namespace Division {
  export type AsObject = {
    id: string,
    name: string,
  }
}

