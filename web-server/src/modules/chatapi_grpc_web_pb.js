/**
 * @fileoverview gRPC-Web generated client stub for chatapi
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.chatapi = require('./chatapi_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.chatapi.ChatapiClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.chatapi.ChatapiPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.chatapi.GetchatRequest,
 *   !proto.chatapi.GoaChatCollection>}
 */
const methodDescriptor_Chatapi_Getchat = new grpc.web.MethodDescriptor(
  '/chatapi.Chatapi/Getchat',
  grpc.web.MethodType.UNARY,
  proto.chatapi.GetchatRequest,
  proto.chatapi.GoaChatCollection,
  /**
   * @param {!proto.chatapi.GetchatRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.chatapi.GoaChatCollection.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.chatapi.GetchatRequest,
 *   !proto.chatapi.GoaChatCollection>}
 */
const methodInfo_Chatapi_Getchat = new grpc.web.AbstractClientBase.MethodInfo(
  proto.chatapi.GoaChatCollection,
  /**
   * @param {!proto.chatapi.GetchatRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.chatapi.GoaChatCollection.deserializeBinary
);


/**
 * @param {!proto.chatapi.GetchatRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.chatapi.GoaChatCollection)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.chatapi.GoaChatCollection>|undefined}
 *     The XHR Node Readable Stream
 */
proto.chatapi.ChatapiClient.prototype.getchat =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/chatapi.Chatapi/Getchat',
      request,
      metadata || {},
      methodDescriptor_Chatapi_Getchat,
      callback);
};


/**
 * @param {!proto.chatapi.GetchatRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.chatapi.GoaChatCollection>}
 *     Promise that resolves to the response
 */
proto.chatapi.ChatapiPromiseClient.prototype.getchat =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/chatapi.Chatapi/Getchat',
      request,
      metadata || {},
      methodDescriptor_Chatapi_Getchat);
};


module.exports = proto.chatapi;

