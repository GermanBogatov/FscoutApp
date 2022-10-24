package handlerSportsman

import (
	"fmt"
	"github.com/GermanBogatov/auth_service/internal/service"
	"github.com/GermanBogatov/auth_service/pkg/logging"
	"github.com/gin-gonic/gin"
)

type HandlerSportsman struct {
	Service service.AuthorizationSportsman
	Logger  logging.Logger
}

func NewHandlerSportsman(service service.AuthorizationSportsman, logger logging.Logger) *HandlerSportsman {
	return &HandlerSportsman{
		Service: service,
		Logger:  logger,
	}
}

func (h *HandlerSportsman) SignUpSportsman(c *gin.Context) {
	fmt.Println("Hello SignUpSportsman")
	id, err := h.Service.CreateSportsman(c)
	if err != nil {
		panic(err)
	}
	fmt.Println("id:", id)
}

func (h *HandlerSportsman) SignInSportsman(c *gin.Context) {

}
