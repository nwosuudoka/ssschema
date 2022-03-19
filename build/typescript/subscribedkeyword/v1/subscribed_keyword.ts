/* eslint-disable */
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "subscribedkeyword.v1";

/** SubscribedKeyword represents a keyword a user subscribed to. */
export interface SubscribedKeyword {
  /** Id represents the id of the record in the database. */
  id: string;
  /** UserId represents the user id that subscribed to the word. */
  userId: string;
  /** Keyword represents the keyword they subscribed to. */
  keyword: string;
}

function createBaseSubscribedKeyword(): SubscribedKeyword {
  return { id: "", userId: "", keyword: "" };
}

export const SubscribedKeyword = {
  encode(message: SubscribedKeyword, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    if (message.keyword !== "") {
      writer.uint32(26).string(message.keyword);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SubscribedKeyword {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubscribedKeyword();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.userId = reader.string();
          break;
        case 3:
          message.keyword = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubscribedKeyword {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      keyword: isSet(object.keyword) ? String(object.keyword) : "",
    };
  },

  toJSON(message: SubscribedKeyword): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.userId !== undefined && (obj.userId = message.userId);
    message.keyword !== undefined && (obj.keyword = message.keyword);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SubscribedKeyword>, I>>(
    object: I
  ): SubscribedKeyword {
    const message = createBaseSubscribedKeyword();
    message.id = object.id ?? "";
    message.userId = object.userId ?? "";
    message.keyword = object.keyword ?? "";
    return message;
  },
};

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
