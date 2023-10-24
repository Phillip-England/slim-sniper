package actionRoutes

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UpdateCem(r *gin.Engine, db *sql.DB) {
	r.POST("/action/cem/update", func(c *gin.Context) {
		timescale := c.PostForm("timescale")
		osat := c.PostForm("osat")
		taste := c.PostForm("taste")
		ace := c.PostForm("ace")
		speed := c.PostForm("speed")
		cleanliness := c.PostForm("cleanliness")
		accuracy := c.PostForm("accuracy")
		fmt.Println(timescale)
		if osat == "" || taste == "" || ace == "" || speed == "" || cleanliness == "" || accuracy == "" {
			c.Redirect(303, "/admin?UpdateCemFormErr=no empty fields")
			return
		}
		c.Redirect(303, "/admin")
	})
}