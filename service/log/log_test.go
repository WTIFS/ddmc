package log

import "testing"

func TestDebug(t *testing.T) {
	type args struct {
		format string
		msg    []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test debug",
			args: args{
				format: "%s: %s",
				msg:    []any{"debug", "hi saber"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.format, tt.args.msg...)
		})
	}
}

func TestErr(t *testing.T) {
	type args struct {
		format string
		msg    []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test err",
			args: args{
				format: "%s: %s",
				msg:    []any{"err", "hi saber"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Err(tt.args.format, tt.args.msg...)
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		format string
		msg    []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test info",
			args: args{
				format: "%s: %s",
				msg:    []any{"info", "hi saber"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.format, tt.args.msg...)
		})
	}
}

func TestLog(t *testing.T) {
	type args struct {
		level  string
		format string
		msg    []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test debug",
			args: args{
				level:  "FATAL",
				format: "%s: %s",
				msg:    []any{"log", "hi saber"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Log(tt.args.level, tt.args.format, tt.args.msg...)
		})
	}
}
