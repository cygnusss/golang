package main

import "fmt"

/*
* Input a list of array  [[1, 2, 3], [1], [1, 2]] return the list of array,
* each array is a combination of one element in each array.
* [[1, 1, 1], [1, 1, 2], [2, 1, 1], [2, 1, 2], [3, 1, 1], [3, 1, 2]]
* Followup: each array in the input list is an iterator, which can only be looped once.
 */

/*
var nums []int
var combinations []int
FOR AofA <- N+1
FOR AofN <- (K+1)*N=NK+N

RULES:
	i != indecies map
	if len==3 => append
indecies = map[int]interface
[ 1, 1, 2, 3]
	BEFORE EACH recursion update indecies
	1; s = 1, i<len
		1, 1; s = 2, i<len
			1, 1, 2; s = 3, i<len !STOP!NEXT
		 	1, 1, 3; s = 3, i==len !STOP!BREAK
		1, 2; s = 2, i<len
		 	1, 2, 1; s = 3, i<len !STOP!NEXT
		 	1, 2, 3; s = 3, i==len !STOP!BREAK
		1, 3; s = 2, i<len
		 	1, 3, 1; s = 3, i<len !STOP!NEXT
			1, 3, 2; s = 3, i==len !STOP!BREAK
... ... ...
*/

func FindCombination(nums []int) [][]int {
	var combs [][]int
	var recurse func(ccomb []int, inds map[int]bool)

	// fmt.Println("Enters findCombinations")

	recurse = func(ccomb []int, inds map[int]bool) {
		// fmt.Println("Enters a loop")
		// fmt.Println("Current combination is:", ccomb)
		// if current combination has 3 items stop
		if len(ccomb) == 3 {
			// fmt.Println("Current combination before appending is:", ccomb)

			combs = append(combs, append([]int(nil), ccomb...))
			// fmt.Println("All current combs are:", combs)
			// fmt.Println("Updated combinations:", combs)
			return
		}

		for i := 0; i < len(nums); i++ {
			// if index exists in map => continue
			if _, ok := inds[i]; ok {
				continue
			}
			// current number
			cnum := nums[i]
			// add index to map
			inds[i] = true

			// fmt.Println("Map at current stage:", inds)

			// add current number to current combination
			ccomb = append(ccomb, cnum)
			// fmt.Println("Current combination of numbers before recursion is:", ccomb, "\n")
			recurse(ccomb, inds)
			// fmt.Println("Current combination of numbers after recursion is:", ccomb)
			// remove current number from current combination
			ccomb = ccomb[:len(ccomb)-1]
			// fmt.Println("Current combination of numbers after popping is:", ccomb, "\n")
			// remove index from map
			delete(inds, i)
		}
		return
	}

	recurse([]int{}, map[int]bool{})

	return combs
}

func main() {
	combinations := FindCombination([]int{1, 2, 3, 4})

	fmt.Printf("Combinations are: %v\n", combinations)
}
