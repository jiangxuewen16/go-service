package main

import (
	"math/rand"
	"io"
	"os"
	"fmt"
	"time"
)
var exceptionSlice = [...]string{"com.carhouse.common.exception.ServiceRuntimeException: 404找不到 — 服务器找不到给定的资源；文档不存在\n",
"java.lang.NullPointerException: 空指针\n"}

func main() {
	for {
		go writeStr()
		time.Sleep(time.Second * 5)
	}
}

func writeStr()  {
	randNum := rand.Intn(1)
	str := exceptionSlice[randNum]
	fileName := "D:/logs/log.log"
	f, err := os.OpenFile(fileName,os.O_APPEND, 0666)

	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	n, err1 := io.WriteString(f, str)
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Printf("写入 %d 个字节\n", n)
}
