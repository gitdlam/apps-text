package text

import "testing"

func TestValidDigits4To15(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		trueFalse bool
	}{{
		args:      args{s: "1234"},
		trueFalse: true,
	},
		{
			args:      args{s: "abcd"},
			trueFalse: false,
		},
		{
			args:      args{s: "1255444555115555"},
			trueFalse: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ValidDigits4To15(tt.args.s) != tt.trueFalse {
				t.Fail()
			}
		})
	}
}
