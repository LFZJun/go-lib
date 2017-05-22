package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	tcpBody := ``
	conn.Write([]byte(tcpBody))
	//lines := strings.Split(tcpBody, "\r\n")
	//for _, line := range lines {
	//	conn.Write([]byte(line + "\r\n"))
	//}
	//conn.Write([]byte("\r\n"))
	conn.SetDeadline(time.Now().Add(time.Second))
	resp, err := ioutil.ReadAll(conn)
	fmt.Println(string(resp))
}
