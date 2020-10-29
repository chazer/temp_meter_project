package devices

import (
	"tmeter/app"
	c "tmeter/app/modules/devices/controllers"
	"tmeter/lib/router"
)

func Init(r *router.Router, f app.AppFactoryInterface) {
	devices := c.NewDevicesController(
		f.GetDevicesService(),
		f.GetAPIProtocol(),
	)

	r.AddRoutes("/devices", devices.Handlers)
}
