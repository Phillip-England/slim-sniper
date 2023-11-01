package componentRoutes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CutomerServiceScoreWidget(r *gin.Engine, db *sql.DB) {
	r.GET("/components/CustomerServiceScoreWidget", func(c *gin.Context) {
		query := "SELECT * FROM cem_score"
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
			// need to return 500 page
			return
		}
		defer rows.Close()
		if rows.Next() {

		}
		widget := fmt.Sprintf(`
			<div class='bg-white rounded m-2 flex flex-col gap-6 p-6'>
				<div class='text-xs items-center justify-between gap-2 flex'>
					<h2 class='self-start text-lg'>CEM Results</h2>
					<div class='flex flex-col'>
						<div class='flex flex-col gap-2'>
							<select class='border p-1'>
								<option>Southroads</option>
								<option>Utica</option>
							</select>
							<select class='border p-1'>
								<option>Current Month</option>
								<option>Three Month Rolling</option>
								<option>Year To Date</option>
							</select>
						</div>
					</div>
				</div>
				<ul class='flex flex-col rounded gap-2 text-xs'>
					<li class='flex flex-row justify-between border-b pb-4'>
						<div class='flex flex-row items-center gap-4'>
							<i class='fa-solid fa-caret-right'></i>
							<p>osat</p>	
						</div>
						<div>
							<div hx-get='/api/cem' hx-swap="outerHTML"></div>
						</div>
					</li>
				</ul>
			</div>
		`, )
        c.Data(http.StatusOK, "text/html", []byte(widget))
    })
}