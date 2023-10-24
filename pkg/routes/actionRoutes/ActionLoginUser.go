package actionRoutes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func ActionLoginUser(r *gin.Engine) {
	r.POST("/action/user/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == os.Getenv("TEAM_USERNAME") && password == os.Getenv("TEAM_PASSWORD") {
			c.Redirect(303, "/team/talent")
			return
		}
		if username == os.Getenv("ADMIN_USERNAME") && password == os.Getenv("ADMIN_PASSWORD") {
			c.Redirect(303, "/admin")
			return
		}
		c.Redirect(303, "/?LoginFormErr=invalid credentails")
	})
}