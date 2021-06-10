package requests

type CallBackStatusRequest []struct {
	Code   string `json:"code"`
	Mobile string `json:"mobile"`
	Time   string `json:"time"`
	TaskId string `json:"taskid"`
}
