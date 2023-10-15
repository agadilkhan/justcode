package main

import "testing"

func Test_generateGoCodeFromJSON(t *testing.T) {
	type args struct {
		jsonStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateGoCodeFromJSON(tt.args.jsonStr); got != tt.want {
				t.Errorf("generateGoCodeFromJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
