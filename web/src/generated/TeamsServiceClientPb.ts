/**
 * @fileoverview gRPC-Web generated client stub for guff.proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  GetTeamsRequest,
  GetTeamsResponse} from './teams_pb';

export class TeamsServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; }) {
    if (!options) options = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetTeams = new grpcWeb.AbstractClientBase.MethodInfo(
    GetTeamsResponse,
    (request: GetTeamsRequest) => {
      return request.serializeBinary();
    },
    GetTeamsResponse.deserializeBinary
  );

  getTeams(
    request: GetTeamsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: GetTeamsResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/guff.proto.TeamsService/GetTeams',
      request,
      metadata || {},
      this.methodInfoGetTeams,
      callback);
  }

}

