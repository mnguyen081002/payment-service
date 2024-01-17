package mappermessage

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"paymentservice/internal/messaging/message"
)

func DecreaseProductQuantityMessage(kmsg kafka.Message) message.DecreaseProductQuantity {
	var msg message.DecreaseProductQuantity
	_ = json.Unmarshal(kmsg.Value, &msg)
	return msg
}
