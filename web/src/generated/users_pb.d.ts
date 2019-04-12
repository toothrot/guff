// package: guff.proto
// file: users.proto

import * as jspb from "google-protobuf";

export class GetCurrentUserRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCurrentUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCurrentUserRequest): GetCurrentUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetCurrentUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCurrentUserRequest;
  static deserializeBinaryFromReader(message: GetCurrentUserRequest, reader: jspb.BinaryReader): GetCurrentUserRequest;
}

export namespace GetCurrentUserRequest {
  export type AsObject = {
  }
}

export class GetCurrentUserResponse extends jspb.Message {
  hasAppStatus(): boolean;
  clearAppStatus(): void;
  getAppStatus(): AppStatus | undefined;
  setAppStatus(value?: AppStatus): void;

  hasUserRoles(): boolean;
  clearUserRoles(): void;
  getUserRoles(): UserRoles | undefined;
  setUserRoles(value?: UserRoles): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCurrentUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCurrentUserResponse): GetCurrentUserResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetCurrentUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCurrentUserResponse;
  static deserializeBinaryFromReader(message: GetCurrentUserResponse, reader: jspb.BinaryReader): GetCurrentUserResponse;
}

export namespace GetCurrentUserResponse {
  export type AsObject = {
    appStatus?: AppStatus.AsObject,
    userRoles?: UserRoles.AsObject,
  }
}

export class UserRoles extends jspb.Message {
  getCanBootstrap(): boolean;
  setCanBootstrap(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserRoles.AsObject;
  static toObject(includeInstance: boolean, msg: UserRoles): UserRoles.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserRoles, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserRoles;
  static deserializeBinaryFromReader(message: UserRoles, reader: jspb.BinaryReader): UserRoles;
}

export namespace UserRoles {
  export type AsObject = {
    canBootstrap: boolean,
  }
}

export class AppStatus extends jspb.Message {
  getCode(): AppStatusCode;
  setCode(value: AppStatusCode): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AppStatus.AsObject;
  static toObject(includeInstance: boolean, msg: AppStatus): AppStatus.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AppStatus, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AppStatus;
  static deserializeBinaryFromReader(message: AppStatus, reader: jspb.BinaryReader): AppStatus;
}

export namespace AppStatus {
  export type AsObject = {
    code: AppStatusCode,
  }
}

export enum AppStatusCode {
  APP_STATUS_UNKNOWN = 0,
  APP_STATUS_OK = 1,
  APP_STATUS_DATABASE_CONNECTION_FAILURE = 2,
  APP_STATUS_NO_DATABASE = 3,
}

