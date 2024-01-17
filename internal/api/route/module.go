package route

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(),
	fx.Invoke(
		NewHealthRoutes,
	),
)
