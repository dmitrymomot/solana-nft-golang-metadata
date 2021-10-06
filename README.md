# solana-nft-golang-metadata

This is a package for getting metadata for solana NFTs.

I followed this guide pretty closely while making it.
https://gist.github.com/creativedrewy/9bce794ff278aae23b64e6dc8f10e906

You can run this to get all nft metadata that a solana wallet has

```./sol-nft -address={solana wallet address that holds nfts} -command=account```

Or get metadata for an individual NFT

```./sol-nft -address={solana nft mint address} -command=nft```

You can also use the methods in `/pkg` as a part of your own solana NFT project.

I originally had this getting all NFT metadata concurrently, but it's very tough to get around rate limiting.
For this reason, if you are building a web app it's best to handle interactions with the solana mainnet on the client side.

