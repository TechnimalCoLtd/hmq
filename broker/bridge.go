package broker

import (
	"github.com/TechnimalCoLtd/hmq/plugins/bridge"
	"go.uber.org/zap"
)

func (b *Broker) Publish(e *bridge.Elements) {
	if b.bridgeMQ == nil {
		return
	}
	if err := b.bridgeMQ.Publish(e); err != nil {
		log.Error("send message to mq error.", zap.Error(err))
	}
}
