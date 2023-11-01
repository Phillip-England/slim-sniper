package apiRoutes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CemScore(r *gin.Engine, db *sql.DB) {
    r.GET("/api/cem", func(c *gin.Context) {
        timescale := c.Query("timescale")
        store := c.Query("store")
        metric := c.Query("metric")
        if timescale == "" || store == "" || metric == "" {
            c.Data(http.StatusOK, "text/html", []byte("<p>0</p>"))
            return
        }
        query := "SELECT " + metric + " FROM cem_score WHERE store = $1 AND timescale = $2"
		rows, err := db.Query(query, store, timescale)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer rows.Close()
		var score float64
		if rows.Next() {
			err := rows.Scan(&score)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
        c.Data(http.StatusOK, "text/html", []byte(fmt.Sprintf("<p>%d</p>", int(score))))
    })
}