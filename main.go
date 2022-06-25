// Copyright (C) 2022 Evgeny Kuznetsov (evgeny@kuznetsov.md)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along tihe this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	MinUint uint = 0
	MaxUint      = ^MinUint
	MaxInt       = int(MaxUint >> 1)
)

func main() {
	firstNC(MaxInt, os.Stdout)
}

// firstNC outputs first num Carmichael numbers to w
func firstNC(num int, w io.Writer) {
	for n, i := 1, 1; i <= num; i++ {
		start := time.Now()
		n = nextC(n)
		t := time.Since(start)
		fmt.Fprintf(w, "%v: %v (in %v)\n", i, n, t)
		n++
	}
}

// nextC finds the lowest Carmichael number >= n
func nextC(n int) int {
	if !isPrime(n) && pFermat(n) {
		return n
	}
	return nextC(n + 1)
}

// isPrime returns true if n is a prime number
func isPrime(n int) bool {
	return (n == smallestDivisor(n))
}

// smallestDivisor returns the smallest divisor of n that is not 1
func smallestDivisor(n int) int {
	return findDivisor(n, 2)
}

// findDivisor searches for the divisor of n >= g
func findDivisor(n, g int) int {
	if (square(g)) > n {
		return n
	}
	if (n % g) == 0 {
		return g
	}
	i := g + 2
	if g == 2 {
		i = 3
	}
	return findDivisor(n, i)
}

// pFermat returns true if n passes Fermat test for primeness for every possible a < n
func pFermat(n int) bool {
	return fermatTestFull(n, n-1)
}

func fermatTestFull(n, a int) bool {
	if a == 0 {
		return true
	}
	if expmod(a, n, n) != a {
		return false
	}
	return fermatTestFull(n, a-1)
}

// expmod returns (base ^ exp) % m
func expmod(base, exp, m int) int {
	if exp == 0 {
		return 1
	}
	if (exp % 2) == 0 {
		return square(expmod(base, exp/2, m)) % m
	}
	return (base * expmod(base, exp-1, m)) % m
}

// square returns n squared
func square(n int) int {
	return n * n
}
