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
		store := c.PostForm("store")
		fmt.Println(timescale, store)
		if osat == "" || taste == "" || ace == "" || speed == "" || cleanliness == "" || accuracy == "" {
			c.Redirect(303, "/admin?UpdateCemFormErr=no empty fields")
			return
		}
		_, deleteErr := db.Exec("DELETE FROM cem_score WHERE timescale = $1 AND store = $2", timescale, store)
		if deleteErr != nil {
			fmt.Println("Error deleting existing entries:", deleteErr)
			c.Redirect(303, "/admin?UpdateCemFormErr=database error")
			return
		}
		_, insertErr := db.Exec("INSERT INTO cem_score (timescale, store, osat, taste, ace, speed, cleanliness, accuracy) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
			timescale, store, osat, taste, ace, speed, cleanliness, accuracy)
		if insertErr != nil {
			fmt.Println("Error inserting new row:", insertErr)
			c.Redirect(303, "/admin?UpdateCemFormErr=database error")
			return
		}
		c.Redirect(303, "/admin")
	})
}