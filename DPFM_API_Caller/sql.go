package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-price-master-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var priceMaster *dpfm_api_output_formatter.PriceMaster
	for _, fn := range accepter {
		switch fn {
		case "PriceMaster":
			c.priceMasterCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
			priceMaster = dpfm_api_output_formatter.ConvertToPriceMasterCreates(subfuncSDC)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		PriceMaster: priceMaster,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var priceMaster *dpfm_api_output_formatter.PriceMaster
	for _, fn := range accepter {
		switch fn {
		case "PriceMaster":
			priceMaster = c.priceMasterUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		PriceMaster: priceMaster,
	}

	return data
}

func (c *DPFMAPICaller) priceMasterCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PriceMaster {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_price_master_price_master_dataの更新
	priceMasterData := subfuncSDC.Message.PriceMaster
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": priceMasterData, "function": "PriceMasterPriceMaster", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToPriceMasterCreates(subfuncSDC)

	return data
}

func (c *DPFMAPICaller) priceMasterUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PriceMaster {
	req := dpfm_api_processing_formatter.ConvertToPriceMasterUpdates(input.PriceMaster)

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "PriceMasterPriceMaster", "runtime_session_id": 123})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("PriceMaster Data cannot insert")
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "PriceMaster Data cannot insert"
		return nil
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToPriceMasterUpdates(req)

	return data
}
