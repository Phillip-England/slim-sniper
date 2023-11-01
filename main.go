package main

import (
	"fmt"
	"log"
	"web-quickstart/pkg/database"
	"web-quickstart/pkg/routes/actionRoutes"
	"web-quickstart/pkg/routes/apiRoutes"
	"web-quickstart/pkg/routes/componentRoutes"
	"web-quickstart/pkg/routes/pageRoutes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

    // dotenv
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    // server config
    r := gin.Default()
    r.Static("/static", "./static")

    // database access
    db, err := database.GetDatabase()
    if err != nil {
        fmt.Println(err)
        log.Panic("Database failed to connect")
    }
    defer db.Close()

    // initializing our db tables
    // core.DeleteTable(db, "cem_score")
    database.InitCemTable(db)

    // page routes
    pageRoutes.PageLogin(r)
    pageRoutes.PageTeamCem(r)
    pageRoutes.PageTeamFinance(r)
    pageRoutes.PageTeamSales(r)
    pageRoutes.PageTeamTalent(r)
    pageRoutes.PageAdminHome(r)

    // action routes
    actionRoutes.ActionLoginUser(r)
    actionRoutes.UpdateCem(r, db)

    // component routes
    componentRoutes.TalentScoresWidget(r, db)
    componentRoutes.CutomerServiceScoreWidget(r, db)
    componentRoutes.SalesResultsWidget(r, db)
    componentRoutes.FinancialResultsWidget(r, db)

    // api routes
    apiRoutes.CemScore(r, db)

    // running
    r.Run()

}

