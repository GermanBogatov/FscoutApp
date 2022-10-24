package handlerAdmin

import (
	"github.com/GermanBogatov/auth_service/internal/service"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/gin-gonic/gin"
)

type HandlerAdmin struct {
	Service service.AuthorizationAdmin
	Logger  logging.Logger
}

func NewHandlerAdmin(service service.AuthorizationAdmin, logger logging.Logger) *HandlerAdmin {
	return &HandlerAdmin{
		Service: service,
		Logger:  logger,
	}
}

func (h *HandlerAdmin) SignUpAdmin(c *gin.Context) {

}

func (h *HandlerAdmin) SignInAdmin(c *gin.Context) {

}
