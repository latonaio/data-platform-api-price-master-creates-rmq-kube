package existence_conf

import "data-platform-api-price-master-creates-rmq-kube/config"

type exconfQueueMapper map[string]string

func NewExconfQueueMapper(c *config.Conf) exconfQueueMapper {
	m := exconfQueueMapper{}
	ecQNames := c.RMQ.QueueToExConf()
	m["SupplyChainRelationshipGeneral"] = ecQNames[0%len(ecQNames)]
	m["Buyer"] = ecQNames[1%len(ecQNames)]
	m["Seller"] = ecQNames[1%len(ecQNames)]
	m["Product"] = ecQNames[2%len(ecQNames)]
	return m
}

func (m exconfQueueMapper) getQueueNameByConfContent(content string) string {
	return m[content]
}
