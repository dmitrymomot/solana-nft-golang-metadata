package main

import (
	"flag"
	"fmt"

	metadata "github.com/solana-nft-golang-metadata/pkg"
)

func main() {
	// Basic flag declarations are available for string,
	// integer, and boolean options. Here we declare a
	// string flag `word` with a default value `"foo"`
	// and a short description. This `flag.String` function
	// returns a string pointer (not a string value);
	// we'll see how to use this pointer below.
	command := flag.String("command", "account", "a get command, either 'account' or 'nft'")
	address := flag.String("address", "", "a solana address")

	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()

	// Here we'll just dump out the parsed options and
	// any trailing positional arguments. Note that we
	// need to dereference the pointers with e.g. `*wordPtr`
	// to get the actual option values.
	fmt.Println("command:", *command)
	fmt.Println("address:", *address)

	if *command == "account" {
		fmt.Println(metadata.AllNFTsForAddress(*address))
	} else if *command == "nft" {
		fmt.Println(metadata.NFTMetadata(*address))
	}
}
