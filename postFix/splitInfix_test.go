package postFix

import (
	"reflect"
	"testing"
)

func TestSplitInfix(t *testing.T) {
	type args struct {
		infix string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Split Infix Success",
			args: args{
				infix: "1+2.5/3*4",
			},
			want: []string{
				"1",
				"+",
				"2.5",
				"/",
				"3",
				"*",
				"4",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitInfix(tt.args.infix)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitInfix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitInfix() got = %v, want %v", got, tt.want)
			}
		})
	}
}
