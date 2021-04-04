package golog

import (
	"testing"
	"time"
)

func TestWriteErLog(t *testing.T) {
	if e := WriteErLog("error"); e != nil {
		t.Error(e)
	}
	t.Log("ok")
}

func TestWriteSuLog(t *testing.T) {
	if e := WriteSuLog("success"); e != nil {
		t.Error(e)
	}
	t.Log("ok")
}

func TestWriteErLog2(t *testing.T) {
	LogDir = "./logs_er"
	Async = true
	for i := 0; i < 10; i++ {
		if e := WriteErLog("error"); e != nil {
			t.Error(e)
		}
		t.Log("ok")
	}
	time.Sleep(time.Second)
}

func TestWriteSuLog2(t *testing.T) {
	LogDir = "./logs_su"
	Async = true
	for i := 0; i < 100; i++ {
		if e := WriteSuLog("success"); e != nil {
			t.Error(e)
		}
	}
	t.Log("ok")
	time.Sleep(time.Second * 3)
}
