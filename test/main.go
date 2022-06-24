package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	matrixgo "github.com/LibertusDio/matrix-go"
)

func main() {
	runtime.GOMAXPROCS(8)
	t1 := time.Now().Unix()
	m := matrixgo.NewIntMatrix(10000, 15000)
	t2 := time.Now().Unix()
	fmt.Printf("create m in %ds\r\n", t2-t1)

	n := matrixgo.NewIntMatrix(20000, 10000)
	t1 = time.Now().Unix()
	fmt.Printf("create n in %ds\r\n", t1-t2)

	m = matrixgo.MIntScalar(func(i int) int {
		return rand.Intn(3)
	}, m)
	t2 = time.Now().Unix()
	fmt.Printf("scalar m in %ds\r\n", t2-t1)

	n = matrixgo.MIntScalar(func(i int) int {
		return rand.Intn(4)
	}, n)
	t1 = time.Now().Unix()
	fmt.Printf("scalar n in %ds\r\n", t1-t2)

	matrixgo.MIntProduct(m, n)
	t2 = time.Now().Unix()
	fmt.Printf("product mn in %ds\r\n", t2-t1)

}
