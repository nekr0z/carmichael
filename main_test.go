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
	"io"
	"testing"
)

func TestNextC(t *testing.T) {
	start := 1
	want := 561
	got := nextC(start)
	if got != want {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func benchmarkC(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		firstNC(i, io.Discard)
	}
}

func BenchmarkC1(b *testing.B)  { benchmarkC(1, b) }
func BenchmarkC2(b *testing.B)  { benchmarkC(2, b) }
func BenchmarkC3(b *testing.B)  { benchmarkC(3, b) }
func BenchmarkC10(b *testing.B) { benchmarkC(10, b) }
func BenchmarkC20(b *testing.B) { benchmarkC(20, b) }
func BenchmarkC50(b *testing.B) { benchmarkC(50, b) }
