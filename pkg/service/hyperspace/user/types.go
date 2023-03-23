package user

import "andromeda/pkg/service/hyperspace/common"

type OfferCondition struct {
	BuyerAddress string  `json:"buyer_address"`
	ActionType   *string `json:"action_type,omitempty"`
}

type OfferParams struct {
	Condition      OfferCondition           `json:"condition"`
	OrderBy        common.OrderConfig       `json:"order_by"`
	PaginationInfo *common.PaginationConfig `json:"pagination_info,omitempty"`
}
