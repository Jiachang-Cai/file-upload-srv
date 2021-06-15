package requests

type UploadRequest struct {
	Path      string   `form:"path,default=."` // 上传路径
	Size      *int     `form:"size"`
	Dimension string   `form:"dimension"`
	Ext       []string `form:"ext"`
}

type FileResponse struct {
	Url      string  `json:"url"`       // 文件网络路径
	Size     string  `json:"size"`      // 文件大小
	Ext      string  `json:"ext"`       // 文件拓展
	Width    string  `json:"width"`     // 宽
	Height   string  `json:"height"`    // 高
	Duration float64 `json:"duration"`  // 文件时长
	MimeType string  `json:"mime_type"` // 文件类型
}
