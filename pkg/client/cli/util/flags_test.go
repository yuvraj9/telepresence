package util

import (
	"strings"
	"testing"
)

func TestGetUnparsedFlagValue(t *testing.T) {
	tests := []struct {
		args    []string
		flag    string
		wantV   string
		wantErr bool
	}{
		{
			[]string{"--name="},
			"--name",
			"",
			true,
		},
		{
			[]string{"--name"},
			"--name",
			"",
			true,
		},
		{
			[]string{"--name", "--other"},
			"--name",
			"",
			true,
		},
		{
			[]string{"--name", "-o"},
			"--name",
			"",
			true,
		},
		{
			[]string{"--name=value"},
			"--name",
			"value",
			false,
		},
		{
			[]string{"--name=-value-"},
			"--name",
			"-value-",
			false,
		},
		{
			[]string{"--name", "value"},
			"--name",
			"value",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(strings.Join(tt.args, "_"), func(t *testing.T) {
			gotV, err := GetUnparsedFlagValue(tt.args, tt.flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUnparsedFlagValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotV != tt.wantV {
				t.Errorf("getUnparsedFlagValue() gotV = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}
