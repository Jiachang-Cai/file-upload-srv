package main

import (
	"flag"
	"fmt"
	"fmtsmsapi/cache"
	"fmtsmsapi/models"
	"net/http"
	"os"
	"runtime"

	"fmtsmsapi/config"
	"fmtsmsapi/router"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
	"github.com/spf13/viper"
)

var (
	c string
)

func init() {
	flag.StringVar(&c, "c", "dev", "配置文件")
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	flag.Parse()
	// init config
	if err := config.Init(c); err != nil {
		panic(err)
	}
	// init db
	models.DB.Init()
	defer models.DB.Close()
	// init redis
	cache.RedisDB.Init()
	// set gin mode
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	// Routes
	router.Load(g)
	fmt.Printf("pid: %d addr: %s\n", os.Getpid(), viper.GetString("addr"))
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	gracehttp.Serve(&http.Server{Addr: viper.GetString("addr"), Handler: g})
}
