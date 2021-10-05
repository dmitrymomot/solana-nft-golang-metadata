package main

import "fmt"

func accountOwnerExample() {
	accountOwner := "9j3Mcte8bwh97SsUBqZgApG5xieGCaXHYKCjFSwxZ14t"
	allNFTs, err := AllNFTsForAddress(accountOwner)
	if err != nil {
		fmt.Println(err)
	}
	for _, metadata := range allNFTs {
		fmt.Println(string(metadata))
	}
}

func mintAddressExample() {
	mintAddress := "3wW42N6Q5JcqbewpEZUHsdjU7XEYMf5p4CZADZmTdfEi"
	metadata, err := NFTMetadata(mintAddress)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(metadata))
}
