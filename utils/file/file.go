package file

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

func GetFileByte(fileObj *multipart.FileHeader) ([]byte, error) {
	src, err := fileObj.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	fileByte, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}
	return fileByte, nil
}

func SaveFile(fileByte []byte, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, bytes.NewReader(fileByte)); err != nil {
		return err
	}
	return nil

}

func GetFileFullUrl(path, name string) string {
	return filepath.Join(viper.GetString("static_url"), path, name)
}
func GetFileFullPath(path, name string) string {
	return filepath.Join(viper.GetString("static_path"), path, name)
}

func CreateSavePath(path string) error {
	// 创建文件夹
	if err := IsNotExistMkDir(filepath.Join(viper.GetString("static_path"), path)); err != nil {
		return err
	}
	return nil
}
