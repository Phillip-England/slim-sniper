package main

import (
	"fmt"
	"log"
	"web-quickstart/pkg/core"
	"web-quickstart/pkg/routes/actionRoutes"
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
    db, err := core.GetDatabase()
    if err != nil {
        fmt.Println(err)
        log.Panic("Database failed to connect")
    }
    defer db.Close()

    // initializing our db tables
    core.DeleteTable(db, "cem_score")

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
    componentRoutes.TalentScoresWidget(r)
    componentRoutes.CutomerServiceScoreWidget(r)
    componentRoutes.SalesResultsWidget(r)
    componentRoutes.FinancialResultsWidget(r)

    // running
    r.Run()

}

