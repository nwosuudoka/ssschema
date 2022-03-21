// package: subscribedkeyword.v1
// file: subscribedkeyword/v1/subscribed_keyword.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class SubscribedKeyword extends jspb.Message { 
    getId(): string;
    setId(value: string): SubscribedKeyword;
    getUserId(): string;
    setUserId(value: string): SubscribedKeyword;
    getKeyword(): string;
    setKeyword(value: string): SubscribedKeyword;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SubscribedKeyword.AsObject;
    static toObject(includeInstance: boolean, msg: SubscribedKeyword): SubscribedKeyword.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SubscribedKeyword, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SubscribedKeyword;
    static deserializeBinaryFromReader(message: SubscribedKeyword, reader: jspb.BinaryReader): SubscribedKeyword;
}

export namespace SubscribedKeyword {
    export type AsObject = {
        id: string,
        userId: string,
        keyword: string,
    }
}
