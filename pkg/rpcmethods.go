package metadata

import (
	"fmt"

	"github.com/ybbus/jsonrpc/v2"
)

func GetTokenAccountsByOwner(rpcClient jsonrpc.RPCClient, accountOwner string) (*TokenAccountsByOwner, error) {
	response, err := rpcClient.Call("getTokenAccountsByOwner", accountOwner, map[string]interface{}{
		"programId": PROGRAM_ID,
	}, map[string]interface{}{
		"encoding": "jsonParsed",
	})

	if err != nil {
		fmt.Println(err)
	}

	var tokenAccountsByOwner *TokenAccountsByOwner
	err = response.GetObject(&tokenAccountsByOwner)

	if err != nil {
		return nil, err
	}

	return tokenAccountsByOwner, nil
}

func GetAccountInfo(rpcClient jsonrpc.RPCClient, account string) (*AccountInfo, error) {
	response, err := rpcClient.Call("getAccountInfo", account, map[string]interface{}{
		"encoding": "base64",
	})

	if err != nil {
		fmt.Println(err)
	}

	var accountInfo *AccountInfo
	err = response.GetObject(&accountInfo)

	if err != nil {
		return nil, err
	}

	return accountInfo, nil
}
