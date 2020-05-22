package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name            string
		args            args
		want            string
		errShouldRaised bool
	}{
		{name: "empty", args: args{val: ""}, want: "", errShouldRaised: false},
		{name: "255 to 0xFF", args: args{val: "255"}, want: "0xFF", errShouldRaised: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.args.val)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			if tt.errShouldRaised {
				if err == nil {
					t.Errorf("should raise err but err is nil")
				}
			} else {
				if err != nil {
					t.Errorf("err = %v", err)
				}
			}
		})
	}
}
