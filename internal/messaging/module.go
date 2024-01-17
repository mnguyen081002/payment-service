package messaging

import (
	"go.uber.org/fx"
)

type Subscriber interface {
	Subscribe()
}

var Module = fx.Options()
