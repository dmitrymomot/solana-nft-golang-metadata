package main

import (
	"fmt"

	"github.com/ybbus/jsonrpc/v2"

	methods "github.com/solana-nft-golang-metadata/jsonrpc-methods"
	responses "github.com/solana-nft-golang-metadata/jsonrpc-responses"
	utils "github.com/solana-nft-golang-metadata/utils"
)

func main() {
	rpcClient := jsonrpc.NewClient("https://api.mainnet-beta.solana.com")

	address := "6PiuodUPFTndjTiy2X3u1znTgbiUsQs7o5pUgKdLW6mk"

	tokenAccountsByOwner, err := methods.GetTokenAccountsByOwner(rpcClient, address)

	if err != nil {
		fmt.Println(err)
	}

	value := tokenAccountsByOwner.Value

	var tokenAccounts []responses.TokenAccount

	for _, tokenAccount := range value {
		tokenAccounts = append(tokenAccounts, tokenAccount.TokenAccount)
	}

	nftAccounts := locateNFTAccounts(tokenAccounts)

	utils.ToJson(nftAccounts)
}

func locateNFTAccounts(tokenAccountsByOwner []responses.TokenAccount) []responses.TokenAccount {
	var nftAccounts []responses.TokenAccount

	for _, tokenAccount := range tokenAccountsByOwner {
		amount := tokenAccount.Account.Data.Parsed.Info.TokenAmount.Amount
		fmt.Println("amount", amount)
		decimals := tokenAccount.Account.Data.Parsed.Info.TokenAmount.Decimals
		fmt.Println("decimals", amount)
		if amount == 1 && decimals == 0 {
			nftAccounts = append(nftAccounts, tokenAccount)
		}
	}
	return nftAccounts
}
