// package: cryptocurrency.v1
// file: cryptocurrency/v1/cryptocurrency.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as tagger_tagger_pb from "../../tagger/tagger_pb";

export class Cryptocurrency extends jspb.Message { 
    getId(): string;
    setId(value: string): Cryptocurrency;
    getName(): string;
    setName(value: string): Cryptocurrency;
    getSymbol(): string;
    setSymbol(value: string): Cryptocurrency;
    getLogo(): string;
    setLogo(value: string): Cryptocurrency;
    clearPlatformsList(): void;
    getPlatformsList(): Array<string>;
    setPlatformsList(value: Array<string>): Cryptocurrency;
    addPlatforms(value: string, index?: number): string;

    hasCoinMarketCap(): boolean;
    clearCoinMarketCap(): void;
    getCoinMarketCap(): CoinmarketCap | undefined;
    setCoinMarketCap(value?: CoinmarketCap): Cryptocurrency;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Cryptocurrency.AsObject;
    static toObject(includeInstance: boolean, msg: Cryptocurrency): Cryptocurrency.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Cryptocurrency, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Cryptocurrency;
    static deserializeBinaryFromReader(message: Cryptocurrency, reader: jspb.BinaryReader): Cryptocurrency;
}

export namespace Cryptocurrency {
    export type AsObject = {
        id: string,
        name: string,
        symbol: string,
        logo: string,
        platformsList: Array<string>,
        coinMarketCap?: CoinmarketCap.AsObject,
    }
}

export class CoinmarketCap extends jspb.Message { 
    getId(): number;
    setId(value: number): CoinmarketCap;

    hasListing(): boolean;
    clearListing(): void;
    getListing(): CoinmarketCapListing | undefined;
    setListing(value?: CoinmarketCapListing): CoinmarketCap;

    hasInfo(): boolean;
    clearInfo(): void;
    getInfo(): CoinmarketCapMetadata | undefined;
    setInfo(value?: CoinmarketCapMetadata): CoinmarketCap;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CoinmarketCap.AsObject;
    static toObject(includeInstance: boolean, msg: CoinmarketCap): CoinmarketCap.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CoinmarketCap, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CoinmarketCap;
    static deserializeBinaryFromReader(message: CoinmarketCap, reader: jspb.BinaryReader): CoinmarketCap;
}

export namespace CoinmarketCap {
    export type AsObject = {
        id: number,
        listing?: CoinmarketCapListing.AsObject,
        info?: CoinmarketCapMetadata.AsObject,
    }
}

export class CoinmarketCapQuote extends jspb.Message { 
    getPrice(): number;
    setPrice(value: number): CoinmarketCapQuote;
    getVolume24h(): number;
    setVolume24h(value: number): CoinmarketCapQuote;
    getPercentChange1h(): number;
    setPercentChange1h(value: number): CoinmarketCapQuote;
    getPercentChange24h(): number;
    setPercentChange24h(value: number): CoinmarketCapQuote;
    getPercentChange7d(): number;
    setPercentChange7d(value: number): CoinmarketCapQuote;
    getMarketCap(): number;
    setMarketCap(value: number): CoinmarketCapQuote;
    getLastUpdated(): string;
    setLastUpdated(value: string): CoinmarketCapQuote;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CoinmarketCapQuote.AsObject;
    static toObject(includeInstance: boolean, msg: CoinmarketCapQuote): CoinmarketCapQuote.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CoinmarketCapQuote, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CoinmarketCapQuote;
    static deserializeBinaryFromReader(message: CoinmarketCapQuote, reader: jspb.BinaryReader): CoinmarketCapQuote;
}

export namespace CoinmarketCapQuote {
    export type AsObject = {
        price: number,
        volume24h: number,
        percentChange1h: number,
        percentChange24h: number,
        percentChange7d: number,
        marketCap: number,
        lastUpdated: string,
    }
}

export class CoinmarketCapListing extends jspb.Message { 

    getQuoteMap(): jspb.Map<string, CoinmarketCapQuote>;
    clearQuoteMap(): void;
    getMaxSupply(): number;
    setMaxSupply(value: number): CoinmarketCapListing;
    getTotalSupply(): number;
    setTotalSupply(value: number): CoinmarketCapListing;
    getCmcRank(): number;
    setCmcRank(value: number): CoinmarketCapListing;
    getCirculatingSupply(): number;
    setCirculatingSupply(value: number): CoinmarketCapListing;
    getLastUpdated(): string;
    setLastUpdated(value: string): CoinmarketCapListing;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CoinmarketCapListing.AsObject;
    static toObject(includeInstance: boolean, msg: CoinmarketCapListing): CoinmarketCapListing.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CoinmarketCapListing, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CoinmarketCapListing;
    static deserializeBinaryFromReader(message: CoinmarketCapListing, reader: jspb.BinaryReader): CoinmarketCapListing;
}

export namespace CoinmarketCapListing {
    export type AsObject = {

        quoteMap: Array<[string, CoinmarketCapQuote.AsObject]>,
        maxSupply: number,
        totalSupply: number,
        cmcRank: number,
        circulatingSupply: number,
        lastUpdated: string,
    }
}

export class CoinmarketCapMetadata extends jspb.Message { 
    getDescription(): string;
    setDescription(value: string): CoinmarketCapMetadata;
    getDateAdded(): string;
    setDateAdded(value: string): CoinmarketCapMetadata;
    getDateLaunched(): string;
    setDateLaunched(value: string): CoinmarketCapMetadata;
    getCategory(): string;
    setCategory(value: string): CoinmarketCapMetadata;
    getLogo(): string;
    setLogo(value: string): CoinmarketCapMetadata;

    hasUrls(): boolean;
    clearUrls(): void;
    getUrls(): CoinmarketCapUrls | undefined;
    setUrls(value?: CoinmarketCapUrls): CoinmarketCapMetadata;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CoinmarketCapMetadata.AsObject;
    static toObject(includeInstance: boolean, msg: CoinmarketCapMetadata): CoinmarketCapMetadata.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CoinmarketCapMetadata, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CoinmarketCapMetadata;
    static deserializeBinaryFromReader(message: CoinmarketCapMetadata, reader: jspb.BinaryReader): CoinmarketCapMetadata;
}

export namespace CoinmarketCapMetadata {
    export type AsObject = {
        description: string,
        dateAdded: string,
        dateLaunched: string,
        category: string,
        logo: string,
        urls?: CoinmarketCapUrls.AsObject,
    }
}

export class CoinmarketCapUrls extends jspb.Message { 
    clearWebsiteList(): void;
    getWebsiteList(): Array<string>;
    setWebsiteList(value: Array<string>): CoinmarketCapUrls;
    addWebsite(value: string, index?: number): string;
    clearTwitterList(): void;
    getTwitterList(): Array<string>;
    setTwitterList(value: Array<string>): CoinmarketCapUrls;
    addTwitter(value: string, index?: number): string;
    clearRedditList(): void;
    getRedditList(): Array<string>;
    setRedditList(value: Array<string>): CoinmarketCapUrls;
    addReddit(value: string, index?: number): string;
    clearSourceCodeList(): void;
    getSourceCodeList(): Array<string>;
    setSourceCodeList(value: Array<string>): CoinmarketCapUrls;
    addSourceCode(value: string, index?: number): string;
    clearChatList(): void;
    getChatList(): Array<string>;
    setChatList(value: Array<string>): CoinmarketCapUrls;
    addChat(value: string, index?: number): string;
    clearTechnicalDocList(): void;
    getTechnicalDocList(): Array<string>;
    setTechnicalDocList(value: Array<string>): CoinmarketCapUrls;
    addTechnicalDoc(value: string, index?: number): string;
    clearMessageBoardList(): void;
    getMessageBoardList(): Array<string>;
    setMessageBoardList(value: Array<string>): CoinmarketCapUrls;
    addMessageBoard(value: string, index?: number): string;
    clearExplorerList(): void;
    getExplorerList(): Array<string>;
    setExplorerList(value: Array<string>): CoinmarketCapUrls;
    addExplorer(value: string, index?: number): string;
    clearAnnouncementList(): void;
    getAnnouncementList(): Array<string>;
    setAnnouncementList(value: Array<string>): CoinmarketCapUrls;
    addAnnouncement(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CoinmarketCapUrls.AsObject;
    static toObject(includeInstance: boolean, msg: CoinmarketCapUrls): CoinmarketCapUrls.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CoinmarketCapUrls, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CoinmarketCapUrls;
    static deserializeBinaryFromReader(message: CoinmarketCapUrls, reader: jspb.BinaryReader): CoinmarketCapUrls;
}

export namespace CoinmarketCapUrls {
    export type AsObject = {
        websiteList: Array<string>,
        twitterList: Array<string>,
        redditList: Array<string>,
        sourceCodeList: Array<string>,
        chatList: Array<string>,
        technicalDocList: Array<string>,
        messageBoardList: Array<string>,
        explorerList: Array<string>,
        announcementList: Array<string>,
    }
}
