package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/components/xloading"
	"web-quickstart/pkg/components/xtext"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageTeamFinance(r *gin.Engine) {
	r.GET("/team/finance", func(c *gin.Context) {
        b := core.NewPageBuilder("CFA Suite | Financial Stewardship")
        b.Add(xlayout.Banner("CFA Suite", "Having a positive influence", xlayout.TeamNavMenu(c.Request.URL.Path)))
		b.Add(xtext.Article("Financial Stewardship", "We can bring in sales, but that does not make us profitable. We are only profitable is our sales are greater than our expenses. The results in this section indicate how profitable we are and give a good idea of where we are spending money."))
		b.Add(xloading.FinanceShadowWidget())
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}