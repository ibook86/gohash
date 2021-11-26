package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var filePath string
	fmt.Print("输入文件路径(按回车键结束)：")
	fmt.Scanln(&filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("文件路径错误", err)
	}

	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	sum := hash.Sum(nil)

	fmt.Print("文件路径：")
	fmt.Printf("%s\n", filePath)

	fmt.Print("md5值：")
	fmt.Printf("%x\n", sum)

	fmt.Println("按'Enter'键退出本程序")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
