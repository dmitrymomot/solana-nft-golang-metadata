package metadata

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/gagliardetto/solana-go"
	borsh "github.com/near/borsh-go"
	"github.com/ybbus/jsonrpc/v2"
)

// AllNFTsForAddress gets metadata for all NFTs owned by a solana address
// Returns a slice of byte slices containing each NFT metadata json
// TODO - implement a way to do this concurrently because waiting until the end is slow
func AllNFTsForAddress(address string) ([]string, error) {
	rpcClient := jsonrpc.NewClient(SOLANA_MAINNET)
	tokenAccountsByOwner, err := GetTokenAccountsByOwner(rpcClient, address)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var NFTs []string

	for _, tokenAccount := range tokenAccountsByOwner.Value {
		if tokenAccount.Account.Data.Parsed.Info.TokenAmount.Amount != 1 &&
			tokenAccount.Account.Data.Parsed.Info.TokenAmount.Decimals != 0 {
			continue
		}
		mint := tokenAccount.Account.Data.Parsed.Info.Mint
		metadata, err := NFTMetadata(mint)
		if err != nil {
			fmt.Println(err)
			continue
		}

		NFTs = append(NFTs, metadata)
	}

	return NFTs, nil
}

// NFTMetadata gets metadata for an NFT using the solana token mint address
// Returns a byte slice containing the NFT metadata json
func NFTMetadata(address string) (string, error) {
	rpcClient := jsonrpc.NewClient(SOLANA_MAINNET)
	programAddress := derivePDA(address)
	accountInfo, err := GetAccountInfo(rpcClient, programAddress.String())
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if accountInfo == nil {
		return "", nil
	}

	if reflect.ValueOf(accountInfo.Value).IsZero() {
		return "", nil
	}

	nftURI := getURI(accountInfo)
	return requestNFTMetadata(nftURI)
}

func derivePDA(mint string) solana.PublicKey {
	mintDecoded := base58.Decode(mint)
	publicKey := solana.MustPublicKeyFromBase58(METADATA_PUBKEY)
	seed := [][]byte{
		[]byte(METAPLEX_SEED),
		base58.Decode(METADATA_PUBKEY),
		mintDecoded,
	}

	PDA, _, err := solana.FindProgramAddress(seed, publicKey)
	if err != nil {
		fmt.Println(err)
	}

	return PDA
}

func getURI(accountInfo *AccountInfo) string {
	var sanitizedURI string
	if accountInfo.Value.Data != nil {
		decoded, err := b64.StdEncoding.DecodeString(accountInfo.Value.Data[0])
		if err != nil {
			fmt.Println(err)
		}
		mm := new(MetaplexMeta)
		borsh.Deserialize(mm, decoded)

		uri := mm.Data.Uri
		sanitizedURI = strings.Replace(uri, "\u0000", "", -1)
	}

	return sanitizedURI
}

func requestNFTMetadata(uri string) (string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	json := string(body)

	return json, nil
}
