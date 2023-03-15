package main

import (
	"fmt"
)

func main() {
	var (
		n, k int
	)

	// Penampung input gem
	gems := [][]int{}
	// Penampung hasil
	result := make([]int, n)

	fmt.Println("Masukkan total gem: ")
	fmt.Scanln(&n)

	fmt.Println("Masukkan Kapasitas Tas: ")
	fmt.Scanln(&k)

	// Mengumpulkan gem
	if n != 0 {
		for i := 1; i <= n; i++ {
			var (
				u, w int
			)
			// Ambil gem
			fmt.Println("Masukkan Gem: ")
			fmt.Scanln(&u)
			fmt.Println("Masukkan Berat Gem: ")
			fmt.Scanln(&w)

			gem := []int{u, w}
			gems = append(gems, gem)
		}
	} else {
		fmt.Println("Masukkan total gem!")
		return
	}

	// Proses pencarian gem
	weight := 0
	for i := 0; i < len(gems); i++ {

		weight += gems[i][1]

		if weight > k {
			break
		}

		result = append(result, gems[i][0])

	}

	// Output
	if len(result) == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println("hasil", result)
	}
}
