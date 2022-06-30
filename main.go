package main

import (
	"fmt"

	"github.com/tarm/serial"
)

func main() {
	var com string //"COM2"
	var baud int   //115200
	fmt.Println("请输入端口号 波特率 使用空格隔开!")
	fmt.Scanf("%s %d", &com, &baud)
	conn := new(SerialConnection)
	err := conn.ConnectToSerial(com, baud)
	if err == nil {
		fmt.Printf("连接端口%v成功!", com)
		conn.Send("xiaoxi!!!") //发送 xiaoxi!!!
		conn.ReadSerial()      //读取消息
	}
}

func (sc *SerialConnection) Send(test string) (int, error) {
	n, err := sc.S.Write([]byte(test)) // 发送 内容 test
	return n, err
}

//连接串口
func (sc *SerialConnection) ConnectToSerial(name string, baud int) error {
	c := &serial.Config{Name: name, Baud: baud} //设置串口名称 波特率等配置信息
	ch := make(chan []byte, 128)
	c2 := make(chan struct{}, 10)
	sc.Ch = &ch
	//打开串口
	s, err := serial.OpenPort(c) //打开串口操作
	if err != nil {              //判断打开是否失败
		return err
	}
	sc.S = s
	sc.StopCh = &c2
	return nil
}

type SerialConnection struct {
	S      *serial.Port
	Ch     *chan []byte
	StopCh *chan struct{}
}

const (
	MAXRWLEN = 128
)

func (sc *SerialConnection) ReadSerial() {
	var num int
	for {
		select {
		case <-(*sc.StopCh):
			return
		default:
			buffer := make([]byte, MAXRWLEN)
			num, _ = (*sc.S).Read(buffer)
			if num > 0 {
				(*sc.Ch) <- buffer
				fmt.Println(string(buffer))
			}
		}
	}
}
