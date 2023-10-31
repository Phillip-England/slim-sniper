package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/components/xloading"
	"web-quickstart/pkg/components/xtext"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageTeamCem(r *gin.Engine) {
	r.GET("/team/cem", func(c *gin.Context) {
        b := core.NewPageBuilder("CFA Suite | Customer Service")
        b.Add(xlayout.Banner("CFA Suite", "Having a positive influence", xlayout.TeamNavMenu(c.Request.URL.Path)))
		b.Add(xtext.Article("Customer Service", "Our customer service is what sets apart from our competitors. Chick-fil-A is known for being a company that cares about people, not just their money. The results in this section help give us an idea about how our guest feel about our business."))
		b.Add(xloading.CemShadowWidget())
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}