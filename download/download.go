package download

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var urldatapath = "https://api.bootcdn.cn/libraries/%s.min.json"
var urlfilepath = "https://cdn.bootcss.com/%s/%s/%s"

//var

/*
Download 下载一个框架下面的所有文件
name:"框架名称"
version:"版本号"
path:"存储位置"
tp:"文件保存类型 0:从新下载" 1.存在即跳过
*/
func Download(name, version, path string, ty int) error {
	up := fmt.Sprintf(urldatapath, name)
	bts, err := getByte(up)
	if err != nil {
		return err
	}
	value := info{}
	err = json.Unmarshal(bts, &value)
	if err != nil {
		return err
	}
	if strings.TrimSpace(version) == "" {
		version = value.Version
	}
	var files []string
	for _, item := range value.Assets {
		if item.Version == version {
			files = item.Files
		}
	}
	for _, item := range files {
		bts, err := getByte(fmt.Sprintf(urlfilepath, name, version, item))
		if err != nil {
			return errors.New("[" + item + "]" + err.Error())
		}
		fp := filepath.Join(path, item)
		if ty == 1 && exists(fp) {
			continue
		}
		err = writeFile(fp, bts)
		if err != nil {
			return errors.New("[" + item + "]" + err.Error())
		}
	}

	return nil
}

func getByte(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("错误的资源名称或请求失败")
	}
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bts, nil
}

func writeFile(path string, bt []byte) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bt, 0666)
}

func exists(path string) bool {
	i, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !i.IsDir()
}

type info struct {
	Assets []struct {
		Version string   `json:"version"`
		Files   []string `json:"files"`
	} `json:"assets"`
	Description string   `json:"description"`
	Homepage    string   `json:"homepage"`
	Keywords    []string `json:"keywords"`
	License     string   `json:"license"`
	Name        string   `json:"name"`
	Repository  struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
	Stars   int    `json:"stars"`
	Version string `json:"version"`
}
