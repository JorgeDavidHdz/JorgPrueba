package postFix

import (
	"reflect"
	"testing"
)

func TestInfixToPostFix(t *testing.T) {
	type args struct {
		elements  []string
		operators Operators
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				elements: []string{
					"1",
					"+",
					"2.5",
					"/",
					"3",
					"*",
					"4",
				},
				operators: Operators{
					Types: map[string]Operation{
						"+": &SumMock{},
						"-": &SubMock{},
						"*": &MulMock{},
						"/": &DivMock{},
					},
				},
			},
			want: []string{
				"1",
				"2.5",
				"3",
				"/",
				"4",
				"*",
				"+",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InfixToPostFix(tt.args.elements, tt.args.operators)
			if (err != nil) != tt.wantErr {
				t.Errorf("InfixToPostFix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InfixToPostFix() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type SubMock struct {
}

func (s *SubMock) Calculate(a, b float64) float64 {
	return a - b
}

type MulMock struct {
}

func (s *MulMock) Calculate(a, b float64) float64 {
	return a * b
}

type DivMock struct {
}

func (s *DivMock) Calculate(a, b float64) float64 {
	return a / b
}
