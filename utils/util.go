package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

//存储类型

//更多存储类型有待扩展
const (
	Version           = "1.0"
	StoreLocal string = "local"
	StoreOss   string = "oss"
)

var (
	BasePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	StoreType   = beego.AppConfig.String("store_type") //存储类型
)


//操作图片显示
func ShowImg(img string, style ...string) (url string) {
	if strings.HasPrefix(img, "https://") || strings.HasPrefix(img, "http://") {
		return img
	}
	img = "/" + strings.TrimLeft(img, "./")
	switch StoreType {
	case StoreOss:
		s := ""
		if len(style) > 0 && strings.TrimSpace(style[0]) != "" {
			s = "/" + style[0]
		}
		url = strings.TrimRight(beego.AppConfig.String("oss::Domain"), "/ ") + img + s
	case StoreLocal:
		url = img
	}

	return url
}
