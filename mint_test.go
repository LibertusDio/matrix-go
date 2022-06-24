package matrixgo

import (
	"reflect"
	"testing"
)

func TestNewIntMatrix(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want MInt
	}{
		{
			name: "test create int matrix",
			args: args{x: 3, y: 4},
			want: MInt{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			name: "test create zero int matrix",
			args: args{x: 0, y: 0},
			want: MInt{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIntMatrix(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMIntScalarCalculation(t *testing.T) {
	type args struct {
		m MInt
		n int
		f func(m MInt, n int) MInt
	}
	tests := []struct {
		name string
		args args
		want MInt
	}{
		{
			name: "test add int matrix",
			args: args{m: NewIntMatrix(3, 4), n: 1, f: MIntScalarAdd},
			want: MInt{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
		{
			name: "test multi int matrix",
			args: args{m: MInt{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, n: 3, f: MIntScalarMul},
			want: MInt{{3, 3, 3}, {3, 3, 3}, {3, 3, 3}, {3, 3, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.f(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MIntScalarF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMIntTranspose(t *testing.T) {
	type args struct {
		m MInt
	}
	tests := []struct {
		name string
		args args
		want MInt
	}{
		{
			name: "test tranpose int matrix",
			args: args{m: MInt{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}},
			want: MInt{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MIntTranspose(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MIntTranspose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMIntMainDiag(t *testing.T) {
	type args struct {
		m MInt
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "test err maindiag int matrix",
			args:    args{m: MInt{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test maindiag int matrix",
			args:    args{m: MInt{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}},
			want:    []int{1, 1, 1, 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MIntMainDiag(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("MIntMainDiag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MIntMainDiag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMIntSum(t *testing.T) {
	type args struct {
		a MInt
		b MInt
	}
	tests := []struct {
		name    string
		args    args
		want    MInt
		wantErr bool
	}{
		{
			name:    "test sum 2 int matrix",
			args:    args{a: MInt{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, b: MInt{{2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}}},
			want:    MInt{{3, 3, 3, 3}, {3, 3, 3, 3}, {3, 3, 3, 3}},
			wantErr: false,
		},
		{
			name:    "test err sum 2 int matrix",
			args:    args{a: MInt{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, b: MInt{{2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MIntSum(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("MIntSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MIntSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMIntProduct(t *testing.T) {
	type args struct {
		a MInt
		b MInt
	}
	tests := []struct {
		name    string
		args    args
		want    MInt
		wantErr bool
	}{
		{
			name:    "test err product 2 int matrix",
			args:    args{a: MInt{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, b: MInt{{2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test product 2 int matrix",
			args:    args{a: MInt{{1, 2, 3}, {4, 5, 6}}, b: MInt{{10, 11}, {20, 21}, {30, 31}}},
			want:    MInt{{140, 146}, {320, 335}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MIntProduct(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("MIntProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MIntProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMIntSubmatrix(t *testing.T) {
	type args struct {
		m MInt
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		want    MInt
		wantErr bool
	}{
		{
			name:    "test sub int matrix",
			args:    args{m: MInt{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 1, y: 1},
			want:    MInt{{1, 3}, {7, 9}},
			wantErr: false,
		},
		{
			name:    "test error max x sub int matrix",
			args:    args{m: MInt{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 3, y: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test error max y sub int matrix",
			args:    args{m: MInt{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 1, y: 3},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test error min x sub int matrix",
			args:    args{m: MInt{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: -1, y: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test error min y sub int matrix",
			args:    args{m: MInt{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 0, y: -1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test error zero matrix sub int matrix",
			args:    args{m: MInt{}, x: 0, y: 0},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test max edge sub int matrix",
			args:    args{m: MInt{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 2, y: 2},
			want:    MInt{{1, 2}, {4, 5}},
			wantErr: false,
		},
		{
			name:    "test min edge sub int matrix",
			args:    args{m: MInt{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 0, y: 0},
			want:    MInt{{5, 6}, {8, 9}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MIntSubmatrix(tt.args.m, tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("MIntSubmatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MIntSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
