/**
 * @fileoverview gRPC-Web generated client stub for guff.proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  ScrapeRequest,
  ScrapeResponse} from './admin_pb';

export class AdminServiceClient {
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

  methodInfoScrape = new grpcWeb.AbstractClientBase.MethodInfo(
    ScrapeResponse,
    (request: ScrapeRequest) => {
      return request.serializeBinary();
    },
    ScrapeResponse.deserializeBinary
  );

  scrape(
    request: ScrapeRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ScrapeResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/guff.proto.AdminService/Scrape',
      request,
      metadata || {},
      this.methodInfoScrape,
      callback);
  }

}

