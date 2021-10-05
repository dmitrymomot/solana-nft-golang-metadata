package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ybbus/jsonrpc/v2"

	b64 "encoding/base64"
	"encoding/json"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/gagliardetto/solana-go"
	sdk "github.com/gagliardetto/solana-go"
	borsh "github.com/near/borsh-go"

	globals "github.com/solana-nft-golang-metadata/globals"
	methods "github.com/solana-nft-golang-metadata/jsonrpc-methods"
	responses "github.com/solana-nft-golang-metadata/jsonrpc-responses"
)

func main() {
	rpcClient := jsonrpc.NewClient("https://api.mainnet-beta.solana.com")

	accountOwner := "9j3Mcte8bwh97SsUBqZgApG5xieGCaXHYKCjFSwxZ14t"

	tokenAccountsByOwner, err := methods.GetTokenAccountsByOwner(rpcClient, accountOwner)

	if err != nil {
		fmt.Println(err)
	}

	value := tokenAccountsByOwner.Value

	nftAccounts := locateNFTAccounts(value)

	programAddresses := derivePDAs(nftAccounts)

	accountInfoList := getAccountInfoList(rpcClient, programAddresses)

	nftURIs := getURIs(accountInfoList)

	serializedNFTs := serializeNFTs(nftURIs)

	fmt.Println(serializedNFTs)
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

func getAccountInfoList(rpcClient jsonrpc.RPCClient, publicKeys []sdk.PublicKey) []responses.GetAccountInfo {
	var accountInfoList []responses.GetAccountInfo
	for _, address := range publicKeys {
		accountInfo, err := methods.GetAccountInfo(rpcClient, address.String())
		if err != nil {
			fmt.Println(err)
		}
		if accountInfo != nil {
			accountInfoList = append(accountInfoList, *accountInfo)
		}
	}
	return accountInfoList
}

func getURIs(accountInfoList []responses.GetAccountInfo) []string {
	var URIs []string
	for _, accountInfo := range accountInfoList {
		if accountInfo.Value.Data != nil {
			decoded, err := b64.StdEncoding.DecodeString(accountInfo.Value.Data[0])
			if err != nil {
				fmt.Println(err)
			}
			mm := new(globals.MetaplexMeta)
			borsh.Deserialize(mm, decoded)

			uri := mm.Data.Uri
			sanitizedURI := strings.Replace(uri, "\u0000", "", -1)

			URIs = append(URIs, sanitizedURI)
		}
	}
	return URIs
}

func serializeNFTs(URIs []string) []globals.MetaplexJSONStructure {

	var AddressNFTMetadata []globals.MetaplexJSONStructure

	for _, uri := range URIs {
		resp, err := http.Get(uri)
		if err != nil {
			fmt.Println(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var nftMetadata = new(globals.MetaplexJSONStructure)

		err = json.Unmarshal(body, &nftMetadata)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(nftMetadata)
		AddressNFTMetadata = append(AddressNFTMetadata, *nftMetadata)
	}

	return AddressNFTMetadata
}
