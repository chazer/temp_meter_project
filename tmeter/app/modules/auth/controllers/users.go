package controllers

import (
	"encoding/json"
	"net/http"
	"tmeter/app/api"
	auth "tmeter/app/modules/auth/services"
	"tmeter/app/modules/auth/views"
	"tmeter/app/modules/users/entities"
	"tmeter/lib/debug"
	"tmeter/lib/middlewares"
	"tmeter/lib/middlewares/wrappers"
)

import "tmeter/lib/router"

type AuthUsersController struct {
	Handlers  *router.Routes
	api       *api.APIProtocol
	auth      *auth.AuthService
	formatter api.FormatterConfig
}

func NewAuthUsersController(protocol *api.APIProtocol, auth *auth.AuthService) *AuthUsersController {
	c := &AuthUsersController{
		Handlers: router.NewRoutes(),
		api:      protocol,
		auth:     auth,
		formatter: api.FormatterConfig{
			DocumentView: views.NewAuthTokenResponseView(),
		},
	}

	c.Handlers.POST(
		"/token",
		// TODO: rewrite it
		middlewares.LogMiddleware(
			// TODO: Add ContentType guard middleware
			wrappers.NewJsonErrorsWrapper().AsMiddleware()(
				c.handlerCreateNewUserToken,
			),
		),
		// BODY: {
		//   user_name: string,
		//   user_email: string,
		// }
	)

	return c
}

type tokenRequestDTO struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

func (c *AuthUsersController) handlerCreateNewUserToken(resp http.ResponseWriter, req *http.Request) {
	var dto tokenRequestDTO
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	debug.Printf("create token for (email=%s; name=%s)", dto.UserEmail, dto.UserName)

	token, err := c.auth.IssueTokenForUser(&entities.User{
		Name:  dto.UserName,
		Email: dto.UserEmail,
	})
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	c.api.WriteEntityDocumentResponse(resp, &c.formatter, *token)
}
