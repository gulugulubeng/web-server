package comment

import (
	"log"
	"os"
)

// 日志输出路径
var loggerPath = "./"

// PLog 项目日志记录
var PLog *log.Logger

// DBLog 数据库记录
var DBLog *log.Logger

// StartLogger 开启日志
func StartLogger(logPath string) {
	loggerPath = logPath

	projectLogger("proLog.txt")
	dbLogger("dbLog.txt")
}

// 项目日志记录
func projectLogger(fileName string) {
	pLogFile, err := os.Create(loggerPath + "\\" + fileName)
	if err != nil {
		log.Fatalln("日志文件创建错误")
		return
	}
	PLog = log.New(pLogFile, "项目异常：", 8)
}

// 数据库记录
func dbLogger(fileName string) {
	dbLogFile, err := os.Create(loggerPath + "\\" + fileName)
	if err != nil {
		log.Fatalln("日志文件创建错误")
		return
	}
	DBLog = log.New(dbLogFile, "数据库异常：", 8)
}
