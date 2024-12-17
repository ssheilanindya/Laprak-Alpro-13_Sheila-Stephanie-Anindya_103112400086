package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("harus bilangan bulat positif.")
		return
	}

	jumlahDigit := 0
	for n > 0 {
		jumlahDigit++
		n = n / 10
	}

	fmt.Println(jumlahDigit)
}
