package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i:=0 ; i< 10 ; i++ {
		size := 999999
		rand.Seed(time.Now().UnixNano())
		expect := rand.Intn(size)
		
		//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		arr := []int{}
		for i := 0; i< size; i++{
			arr = append(arr, i)
		}
		start := time.Now();
		res := BinarySearch(expect, arr)
		fmt.Printf("arr[%v] = %v used %v \n", res, arr[res], time.Since(start))
	}
	
}

func BinarySearch(expect int, arr []int) int {
	low := 0
	high := len(arr) - 1	

	for low <= high {	
		mid := (low + high) / 2	
		if expect == arr[mid] {
			return mid
		} 
		if expect > arr[mid]{
			low = mid + 1	
		}else{
			high = mid - 1
		}
	}
	return 0

}

