package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToPriceMasterUpdates(priceMaster dpfm_api_input_reader.PriceMaster) *PriceMasterUpdates {
	data := priceMaster

	return &PriceMasterUpdates{
		ConditionType:              data.ConditionType,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
		ConditionRateValue:         data.ConditionRateValue,
		ConditionRateValueUnit:     data.ConditionRateValueUnit,
		ConditionRateRatio:         data.ConditionRateRatio,
		ConditionRateRatioUnit:     data.ConditionRateRatioUnit,
		ConditionIsDeleted:         data.ConditionIsDeleted,
	}
}
