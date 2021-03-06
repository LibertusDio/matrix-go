package matrixgo

import (
	"reflect"
	"testing"
)

func TestNewFloatMatrix(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want MFloat
	}{
		{
			name: "test create float matrix",
			args: args{x: 3, y: 4},
			want: MFloat{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			name: "test create zero float matrix",
			args: args{x: 0, y: 0},
			want: MFloat{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFloatMatrix(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFloatMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMFloatScalarCalculation(t *testing.T) {
	type args struct {
		m MFloat
		n float64
		f func(m MFloat, n float64) MFloat
	}
	tests := []struct {
		name string
		args args
		want MFloat
	}{
		{
			name: "test add float matrix",
			args: args{m: NewFloatMatrix(3, 4), n: 1, f: MFloatScalarAdd},
			want: MFloat{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
		{
			name: "test multi float matrix",
			args: args{m: MFloat{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, n: 3, f: MFloatScalarMul},
			want: MFloat{{3, 3, 3}, {3, 3, 3}, {3, 3, 3}, {3, 3, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.f(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MFloatScalarF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMFloatTranspose(t *testing.T) {
	type args struct {
		m MFloat
	}
	tests := []struct {
		name string
		args args
		want MFloat
	}{
		{
			name: "test tranpose float matrix",
			args: args{m: MFloat{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}},
			want: MFloat{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MFloatTranspose(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MFloatTranspose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMFloatMainDiag(t *testing.T) {
	type args struct {
		m MFloat
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name:    "test err maindiag float matrix",
			args:    args{m: MFloat{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test maindiag float matrix",
			args:    args{m: MFloat{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}},
			want:    []float64{1, 1, 1, 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MFloatMainDiag(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("MFloatMainDiag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MFloatMainDiag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMFloatSum(t *testing.T) {
	type args struct {
		a MFloat
		b MFloat
	}
	tests := []struct {
		name    string
		args    args
		want    MFloat
		wantErr bool
	}{
		{
			name:    "test sum 2 float matrix",
			args:    args{a: MFloat{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, b: MFloat{{2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}}},
			want:    MFloat{{3, 3, 3, 3}, {3, 3, 3, 3}, {3, 3, 3, 3}},
			wantErr: false,
		},
		{
			name:    "test err sum 2 float matrix",
			args:    args{a: MFloat{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, b: MFloat{{2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MFloatSum(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("MFloatSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MFloatSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMFloatProduct(t *testing.T) {
	type args struct {
		a MFloat
		b MFloat
	}
	tests := []struct {
		name    string
		args    args
		want    MFloat
		wantErr bool
	}{
		{
			name:    "test err product 2 float matrix",
			args:    args{a: MFloat{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, b: MFloat{{2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test product 2 float matrix",
			args:    args{a: MFloat{{1, 2, 3}, {4, 5, 6}}, b: MFloat{{10, 11}, {20, 21}, {30, 31}}},
			want:    MFloat{{140, 146}, {320, 335}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MFloatProduct(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("MFloatProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MFloatProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMFloatSubmatrix(t *testing.T) {
	type args struct {
		m MFloat
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		want    MFloat
		wantErr bool
	}{
		{
			name:    "test sub float matrix",
			args:    args{m: MFloat{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 1, y: 1},
			want:    MFloat{{1, 3}, {7, 9}},
			wantErr: false,
		},
		{
			name:    "test error max x sub float matrix",
			args:    args{m: MFloat{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 3, y: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test error max y sub float matrix",
			args:    args{m: MFloat{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 1, y: 3},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test error min x sub float matrix",
			args:    args{m: MFloat{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: -1, y: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test error min y sub float matrix",
			args:    args{m: MFloat{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 0, y: -1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test error zero matrix sub float matrix",
			args:    args{m: MFloat{}, x: 0, y: 0},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test max edge sub float matrix",
			args:    args{m: MFloat{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 2, y: 2},
			want:    MFloat{{1, 2}, {4, 5}},
			wantErr: false,
		},
		{
			name:    "test min edge sub float matrix",
			args:    args{m: MFloat{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, x: 0, y: 0},
			want:    MFloat{{5, 6}, {8, 9}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MFloatSubmatrix(tt.args.m, tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("MFloatSubmatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MFloatSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
