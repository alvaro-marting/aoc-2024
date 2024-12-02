package pkg

import (
	"iter"
	"log"
	"slices"
	"strconv"
)

func MapStringSliceToInt(s []string) []int {
	return slices.Collect(mapSlice(slices.Values(s), func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return i
	}))
}

func mapSlice[T, U any](seq iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for a := range seq {
			if !yield(f(a)) {
				return
			}
		}
	}
}
