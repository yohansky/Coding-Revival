package main

import "fmt"

func main() {
	price := 150000
	quantity := 1
	discount := 0
	total := price * quantity

	fmt.Println("Jumlah Produk : ",quantity)
	fmt.Println("Total Harga : ",total)

	if total >= 200000 {
		discount = total * 10 / 100
		fmt.Println("Selamat, Anda mendapatkan diskon 10%")
		fmt.Println("Nominal diskon:", discount)
	} else {
		fmt.Println("Maaf, Anda belum mendapatkan diskon")
	}

	finalTotal := total - discount
	fmt.Println("Total setelah diskon:", finalTotal)
}