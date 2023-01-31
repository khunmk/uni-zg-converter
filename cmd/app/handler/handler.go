package handler

import (
	"embed"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	R   *gin.Engine
	Res embed.FS
}

type HConfig struct {
	R   *gin.Engine
	Res embed.FS
}

func NewHandler(c *HConfig) *Handler {
	return &Handler{
		R:   c.R,
		Res: c.Res,
	}
}

func (h *Handler) Register() {
	homeCtr := newHomeController(h)
	homeCtr.Register()
}

func (h *Handler) Destroy() {
	//
}
