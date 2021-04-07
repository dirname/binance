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
			DPanic(tt.args.format, tt.args.args)
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
			Debug(tt.args.format, tt.args.args)
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
			Error(tt.args.format, tt.args.args)
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
			Info(tt.args.format, tt.args.args)
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
			SetLevel(tt.args.level)
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
			Warn(tt.args.format, tt.args.args)
		})
	}
}
