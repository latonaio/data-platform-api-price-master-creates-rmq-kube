package requests

type PriceMaster struct {
	SupplyChainRelationshipID  *int     `json:"SupplyChainRelationshipID"`
	Buyer                      *int     `json:"Buyer"`
	Seller                     *int     `json:"Seller"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionValidityEndDate   *string  `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate *string  `json:"ConditionValidityStartDate"`
	Product                    string   `json:"Product"`
	ConditionType              string   `json:"ConditionType"`
	CreationDate               *string  `json:"CreationDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     string   `json:"ConditionRateValueUnit"`
	ConditionRateRatio         *float32 `json:"ConditionRateRatio"`
	ConditionRateRatioUnit     string   `json:"ConditionRateRatioUnit"`
	BaseUnit                   string   `json:"BaseUnit"`
	ConditionIsDeleted         *bool    `json:"ConditionIsDeleted"`
}
