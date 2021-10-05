package methods

import (
	"fmt"

	"github.com/ybbus/jsonrpc/v2"

	globals "github.com/solana-nft-golang-metadata/globals"
	responses "github.com/solana-nft-golang-metadata/jsonrpc-responses"
)

func GetTokenAccountsByOwner(rpcClient jsonrpc.RPCClient, accountOwner string) (*responses.GetTokenAccountsByOwner, error) {
	response, err := rpcClient.Call("getTokenAccountsByOwner", accountOwner, map[string]interface{}{
		"programId": globals.PROGRAM_ID,
	}, map[string]interface{}{
		"encoding": "jsonParsed",
	})

	if err != nil {
		fmt.Println(err)
	}

	var tokenAccountsByOwner *responses.GetTokenAccountsByOwner
	err = response.GetObject(&tokenAccountsByOwner)

	if err != nil {
		return nil, err
	}

	return tokenAccountsByOwner, nil
}

func GetAccountInfo(rpcClient jsonrpc.RPCClient, account string) (*responses.GetAccountInfo, error) {
	response, err := rpcClient.Call("getAccountInfo", account, map[string]interface{}{
		"encoding": "base64",
	})

	if err != nil {
		fmt.Println(err)
	}

	var accountInfo *responses.GetAccountInfo
	err = response.GetObject(&accountInfo)

	if err != nil {
		return nil, err
	}

	return accountInfo, nil
}
