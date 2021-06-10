package handlers

import (
	"fmt"
	"fmtsmsapi/utils/app"
	"fmtsmsapi/utils/fmtsms"

	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
	"github.com/spf13/viper"
)

func AccountBalance(c *gin.Context) {

	client := fmtsms.NewClient(viper.GetString("appid"), viper.GetString("apikey"))
	fmt.Println(client.AppId);
	resp,err:=client.GetBalance()
	if err != nil {
		log.Error(app.LogError(c, err))
		c.String(200, app.LogError(c, err))
		return
	}
	app.Response(c,map[string]interface{}{
		"balance":resp.Data.Balance,
	})
}
