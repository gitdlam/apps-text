package text

import (
	"testing"
)

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

func TestEWMEncodeUUID(t *testing.T) {
	data := map[string]string{
		"005056B31E4F1EE48398F7FA866FCB15": "051MinvF7kI3cFVwXc}B5G",
		"005056B33BF11EE3BAC398355F449E72": "051Mipln7kEwmvWrNqIUSW",
		"005056B3483B1ED487A6145349669844": "051MiqWx7jI7fXHJIMQOH0",
		"4C419E52B069732EE10000000A141F1F": "J46UKh1fSoxX00002XGV7m",
		"91AD714ECF7C7571E10000000A141F31": "aQrnJizyTN7X00002XGVCG",
	}

	for k, v := range data {
		t.Run("", func(t *testing.T) {
			if got := EWMEncodeUUID(k); got != v {
				t.Errorf("EWMDecodeUUID() = %v, want %v", got, v)
			}
		})
	}
}

func TestEWMDecodeUUID(t *testing.T) {

	data := map[string]string{
		"005056B31E4F1EE48398F7FA866FCB15": "051MinvF7kI3cFVwXc}B5G",
		"005056B33BF11EE3BAC398355F449E72": "051Mipln7kEwmvWrNqIUSW",
		"005056B3483B1ED487A6145349669844": "051MiqWx7jI7fXHJIMQOH0",
		"4C419E52B069732EE10000000A141F1F": "J46UKh1fSoxX00002XGV7m",
		"91AD714ECF7C7571E10000000A141F31": "aQrnJizyTN7X00002XGVCG",
	}

	for k, v := range data {
		t.Run("", func(t *testing.T) {
			if got := EWMDecodeUUID(v); got != k {
				t.Errorf("EWMDecodeUUID() = %v, want %v", got, k)
			}
		})
	}
}
