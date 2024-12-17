package main

import "fmt"

func main() {
	var target, donasi, totalDonasi, donatur int

	fmt.Print("target donasi: ")
	fmt.Scan(&target)

	for totalDonasi < target {
		donatur++

		fmt.Printf("donasi dari donatur %d: ", donatur)
		fmt.Scan(&donasi)

		totalDonasi += donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donatur, donasi, totalDonasi)
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, donatur)
}
