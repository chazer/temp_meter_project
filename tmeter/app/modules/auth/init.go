package auth

import (
	"tmeter/app"
	"tmeter/app/modules/auth/controllers"
	"tmeter/lib/router"
)

func Init(r *router.Router, f app.AppFactoryInterface) {
	users := controllers.NewAuthUsersController(
		f.GetAPIProtocol(),
		f.GetAuthService(),
	)

	devices := controllers.NewAuthDevicesController(
		f.GetAPIProtocol(),
		f.GetAuthService(),
		f.GetDevicesService(),
	)

	r.AddRoutes("/auth/user", users.Handlers)
	r.AddRoutes("/auth/device", devices.Handlers)
}
