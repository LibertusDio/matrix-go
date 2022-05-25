package matrixgo

import (
	"errors"
	"sync"
)

type MInt [][]int

func NewIntMatrix(x, y int) MInt {
	mint := make([][]int, x)
	for i := 0; i < x; i++ {
		mint[i] = make([]int, y)
	}
	return (MInt)(mint)
}

func MIntScalar(f func(i int) int, m MInt) MInt {
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

func MIntScalarAdd(m MInt, n int) MInt {
	return MIntScalar(func(i int) int { return i + n }, m)
}

func MIntScalarMul(m MInt, n int) MInt {
	return MIntScalar(func(i int) int { return i * n }, m)
}

func MIntTranspose(m MInt) MInt {
	var x, y int
	x = len(m)
	if x > 0 {
		y = len(m[0])
	}
	nm := NewIntMatrix(y, x)
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

func MIntMainDiag(m MInt) ([]int, error) {
	var x, y int
	x = len(m)
	if x > 0 {
		y = len(m[0])
	}
	if x != y {
		return nil, errors.New("not square")
	}
	nm := make([]int, x)
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

func MIntSum(a MInt, b MInt) (MInt, error) {
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
	nm := NewIntMatrix(xa, ya)
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
