package main

import "fmt"

func main() {

	coll := []int{1, 2, 3}
	// increment
	res := Map(coll, func(n int) int {
		return n + 1
	})
	fmt.Println(res)

	//stringify
	res2 := Map(coll, func(n int) string {
		return fmt.Sprintf("str_%d", n)
	})
	fmt.Println(res2)

}

func Map[T, R any](coll []T, f func(T) R) []R {
	res := make([]R, len(coll))
	for i, el := range coll {
		res[i] = f(el)
	}
	return res
}