package main

import (
	"encoding/json"
	"fmt"

	"github.com/ybbus/jsonrpc/v2"
)

func main() {
	fmt.Println("GO")
	rpcClient := jsonrpc.NewClient("https://api.mainnet-beta.solana.com")
	response, err := rpcClient.Call("getTokenAccountsByOwner", "rzD4a16cyZ4Ruo8xb3HjbB19pk6MZRqbifHHkznt5VD", map[string]interface{}{
		"programId": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
	}, map[string]interface{}{
		"encoding": "jsonParsed",
	})
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.MarshalIndent(response.Result, "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
