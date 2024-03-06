package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// CreateDir 检查并生成存放图片的目录
func CreateDir(SavePath, subPath string) (string, error) {
	var dirs string
	if subPath == "" {
		dirs = fmt.Sprintf("%s/", SavePath)
	} else {
		dirs = fmt.Sprintf("%s/%s/", SavePath, time.Now().Format(subPath))
	}
	_, err := os.Stat(dirs)
	if err != nil {
		err = os.MkdirAll(dirs, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return dirs, nil
}

// GetAllFile 递归获取指定目录下的所有文件名
func GetAllFile(pathname string) ([]string, error) {
	result := []string{}

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}
	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFile(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			result = append(result, fullname)
		}
	}
	return result, nil
}

// PathExists 判断一个文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// RemoveAll 移除目录下的所有文件
func RemoveAll(dirPath string) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println(err)
	}
}
