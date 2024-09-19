package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/icodealot/noaa"
)

type coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type points struct {
	Time    string `json:"Time"`
	Summary string `json:"summary"`
	Detail  string `json:"details"`
}

func getWeather(ctx *gin.Context) {
	coord := &coordinates{}
	err := ctx.BindJSON(&coord)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	point := &points{}
	period, err := noaa.Forecast(coord.Latitude, coord.Longitude)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	point.Time = period.Periods[0].Name
	point.Summary = period.Periods[0].Summary
	point.Detail = period.Periods[0].Details

	ctx.JSON(http.StatusOK, point)
 	}

func main() {
	noaa.SetConfig(noaa.GetConfig())

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	server.GET("/weather", getWeather)
	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(0)
}
