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

type MetaplexJSONStructure struct {
	AnimationURL string `json:"animation_url"`
	Attributes   []struct {
		TraitType string `json:"trait_type"`
		Value     string `json:"value"`
	} `json:"attributes"`
	Collection struct {
		Family string `json:"family"`
		Name   string `json:"name"`
	} `json:"collection"`
	Description string `json:"description"`
	ExternalURL string `json:"external_url"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Properties  struct {
		Category string `json:"category"`
		Creators []struct {
			Address string `json:"address"`
			Share   int    `json:"share"`
		} `json:"creators"`
		Files []struct {
			Cdn  bool   `json:"cdn,omitempty"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"files"`
	} `json:"properties"`
	SellerFeeBasisPoints int    `json:"seller_fee_basis_points"`
	Symbol               string `json:"symbol"`
}