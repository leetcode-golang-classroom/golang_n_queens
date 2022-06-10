package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	n := 4
	for idx := 0; idx < b.N; idx++ {
		solveNQueens(n)
	}
}
func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "n=4",
			args: args{n: 4},
			want: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			name: "n=1",
			args: args{n: 1},
			want: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
