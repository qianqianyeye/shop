package sadmin

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"git.jiaxianghudong.com/go/logs"
	"github.com/satori/go.uuid"
	"shop/admin/config"
)

const (
	ImageSavePath = "upload/images/"
	PrefixUrl = "localhost"
	RuntimeRootPath="runtime/"
	ImageMaxSize=5* 1024 * 1024  //5MB
	ImageAllowExts =".jpg,.jpeg,.png"
)
func GetImageFullUrl(name string) string {
	return fmt.Sprintf(PrefixUrl+":%d", config.GetListen()) + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	//fileName := strings.TrimSuffix(name, ext)
	fileName :=	uuid.Must(uuid.NewV4()).String()

	return fileName + ext
}

func GetImagePath() string {
	return ImageSavePath
}

func GetImageFullPath() string {
	return RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := GetExt(fileName)
	allows := strings.Split(ImageAllowExts,",")
	for _, allowExt := range allows {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err :=GetSize(f)
	if err != nil {
		logs.Error(err)
		return false
	}

	return size <= ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
