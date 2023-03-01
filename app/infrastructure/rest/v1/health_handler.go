package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type health struct {
	ArtifactVersion string `json:"artifact_version"`
	Status          string `json:"status"`
}

func NewHealth(g *echo.Group) *health {
	h := &health{ArtifactVersion: "1.0.0"}
	h.healthCheck(g)
	return h
}

func (h *health) healthCheck(e *echo.Group) {
	e.GET("/health", h.healthHandler)
}

func (h *health) healthHandler(c echo.Context) error {
	h.Status = "OK"
	return c.JSON(http.StatusOK, h)
}
