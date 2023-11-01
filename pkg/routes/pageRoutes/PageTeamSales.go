package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/components/xloading"
	"web-quickstart/pkg/components/xtext"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageTeamSales(r *gin.Engine) {
	r.GET("/team/sales", func(c *gin.Context) {
        b := core.NewPageBuilder("CFA Suite | Sales & Brand Growth")
        b.Add(xlayout.Banner("CFA Suite", "Having a positive influence", xlayout.TeamNavMenu(c.Request.URL.Path)))
		b.Add(xtext.Article("Sales & Brand Growth", "Our people are great, but we wouldn't have this opportunity without our guests. If they did not come and make purchases, we would be employeed elsewhere. The results in this section help us keep an eye on how much money our business is generating."))
		b.Add(xloading.SalesShadowWidget())
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}