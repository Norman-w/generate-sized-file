package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前用户的桌面路径
	desktopPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("获取用户主目录失败:", err)
		return
	}
	desktopPath = filepath.Join(desktopPath, "Desktop")

	// 创建一个23*1024*1024字节的文件
	filePath := filepath.Join(desktopPath, "23m+1b.txt")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败:", err)
		}
	}(file)

	// 写入23*1024*1024 +1字节的数据
	data := make([]byte, 23*1024*1024+1)
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	fmt.Println("文件创建成功:", filePath)
}
