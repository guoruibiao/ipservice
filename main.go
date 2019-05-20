package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ipipdotnet/ipdb-go"
)

func main() {
	app := gin.Default()

	db, err := ipdb.NewDistrict(ipPath)
	if err != nil {
		log.Fatal(err)
	}
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	app.GET("/reload", func(ctx *gin.Context) {
		err := db.Reload(ipPath)
		result := "ok"
		if err != nil {
			log.Fatalf("falied to reload db")
			result = "failed"
		}
		ctx.JSON(http.StatusOK, gin.H{"ret": result})
	})

	app.GET("/ip/info/:ip", func(ctx *gin.Context) {
		ip := IpNToA(ctx.Param("ip")).String()
		info, err := db.FindInfo(ip, "CN")
		if err != nil {
			log.Fatalf("failed to find %s info", ip)
		}
		hostname, _ := os.Hostname()
		ctx.JSON(http.StatusOK, gin.H{
			"host":       hostname,
			"build time": db.BuildTime(),
			"languages":  db.Languages(),
			"fields":     db.Fields(),
			"info":       info,
		})
	})

	app.Run(":8080")
}
