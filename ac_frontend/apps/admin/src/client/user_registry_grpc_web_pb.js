/**
 * @fileoverview gRPC-Web generated client stub for api.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.api = {};
proto.api.v1 = require('./user_registry_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.v1.UserRegistryClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

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
proto.api.v1.UserRegistryPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

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
 *   !proto.api.v1.GetUserRequest,
 *   !proto.api.v1.GetUserResponse>}
 */
const methodDescriptor_UserRegistry_GetUser = new grpc.web.MethodDescriptor(
  '/api.v1.UserRegistry/GetUser',
  grpc.web.MethodType.UNARY,
  proto.api.v1.GetUserRequest,
  proto.api.v1.GetUserResponse,
  /**
   * @param {!proto.api.v1.GetUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.GetUserResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.v1.GetUserRequest,
 *   !proto.api.v1.GetUserResponse>}
 */
const methodInfo_UserRegistry_GetUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.v1.GetUserResponse,
  /**
   * @param {!proto.api.v1.GetUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.GetUserResponse.deserializeBinary
);


/**
 * @param {!proto.api.v1.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.v1.GetUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.v1.GetUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.v1.UserRegistryClient.prototype.getUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.v1.UserRegistry/GetUser',
      request,
      metadata || {},
      methodDescriptor_UserRegistry_GetUser,
      callback);
};


/**
 * @param {!proto.api.v1.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.v1.GetUserResponse>}
 *     Promise that resolves to the response
 */
proto.api.v1.UserRegistryPromiseClient.prototype.getUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.v1.UserRegistry/GetUser',
      request,
      metadata || {},
      methodDescriptor_UserRegistry_GetUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.v1.CreateUserRequest,
 *   !proto.api.v1.CreateUserResponse>}
 */
const methodDescriptor_UserRegistry_CreateUser = new grpc.web.MethodDescriptor(
  '/api.v1.UserRegistry/CreateUser',
  grpc.web.MethodType.UNARY,
  proto.api.v1.CreateUserRequest,
  proto.api.v1.CreateUserResponse,
  /**
   * @param {!proto.api.v1.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.CreateUserResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.v1.CreateUserRequest,
 *   !proto.api.v1.CreateUserResponse>}
 */
const methodInfo_UserRegistry_CreateUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.v1.CreateUserResponse,
  /**
   * @param {!proto.api.v1.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.CreateUserResponse.deserializeBinary
);


/**
 * @param {!proto.api.v1.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.v1.CreateUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.v1.CreateUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.v1.UserRegistryClient.prototype.createUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.v1.UserRegistry/CreateUser',
      request,
      metadata || {},
      methodDescriptor_UserRegistry_CreateUser,
      callback);
};


/**
 * @param {!proto.api.v1.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.v1.CreateUserResponse>}
 *     Promise that resolves to the response
 */
proto.api.v1.UserRegistryPromiseClient.prototype.createUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.v1.UserRegistry/CreateUser',
      request,
      metadata || {},
      methodDescriptor_UserRegistry_CreateUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.v1.LoginUserRequest,
 *   !proto.api.v1.LoginUserResponse>}
 */
const methodDescriptor_UserRegistry_LoginUser = new grpc.web.MethodDescriptor(
  '/api.v1.UserRegistry/LoginUser',
  grpc.web.MethodType.UNARY,
  proto.api.v1.LoginUserRequest,
  proto.api.v1.LoginUserResponse,
  /**
   * @param {!proto.api.v1.LoginUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.LoginUserResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.v1.LoginUserRequest,
 *   !proto.api.v1.LoginUserResponse>}
 */
const methodInfo_UserRegistry_LoginUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.v1.LoginUserResponse,
  /**
   * @param {!proto.api.v1.LoginUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.LoginUserResponse.deserializeBinary
);


/**
 * @param {!proto.api.v1.LoginUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.v1.LoginUserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.v1.LoginUserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.v1.UserRegistryClient.prototype.loginUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.v1.UserRegistry/LoginUser',
      request,
      metadata || {},
      methodDescriptor_UserRegistry_LoginUser,
      callback);
};


/**
 * @param {!proto.api.v1.LoginUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.v1.LoginUserResponse>}
 *     Promise that resolves to the response
 */
proto.api.v1.UserRegistryPromiseClient.prototype.loginUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.v1.UserRegistry/LoginUser',
      request,
      metadata || {},
      methodDescriptor_UserRegistry_LoginUser);
};


module.exports = proto.api.v1;

