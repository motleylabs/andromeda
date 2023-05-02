package nft

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"andromeda/pkg/service/web3"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type HSBuyParams struct {
	Buyer        string  `json:"buyer_address"`
	Price        float64 `json:"price"`
	TokenAddress string  `json:"token_address"`
	BuyerBroker  string  `json:"buyer_broker"`
}

type HSBuyError struct {
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
}

type HSBuyResponse struct {
	Data      []byte      `json:"data"`
	StdBuffer []byte      `json:"stdBuffer"`
	Error     *HSBuyError `json:"error"`
}

type SolanaFMResponseV0Signed struct {
	Data []byte `json:"data"`
}

type SolanaFMResponseV0 struct {
	TxSigned SolanaFMResponseV0Signed `json:"txSigned"`
}

type SolanaFMResponse struct {
	Error *string            `json:"error"`
	V0    SolanaFMResponseV0 `json:"v0"`
}

const SOLANA_FM_HYPER = "https://hyper.solana.fm/v3/instructions/buy_now"

func GetMagicEdenBuffer(params *types.BuyParams) (*types.BuyRes, error) {
	price, err := strconv.ParseInt(params.Price, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("price param is not valid")
	}

	v := url.Values{}
	v.Set("network", "mainnet")
	v.Set("seller", params.Seller)
	v.Set("buyer", params.Buyer)
	v.Set("auctionHouseAddress", params.AuctionHouseAddress)
	v.Set("tokenMint", params.Mint)
	v.Set("price", fmt.Sprintf("%v", float64(price)/common.LAMPORTS_PER_SOL))
	v.Set("sellerExpiry", "-1")

	ata, _, err := web3.GetATA(params.Seller, params.Mint)
	if err != nil {
		return nil, err
	}

	v.Set("tokenATA", ata.String())

	res, err := request.ProcessBrowserGet(fmt.Sprintf("%v?%v", SOLANA_FM_HYPER, v.Encode()))
	if err != nil {
		return nil, err
	}

	var buyRes SolanaFMResponse
	if err := json.Unmarshal(res, &buyRes); err != nil {
		return nil, err
	}

	if buyRes.Error != nil {
		return nil, fmt.Errorf(*buyRes.Error)
	}

	if len(buyRes.V0.TxSigned.Data) > 0 {
		return &types.BuyRes{
			Buffer: buyRes.V0.TxSigned.Data,
		}, nil
	}

	return nil, fmt.Errorf("no data returned")
}

func GetBuyNowTx(params *types.BuyParams) (*types.BuyRes, error) {
	if params.AuctionHouseAddress == "E8cU1WiRWjanGxmn96ewBgk9vPTcL6AEZ1t6F6fkgUWe" {
		return GetMagicEdenBuffer(params)
	}

	price, err := strconv.ParseInt(params.Price, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("price param is not valid")
	}

	hsParams := HSBuyParams{
		Buyer:        params.Buyer,
		BuyerBroker:  params.BuyerBroker,
		Price:        float64(price) / common.LAMPORTS_PER_SOL,
		TokenAddress: params.Mint,
	}
	payload, err := json.Marshal(hsParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/create-buy-tx", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var buyRes HSBuyResponse
	if err := json.Unmarshal(res, &buyRes); err != nil {
		return nil, err
	}

	if buyRes.Error != nil {
		return nil, fmt.Errorf(buyRes.Error.Message)
	}

	if len(buyRes.StdBuffer) > 0 {
		return &types.BuyRes{
			Buffer: buyRes.StdBuffer,
			Data:   buyRes.Data,
		}, nil
	}

	return nil, fmt.Errorf("no data returned")
}
