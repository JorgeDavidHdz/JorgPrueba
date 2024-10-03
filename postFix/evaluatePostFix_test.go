package postFix

import "testing"

func TestEvaluatePostFix(t *testing.T) {
	type args struct {
		postFix []string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "",
			args: args{
				postFix: []string{
					"1",
					"2.5",
					"3",
					"/",
					"4",
					"*",
					"+",
				},
			},
			want:    4.333333333333334,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvaluatePostFix(tt.args.postFix)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvaluatePostFix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EvaluatePostFix() got = %v, want %v", got, tt.want)
			}
		})
	}
}
