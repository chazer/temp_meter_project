package auth

import (
	"tmeter/app"
	"tmeter/app/modules/auth/controllers"
	"tmeter/lib/router"
)

func Init(r *router.Router, f app.AppFactoryInterface) {
	auth := controllers.NewAuthController(
		f.GetAPIProtocol(),
		f.GetAuthService(),
	)

	r.AddRoutes("/auth", auth.Handlers)
}
