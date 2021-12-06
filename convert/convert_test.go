package convert

import (
	"reflect"
	"testing"
)

func Test_findAlts(t *testing.T) {
	type args struct {
		input        string
		translations map[string]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findAlts(tt.args.input, tt.args.translations)
		})
	}
}

func TestRun(t *testing.T) {
	type args struct {
		path *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_process(t *testing.T) {
	type args struct {
		path string
		dt   map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    *load
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := process(tt.args.path, tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}
