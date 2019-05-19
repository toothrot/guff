import * as jspb from "google-protobuf"

export class GetTeamsRequest extends jspb.Message {
  getDivisionId(): string;
  setDivisionId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTeamsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTeamsRequest): GetTeamsRequest.AsObject;
  static serializeBinaryToWriter(message: GetTeamsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTeamsRequest;
  static deserializeBinaryFromReader(message: GetTeamsRequest, reader: jspb.BinaryReader): GetTeamsRequest;
}

export namespace GetTeamsRequest {
  export type AsObject = {
    divisionId: string,
  }
}

export class GetTeamsResponse extends jspb.Message {
  getTeamsList(): Array<Team>;
  setTeamsList(value: Array<Team>): void;
  clearTeamsList(): void;
  addTeams(value?: Team, index?: number): Team;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTeamsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetTeamsResponse): GetTeamsResponse.AsObject;
  static serializeBinaryToWriter(message: GetTeamsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTeamsResponse;
  static deserializeBinaryFromReader(message: GetTeamsResponse, reader: jspb.BinaryReader): GetTeamsResponse;
}

export namespace GetTeamsResponse {
  export type AsObject = {
    teamsList: Array<Team.AsObject>,
  }
}

export class Team extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getDivisionId(): string;
  setDivisionId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Team.AsObject;
  static toObject(includeInstance: boolean, msg: Team): Team.AsObject;
  static serializeBinaryToWriter(message: Team, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Team;
  static deserializeBinaryFromReader(message: Team, reader: jspb.BinaryReader): Team;
}

export namespace Team {
  export type AsObject = {
    id: string,
    name: string,
    divisionId: string,
  }
}

