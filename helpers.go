package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/gagliardetto/solana-go"
	borsh "github.com/near/borsh-go"
	globals "github.com/solana-nft-golang-metadata/globals"
	responses "github.com/solana-nft-golang-metadata/jsonrpc-responses"
)

func derivePDA(mint string) solana.PublicKey {
	mintDecoded := base58.Decode(mint)
	publicKey := solana.MustPublicKeyFromBase58(globals.METADATA_PUBKEY)
	seed := [][]byte{
		[]byte(globals.METAPLEX_SEED),
		base58.Decode(globals.METADATA_PUBKEY),
		mintDecoded,
	}

	PDA, _, err := solana.FindProgramAddress(seed, publicKey)
	if err != nil {
		fmt.Println(err)
	}

	return PDA
}

func getURI(accountInfo *responses.GetAccountInfo) string {
	var sanitizedURI string
	if accountInfo.Value.Data != nil {
		decoded, err := b64.StdEncoding.DecodeString(accountInfo.Value.Data[0])
		if err != nil {
			fmt.Println(err)
		}
		mm := new(globals.MetaplexMeta)
		borsh.Deserialize(mm, decoded)

		uri := mm.Data.Uri
		sanitizedURI = strings.Replace(uri, "\u0000", "", -1)
	}

	return sanitizedURI
}

func requestNFTMetadata(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body, nil
}
