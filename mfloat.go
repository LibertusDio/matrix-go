package matrixgo

import (
	"errors"
	"sync"
)

type MFloat [][]float64

func NewFloatMatrix(x, y int) MFloat {
	mfloat64 := make([][]float64, y)
	for i := 0; i < y; i++ {
		mfloat64[i] = make([]float64, x)
	}
	return (MFloat)(mfloat64)
}

func MFloatScalar(f func(i float64) float64, m MFloat) MFloat {
	var x, y int
	x = len(m)
	if x > 0 {
		y = len(m[0])
	}
	var wg sync.WaitGroup
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				m[i][j] = f(m[i][j])
			}(i, j)
		}
	}
	wg.Wait()
	return m
}

func MFloatScalarAdd(m MFloat, n float64) MFloat {
	return MFloatScalar(func(i float64) float64 { return i + n }, m)
}

func MFloatScalarMul(m MFloat, n float64) MFloat {
	return MFloatScalar(func(i float64) float64 { return i * n }, m)
}

func MFloatTranspose(m MFloat) MFloat {
	var x, y int
	y = len(m)
	if y > 0 {
		x = len(m[0])
	}
	nm := NewFloatMatrix(y, x)
	var wg sync.WaitGroup
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				nm[i][j] = m[j][i]
			}(i, j)
		}
	}
	wg.Wait()
	return nm
}

func MFloatMainDiag(m MFloat) ([]float64, error) {
	var x, y int
	x = len(m)
	if x > 0 {
		y = len(m[0])
	}
	if x != y {
		return nil, errors.New("not square")
	}
	nm := make([]float64, x)
	var wg sync.WaitGroup
	for i := 0; i < x; i++ {
		nm[i] = m[i][i]
	}
	wg.Wait()
	return nm, nil
}

func MFloatSum(a MFloat, b MFloat) (MFloat, error) {
	var ya, xa int
	ya = len(a)
	if ya > 0 {
		xa = len(a[0])
	}
	var yb, xb int
	yb = len(b)
	if yb > 0 {
		xb = len(b[0])
	}
	if ya != yb || xa != xb {
		return nil, errors.New("not balance")
	}
	nm := NewFloatMatrix(xa, ya)
	var wg sync.WaitGroup
	for i := 0; i < ya; i++ {
		for j := 0; j < xa; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				nm[i][j] = a[i][j] + b[i][j]
			}(i, j)
		}
	}
	wg.Wait()
	return nm, nil
}

func MFloatProduct(a MFloat, b MFloat) (MFloat, error) {
	var ya, xa int
	ya = len(a)
	if ya > 0 {
		xa = len(a[0])
	}
	var yb, xb int
	yb = len(b)
	if yb > 0 {
		xb = len(b[0])
	}
	if ya != xb {
		return nil, errors.New("incompatible")
	}
	nm := NewFloatMatrix(ya, xb)
	n := xa
	var wg sync.WaitGroup
	for i := 0; i < xb; i++ {
		for j := 0; j < ya; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sum := 0.0
				for k := 0; k < n; k++ {
					sum += a[j][k] * b[k][i]
				}
				nm[j][i] = sum
			}(i, j)
		}
	}
	wg.Wait()
	return nm, nil
}

func MFloatSubmatrix(a MFloat, x, y int) (MFloat, error) {
	var xa, ya int
	xa = len(a)
	if xa > 0 {
		ya = len(a[0])
	}
	if xa <= 0 || ya <= 0 || x < 0 || y < 0 || x >= xa || y >= ya {
		return nil, errors.New("invalid")
	}

	n := NewFloatMatrix(xa-1, ya-1)
	for i := 0; i < xa; i++ {
		for j := 0; j < ya; j++ {
			xn, yn := i, j
			if xn == x || yn == y {
				continue
			}
			if xn > x {
				xn = i - 1
			}
			if yn > y {
				yn = j - 1
			}
			n[xn][yn] = a[i][j]
		}
	}
	return n, nil
}
