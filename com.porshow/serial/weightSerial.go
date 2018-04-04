package serial

import (
	"github.com/tarm/serial"
	"log"
	"readWeight/com.porshow/config"
	"runtime/debug"
)

/**
串口指针
 */
var s *serial.Port

func init() {
	c := &serial.Config{Name: config.M[config.WeightMachine], Baud: 9600}
	var port, err = serial.OpenPort(c)
	s = port
	if err != nil {
		log.Fatal(err, "\n", string(debug.Stack()))
	}
}

/**
监听串口信息
 */
func Listening() {
	for {
		buf := make([]byte, 10)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		string1 := string(buf)
		log.Printf("%q", buf[:n])
		log.Println(string1)
	}
}
