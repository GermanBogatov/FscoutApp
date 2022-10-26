package handlerSportsman

import (
	"fmt"
	"github.com/GermanBogatov/auth_service/internal/model/modelSportsman"
	"github.com/GermanBogatov/auth_service/internal/service"
	"github.com/GermanBogatov/auth_service/pkg/jwt"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerSportsman struct {
	Service service.AuthorizationSportsman
	Logger  logging.Logger
	Helper  jwt.Helper
}

func NewHandlerSportsman(service service.AuthorizationSportsman, logger logging.Logger, helper jwt.Helper) *HandlerSportsman {
	return &HandlerSportsman{
		Service: service,
		Logger:  logger,
		Helper:  helper,
	}
}

func (h *HandlerSportsman) SignUpSportsman(c *gin.Context) {

	var sportsman modelSportsman.SportsmanDTO

	if err := c.BindJSON(&sportsman); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	//TODO validate!!!

	uuid, err := h.Service.CreateSportsman(c, sportsman)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"uuid": uuid,
	})
}

func (h *HandlerSportsman) SignInSportsman(c *gin.Context) {

	var sportsmanSign modelSportsman.SignInDTO

	if err := c.BindJSON(&sportsmanSign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	//TODO validate!!!
	sportsman, err := h.Service.GetSportsman(c.Request.Context(), sportsmanSign)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("authDTO cheeck:", sportsman)
	//TODO jwt token

	token, refreshToken, err := h.Helper.GenerateAccessToken(sportsman)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	t1, t2, _ := h.Helper.UpdateRefreshToken(refreshToken)
	fmt.Println("refresh: ", t1, t2)

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  token,
		"refresh_token": refreshToken,
	})
}

// получение всех данных спортсмена после аутент
func (h *HandlerSportsman) GetSportsman(c *gin.Context) {

	var sportsmanSign modelSportsman.SignInDTO

	if err := c.BindJSON(&sportsmanSign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	sportsman, err := h.Service.GetSportsman(c.Request.Context(), sportsmanSign)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, sportsman)
}
