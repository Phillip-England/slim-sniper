package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xform"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageAdminHome(r *gin.Engine) {
	r.GET("/admin", func(c *gin.Context) {
		cemFormErr := c.Query("UpdateCemFormErr")
        b := core.NewPageBuilder("CFA Suite | Admin Home")
		b.Add(xlayout.Banner("CFA Suite", ""))
		b.Add(xform.UpdateCemForm(cemFormErr))
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}