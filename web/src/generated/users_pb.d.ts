import * as jspb from "google-protobuf"

export class GetCurrentUserRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCurrentUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCurrentUserRequest): GetCurrentUserRequest.AsObject;
  static serializeBinaryToWriter(message: GetCurrentUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCurrentUserRequest;
  static deserializeBinaryFromReader(message: GetCurrentUserRequest, reader: jspb.BinaryReader): GetCurrentUserRequest;
}

export namespace GetCurrentUserRequest {
  export type AsObject = {
  }
}

export class GetCurrentUserResponse extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): void;

  getGoogleOauthConfig(): GoogleOAuthConfig | undefined;
  setGoogleOauthConfig(value?: GoogleOAuthConfig): void;
  hasGoogleOauthConfig(): boolean;
  clearGoogleOauthConfig(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCurrentUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCurrentUserResponse): GetCurrentUserResponse.AsObject;
  static serializeBinaryToWriter(message: GetCurrentUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCurrentUserResponse;
  static deserializeBinaryFromReader(message: GetCurrentUserResponse, reader: jspb.BinaryReader): GetCurrentUserResponse;
}

export namespace GetCurrentUserResponse {
  export type AsObject = {
    email: string,
    googleOauthConfig?: GoogleOAuthConfig.AsObject,
  }
}

export class GoogleOAuthConfig extends jspb.Message {
  getClientId(): string;
  setClientId(value: string): void;

  getLoginurl(): string;
  setLoginurl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GoogleOAuthConfig.AsObject;
  static toObject(includeInstance: boolean, msg: GoogleOAuthConfig): GoogleOAuthConfig.AsObject;
  static serializeBinaryToWriter(message: GoogleOAuthConfig, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GoogleOAuthConfig;
  static deserializeBinaryFromReader(message: GoogleOAuthConfig, reader: jspb.BinaryReader): GoogleOAuthConfig;
}

export namespace GoogleOAuthConfig {
  export type AsObject = {
    clientId: string,
    loginurl: string,
  }
}

