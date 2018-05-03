package main

import (
	"math/rand"
	"io"
	"os"
	"fmt"
	"time"
)

var exceptionSlice = [...]string{"com.carhouse.common.exception.ServiceRuntimeException: 404找不到 — 服务器找不到给定的资源；文档不存在\n\r",
	"java.lang.NullPointerException: 空指针\n\t"}

func main() {

	fileName := "D:/logs/log.log"
	f, err := os.OpenFile(fileName, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	for {
		randNum := rand.Intn(2)
		str := exceptionSlice[randNum]
		writeStr(f,str)
		time.Sleep(time.Millisecond * 100)
	}
}

func writeStr(f io.Writer, str string) error {

	_, err := io.WriteString(f, str)
	if err != nil {
		return err
	}
	fmt.Sprintf("写入 【%d】\n", str)
	return nil
}
