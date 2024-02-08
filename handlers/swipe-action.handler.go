package handlers

import (
	"fmt"
	"test_dealls/helpers"
	"test_dealls/requests"
	"test_dealls/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SwipeActionHandler interface {
	SwipeAction(ctx *gin.Context)
}

type swipeActionHandler struct {
	swipeActionService services.SwipeActionService
}

func NewSwipeActionHandler(swipeActionService services.SwipeActionService) SwipeActionHandler {
	return &swipeActionHandler{
		swipeActionService: swipeActionService,
	}
}

func (receiver *swipeActionHandler) SwipeAction (ctx *gin.Context) {
	
	data := new(requests.SwipeActionRequest)
	var formatResponse helpers.TypeReturnResponse

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
        if validationErr, ok := err.(validator.ValidationErrors); ok {
            validation := helpers.ValidationCustomMessage(validationErr)
			formatResponse = helpers.ReturnResponse(false, "validation", nil, "", validation)
			} else {
			formatResponse = helpers.ReturnResponse(false, "server", nil, "", validationErr.Error())
		}

		ctx.JSON(formatResponse.Code, formatResponse)
		return
	}

	swipeAction, _ := receiver.swipeActionService.SwipeActionUser(ctx, *data)
	
	formatResponse = helpers.ReturnResponse(swipeAction.Status, swipeAction.Type, swipeAction.Data, swipeAction.Message, swipeAction.Error)
	fmt.Println("formatResponse : ", formatResponse)
	ctx.JSON(formatResponse.Code, formatResponse)
	return
}