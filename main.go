package main

import (
	"fmt"

	"github.com/ybbus/jsonrpc/v2"

	base58 "github.com/btcsuite/btcutil/base58"
	sdk "github.com/gagliardetto/solana-go"

	globals "github.com/solana-nft-golang-metadata/globals"
	methods "github.com/solana-nft-golang-metadata/jsonrpc-methods"
	responses "github.com/solana-nft-golang-metadata/jsonrpc-responses"
)

func main() {
	rpcClient := jsonrpc.NewClient("https://api.mainnet-beta.solana.com")

	accountOwner := "CDwhBZtT72QESpACGtf2mrvfT1xdCmdMiTMPMwb78sn7"

	tokenAccountsByOwner, err := methods.GetTokenAccountsByOwner(rpcClient, accountOwner)

	if err != nil {
		fmt.Println(err)
	}

	value := tokenAccountsByOwner.Value

	nftAccounts := locateNFTAccounts(value)

	// Base58 encoding of the mint address
	// Base58 econding of the metadata public key
	// metaplex seed constant

	derivePDAs(nftAccounts)

}

func locateNFTAccounts(value []responses.Value) []responses.Value {
	var nftAccounts []responses.Value

	for _, account := range value {
		amount := account.Account.Data.Parsed.Info.TokenAmount.Amount
		decimals := account.Account.Data.Parsed.Info.TokenAmount.Decimals
		if amount == 1 && decimals == 0 {
			nftAccounts = append(nftAccounts, account)
		}
	}
	return nftAccounts
}

func derivePDAs(nftAccounts []responses.Value) {
	for _, value := range nftAccounts {
		mint := value.Account.Data.Parsed.Info.Mint

		base58Mint := base58.Decode(mint)

		publicKey := sdk.PublicKeyFromBytes(base58Mint)

		seed := [][]byte{
			[]byte(base58.Decode(globals.METADATA_ACCOUNT_PUBKEY)),
			[]byte(globals.METADATA_NAME),
		}

		fmt.Println(sdk.FindProgramAddress(seed, publicKey))
	}
}
