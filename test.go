package main

import (
	"github.com/hpcloud/tail"
	"fmt"
	//"golang.org/x/text/encoding/simplifiedchinese"
	"strings"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)


var ExceptionSlice = [...]string{"com.carhouse.common.exception.ServiceRuntimeException", "java.lang.NullPointerException"}

func main() {
	t, err := tail.TailFile("D:/logs/log.log", tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
	}

	for line := range t.Lines {
		//var str = ConvertByte2String([]byte(line.Text), GB18030)

		for _, exception := range ExceptionSlice {
			if strings.Contains(line.Text, exception) {
				var str = fmt.Sprintf("%q:::::Line:%d:::::Time:%T", line.Text, 100, line.Time)
				fmt.Println(str)
			}
		}
	}
}

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}