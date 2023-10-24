package excel

import "testing"

func TestIndex2ExcelRow(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{index: 0},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index2ExcelRow(tt.args.index); got != tt.want {
				t.Errorf("Index2ExcelRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
