package conclusion

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			name: "example 2",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}

	for idx, f := range []func([][]int) [][]int{
		updateMatrix,
		updateMatrix1,
		updateMatrix2,
	} {
		t.Run(fmt.Sprintf("func#%d", idx), func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := f(copyMatrix(tt.args.mat)); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}

func copyMatrix(mat [][]int) [][]int {
	var out = make([][]int, len(mat))
	for i := range mat {
		out[i] = make([]int, len(mat[i]))
		for j := range mat {
			out[i][j] = mat[i][j]
		}
	}
	return out
}
