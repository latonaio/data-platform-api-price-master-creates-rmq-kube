package dpfm_api_input_reader

type EC_MC struct {
}

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	PriceMaster         PriceMaster `json:"PriceMaster"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	OrderID             *int        `json:"order_id"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type PriceMaster struct {
	BusinessPartner            *int     `json:"BusinessPartner"`
	ConditionRecordCategory    string   `json:"ConditionRecordCategory"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              string   `json:"ConditionType"`
	ConditionValidityEndDate   *string  `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate *string  `json:"ConditionValidityStartDate"`
	Product                    string   `json:"Product"`
	Customer                   *int     `json:"Customer"`
	Supplier                   *int     `json:"Supplier"`
	CreationDate               *string  `json:"CreationDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     string   `json:"ConditionRateValueUnit"`
	ConditionRateRatio         *float32 `json:"ConditionRateRatio"`
	ConditionRateRatioUnit     string   `json:"ConditionRateRatioUnit"`
	ConditionCurrency          string   `json:"ConditionCurrency"`
	BaseUnit                   string   `json:"BaseUnit"`
	ConditionIsDeleted         *bool    `json:"ConditionIsDeleted"`
}
