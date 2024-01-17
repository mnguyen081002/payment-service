package messaging

import (
	"go.uber.org/fx"
	"paymentservice/internal/messaging/producer"
	"paymentservice/internal/messaging/subscriber"
)

type Subscriber interface {
	Subscribe()
}

var Module = fx.Options(
	fx.Provide(producer.NewCmsProductProducer),
	fx.Invoke(subscriber.NewUpdateProductSubscribe),
)
