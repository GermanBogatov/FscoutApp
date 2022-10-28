package handler

import (
	"github.com/GermanBogatov/auth_service/internal/handler/handlerAdmin"
	"github.com/GermanBogatov/auth_service/internal/handler/handlerScout"
	"github.com/GermanBogatov/auth_service/internal/handler/handlerSportsman"
	"github.com/GermanBogatov/auth_service/internal/service"
	"github.com/GermanBogatov/auth_service/pkg/jwt"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/gin-gonic/gin"
)

type HandlerSportsman interface {
	SignUpSportsman(c *gin.Context)
	SignInSportsman(c *gin.Context)
	GetSportsman(c *gin.Context)
	RefreshTokenSportsman(c *gin.Context)
}
type HandlerScout interface {
	SignUpScout(c *gin.Context)
	SignInScout(c *gin.Context)
	GetScout(c *gin.Context)
	RefreshTokenScout(c *gin.Context)
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

func NewHandler(services *service.Service, logger logging.Logger, helper jwt.Helper) (*Handler, error) {
	return &Handler{
		HandlerSportsman: handlerSportsman.NewHandlerSportsman(services.AuthorizationSportsman, logger, helper),
		HandlerScout:     handlerScout.NewHandlerScout(services.AuthorizationScout, logger, helper),
		HandlerAdmin:     handlerAdmin.NewHandlerAdmin(services.AuthorizationAdmin, logger, helper),
	}, nil
}

func (h *Handler) InitRoutes() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		sportsman := auth.Group("/sportsman")
		{
			sportsman.POST("/sign-up", h.SignUpSportsman)
			sportsman.POST("/sign-in", h.SignInSportsman)
			sportsman.POST("/refresh=:refresh_token", h.RefreshTokenSportsman)
		}

		scout := auth.Group("/scout")
		{
			scout.POST("/sign-up", h.SignUpScout)
			scout.POST("/sign-in", h.SignInScout)
			scout.POST("/refresh=:refresh_token", h.RefreshTokenScout)
		}
		admin := router.Group("/fscout/admin")
		{
			admin.POST("/sign-up", h.SignUpAdmin)
			admin.POST("/sign-in", h.SignInAdmin)
			//	auth.POST("/refresh=:refresh_token", h.refresh)
		}

	}

	userSportsman := router.Group("/sportsman")
	{
		userSportsman.Use(jwt.MiddlewareSportsman())
		userSportsman.GET("/", h.GetSportsman)
	}

	userScout := router.Group("/scout")
	{
		userScout.Use(jwt.MiddlewareScout())
		userScout.GET("/", h.GetScout)
	}
	return router
}
