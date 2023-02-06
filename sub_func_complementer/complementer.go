package sub_func_complementer

import (
	"context"
	dpfm_api_input_reader "data-platform-api-price-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-price-master-creates-rmq-kube/config"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type SubFuncComplementer struct {
	ctx context.Context
	c   *config.Conf
	rmq *rabbitmq.RabbitmqClient
}

func NewSubFuncComplementer(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient) *SubFuncComplementer {
	return &SubFuncComplementer{
		ctx: ctx,
		c:   c,
		rmq: rmq,
	}
}

func (c *SubFuncComplementer) ComplementPriceMaster(input *dpfm_api_input_reader.SDC, subfuncSDC *SDC, l *logger.Logger) error {
	s := &SDC{}
	res, err := c.rmq.SessionKeepRequest(nil, c.c.RMQ.QueueToSubFunc()["PriceMaster"], input)
	if err != nil {
		return err
	}
	res.Success()

	err = json.Unmarshal(res.Raw(), s)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(s.Message)
	msg := &Message{}
	err = json.Unmarshal(b, msg)
	if err != nil {
		return err
	}
	subfuncSDC.SubfuncResult = s.SubfuncResult
	subfuncSDC.SubfuncError = s.SubfuncError

	subfuncSDC.Message.PriceMaster = msg.PriceMaster

	return err
}

func getBoolPtr(b bool) *bool {
	return &b
}
