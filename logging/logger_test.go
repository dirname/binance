package logging

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestDPanic(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "DPanic Logger", args: args{
			format: "%s",
			args:   []interface{}{"DPanic Logger"},
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Debug Logger", args: args{
			format: "%s",
			args:   []interface{}{"Debug Logger"},
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Error Logger", args: args{
			format: "%s",
			args:   []interface{}{"Error Logger"},
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestFatal(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Fatal Logger", args: args{
			format: "%s",
			args:   []interface{}{"Fatal Logger"},
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Info Logger", args: args{
			format: "%s",
			args:   []interface{}{"Info Logger"},
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPanic(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Panic Logger", args: args{
			format: "%s",
			args:   []interface{}{"Panic Logger"},
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSetLevel(t *testing.T) {
	type args struct {
		level zapcore.Level
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Set Logger Level", args: args{
			level: zapcore.InfoLevel,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Warn Logger", args: args{
			format: "%s",
			args:   []interface{}{"Warn Logger"},
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
