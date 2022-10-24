package handler

import (
	"github.com/GermanBogatov/auth_service/internal/handler/handlerAdmin"
	"github.com/GermanBogatov/auth_service/internal/handler/handlerScout"
	"github.com/GermanBogatov/auth_service/internal/handler/handlerSportsman"
	"github.com/GermanBogatov/auth_service/internal/service"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/gin-gonic/gin"
)

type HandlerSportsman interface {
	SignUpSportsman(c *gin.Context)
	SignInSportsman(c *gin.Context)
}
type HandlerScout interface {
	SignUpScout(c *gin.Context)
	SignInScout(c *gin.Context)
}
type HandlerAdmin interface {
	SignUpAdmin(c *gin.Context)
	SignInAdmin(c *gin.Context)
}
type Handler struct {
	HandlerSportsman
	HandlerScout
	HandlerAdmin
}

func NewHandler(services *service.Service, logger logging.Logger) (*Handler, error) {
	return &Handler{
		HandlerSportsman: handlerSportsman.NewHandlerSportsman(services.AuthorizationSportsman, logger),
		HandlerScout:     handlerScout.NewHandlerScout(services.AuthorizationScout, logger),
		HandlerAdmin:     handlerAdmin.NewHandlerAdmin(services.AuthorizationAdmin, logger),
	}, nil
}

func (h *Handler) InitRoutes() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	authSportsman := router.Group("/auth/sportsman")
	{
		authSportsman.POST("/sign-up", h.SignUpSportsman)
		authSportsman.POST("/sign-in", h.SignInSportsman)
		//	auth.POST("/refresh=:refresh_token", h.refresh)
	}
	authScout := router.Group("/auth/scout")
	{
		authScout.POST("/sign-up", h.SignUpScout)
		authScout.POST("/sign-in", h.SignInScout)
		//	auth.POST("/refresh=:refresh_token", h.refresh)
	}
	authAdmin := router.Group("/auth/fscout/admin")
	{
		authAdmin.POST("/sign-up", h.SignUpAdmin)
		authAdmin.POST("/sign-in", h.SignInAdmin)
		//	auth.POST("/refresh=:refresh_token", h.refresh)
	}

	return router
}
