package examples

import (
	"fmt"

	metadata "github.com/based58/solana-nft-golang-metadata/pkg"
)

func accountOwnerExample() {
	accountOwner := "9j3Mcte8bwh97SsUBqZgApG5xieGCaXHYKCjFSwxZ14t"
	allNFTs, err := metadata.AllNFTsForAddress(accountOwner)
	if err != nil {
		fmt.Println(err)
	}
	for _, metadata := range allNFTs {
		fmt.Println(string(metadata))
	}
}

func mintAddressExample() {
	mintAddress := "9qtN3RDr8sykdzmJoDjYbAAQmjoQcAfsqj8ipT6rh413"
	metadata, err := metadata.NFTMetadata(mintAddress)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(metadata))
}
