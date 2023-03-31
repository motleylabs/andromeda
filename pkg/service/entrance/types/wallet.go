package types

type CollectedCollection struct {
	Name           string  `json:"name"`
	Image          string  `json:"image"`
	EstimatedValue string  `json:"estimatedValue"`
	NFTsOwned      int     `json:"nftsOwned"`
	Symbol         string  `json:"symbol"`
	FloorPrice     float64 `json:"floorPrice"`
}

type UserNFT struct {
	Collections []CollectedCollection `json:"collections"`
	NFTs        []NFT                 `json:"nfts"`
}
