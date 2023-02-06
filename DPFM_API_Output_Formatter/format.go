package dpfm_api_output_formatter

import (
	dpfm_api_processing_formatter "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-price-master-creates-rmq-kube/sub_func_complementer"
)

func ConvertToPriceMasterCreates(subfuncSDC *sub_func_complementer.SDC) *PriceMaster {
	data := subfuncSDC.Message.PriceMaster

	priceMaster := &PriceMaster{
		SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
		Buyer:                      data.Buyer,
		Seller:                     data.Seller,
		ConditionRecord:            data.ConditionRecord,
		ConditionSequentialNumber:  data.ConditionSequentialNumber,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
		Product:                    data.Product,
		ConditionType:              data.ConditionType,
		CreationDate:               data.CreationDate,
		ConditionRateValue:         data.ConditionRateValue,
		ConditionRateValueUnit:     data.ConditionRateValueUnit,
		ConditionRateRatio:         data.ConditionRateRatio,
		ConditionRateRatioUnit:     data.ConditionRateRatioUnit,
		BaseUnit:                   data.BaseUnit,
		ConditionIsDeleted:         data.ConditionIsDeleted,
	}

	return priceMaster
}

func ConvertToPriceMasterUpdates(priceMasterUpdates *dpfm_api_processing_formatter.PriceMasterUpdates) *PriceMaster {
	data := priceMasterUpdates

	priceMaster := &PriceMaster{
		ConditionType:              data.ConditionType,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
		ConditionRateValue:         data.ConditionRateValue,
		ConditionRateValueUnit:     data.ConditionRateValueUnit,
		ConditionRateRatio:         data.ConditionRateRatio,
		ConditionRateRatioUnit:     data.ConditionRateRatioUnit,
		ConditionIsDeleted:         data.ConditionIsDeleted,
	}

	return priceMaster
}
