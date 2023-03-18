package mlog

import (
	"io"
	"log"
	"mmo_server/utils/globalConfig"
	"os"
)

type MLogger struct {
	l *log.Logger
}

func (ml *MLogger) Println(v ...any) {
	if !globalConfig.ProjectCfg.DevelopmentMode {
		return
	}
	ml.l.Println(v...)
}
func (ml *MLogger) Print(v ...any) {
	if !globalConfig.ProjectCfg.DevelopmentMode {
		return
	}
	ml.l.Print(v...)
}
func (ml *MLogger) Printf(format string, v ...any) {
	if !globalConfig.ProjectCfg.DevelopmentMode {
		return
	}
	ml.l.Printf(format, v...)
}

var (
	Trace   *MLogger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init() {

	Trace = &MLogger{
		l: log.New(os.Stdout,
			"[TRACE]>> ",
			log.Ldate|log.Ltime),
	}

	Info = log.New(os.Stdout,
		"[INFO]>>",
		log.Ldate|log.Ltime)

	Warning = log.New(os.Stdout,
		"[Warning]>>",
		log.Ldate|log.Ltime|log.Lshortfile)

	file, err := os.OpenFile("../Logs/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("无法打开Log文件: ", err)
	}
	Error = log.New(io.MultiWriter(file, os.Stderr),
		"[Error]>>",
		log.Ldate|log.Ltime|log.Lshortfile)
}
