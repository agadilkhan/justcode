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
		{name: "test1",
			args: struct{ jsonStr string }{jsonStr: `{
		"key1": "value1",
		"key2": {
			"subkey1": "subvalue1"
		}
	}`}, want: `result := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"subkey1": "subvalue1",
		},
	}
	`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateGoCodeFromJSON(tt.args.jsonStr); got != tt.want {
				t.Errorf("generateGoCodeFromJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
