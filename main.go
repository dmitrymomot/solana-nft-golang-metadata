package main

import (
	"fmt"

	"github.com/ybbus/jsonrpc/v2"

	b64 "encoding/base64"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/gagliardetto/solana-go"
	sdk "github.com/gagliardetto/solana-go"
	borsh "github.com/near/borsh-go"

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

	programAddresses := derivePDAs(nftAccounts)

	var accountInfoList []responses.GetAccountInfo

	for _, address := range programAddresses {
		accountInfo, err := methods.GetAccountInfo(rpcClient, address.String())
		if err != nil {
			fmt.Println(err)
		}
		accountInfoList = append(accountInfoList, *accountInfo)
	}

	utils.ToJson(accountInfoList)

	for _, accountInfo := range accountInfoList {
		y := new(globals.MetaplexMeta)
		decoded, err := b64.StdEncoding.DecodeString(accountInfo.Value.Data[0])
		if err != nil {
			fmt.Println(err)
		}
		borsh.Deserialize(y, decoded)
		fmt.Println(y)
	}
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
