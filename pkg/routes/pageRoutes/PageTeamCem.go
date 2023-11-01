package pageRoutes

import (
	"fmt"
	"net/http"
	"net/url"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/components/xtext"
	"web-quickstart/pkg/components/xwidget"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageTeamCem(r *gin.Engine) {

	r.POST("/team/cem", func(c *gin.Context) {
		formStore := c.PostForm("store")
		formTimescale := c.PostForm("timescale")
		redirectUrl := fmt.Sprintf("/team/cem?timescale=%s&store=%s", url.QueryEscape(formTimescale), url.QueryEscape(formStore))
		c.Redirect(303, redirectUrl)
	})

	r.GET("/team/cem", func(c *gin.Context) {
		timescale := c.Query("timescale")
		store := c.Query("store")
		if timescale == "" {
			timescale = "current month"
		}
		if store == "" {
			store = "southroads"
		}
        b := core.NewPageBuilder("CFA Suite | Customer Service")
        b.Add(xlayout.Banner("CFA Suite", "Having a positive influence", xlayout.TeamNavMenu(c.Request.URL.Path)))
		b.Add(xtext.Article("Customer Service", "Our customer service is what sets apart from our competitors. Chick-fil-A is known for being a company that cares about people, not just their money. The results in this section help give us an idea about how our guest feel about our business."))
		b.Add(xwidget.CemScores(timescale, store))
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}