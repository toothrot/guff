/**
 * @fileoverview gRPC-Web generated client stub for guff.proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {GetCurrentUserRequest, GetCurrentUserResponse} from './users_pb';

export class UsersServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor(hostname: string,
              credentials: null | { [index: string]: string; },
              options: null | { [index: string]: string; }) {
    if (!options) options = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetCurrentUser = new grpcWeb.AbstractClientBase.MethodInfo(
    GetCurrentUserResponse,
    (request: GetCurrentUserRequest) => {
      return request.serializeBinary();
    },
    GetCurrentUserResponse.deserializeBinary
  );

  getCurrentUser(
    request: GetCurrentUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: GetCurrentUserResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
      '/guff.proto.UsersService/GetCurrentUser',
      request,
      metadata || {},
      this.methodInfoGetCurrentUser,
      callback);
  }

}

