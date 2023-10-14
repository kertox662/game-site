// @generated by protoc-gen-es v1.3.3 with parameter "target=ts"
// @generated from file proto/games/tictactoe/tictactoe.proto (package tictactoe, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Board } from "./board_pb.js";
import { GameMetadata, Move } from "./model_pb.js";

/**
 * @generated from message tictactoe.CreateGameRequest
 */
export class CreateGameRequest extends Message<CreateGameRequest> {
  /**
   * @generated from field: int32 max_players = 1;
   */
  maxPlayers = 0;

  /**
   * @generated from field: int32 board_size = 2;
   */
  boardSize = 0;

  /**
   * @generated from field: int32 connect_target = 3;
   */
  connectTarget = 0;

  constructor(data?: PartialMessage<CreateGameRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tictactoe.CreateGameRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "max_players", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "board_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "connect_target", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateGameRequest {
    return new CreateGameRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateGameRequest {
    return new CreateGameRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateGameRequest {
    return new CreateGameRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateGameRequest | PlainMessage<CreateGameRequest> | undefined, b: CreateGameRequest | PlainMessage<CreateGameRequest> | undefined): boolean {
    return proto3.util.equals(CreateGameRequest, a, b);
  }
}

/**
 * @generated from message tictactoe.CreateGameResponse
 */
export class CreateGameResponse extends Message<CreateGameResponse> {
  /**
   * @generated from field: string game_id = 1;
   */
  gameId = "";

  constructor(data?: PartialMessage<CreateGameResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tictactoe.CreateGameResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "game_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateGameResponse {
    return new CreateGameResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateGameResponse {
    return new CreateGameResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateGameResponse {
    return new CreateGameResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateGameResponse | PlainMessage<CreateGameResponse> | undefined, b: CreateGameResponse | PlainMessage<CreateGameResponse> | undefined): boolean {
    return proto3.util.equals(CreateGameResponse, a, b);
  }
}

/**
 * @generated from message tictactoe.GetGameDataRequest
 */
export class GetGameDataRequest extends Message<GetGameDataRequest> {
  /**
   * @generated from field: string game_id = 1;
   */
  gameId = "";

  constructor(data?: PartialMessage<GetGameDataRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tictactoe.GetGameDataRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "game_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGameDataRequest {
    return new GetGameDataRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGameDataRequest {
    return new GetGameDataRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGameDataRequest {
    return new GetGameDataRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetGameDataRequest | PlainMessage<GetGameDataRequest> | undefined, b: GetGameDataRequest | PlainMessage<GetGameDataRequest> | undefined): boolean {
    return proto3.util.equals(GetGameDataRequest, a, b);
  }
}

/**
 * @generated from message tictactoe.GetGameDataResponse
 */
export class GetGameDataResponse extends Message<GetGameDataResponse> {
  /**
   * @generated from field: tictactoe.Board data = 1;
   */
  data?: Board;

  /**
   * @generated from field: tictactoe.GameMetadata metadata = 2;
   */
  metadata?: GameMetadata;

  constructor(data?: PartialMessage<GetGameDataResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tictactoe.GetGameDataResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: Board },
    { no: 2, name: "metadata", kind: "message", T: GameMetadata },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGameDataResponse {
    return new GetGameDataResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGameDataResponse {
    return new GetGameDataResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGameDataResponse {
    return new GetGameDataResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetGameDataResponse | PlainMessage<GetGameDataResponse> | undefined, b: GetGameDataResponse | PlainMessage<GetGameDataResponse> | undefined): boolean {
    return proto3.util.equals(GetGameDataResponse, a, b);
  }
}

/**
 * @generated from message tictactoe.MakeMoveRequest
 */
export class MakeMoveRequest extends Message<MakeMoveRequest> {
  /**
   * @generated from field: string game_id = 1;
   */
  gameId = "";

  /**
   * @generated from field: int32 player = 2;
   */
  player = 0;

  /**
   * @generated from field: tictactoe.Move move = 3;
   */
  move?: Move;

  constructor(data?: PartialMessage<MakeMoveRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tictactoe.MakeMoveRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "game_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "player", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "move", kind: "message", T: Move },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MakeMoveRequest {
    return new MakeMoveRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MakeMoveRequest {
    return new MakeMoveRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MakeMoveRequest {
    return new MakeMoveRequest().fromJsonString(jsonString, options);
  }

  static equals(a: MakeMoveRequest | PlainMessage<MakeMoveRequest> | undefined, b: MakeMoveRequest | PlainMessage<MakeMoveRequest> | undefined): boolean {
    return proto3.util.equals(MakeMoveRequest, a, b);
  }
}

/**
 * @generated from message tictactoe.MakeMoveResponse
 */
export class MakeMoveResponse extends Message<MakeMoveResponse> {
  constructor(data?: PartialMessage<MakeMoveResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tictactoe.MakeMoveResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MakeMoveResponse {
    return new MakeMoveResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MakeMoveResponse {
    return new MakeMoveResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MakeMoveResponse {
    return new MakeMoveResponse().fromJsonString(jsonString, options);
  }

  static equals(a: MakeMoveResponse | PlainMessage<MakeMoveResponse> | undefined, b: MakeMoveResponse | PlainMessage<MakeMoveResponse> | undefined): boolean {
    return proto3.util.equals(MakeMoveResponse, a, b);
  }
}

/**
 * @generated from message tictactoe.ListGamesRequest
 */
export class ListGamesRequest extends Message<ListGamesRequest> {
  constructor(data?: PartialMessage<ListGamesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tictactoe.ListGamesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListGamesRequest {
    return new ListGamesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListGamesRequest {
    return new ListGamesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListGamesRequest {
    return new ListGamesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListGamesRequest | PlainMessage<ListGamesRequest> | undefined, b: ListGamesRequest | PlainMessage<ListGamesRequest> | undefined): boolean {
    return proto3.util.equals(ListGamesRequest, a, b);
  }
}

/**
 * @generated from message tictactoe.ListGamesResponse
 */
export class ListGamesResponse extends Message<ListGamesResponse> {
  /**
   * @generated from field: repeated tictactoe.GameMetadata games = 1;
   */
  games: GameMetadata[] = [];

  constructor(data?: PartialMessage<ListGamesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tictactoe.ListGamesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "games", kind: "message", T: GameMetadata, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListGamesResponse {
    return new ListGamesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListGamesResponse {
    return new ListGamesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListGamesResponse {
    return new ListGamesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListGamesResponse | PlainMessage<ListGamesResponse> | undefined, b: ListGamesResponse | PlainMessage<ListGamesResponse> | undefined): boolean {
    return proto3.util.equals(ListGamesResponse, a, b);
  }
}

