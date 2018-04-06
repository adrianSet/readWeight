package log1

import (
	"log"
	"os"
	"io"
)

var (
	//日志处理器
	Logger *log.Logger
)


func init(){

	log.SetPrefix("trace:")
	log.SetFlags(log.Ldate | log.Llongfile)
	file,error:=os.OpenFile("error.text",os.O_APPEND|os.O_CREATE|os.O_WRONLY ,0666)
	if error!=nil{
		log.Fatalln(error)
	}

	Logger= log.New(io.MultiWriter(file, os.Stdout), "TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}
