package pet

import (
	"go-hexagonal/api/common"
	"go-hexagonal/api/paginator"
	"go-hexagonal/api/v1/pet/request"
	"go-hexagonal/api/v1/pet/response"
	"go-hexagonal/business/pet"
	"strconv"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service pet.Service
}

//NewController Construct item API controller
func NewController(service pet.Service) *Controller {
	return &Controller{
		service,
	}
}

//GetItemByID Get item by ID echo handler
func (controller *Controller) FindPetByID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	pet, err := controller.service.FindPetByID(id, int(userID))
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetPetResponse(*pet)

	return c.JSON(common.NewSuccessResponse(response))
}

//FindAllPet Find All Pet with pagination handler
func (controller *Controller) FindAllPet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	pageQueryParam := c.QueryParam("page")
	rowPerPageQueryParam := c.QueryParam("row_per_page")

	skip, page, rowPerPage := paginator.CreatePagination(pageQueryParam, rowPerPageQueryParam)

	pets, err := controller.service.FindAllPet(int(userID), skip, rowPerPage)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllPetResponse(pets, page, rowPerPage)

	return c.JSON(common.NewSuccessResponse(response))
}

// InsertPet Create new pet echo handler
func (controller *Controller) InsertPet(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	insertPetRequest := new(request.InsertPetRequest)
	if err := c.Bind(insertPetRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertPet(*insertPetRequest.ToUpsertPetSpec(int(userID)), "creator")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

// UpdatePet update existing pet
func (controller *Controller) UpdatePet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	updatePetRequest := new(request.UpdatePetRequest)
	if err := c.Bind(updatePetRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdatePet(id, int(userID), updatePetRequest.Name, "modifier", updatePetRequest.Version)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
