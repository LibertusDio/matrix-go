package matrixgo

import (
	"errors"
	"sync"
)

type MFloat [][]float64

func NewFloatMatrix(x, y int) MFloat {
	mfloat64 := make([][]float64, x)
	for i := 0; i < x; i++ {
		mfloat64[i] = make([]float64, y)
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
	x = len(m)
	if x > 0 {
		y = len(m[0])
	}
	nm := NewFloatMatrix(y, x)
	var wg sync.WaitGroup
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				nm[j][i] = m[i][j]
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
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			nm[i] = m[i][i]
		}(i)
	}
	wg.Wait()
	return nm, nil
}

func MFloatSum(a MFloat, b MFloat) (MFloat, error) {
	var xa, ya int
	xa = len(a)
	if xa > 0 {
		ya = len(a[0])
	}
	var xb, yb int
	xb = len(b)
	if xb > 0 {
		yb = len(b[0])
	}
	if xa != xb || ya != yb {
		return nil, errors.New("not balance")
	}
	nm := NewFloatMatrix(xa, ya)
	var wg sync.WaitGroup
	for i := 0; i < xa; i++ {
		for j := 0; j < ya; j++ {
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
