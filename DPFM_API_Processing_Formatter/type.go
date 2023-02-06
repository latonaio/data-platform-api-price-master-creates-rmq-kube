package dpfm_api_processing_formatter

type PriceMasterUpdates struct {
	ConditionType              string   `json:"ConditionType"`
	ConditionValidityEndDate   *string  `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate *string  `json:"ConditionValidityStartDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     string   `json:"ConditionRateValueUnit"`
	ConditionRateRatio         *float32 `json:"ConditionRateRatio"`
	ConditionRateRatioUnit     string   `json:"ConditionRateRatioUnit"`
	ConditionIsDeleted         *bool    `json:"ConditionIsDeleted"`
}
