package controller

import (
	"net/http"
	"strings"

	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/model/dto"
	"enigmacamp.com/be-enigma-laundry/usecase"
	"enigmacamp.com/be-enigma-laundry/utils/common"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	registerHandler(ctx *gin.Context)
	loginHandler(ctx *gin.Context)
	refreshTokenHandler(ctx *gin.Context)
	Route()
}

type authController struct {
	uc         usecase.AuthUseCase
	rg         *gin.RouterGroup
	jwtService common.JwtToken
}

func (a *authController) registerHandler(ctx *gin.Context) {
	var payload model.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newResponse, err := a.uc.Register(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Created", newResponse)
}

func (a *authController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	resPayload, err := a.uc.Login(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", resPayload)
}

func (a *authController) refreshTokenHandler(ctx *gin.Context) {
	tokenString := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", -1)
	newToken, err := a.jwtService.RefreshToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	common.SendCreateResponse(ctx, "Ok", newToken)
}

func (a *authController) Route() {
	ug := a.rg.Group("/auth")
	ug.POST("/register", a.registerHandler)
	ug.POST("/login", a.loginHandler)
	ug.GET("/refresh-token", a.refreshTokenHandler)
}

func NewAuthController(uc usecase.AuthUseCase, rg *gin.RouterGroup, jwtService common.JwtToken) *authController {
	return &authController{uc: uc, rg: rg, jwtService: jwtService}
}
