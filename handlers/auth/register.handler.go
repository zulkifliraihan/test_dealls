package handlers

import (
	"fmt"
	"test_dealls/helpers"
	"test_dealls/requests"
	authService "test_dealls/services/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterHandler interface {
	Register(ctx *gin.Context)
}

type registerHandler struct {
	registerService authService.RegisterService
}

func NewRegisterHandler(registerService authService.RegisterService) RegisterHandler {
	return &registerHandler{
		registerService: registerService,
	}
}

func (receiver *registerHandler) Register (ctx *gin.Context) {
	
	data := new(requests.RegisterRequest)
	var formatResponse helpers.TypeReturnResponse

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
        if validationErr, ok := err.(validator.ValidationErrors); ok {
            validation := helpers.ValidationCustomMessage(validationErr)
			fmt.Println(validation)
			formatResponse = helpers.ReturnResponse(false, "validation", nil, "", validation)

		} else {
			formatResponse = helpers.ReturnResponse(false, "server", nil, "", nil)

		}

		ctx.JSON(formatResponse.Code, formatResponse)
		return
	}

	users, _ := receiver.registerService.RegisterUser(*data)
	
	formatResponse = helpers.ReturnResponse(users.Status, users.Type, users.Data, users.Message, users.Error)

	ctx.JSON(formatResponse.Code, formatResponse)
	return
}