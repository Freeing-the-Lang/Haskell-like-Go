package main

import (
	"fmt"
)

// ===== Functional Core =====

func Map[T any, R any](xs []T, f func(T) R) []R {
	out := make([]R, len(xs))
	for i, v := range xs {
		out[i] = f(v)
	}
	return out
}

func Filter[T any](xs []T, f func(T) bool) []T {
	out := []T{}
	for _, v := range xs {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func Foldl[T any, R any](xs []T, init R, f func(R, T) R) R {
	acc := init
	for _, v := range xs {
		acc = f(acc, v)
	}
	return acc
}

// ===== Entry Point =====

func main() {
	nums := []int{1, 2, 3, 4, 5}

	squares := Map(nums, func(x int) int { return x * x })
	filtered := Filter(squares, func(x int) bool { return x > 5 })
	sum := Foldl(filtered, 0, func(acc, x int) int { return acc + x })

	fmt.Println("🟣 squares :", squares)
	fmt.Println("🟢 filtered:", filtered)
	fmt.Println("🟡 sum     :", sum)
}
