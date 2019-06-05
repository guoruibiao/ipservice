package main

import (
	"log"
	"net/http"
	"os"

	"github.com/guoruibiao/ipservice/utils"
	"github.com/guoruibiao/ipservice/constants"
	"github.com/gin-gonic/gin"
	"github.com/ipipdotnet/ipdb-go"
	"strings"
	cache2 "github.com/guoruibiao/ipservice/cache"
	)

func main() {
	app := gin.Default()
	// new instance of cache
	cacher, err := cache2.NewLRUCache(1024)
	if err != nil {
		log.Fatal(err)
	}

	db, err := ipdb.NewDistrict(constants.IPDB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	app.GET("/reload", func(ctx *gin.Context) {
		err := db.Reload(constants.IPDB_PATH)
		result := "ok"
		if err != nil {
			log.Fatalf("falied to reload db")
			result = "failed"
		}
		ctx.JSON(http.StatusOK, gin.H{"ret": result})
	})

	app.GET("/ip/info/:ip", func(ctx *gin.Context) {
		rawip := ctx.Param("ip")
		var ip string
		if len(strings.Split(rawip, ".")) != 4 {
			ip = utils.IpNToA(rawip).String()
		}else{
			ip = rawip
		}

		info, err := db.FindInfo(ip, "CN")
		if err != nil {
			//log.Fatalf("failed to find %s info", ip)
		}else{
			var et *constants.Entry
			if info == nil || info.CityName == "" {
				// 拿cache中的数据
				if entry, err := cacher.Get(ip); err == nil {
					info.RegionName = entry.RegionName
					info.CityName = entry.CityName
				}else{
					// 缓存没有，走官方接口
					et, err = utils.GetOfficialData(ip)
					if err == nil {
						info.RegionName = et.RegionName
						info.CityName = et.CityName
					}
				}
				//fmt.Println("OFFICIAL DATA: ", et, "--info: ", info)
			}else {
				// 添加cache
				et = &constants.Entry{
					RegionName: info.RegionName,
					CityName:   info.CityName,
				}
			}
			cacher.Cache(ip, et)
		}
		//fmt.Println(rawip, ip, info, cacher.PrintLRUCacheContainer())
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
