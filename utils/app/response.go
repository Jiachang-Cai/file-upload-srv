package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ListResponseBase struct {
	Total  uint64      `json:"total"`
	ErrNo  string      `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data,omitempty"`
	Sum    interface{} `json:"sum,omitempty"`
}
type ResponseBase struct {
	State string      `json:"state"`
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
}

func ErrorResponse(c *gin.Context, errMsg string) {
	if viper.GetString("runmode") == "debug" && c.GetString("error") != "" {
		errMsg = c.GetString("error")
	}
	c.JSON(http.StatusOK, ResponseBase{
		State: "0",
		Msg:   errMsg,
		Code:  "400",
	})
}

func ErrorNoResponse(c *gin.Context, code, errMsg string) {
	if viper.GetString("runmode") == "debug" && c.GetString("error") != "" {
		errMsg = c.GetString("error")
	}
	c.JSON(http.StatusOK, ResponseBase{
		State: "0",
		Code:  code,
		Msg:   errMsg,
	})
}

func Response(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseBase{
		State: "1",
		Data:  data,
		Code:  "200",
	})
}

func ResponseMsg(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, ResponseBase{
		State: "0",
		Msg:   msg,
		Data:  data,
		Code:  "200",
	})
}

func ListResponse(c *gin.Context, data interface{}, total uint64) {
	c.JSON(http.StatusOK, ListResponseBase{
		Total: total,
		ErrNo: "0",
		Data:  data,
	})
}

func LogError(c *gin.Context, err ...error) string {
	errString := fmt.Sprintf("uri:%v method:%v client_ip:%v request:%+v err:%+v", c.Request.RequestURI, c.Request.Method, c.ClientIP(), c.GetString("bodyStr"), err)
	c.Set("error", errString)
	return errString
}
