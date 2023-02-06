package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) productExistenceConf(input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	if &input.PriceMaster.Product != nil {
		if len(input.PriceMaster.Product) == 0 {
			*exconfErrMsg = "cannot specify null keys"
			return
		}
	}
	res, err := c.productMasterGeneralExistenceConfRequest(*&input.PriceMaster.Product, input, existenceMap, mtx, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
	}
	if res != "" {
		*exconfErrMsg = res
	}
}

func (c *ExistenceConf) productMasterGeneralExistenceConfRequest(product string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	key := "Product"
	keys := newResult(map[string]interface{}{
		"Product": product,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.ProductMasterGeneralReturn.Product = product

	exist, err = c.exconfRequest(req, key, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}
