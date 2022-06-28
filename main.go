package main

import (
	"log"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "COM45", Baud: 115200} //设置串口名称 波特率等配置信息
	s, err := serial.OpenPort(c)                     //打开串口操作
	if err != nil {                                  //判断打开是否失败
		log.Fatal(err) //打印失败信息
	}
	n, err := s.Write([]byte("test")) // 发送 内容 test
	if err != nil {
		log.Fatal(err, n) //打印发送成功或者失败信息
	}
	buf := make([]byte, 128) //定义一个空字节用于接收消息
	n, err = s.Read(buf)     //读取串口发送的内容
	if err != nil {          //判断读取是否失败了
		log.Fatal(err, n) //打印失败的信息
	}
	log.Printf("%q", buf[:n]) //打印读取到的内容
}
