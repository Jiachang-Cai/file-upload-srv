package handlers

import (
	"fmt"
	"strings"

	"fmtsmsapi/handlers/requests"
	"fmtsmsapi/models"
	"fmtsmsapi/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
)

func CallBackStatus(c *gin.Context) {
	var r requests.CallBackStatusRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.Error(app.LogError(c, err))
		c.String(200, app.LogError(c, err))
		return
	}
	go func(r requests.CallBackStatusRequest) {
		// 批量修改
		updList := make([]string, 0)
		for _, item := range r {
			if item.Code == "DELIVRD" {
				updList = append(updList, fmt.Sprintf("update fmt_sms_send_log set status=1,send_time='%s' where task_id='%s' and phone='%s';", item.Time, item.TaskId, item.Mobile))
			} else {
				updList = append(updList, fmt.Sprintf("update fmt_sms_send_log set status=2,send_time='%s' where task_id='%s' and phone='%s';", item.Time, item.TaskId, item.Mobile))
			}
		}
		if len(updList) > 0 {
			// 执行sql
			batchSql := strings.Join(updList, "\n")
			if err := models.ExecSql(batchSql); err != nil {
				log.Error(app.LogError(c, err))
			}
		}

	}(r)
	// 加入数据到队列
	//client := cache.RedisDB.Game
	//b, err := json.Marshal(r)
	//if err != nil {
	//	log.Error(app.LogError(c, err))
	//	c.String(200, app.LogError(c, err))
	//	return
	//}
	//if err := client.RPush(c, "fmt_sms_callback_data", string(b)).Err(); err != nil {
	//	log.Error(app.LogError(c, err))
	//	c.String(200, app.LogError(c, err))
	//	return
	//}
	c.String(200, "SUCCESS")
	return
}

func CallBackTmpl(c *gin.Context) {

}

func CallBackTest(c *gin.Context) {
	for i := 0; i < 100000; i++ {
		go func(i int) {
			sql := fmt.Sprintf("insert into fmt_sms_send_log(task_id) values('%d')", i)
			if err := models.ExecSql(sql); err != nil {
				log.Error(app.LogError(c, err))
			}
		}(i)
	}
}
