package auth

import (
	"go-hexagonal/api/common"
	"go-hexagonal/api/v1/auth/request"
	"go-hexagonal/api/v1/auth/response"
	"go-hexagonal/business/auth"

	echo "github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service auth.Service
}

//NewController Construct item API controller
func NewController(service auth.Service) *Controller {
	return &Controller{
		service,
	}
}

//Login by given username and password will return JWT token
func (controller *Controller) Login(c echo.Context) error {
	loginRequest := new(request.LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	token, err := controller.service.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewLoginResponse(token)

	return c.JSON(common.NewSuccessResponse(response))
}
