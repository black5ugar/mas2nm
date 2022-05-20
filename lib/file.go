package lib

import "os"

//
// 文件（夹）创建相关
//
func DirExists(path string) (string, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, nil
}
