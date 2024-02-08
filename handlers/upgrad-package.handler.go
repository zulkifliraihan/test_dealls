package handlers

import (
	"test_dealls/helpers"
	"test_dealls/services"

	"github.com/gin-gonic/gin"
)

type UpgradePackageHandler interface {
	UpgradePackage(ctx *gin.Context)
}

type upgradePackageHandler struct {
	upgradePackageService services.UpgradePackageService
}

func NewUpgradePackageHandler(upgradePackageService services.UpgradePackageService) UpgradePackageHandler {
	return &upgradePackageHandler{
		upgradePackageService: upgradePackageService,
	}
}

func (receiver *upgradePackageHandler) UpgradePackage (ctx *gin.Context) {
	
	var formatResponse helpers.TypeReturnResponse

	upgradePackage, _ := receiver.upgradePackageService.UpgradePackageUser(ctx)
	
	formatResponse = helpers.ReturnResponse(upgradePackage.Status, upgradePackage.Type, upgradePackage.Data, upgradePackage.Message, upgradePackage.Error)

	ctx.JSON(formatResponse.Code, formatResponse)
	return
}