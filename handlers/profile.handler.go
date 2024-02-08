package handlers

import (
	"test_dealls/helpers"
	"test_dealls/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler interface {
	Profile(ctx *gin.Context)
}

type profileHandler struct {
	profileService services.ProfileService
}

func NewProfileHandler(profileService services.ProfileService) ProfileHandler {
	return &profileHandler{
		profileService: profileService,
	}
}

func (receiver *profileHandler) Profile (ctx *gin.Context) {
	
	var formatResponse helpers.TypeReturnResponse

	profile, _ := receiver.profileService.ProfileUser(ctx)
	
	formatResponse = helpers.ReturnResponse(profile.Status, profile.Type, profile.Data, profile.Message, profile.Error)

	ctx.JSON(formatResponse.Code, formatResponse)
	return
}