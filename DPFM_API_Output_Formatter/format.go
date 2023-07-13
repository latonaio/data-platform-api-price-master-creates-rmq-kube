package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToPriceMasterCreates(sdc *dpfm_api_input_reader.SDC) (*PriceMaster, error) {
	data := sdc.PriceMaster

	priceMaster, err := TypeConverter[*PriceMaster](data)
	if err != nil {
		return nil, err
	}

	return priceMaster, nil
}

func ConvertToPriceMasterUpdates(priceMasterData dpfm_api_input_reader.PriceMaster) (*PriceMaster, error) {
	data := priceMasterData

	priceMaster, err := TypeConverter[*PriceMaster](data)
	if err != nil {
		return nil, err
	}

	return priceMaster, nil
}
