package matrixgo

import (
	"errors"
	"sync"
)

type MInt [][]int

func NewIntMatrix(x, y int) MInt {
	mint := make([][]int, y)
	for i := 0; i < y; i++ {
		mint[i] = make([]int, x)
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
	y = len(m)
	if y > 0 {
		x = len(m[0])
	}
	nm := NewIntMatrix(y, x)
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
	nm := NewIntMatrix(xa, ya)
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

func MIntProduct(a MInt, b MInt) (MInt, error) {
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
	nm := NewIntMatrix(ya, xb)
	n := xa
	var wg sync.WaitGroup
	for i := 0; i < xb; i++ {
		for j := 0; j < ya; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sum := 0
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

func MIntSubmatrix(a MInt, x, y int) (MInt, error) {
	var xa, ya int
	xa = len(a)
	if xa > 0 {
		ya = len(a[0])
	}
	if xa <= 0 || ya <= 0 || x < 0 || y < 0 || x >= xa || y >= ya {
		return nil, errors.New("invalid")
	}

	n := NewIntMatrix(xa-1, ya-1)
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
