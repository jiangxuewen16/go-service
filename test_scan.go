package main

import (
	"net"
	"time"
	"fmt"
	"encoding/json"
)

type FileClientTransfer struct {
	conn *net.Conn `json:"conn"` //socket连接

	label string `json:"label"` //master-主服务端(中心服务器)  slave-从服务器

	FileName      string `json:"file_name"`       //待发送文件名称
	MergeFileName string `json:"merge_file_name"` //待合并文件名称
	Coroutine     int    `json:"coroutine"`       //协程数量或拆分文件的数量
	BufSize       int    `json:"buf_size"`        //单次发送数据的大小
}

func main() {
	conn,_ := net.Dial("tcp","127.0.0.1:8001")

	f := &FileClientTransfer{FileName:"adada.jpg",MergeFileName:"1111.png",Coroutine:20,BufSize:40}

	s,_ := json.Marshal(f)


	str := "HEAD / SOCKET/1.0\r\n\r\nRequst-router:/file/transfer\r\n\r\nStatus-code:200\r\n\r\nContent-type:application/json\r\n\r\nAuthentication:\r\n\r\nBody:"  + string(s)

	for {
		go conn.Write([]byte(str))
		go func() {
			b := make([]byte,1000)
			conn.Read(b)
			fmt.Println(string(b))
		}()
		time.Sleep(3 * time.Second)
	}
	//fmt.Println(ioutil.ReadAll(conn))
}
