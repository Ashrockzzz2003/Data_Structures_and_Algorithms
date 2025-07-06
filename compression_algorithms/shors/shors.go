package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gcd(m, n int) int {
	for n != 0 {
		m, n = n, m%n
	}
	return m
}

func order(a, N int) int {
	for i := 1; i < N; i++ {
		if mExpo(a, i, N) == 1 {
			return i
		}
	}
	return -1
}
func mExpo(b, e, m int) int {
	total := 1
	b = b % m
	for e > 0 {
		if e%2 == 1 {
			total = (total * b) % m
		}
		e = e >> 1
		b = (b * b) % m
	}
	return total
}

func Shors(N int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(N-1) + 1
	g := gcd(a, N)
	if g != 1 {
		return g, N / g
	}

	if N%2 == 0 {
		return 2, N / 2
	}

	o := order(a, N)
	if o == -1 {
		return 0, 0
	}

	if o%2 == 0 {
		x1 := mExpo(a, o/2, N)
		if x1 != N-1 {
			f1 := gcd(a-1, N)
			f2 := N / f1
			return f1, f2
		}
	}
	return 0, 0
}

func main() {
	N := 12
	//N:= 16
	//N:=342
	//N:= 100
	fmt.Println("Trying to factorize N:", N)

	for {
		p, q := Shors(N)
		if p != 0 && q != 0 {
			fmt.Printf("Factors of %d: %d and %d\n", N, p, q)
			break
		} else {
			fmt.Println("No factors found, retrying...")
		}
	}
}
