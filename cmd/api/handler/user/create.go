package user

import (
	"github.com/gin-gonic/gin"
	ucs "github.com/imjowend/microservicio-usuarios-y-autenticacion/internal/core"
	"github.com/imjowend/microservicio-usuarios-y-autenticacion/internal/core/user"

	"net/http"
)

type RestHandler struct {
	ucs ucs.UseCasePort
}

func NewRestHandler(ucs ucs.UseCasePort) *RestHandler {
	return &RestHandler{
		ucs: ucs,
	}
}

func (h *RestHandler) CreateUser(c *gin.Context) {

	if err := h.ucs.CreateUser(c.Request.Context(), *user.User); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Report successfully created"})
}
