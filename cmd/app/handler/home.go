package handler

import (
	"embed"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type homeController struct {
	R   *gin.Engine
	Res embed.FS
}

func newHomeController(h *Handler) *homeController {
	return &homeController{
		R:   h.R,
		Res: h.Res,
	}
}

func (ctr *homeController) Register() {
	ctr.R.GET("/", ctr.getHome)
	ctr.R.GET("/wasm_exec.js", ctr.wasmJs)
	ctr.R.GET("/main.js", ctr.mainJs)
	ctr.R.StaticFile("/main.wasm", "cmd/app/static/main.wasm")
}

func (ctr *homeController) getHome(c *gin.Context) {
	tpl, err := template.ParseFS(ctr.Res, "static/index.html")
	if err != nil {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
		return
	}

	uniTxt := "သီဟိုဠ်မှ ဉာဏ်ကြီးရှင်သည် အာယုဝဍ္ဎနဆေးညွှန်းစာကို ဇလွန်ဈေးဘေး ဗာဒံပင်ထက် အဓိဋ္ဌာန်လျက် ဂဃနဏဖတ်ခဲ့သည်။"
	zgTxt := "သီဟိုဠ္မွ ဉာဏ္ႀကီးရွင္သည္ အာယုဝၯနေဆးၫႊန္းစာကို ဇလြန္ေဈးေဘး ဗာဒံပင္ထက္ အဓိ႒ာန္လ်က္ ဂဃနဏဖတ္ခဲ့သည္။"

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	if err := tpl.Execute(c.Writer, gin.H{
		"uniTxt": uniTxt,
		"zgTxt":  zgTxt,
	}); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
		return
	}
}

func (ctr *homeController) wasmJs(c *gin.Context) {
	tpl, err := template.ParseFS(ctr.Res, "static/wasm_exec.js")
	if err != nil {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	if err := tpl.Execute(c.Writer, nil); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
		return
	}
}

func (ctr *homeController) mainJs(c *gin.Context) {
	tpl, err := template.ParseFS(ctr.Res, "static/main.js")
	if err != nil {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	if err := tpl.Execute(c.Writer, nil); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
		return
	}
}
