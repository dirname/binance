package logging

import (
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestEnable(t *testing.T) {
	type args struct {
		enable bool
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "TestEnable", args: args{true}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Enable(tt.args.enable)
		})
	}
}

func TestGetInstance(t *testing.T) {
	tests := []struct {
		name string
		want *PerformanceLogger
	}{
		{"TestGetInstance", nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Enable(true)
			if got := GetInstance(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerformanceLogger_Start(t *testing.T) {
	type fields struct {
		logger *log.Logger
		enable bool
		file   *os.File
		index  int
		start  time.Time
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"TestPerformanceLogger_Start", fields{
			logger: nil,
			enable: false,
			file:   nil,
			index:  0,
			start:  time.Time{},
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Enable(true)
			p := &PerformanceLogger{
				logger: tt.fields.logger,
				enable: tt.fields.enable,
				file:   tt.fields.file,
				index:  tt.fields.index,
				start:  tt.fields.start,
			}
			p.Start()
		})
	}
}

func TestPerformanceLogger_StopAndLog(t *testing.T) {
	type fields struct {
		logger *log.Logger
		enable bool
		file   *os.File
		index  int
		start  time.Time
	}
	type args struct {
		method string
		url    string
	}
	fileName := time.Now().Format("20060102_150405.txt")
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal("Failed to open file: ", fileName)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestPerformanceLogger_StopAndLog", fields{
			logger: log.New(file, "", 0),
			enable: false,
			file:   nil,
			index:  1,
			start:  time.Time{},
		}, args{
			method: "test",
			url:    "test",
		}},
		{"TestPerformanceLogger_StopAndLog", fields{
			logger: log.New(file, "", 0),
			enable: false,
			file:   nil,
			index:  1,
			start:  time.Time{},
		}, args{
			method: "test",
			url:    "https://google.com?test",
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Enable(true)
			p := &PerformanceLogger{
				logger: tt.fields.logger,
				enable: tt.fields.enable,
				file:   tt.fields.file,
				index:  tt.fields.index,
				start:  tt.fields.start,
			}
			p.StopAndLog(tt.args.method, tt.args.url)
		})
	}
}
