package main

import (
	"fmt"
	"net"
	"encoding/json"
)

type FileClientTransfer struct {
	Conn *net.Conn `json:"conn"` //socket连接

	label string `json:"label"` //master-主服务端(中心服务器)  slave-从服务器

	FileName      string `json:"file_name"`       //待发送文件名称
	MergeFileName string `json:"merge_file_name"` //待合并文件名称
	Coroutine     int    `json:"coroutine"`       //协程数量或拆分文件的数量
	BufSize       int    `json:"buf_size"`        //单次发送数据的大小
}

func main() {
	a := `{"conn":null,"file_name":"adada.jpg","merge_file_name":"","coroutine":0,"buf_size":0}`
	f := &FileClientTransfer{}
	json.Unmarshal([]byte(a), f)

	fmt.Println(f)
}
