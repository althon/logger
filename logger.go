package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var info *log.Logger
var errInfo *log.Logger
func Init(){
	file,err:=os.OpenFile("info.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		fmt.Println("[Warning] cannot create info.log file")
	}else{
		info = log.New(io.MultiWriter(file,os.Stderr),"[info] ",log.Ldate | log.Ltime)
	}

	file,err =os.OpenFile("error.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		fmt.Println("[Warning] cannot create error.log file")
	}else{
		errInfo = log.New(io.MultiWriter(file,os.Stderr),"[error] ",log.Ldate | log.Ltime)
	}
}

func Println(a ...interface{}) {
	if info!=nil{
		info.Println(a...)
	}
}

func Printf(format string, v ...interface{}){
	if info!=nil{
		info.Printf(format,v...)
	}
}

func Errorln(a ...interface{}) {
	if errInfo!=nil{
		errInfo.Println(a...)
	}
}

func Errorf(format string, v ...interface{}) {
	if errInfo!=nil{
		errInfo.Printf(format,v...)
	}
}
