package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func hitungjarakArr(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0
	}

	spacing := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != spacing {
			return false, 0
		}
	}
	return true, spacing
}

func main() {
	var n int
	fmt.Println(" jumlah elemen:")
	fmt.Scan(&n)

	fmt.Println(":")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i] < 0 {
			return
		}
	}

	insertionSort(arr)
	hitungjarak, jarak := hitungjarakArr(arr)

	fmt.Println("Hasil Pengurutan:", arr)
	if hitungjarak {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
