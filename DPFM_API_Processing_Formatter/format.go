package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToPriceMasterUpdates(priceMaster dpfm_api_input_reader.PriceMaster) *PriceMasterUpdates {
	data := priceMaster

	return &PriceMasterUpdates{
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			Product:                    data.Product,
			ConditionValidityStartDate: data.ConditionValidityStartDate,
			ConditionValidityEndDate:   data.ConditionValidityEndDate,
			ConditionType:              data.ConditionType,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionRateValueUnit:     data.ConditionRateValueUnit,
			ConditionScaleQuantity:     data.ConditionScaleQuantity,
	}
}
