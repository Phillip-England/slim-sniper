package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/components/xloading"
	"web-quickstart/pkg/components/xtext"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageTeamTalent(r *gin.Engine) {
	r.GET("/team/talent", func(c *gin.Context) {
        b := core.NewPageBuilder("CFA Suite | Talent Engagement")
        b.Add(xlayout.Banner("CFA Suite", "Having a positive influence", xlayout.TeamNavMenu(c.Request.URL.Path)))
		b.Add(xtext.Article("Talent Engagement", "All of the metrics in this section revolve around the people in our business. These results are important because we believe people are a critical success factor."))
		b.Add(xloading.TalentShadowWidget())
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}