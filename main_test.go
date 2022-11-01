package main



import "testing"

func TestBinarySearch(t *testing.T){

	expect := 9
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	
    res := BinarySearch(expect, arr)
   

    if expect != arr[res] {
        t.Errorf("Fail")
    }
}

func TestBinarySearch2(t *testing.T){

	expect := 3
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	
    res := BinarySearch(expect, arr)
   

    if expect != arr[res] {
        t.Errorf("Fail")
    }
}
