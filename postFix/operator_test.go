package postFix

import "testing"

func TestOperators_IsOperator(t *testing.T) {
	type fields struct {
		Types map[string]Operation
	}
	type args struct {
		element string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "",
			fields: fields{
				Types: map[string]Operation{
					"+": &SumMock{},
				},
			},
			args: args{
				element: "-",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operators{
				Types: tt.fields.Types,
			}
			if got := o.IsOperator(tt.args.element); got != tt.want {
				t.Errorf("IsOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

type SumMock struct {
}

func (s *SumMock) Calculate(a, b float64) float64 {
	return a + b
}
