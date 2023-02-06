package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *ExistenceConf) sellerExistenceConf(input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	if input.PriceMaster.Seller == nil {
		*exconfErrMsg = "cannot specify null keys"
		return
	}
	res, err := c.bpGeneralExistenceConfRequest(*input.PriceMaster.Seller, input, existenceMap, mtx, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
	}
	if res != "" {
		*exconfErrMsg = res
	}
}
