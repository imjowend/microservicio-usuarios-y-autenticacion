package handler

import (
	"github.com/gin-gonic/gin"
	ucs "github.com/imjowend/microservicio-usuarios-y-autenticacion/internal/core"
	usm "github.com/imjowend/microservicio-usuarios-y-autenticacion/internal/core/user"
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

func (h *RestHandler) FakeCreateUser(c *gin.Context) {
	var newUser = usm.User{
		ID:       "",
		Username: "",
		Email:    "",
		Password: "",
	}

	if err := h.ucs.CreateUser(c.Request.Context(), &newUser); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "User successfully created"})
}
