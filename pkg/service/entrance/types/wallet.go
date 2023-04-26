package types

type CollectedCollection struct {
	ID             string `json:"id"`
	Slug           string `json:"slug"`
	Name           string `json:"name"`
	Image          string `json:"image"`
	NFTsOwned      int    `json:"nftsOwned"`
	EstimatedValue string `json:"estimatedValue"`
	FloorPrice     string `json:"floorPrice"`
}

type UserNFT struct {
	Collections []CollectedCollection `json:"collections"`
	NFTs        []NFT                 `json:"nfts"`
}
