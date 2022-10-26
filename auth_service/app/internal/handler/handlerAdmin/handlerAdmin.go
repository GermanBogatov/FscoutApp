package handlerAdmin

import (
	"github.com/GermanBogatov/auth_service/internal/service"
	"github.com/GermanBogatov/auth_service/pkg/jwt"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/gin-gonic/gin"
)

type HandlerAdmin struct {
	Service service.AuthorizationAdmin
	Logger  logging.Logger
	Helper  jwt.Helper
}

func NewHandlerAdmin(service service.AuthorizationAdmin, logger logging.Logger, helper jwt.Helper) *HandlerAdmin {
	return &HandlerAdmin{
		Service: service,
		Logger:  logger,
		Helper:  helper,
	}
}

func (h *HandlerAdmin) SignUpAdmin(c *gin.Context) {

}

func (h *HandlerAdmin) SignInAdmin(c *gin.Context) {

}
