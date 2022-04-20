package bark

import "testing"

func TestBark(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "bark test",
			args: args{
				msg: "unit test",
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bark(tt.args.msg); got != tt.want {
				t.Errorf("Bark() = %v, want %v", got, tt.want)
			}
		})
	}
}
