package handlerScout

import (
	"fmt"
	"github.com/GermanBogatov/auth_service/internal/model"
	"github.com/GermanBogatov/auth_service/internal/model/modelScout"
	"github.com/GermanBogatov/auth_service/internal/service"
	"github.com/GermanBogatov/auth_service/pkg/jwt"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerScout struct {
	Service service.AuthorizationScout
	Logger  logging.Logger
	Helper  jwt.Helper
}

func NewHandlerScout(service service.AuthorizationScout, logger logging.Logger, helper jwt.Helper) *HandlerScout {
	return &HandlerScout{
		Service: service,
		Logger:  logger,
		Helper:  helper,
	}
}

func (h *HandlerScout) SignUpScout(c *gin.Context) {

	var scout modelScout.ScoutDTO

	if err := c.BindJSON(&scout); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	//TODO validate!!!

	uuid, err := h.Service.CreateScout(c, scout)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"uuid": uuid,
	})

}

func (h *HandlerScout) SignInScout(c *gin.Context) {

	var scoutSign model.SignInDTO

	if err := c.BindJSON(&scoutSign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	//TODO validate!!!

	scout, err := h.Service.SignInScout(c.Request.Context(), scoutSign)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, refreshToken, err := h.Helper.GenerateAccessToken(scout)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  token,
		"refresh_token": refreshToken,
	})

}

func (h *HandlerScout) GetScout(c *gin.Context) {
	fmt.Println("KooL!")

}

func (h *HandlerScout) RefreshTokenScout(c *gin.Context) {
	refresh := c.Param("refresh_token")

	token, refreshToken, err := h.Helper.UpdateRefreshToken(refresh)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  token,
		"refresh_token": refreshToken,
	})
}
