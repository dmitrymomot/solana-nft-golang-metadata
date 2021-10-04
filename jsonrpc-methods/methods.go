package methods

import (
	"fmt"

	"github.com/ybbus/jsonrpc/v2"

	responses "github.com/solana-nft-golang-metadata/jsonrpc-responses"
)

const ProgramId = "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"

func GetTokenAccountsByOwner(rpcClient jsonrpc.RPCClient, accountOwner string) (*responses.GetTokenAccountsByOwner, error) {
	response, err := rpcClient.Call("getTokenAccountsByOwner", accountOwner, map[string]interface{}{
		"programId": ProgramId,
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
