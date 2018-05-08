package main

import "fmt"

func FindPermutations(arrays [][]int) [][]int {
	var recurse func([]int, int)
	combs := [][]int{}

	recurse = func(ccomb []int, indx int) {
		if indx == len(arrays) {
			duplicate := append([]int(nil), ccomb...)
			combs = append(combs, duplicate)
			return
		}

		fmt.Println("Current slice:", arrays[indx])

		for i := 0; i < len(arrays[indx]); i++ {
			cnum := arrays[indx][i] // Current number in current array
			ccomb = append(ccomb, cnum)
			recurse(ccomb, indx+1)
			ccomb = append(ccomb[:len(ccomb)-1])
		}

	}

	recurse([]int{}, 0)

	return combs
}

func main() {
	data := [][]int{[]int{1, 2, 3}, []int{3, 8}, []int{3}}

	permutations := FindPermutations(data)

	fmt.Printf("Permutations are: %v\n", permutations)
}
