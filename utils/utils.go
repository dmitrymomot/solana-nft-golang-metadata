package utils

import (
	"encoding/json"
	"fmt"
)

func ToJson(a interface{}) {
	json, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json))
}
