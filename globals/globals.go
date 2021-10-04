package globals

import (
	sdk "github.com/gagliardetto/solana-go"
)

const METAPLEX_SEED = "metadata"
const METADATA_PUBKEY = "metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s"

type MetaplexMeta struct {
	Key              byte
	Update_authority sdk.PublicKey
	Mint             sdk.PublicKey
	Data             MetaplexData
}

type MetaplexData struct {
	Name   string
	Symbol string
	Uri    string
}

// data class MetaplexMeta(
//     @FieldOrder(0) val key: Byte,
//     @FieldOrder(1) val update_authority: PublicKey,
//     @FieldOrder(2) val mint: PublicKey,
//     @FieldOrder(3) val data: MetaplexData
// ) : BorshCodable

// data class MetaplexData(
//     @FieldOrder(0) val name: String,
//     @FieldOrder(1) val symbol: String,
//     @FieldOrder(2) val uri: String
// ) : BorshCodable
