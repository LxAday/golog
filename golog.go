package golog

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var (
	LogDir    = "./log" //log save directory
	Partition = "day"   //journal format，day，month，year
	Async     = false   //is sync,error free return
	Skip      = 1       //step
)

//WriteSuLog Write Success Log
//msg message
func WriteSuLog(msg string) error {
	_, fName, line, _ := runtime.Caller(Skip)
	if !Async {
		return writeSuLog(msg, fName, line)
	}
	go func() {
		_ = writeSuLog(msg, fName, line)
	}()
	return nil
}

//WriteErLog Write error log
//msg message
func WriteErLog(msg string) error {
	_, fName, line, _ := runtime.Caller(Skip)
	if !Async {
		return writeErLog(msg, fName, line)
	}
	go func() {
		_ = writeErLog(msg, fName, line)
	}()
	return nil
}

//writeErLog
func writeErLog(msg, fName string, line int) error {
	path, e := createLogDir()
	if e != nil {
		return e
	}
	path = filepath.Join(path, "error.log")
	return writeLog(path, msg, "error", fName, line)
}

//writeSuLog
func writeSuLog(msg, fName string, line int) error {
	path, e := createLogDir()
	if e != nil {
		return e
	}
	path = filepath.Join(path, "succeed.log")
	return writeLog(path, msg, "succeed", fName, line)
}

//writeLog Written to the log
func writeLog(path, msg, level, fName string, line int) error {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	file, e := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if e != nil {
		return e
	}
	defer file.Close()
	_, e = file.WriteString(fmt.Sprintf("[%s] [time:%s] [msg:%s] [file:%s] [line:%d]\n", level, time.Now().Format("2006-01-02 15:04:05"), msg, fName, line))
	if e != nil {
		return e
	}
	if e = file.Sync(); e != nil {
		return e
	}
	return nil
}

//createLogDir
func createLogDir() (string, error) {
	folderName := ""
	t := time.Now()

	switch Partition {
	case "day":
		folderName = t.Format("2006-01-02")
	case "month":
		folderName = t.Format("2006-01")
	case "year":
		folderName = t.Format("2006")
	default:
		return "", errors.New("the log save format is invalid")
	}

	dir := filepath.Join(LogDir, folderName)
	if e := os.MkdirAll(dir, 0666); e != nil {
		return "", e
	}
	return dir, nil
}
