package handlerScout

import (
	"github.com/GermanBogatov/auth_service/internal/service"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/gin-gonic/gin"
)

type HandlerScout struct {
	Service service.AuthorizationScout
	Logger  logging.Logger
}

func NewHandlerScout(service service.AuthorizationScout, logger logging.Logger) *HandlerScout {
	return &HandlerScout{
		Service: service,
		Logger:  logger,
	}
}

func (h *HandlerScout) SignUpScout(c *gin.Context) {

}

func (h *HandlerScout) SignInScout(c *gin.Context) {

}
