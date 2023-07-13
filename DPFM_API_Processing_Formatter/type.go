package dpfm_api_processing_formatter

type PriceMasterUpdates struct {
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	Product                    string   `json:"Product"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionType              string   `json:"ConditionType"`
	ConditionRateValue         float32  `json:"ConditionRateValue"`
	ConditionRateValueUnit     int      `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     int      `json:"ConditionScaleQuantity"`
}
