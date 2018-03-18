package main

func maxSubArray(array []int) int {
	C, G := array[0], array[0]

	for i := 1; i < len(array); i++ {
		temp := C + array[i]

		if temp > array[i] {
			C = temp
		} else {
			C = array[i]
		}

		if C > G {
			G = C
		}
	}

	return G
}

func main() {

}
