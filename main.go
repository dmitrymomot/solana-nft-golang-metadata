package main

import (
	"fmt"
	"reflect"

	"github.com/ybbus/jsonrpc/v2"

	globals "github.com/solana-nft-golang-metadata/globals"
	methods "github.com/solana-nft-golang-metadata/jsonrpc-methods"
)

func main() {
	// accountOwnerExample()

	mintAddressExample()
}

// AllNFTsForAddress gets metadata for all NFTs owned by a solana address
// Returns a slice of byte slices containing each NFT metadata json
// TODO - implement a way to do this concurrently because waiting until the end is slow
func AllNFTsForAddress(address string) ([][]byte, error) {
	rpcClient := jsonrpc.NewClient(globals.SOLANA_MAINNET)
	tokenAccountsByOwner, err := methods.GetTokenAccountsByOwner(rpcClient, address)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var NFTs [][]byte
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
func NFTMetadata(address string) ([]byte, error) {
	rpcClient := jsonrpc.NewClient(globals.SOLANA_MAINNET)
	programAddress := derivePDA(address)
	accountInfo, err := methods.GetAccountInfo(rpcClient, programAddress.String())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if accountInfo == nil {
		fmt.Println("account info is nil")
		return nil, nil
	}

	if reflect.ValueOf(accountInfo.Value).IsZero() {
		fmt.Println("account info value is nil")
		return nil, nil
	}

	nftURI := getURI(accountInfo)
	return requestNFTMetadata(nftURI)
}
