# solana-nft-golang-metadata

This is a package for getting metadata for solana NFTs.

You can run this to get all nft metadata that a solana wallet has

```./sol-nft -address={solana wallet address that holds nfts} -command=account```

Or get metadata for an individual NFT

```./sol-nft -address={solana nft mint address} -command=nft```

You can also use the methods in `/pkg` as a part of your own solana NFT project.
