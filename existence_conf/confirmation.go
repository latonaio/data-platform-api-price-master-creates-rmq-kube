package existence_conf

import (
	"context"
	dpfm_api_input_reader "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-price-master-creates-rmq-kube/config"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ExistenceConf struct {
	ctx context.Context

	c             *config.Conf
	queueToMapper exconfQueueMapper
	rmq           *rabbitmq.RabbitmqClient
}

func NewExistenceConf(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient) *ExistenceConf {
	return &ExistenceConf{
		ctx:           ctx,
		c:             c,
		queueToMapper: NewExconfQueueMapper(c),
		rmq:           rmq,
	}
}

// Confirm returns existenceMap, allExist, err
func (c *ExistenceConf) Conf(data *dpfm_api_input_reader.SDC, ssdc *dpfm_api_output_formatter.SDC, l *logger.Logger) (allExist bool, errs []error) {
	var resMsg string
	existenceMap := make([]bool, 0, 12)
	wg := sync.WaitGroup{}
	mtx := &sync.Mutex{}

	wg.Add(6)
	go c.supplyChainRelationshipGeneralExistenceConf(data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
	go c.buyerExistenceConf(data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
	go c.sellerExistenceConf(data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
	go c.productExistenceConf(data, &existenceMap, &resMsg, &errs, mtx, &wg, l)

	wg.Wait()

	if len(errs) != 0 {
		return false, errs
	}

	ssdc.ExconfResult = getBoolPtr(true)

	if len(resMsg) != 0 {
		ssdc.ExconfResult = getBoolPtr(false)
		ssdc.ExconfError = resMsg
		return false, nil
	}

	return true, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func jsonTypeConversion[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
func confKeyExistence(res map[string]interface{}) bool {
	if res == nil {
		return false
	}
	raw, ok := res["ExistenceConf"]
	exist := fmt.Sprintf("%v", raw)
	if ok {
		return strings.ToLower(exist) == "true"
	}

	return false
}
func (c *ExistenceConf) exconfRequest(req interface{}, key string, log *logger.Logger) (bool, error) {
	res, err := c.rmq.SessionKeepRequest(nil, c.queueToMapper[key], req)
	if err != nil {
		return false, xerrors.Errorf("response error: %w", err)
	}
	res.Success()
	exist := confKeyExistence(res.Data())
	log.Info(res.Data())
	return exist, nil
}

type result struct {
	keys map[string]interface{}
}

func newResult(keys map[string]interface{}) *result {
	return &result{
		keys: keys,
	}
}

func (r *result) fail() string {
	txt := ""
	for k, v := range r.keys {
		txt = fmt.Sprintf("%s%s:%v, ", k, v)
	}
	txt = fmt.Sprintf("%s does not exist", txt)
	return txt
}

func plantNotExist() string {
	return "plant data does not exist."
}
