package main

import (
	"fmt"

	"github.com/ybbus/jsonrpc/v2"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/gagliardetto/solana-go"
	sdk "github.com/gagliardetto/solana-go"

	globals "github.com/solana-nft-golang-metadata/globals"
	methods "github.com/solana-nft-golang-metadata/jsonrpc-methods"
	responses "github.com/solana-nft-golang-metadata/jsonrpc-responses"
	"github.com/solana-nft-golang-metadata/utils"
)

func main() {
	rpcClient := jsonrpc.NewClient("https://api.mainnet-beta.solana.com")

	accountOwner := "HbzndYqJhaCQdH1NqYVbjZqsW1dYwN4mbEw1icouTHdy"

	tokenAccountsByOwner, err := methods.GetTokenAccountsByOwner(rpcClient, accountOwner)

	if err != nil {
		fmt.Println(err)
	}

	value := tokenAccountsByOwner.Value

	nftAccounts := locateNFTAccounts(value)

	utils.ToJson(nftAccounts)

	// Base58 encoding of the mint address
	// Base58 econding of the metadata public key
	// metaplex seed constant

	programAddresses := derivePDAs(nftAccounts)

	utils.ToJson(programAddresses)
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

func derivePDAs(nftAccounts []responses.Value) []solana.PublicKey {
	var PDAs []solana.PublicKey

	for _, value := range nftAccounts {
		mint := value.Account.Data.Parsed.Info.Mint

		mintDecoded := base58.Decode(mint)

		publicKey := sdk.MustPublicKeyFromBase58(globals.METADATA_PUBKEY)

		seed := [][]byte{
			[]byte(globals.METAPLEX_SEED),
			base58.Decode(globals.METADATA_PUBKEY),
			mintDecoded,
		}
		PDA, _, err := sdk.FindProgramAddress(seed, publicKey)
		if err != nil {
			fmt.Println(err)
		}
		PDAs = append(PDAs, PDA)
	}

	return PDAs
}
