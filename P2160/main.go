package main

import "sort"

func main() {

}

func minimumSum(num int) int {
	arr := []int{
		num % 10,
		num / 10 % 10,
		num / 100 % 10,
		num / 1000,
	}
	sort.Ints(arr)
	return arr[0]*10 + arr[1]*10 + arr[2] + arr[3]
}
