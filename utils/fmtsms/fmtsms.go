package fmtsms

import (
	"crypto"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func NewClient(AppId, ApiKey string) *FmtSms {
	return &FmtSms{
		AppId:  AppId,
		ApiKey: ApiKey,
	}
}

type FmtSms struct {
	ApiKey string
	AppId  string
}

type FmtSmsRsp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type FmtSmsRspNew struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Balance int `json:"balance"`
	} `json:"data"`
}

func (c *FmtSms) Send(data map[string]interface{}) (*FmtSmsRsp, error) {
	dataByte, _ := json.Marshal(data)
	encoder := base64.StdEncoding.EncodeToString(dataByte)
	request := gorequest.New().Timeout(3 * time.Minute)
	var rspData FmtSmsRsp
	encoded := url.Values{}
	for k, v := range c.Params() {
		encoded.Set(k, v)
	}
	reqUrl := "http://39.98.184.179:8081/api/v2/mms/templateSend" + "?" + encoded.Encode()
	resp, _, errs := request.Post(reqUrl).
		Set("Content-Type", "text/plain").
		Send(encoder).EndStruct(&rspData)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("错误响应code url:%s data:%+v", reqUrl, data))
	}
	if len(errs) > 0 {
		errStr := fmt.Sprintf("%+v", errs)
		return nil, errors.New(errStr)
	}
	if rspData.Code != "T" {
		return nil, errors.New(fmt.Sprintf("错误响应 url:%s data:%+v rsp:%+v", reqUrl, data, rspData))
	}
	return &rspData, nil
}

func (c *FmtSms) BatchSend() {

}

func (c *FmtSms) GetBalance() (*FmtSmsRspNew, error) {
	encoded := url.Values{}
	for k, v := range c.ParamsNew() {
		encoded.Set(k, v)
	}
	var rspData FmtSmsRspNew

	request := gorequest.New().Timeout(3 * time.Minute)
	reqUrl := "http://39.98.184.179:8081/api/sms/balance" + "?" + encoded.Encode()
	resp, _, errs := request.Get(reqUrl).Send("&" + encoded.Encode()).EndStruct(&rspData)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("错误响应code url:%s", reqUrl))
	}
	if len(errs) > 0 {
		errStr := fmt.Sprintf("%+v", errs)
		return nil, errors.New(errStr)
	}
	if rspData.Code != "T" {
		return nil, errors.New(fmt.Sprintf("错误响应 url:%s  rsp:%+v", reqUrl, rspData))
	}
	return &rspData, nil
}

func (c *FmtSms) Sign(timestamp int64) string {
	str := fmt.Sprintf("%s%s%d%s", c.ApiKey, c.AppId, timestamp, c.ApiKey)
	h := crypto.MD5.New()
	h.Write([]byte(str))
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

func (c *FmtSms) Params() map[string]string {
	timestamp := time.Now().Unix() * 1000
	data := map[string]string{
		"appid":     c.AppId,
		"timestamp": strconv.FormatInt(timestamp, 10),
		"sign":      c.Sign(timestamp),
	}
	return data
}

func (c *FmtSms) ParamsNew() map[string]string {
	timestamp := time.Now().Unix() * 1000
	data := map[string]string{
		"userid": c.AppId,
		"ts":     strconv.FormatInt(timestamp, 10),
		"sign":   c.SignNew(timestamp),
	}
	return data
}

func (c *FmtSms) SignNew(timestamp int64) string {
	str := fmt.Sprintf("%s%d%s", c.AppId, timestamp, c.ApiKey)
	h := crypto.MD5.New()
	h.Write([]byte(str))
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}
