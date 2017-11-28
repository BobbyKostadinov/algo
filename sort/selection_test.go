package sort

import (
	"reflect"
	"testing"
)

func TestSelection(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Selection(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection() = %v, want %v", got, tt.want)
			}
		})
	}
}
