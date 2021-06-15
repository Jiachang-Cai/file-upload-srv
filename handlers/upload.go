package handlers

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
	"github.com/spf13/viper"

	"file-upload-srv/handlers/requests"
	"file-upload-srv/utils/app"
	"file-upload-srv/utils/ffprobe"
	ufile "file-upload-srv/utils/file"
	uhash "file-upload-srv/utils/hash"
)

func UploadAdd(c *gin.Context) {
	var r requests.UploadRequest
	if err := c.ShouldBind(&r); err != nil {
		log.Error(app.LogError(c, err))
		app.ErrorResponse(c, "获取参数失败")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Error(app.LogError(c, err))
		app.ErrorResponse(c, "获取文件对象错误")
		return
	}
	fileByte, err := ufile.GetFileByte(file)
	if err != nil {
		log.Error(app.LogError(c, err))
		app.ErrorResponse(c, "获取文件内容失败")
		return
	}
	fileType := http.DetectContentType(fileByte)
	fileMd5 := uhash.HashByteMd5(fileByte)
	fileExt := strings.ToLower(path.Ext(file.Filename))
	fileSize := len(fileByte)
	// 验证大小
	if r.Size != nil {
		if fileSize > *r.Size {
			app.ErrorResponse(c, "上传大小不符合")
			return
		}
	}
	dst := r.Path + "/" + fileMd5 + fileExt
	// 保存文件
	if err := ufile.SaveFile(fileByte, dst); err != nil {
		log.Error("save file", err)
		app.ErrorResponse(c, "文件存储失败")
		return
	}
	var fileWidth, fileHeight int
	var fileDuration time.Duration
	switch strings.Split(fileType, "/")[0] {
	case "image":
		imageConfig, _, err := image.DecodeConfig(bytes.NewReader(fileByte))
		if err != nil {
			log.Error(app.LogError(c, err))
			app.ErrorResponse(c, "图片解析失败")
			return
		}
		fileWidth = imageConfig.Width
		fileHeight = imageConfig.Height
	case "video":
		// 需要环境安装 ffmpeg
		data, err := ffprobe.GetProbeData(dst, 1000*time.Millisecond)
		if err != nil {
			log.Error("ffprobe.GetProbeData", err)
			app.ErrorResponse(c, "视频解析失败")
			return
		}
		fileDuration = data.Format.Duration()
		firstVideo := data.GetFirstVideoStream()
		fileWidth = firstVideo.Width
		fileHeight = firstVideo.Height
	}
	// 验证尺寸
	if r.Dimension != "" {
		if strings.ToLower(r.Dimension) != fmt.Sprintf("%vx%v", fileWidth, fileHeight) {
			log.Error(app.LogError(c, err))
			app.ErrorResponse(c, "上传尺寸不符合")
			return
		}
	}
	// 验证拓展
	if len(r.Ext) > 0 {
		if !InArray(fileExt, r.Ext) {
			app.ErrorResponse(c, "上传格式不符合")
			return
		}
	}
	fileUrl := strings.Replace(viper.GetString("domain")+"/"+dst, "./", "", -1)
	app.Response(c, requests.FileResponse{
		Url:      fileUrl,
		Size:     fmt.Sprintf("%v", fileSize),
		Ext:      fileExt,
		Width:    fmt.Sprintf("%v", fileWidth),
		Height:   fmt.Sprintf("%v", fileHeight),
		Duration: float64(fileDuration) / float64(time.Second),
		MimeType: fileType,
	})
	return
}

func InArray(val string, array []string) bool {
	for _, i := range array {
		if val == strings.ToLower(i) {
			return true
		}
	}
	return false
}
