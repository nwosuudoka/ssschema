/* eslint-disable */
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "cryptocurrency.v1";

/** Cryptocurrency represents the schema for a cryptocurrency */
export interface Cryptocurrency {
  /** Id represents the id of the record in the database. */
  id: number;
  /** Name represents the name of the currency. */
  name: string;
  /** Symbol represnets the symbol of the currency. */
  symbol: string;
  /** Logo represents the logo of the cryptocurrency. */
  logo: string;
  /** Platforms represents the platforms where this currencies are. */
  platforms: string[];
  /** Coinmarketcap represents the coin market cap data. */
  coinMarketCap: CoinmarketCap | undefined;
}

/** CoinmarketCap represents the coin market cap information. */
export interface CoinmarketCap {
  /** Id represnets the coinmarketcap id. */
  id: number;
  /** Info holds coinmarketcap info. */
  info: CoinmarketCapInfo | undefined;
}

/** CoinmarketCapQuote quote. */
export interface CoinmarketCapQuote {
  /** Price represnets the price change. */
  price: string;
  /** Volume24h. */
  volume24h: string;
  /** PercentChange1h */
  percentChange1h: string;
  /** PercentChange24h */
  percentChange24h: string;
  /** PercentChange7d */
  percentChange7d: string;
  /** MarketCap */
  marketCap: string;
  /** LastUpdated */
  lastUpdated: string;
}

/** CoinmarketCapInfo holds the info of the coin. */
export interface CoinmarketCapInfo {
  /** Quote represnets the quote of the cryptocurrency in different fiat currencies. */
  quote: { [key: string]: CoinmarketCapQuote };
  /** MaxSupply represents the max supply this currency can have. */
  maxSupply: number;
  /** TotalSupply represents the total supply this currency can have. */
  totalSupply: number;
  /** CMCRank represents the cmc rank of the coin. */
  cmcRank: number;
  /** CirculatingSupply is how many is currenly in circulation. */
  circulatingSupply: number;
  /** LastUpdated was the last time this was updated. */
  lastUpdated: string;
}

export interface CoinmarketCapInfo_QuoteEntry {
  key: string;
  value: CoinmarketCapQuote | undefined;
}

/** CoinmarketCapMetadata represents the md of the coin. */
export interface CoinmarketCapMetadata {
  /** Website represents the website of the cmc data. */
  website: string[];
  /** Twitter represents the twitter info of the cmc data. */
  twitter: string[];
  /** Reddit represents the reddit info of the cmc data. */
  reddit: string[];
  /** SourceCode represents the source_code info of the cmc data. */
  sourceCode: string[];
  /** Chat represents the chat infomration of the currency. */
  chat: string[];
  /** TechnicalDoc represents the technical doc of the data. */
  technicalDoc: string[];
}

function createBaseCryptocurrency(): Cryptocurrency {
  return {
    id: 0,
    name: "",
    symbol: "",
    logo: "",
    platforms: [],
    coinMarketCap: undefined,
  };
}

export const Cryptocurrency = {
  encode(message: Cryptocurrency, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.symbol !== "") {
      writer.uint32(26).string(message.symbol);
    }
    if (message.logo !== "") {
      writer.uint32(34).string(message.logo);
    }
    for (const v of message.platforms) {
      writer.uint32(42).string(v!);
    }
    if (message.coinMarketCap !== undefined) {
      CoinmarketCap.encode(
        message.coinMarketCap,
        writer.uint32(50).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Cryptocurrency {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCryptocurrency();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.int32();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.symbol = reader.string();
          break;
        case 4:
          message.logo = reader.string();
          break;
        case 5:
          message.platforms.push(reader.string());
          break;
        case 6:
          message.coinMarketCap = CoinmarketCap.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Cryptocurrency {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      logo: isSet(object.logo) ? String(object.logo) : "",
      platforms: Array.isArray(object?.platforms)
        ? object.platforms.map((e: any) => String(e))
        : [],
      coinMarketCap: isSet(object.coinMarketCap)
        ? CoinmarketCap.fromJSON(object.coinMarketCap)
        : undefined,
    };
  },

  toJSON(message: Cryptocurrency): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.name !== undefined && (obj.name = message.name);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.logo !== undefined && (obj.logo = message.logo);
    if (message.platforms) {
      obj.platforms = message.platforms.map((e) => e);
    } else {
      obj.platforms = [];
    }
    message.coinMarketCap !== undefined &&
      (obj.coinMarketCap = message.coinMarketCap
        ? CoinmarketCap.toJSON(message.coinMarketCap)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Cryptocurrency>, I>>(
    object: I
  ): Cryptocurrency {
    const message = createBaseCryptocurrency();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.symbol = object.symbol ?? "";
    message.logo = object.logo ?? "";
    message.platforms = object.platforms?.map((e) => e) || [];
    message.coinMarketCap =
      object.coinMarketCap !== undefined && object.coinMarketCap !== null
        ? CoinmarketCap.fromPartial(object.coinMarketCap)
        : undefined;
    return message;
  },
};

function createBaseCoinmarketCap(): CoinmarketCap {
  return { id: 0, info: undefined };
}

export const CoinmarketCap = {
  encode(message: CoinmarketCap, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.info !== undefined) {
      CoinmarketCapInfo.encode(message.info, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CoinmarketCap {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCoinmarketCap();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.int32();
          break;
        case 2:
          message.info = CoinmarketCapInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CoinmarketCap {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      info: isSet(object.info)
        ? CoinmarketCapInfo.fromJSON(object.info)
        : undefined,
    };
  },

  toJSON(message: CoinmarketCap): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.info !== undefined &&
      (obj.info = message.info
        ? CoinmarketCapInfo.toJSON(message.info)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CoinmarketCap>, I>>(
    object: I
  ): CoinmarketCap {
    const message = createBaseCoinmarketCap();
    message.id = object.id ?? 0;
    message.info =
      object.info !== undefined && object.info !== null
        ? CoinmarketCapInfo.fromPartial(object.info)
        : undefined;
    return message;
  },
};

function createBaseCoinmarketCapQuote(): CoinmarketCapQuote {
  return {
    price: "",
    volume24h: "",
    percentChange1h: "",
    percentChange24h: "",
    percentChange7d: "",
    marketCap: "",
    lastUpdated: "",
  };
}

export const CoinmarketCapQuote = {
  encode(
    message: CoinmarketCapQuote,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.price !== "") {
      writer.uint32(10).string(message.price);
    }
    if (message.volume24h !== "") {
      writer.uint32(18).string(message.volume24h);
    }
    if (message.percentChange1h !== "") {
      writer.uint32(26).string(message.percentChange1h);
    }
    if (message.percentChange24h !== "") {
      writer.uint32(34).string(message.percentChange24h);
    }
    if (message.percentChange7d !== "") {
      writer.uint32(42).string(message.percentChange7d);
    }
    if (message.marketCap !== "") {
      writer.uint32(50).string(message.marketCap);
    }
    if (message.lastUpdated !== "") {
      writer.uint32(58).string(message.lastUpdated);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CoinmarketCapQuote {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCoinmarketCapQuote();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.price = reader.string();
          break;
        case 2:
          message.volume24h = reader.string();
          break;
        case 3:
          message.percentChange1h = reader.string();
          break;
        case 4:
          message.percentChange24h = reader.string();
          break;
        case 5:
          message.percentChange7d = reader.string();
          break;
        case 6:
          message.marketCap = reader.string();
          break;
        case 7:
          message.lastUpdated = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CoinmarketCapQuote {
    return {
      price: isSet(object.price) ? String(object.price) : "",
      volume24h: isSet(object.volume24h) ? String(object.volume24h) : "",
      percentChange1h: isSet(object.percentChange1h)
        ? String(object.percentChange1h)
        : "",
      percentChange24h: isSet(object.percentChange24h)
        ? String(object.percentChange24h)
        : "",
      percentChange7d: isSet(object.percentChange7d)
        ? String(object.percentChange7d)
        : "",
      marketCap: isSet(object.marketCap) ? String(object.marketCap) : "",
      lastUpdated: isSet(object.lastUpdated) ? String(object.lastUpdated) : "",
    };
  },

  toJSON(message: CoinmarketCapQuote): unknown {
    const obj: any = {};
    message.price !== undefined && (obj.price = message.price);
    message.volume24h !== undefined && (obj.volume24h = message.volume24h);
    message.percentChange1h !== undefined &&
      (obj.percentChange1h = message.percentChange1h);
    message.percentChange24h !== undefined &&
      (obj.percentChange24h = message.percentChange24h);
    message.percentChange7d !== undefined &&
      (obj.percentChange7d = message.percentChange7d);
    message.marketCap !== undefined && (obj.marketCap = message.marketCap);
    message.lastUpdated !== undefined &&
      (obj.lastUpdated = message.lastUpdated);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CoinmarketCapQuote>, I>>(
    object: I
  ): CoinmarketCapQuote {
    const message = createBaseCoinmarketCapQuote();
    message.price = object.price ?? "";
    message.volume24h = object.volume24h ?? "";
    message.percentChange1h = object.percentChange1h ?? "";
    message.percentChange24h = object.percentChange24h ?? "";
    message.percentChange7d = object.percentChange7d ?? "";
    message.marketCap = object.marketCap ?? "";
    message.lastUpdated = object.lastUpdated ?? "";
    return message;
  },
};

function createBaseCoinmarketCapInfo(): CoinmarketCapInfo {
  return {
    quote: {},
    maxSupply: 0,
    totalSupply: 0,
    cmcRank: 0,
    circulatingSupply: 0,
    lastUpdated: "",
  };
}

export const CoinmarketCapInfo = {
  encode(message: CoinmarketCapInfo, writer: Writer = Writer.create()): Writer {
    Object.entries(message.quote).forEach(([key, value]) => {
      CoinmarketCapInfo_QuoteEntry.encode(
        { key: key as any, value },
        writer.uint32(10).fork()
      ).ldelim();
    });
    if (message.maxSupply !== 0) {
      writer.uint32(21).float(message.maxSupply);
    }
    if (message.totalSupply !== 0) {
      writer.uint32(29).float(message.totalSupply);
    }
    if (message.cmcRank !== 0) {
      writer.uint32(37).float(message.cmcRank);
    }
    if (message.circulatingSupply !== 0) {
      writer.uint32(45).float(message.circulatingSupply);
    }
    if (message.lastUpdated !== "") {
      writer.uint32(50).string(message.lastUpdated);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CoinmarketCapInfo {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCoinmarketCapInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = CoinmarketCapInfo_QuoteEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry1.value !== undefined) {
            message.quote[entry1.key] = entry1.value;
          }
          break;
        case 2:
          message.maxSupply = reader.float();
          break;
        case 3:
          message.totalSupply = reader.float();
          break;
        case 4:
          message.cmcRank = reader.float();
          break;
        case 5:
          message.circulatingSupply = reader.float();
          break;
        case 6:
          message.lastUpdated = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CoinmarketCapInfo {
    return {
      quote: isObject(object.quote)
        ? Object.entries(object.quote).reduce<{
            [key: string]: CoinmarketCapQuote;
          }>((acc, [key, value]) => {
            acc[key] = CoinmarketCapQuote.fromJSON(value);
            return acc;
          }, {})
        : {},
      maxSupply: isSet(object.maxSupply) ? Number(object.maxSupply) : 0,
      totalSupply: isSet(object.totalSupply) ? Number(object.totalSupply) : 0,
      cmcRank: isSet(object.cmcRank) ? Number(object.cmcRank) : 0,
      circulatingSupply: isSet(object.circulatingSupply)
        ? Number(object.circulatingSupply)
        : 0,
      lastUpdated: isSet(object.lastUpdated) ? String(object.lastUpdated) : "",
    };
  },

  toJSON(message: CoinmarketCapInfo): unknown {
    const obj: any = {};
    obj.quote = {};
    if (message.quote) {
      Object.entries(message.quote).forEach(([k, v]) => {
        obj.quote[k] = CoinmarketCapQuote.toJSON(v);
      });
    }
    message.maxSupply !== undefined && (obj.maxSupply = message.maxSupply);
    message.totalSupply !== undefined &&
      (obj.totalSupply = message.totalSupply);
    message.cmcRank !== undefined && (obj.cmcRank = message.cmcRank);
    message.circulatingSupply !== undefined &&
      (obj.circulatingSupply = message.circulatingSupply);
    message.lastUpdated !== undefined &&
      (obj.lastUpdated = message.lastUpdated);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CoinmarketCapInfo>, I>>(
    object: I
  ): CoinmarketCapInfo {
    const message = createBaseCoinmarketCapInfo();
    message.quote = Object.entries(object.quote ?? {}).reduce<{
      [key: string]: CoinmarketCapQuote;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = CoinmarketCapQuote.fromPartial(value);
      }
      return acc;
    }, {});
    message.maxSupply = object.maxSupply ?? 0;
    message.totalSupply = object.totalSupply ?? 0;
    message.cmcRank = object.cmcRank ?? 0;
    message.circulatingSupply = object.circulatingSupply ?? 0;
    message.lastUpdated = object.lastUpdated ?? "";
    return message;
  },
};

function createBaseCoinmarketCapInfo_QuoteEntry(): CoinmarketCapInfo_QuoteEntry {
  return { key: "", value: undefined };
}

export const CoinmarketCapInfo_QuoteEntry = {
  encode(
    message: CoinmarketCapInfo_QuoteEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      CoinmarketCapQuote.encode(
        message.value,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): CoinmarketCapInfo_QuoteEntry {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCoinmarketCapInfo_QuoteEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = CoinmarketCapQuote.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CoinmarketCapInfo_QuoteEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value)
        ? CoinmarketCapQuote.fromJSON(object.value)
        : undefined,
    };
  },

  toJSON(message: CoinmarketCapInfo_QuoteEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? CoinmarketCapQuote.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CoinmarketCapInfo_QuoteEntry>, I>>(
    object: I
  ): CoinmarketCapInfo_QuoteEntry {
    const message = createBaseCoinmarketCapInfo_QuoteEntry();
    message.key = object.key ?? "";
    message.value =
      object.value !== undefined && object.value !== null
        ? CoinmarketCapQuote.fromPartial(object.value)
        : undefined;
    return message;
  },
};

function createBaseCoinmarketCapMetadata(): CoinmarketCapMetadata {
  return {
    website: [],
    twitter: [],
    reddit: [],
    sourceCode: [],
    chat: [],
    technicalDoc: [],
  };
}

export const CoinmarketCapMetadata = {
  encode(
    message: CoinmarketCapMetadata,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.website) {
      writer.uint32(10).string(v!);
    }
    for (const v of message.twitter) {
      writer.uint32(18).string(v!);
    }
    for (const v of message.reddit) {
      writer.uint32(26).string(v!);
    }
    for (const v of message.sourceCode) {
      writer.uint32(34).string(v!);
    }
    for (const v of message.chat) {
      writer.uint32(42).string(v!);
    }
    for (const v of message.technicalDoc) {
      writer.uint32(50).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CoinmarketCapMetadata {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCoinmarketCapMetadata();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.website.push(reader.string());
          break;
        case 2:
          message.twitter.push(reader.string());
          break;
        case 3:
          message.reddit.push(reader.string());
          break;
        case 4:
          message.sourceCode.push(reader.string());
          break;
        case 5:
          message.chat.push(reader.string());
          break;
        case 6:
          message.technicalDoc.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CoinmarketCapMetadata {
    return {
      website: Array.isArray(object?.website)
        ? object.website.map((e: any) => String(e))
        : [],
      twitter: Array.isArray(object?.twitter)
        ? object.twitter.map((e: any) => String(e))
        : [],
      reddit: Array.isArray(object?.reddit)
        ? object.reddit.map((e: any) => String(e))
        : [],
      sourceCode: Array.isArray(object?.sourceCode)
        ? object.sourceCode.map((e: any) => String(e))
        : [],
      chat: Array.isArray(object?.chat)
        ? object.chat.map((e: any) => String(e))
        : [],
      technicalDoc: Array.isArray(object?.technicalDoc)
        ? object.technicalDoc.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: CoinmarketCapMetadata): unknown {
    const obj: any = {};
    if (message.website) {
      obj.website = message.website.map((e) => e);
    } else {
      obj.website = [];
    }
    if (message.twitter) {
      obj.twitter = message.twitter.map((e) => e);
    } else {
      obj.twitter = [];
    }
    if (message.reddit) {
      obj.reddit = message.reddit.map((e) => e);
    } else {
      obj.reddit = [];
    }
    if (message.sourceCode) {
      obj.sourceCode = message.sourceCode.map((e) => e);
    } else {
      obj.sourceCode = [];
    }
    if (message.chat) {
      obj.chat = message.chat.map((e) => e);
    } else {
      obj.chat = [];
    }
    if (message.technicalDoc) {
      obj.technicalDoc = message.technicalDoc.map((e) => e);
    } else {
      obj.technicalDoc = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CoinmarketCapMetadata>, I>>(
    object: I
  ): CoinmarketCapMetadata {
    const message = createBaseCoinmarketCapMetadata();
    message.website = object.website?.map((e) => e) || [];
    message.twitter = object.twitter?.map((e) => e) || [];
    message.reddit = object.reddit?.map((e) => e) || [];
    message.sourceCode = object.sourceCode?.map((e) => e) || [];
    message.chat = object.chat?.map((e) => e) || [];
    message.technicalDoc = object.technicalDoc?.map((e) => e) || [];
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
