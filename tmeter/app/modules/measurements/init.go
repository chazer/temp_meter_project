package measurements

import (
	"tmeter/app"
	"tmeter/app/modules/measurements/controllers"
	"tmeter/lib/router"
)

func Init(r *router.Router, f app.AppFactoryInterface) {
	metrics := controllers.NewDeviceTempController(
		f.GetAPIProtocol(),
		f.GetAuthService(),
		f.GetDevicesService(),
		f.GetMeasurementsService(),
	)

	r.AddRoutes("/measurements", metrics.Handlers)
}
