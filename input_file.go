package main

import (
	"math/rand"
	"io"
	"os"
	"fmt"
	"time"
	"flag"
)

var fileName = flag.String("f", "/home/logs/golang/test.log", "文件路径+文件名称")
var sleepTime = flag.Int("t", 10, "写入间隔时间，单位是毫秒")
var stringType = flag.Int("T", 0, "写入类型，0-随机，1-写入第一个，2-写入第二个")
var exceptionSlice = [...]string{"com.carhouse.common.exception.ServiceRuntimeException: 404找不到 — 服务器找不到给定的资源；文档不存在\n",
	"java.lang.NullPointerException: 空指针\n"}

func main() {
	flag.Parse()
	if len(*fileName) == 0 {
		fmt.Println("文件名称不能为空")
		os.Exit(0)
	}
	if str := exceptionSlice[*stringType]; len(str) <= 0 {
		fmt.Println("请填写有效数字")
		os.Exit(0)
	}
	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer f.Close()
	for {
		if *sleepTime <= 0 {
			fmt.Println("写入间隔时间必须大于0")
			os.Exit(0)
		}
		time.Sleep(time.Millisecond * time.Duration(*sleepTime))

		randNum := rand.Intn(len(exceptionSlice))
		str := exceptionSlice[randNum]
		err := writeStr(f, str)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}

func writeStr(f io.Writer, str string) error {

	_, err := io.WriteString(f, str)
	if err != nil {
		return err
	}
	fmt.Printf("写入 【%d】\n", str)
	return nil
}
