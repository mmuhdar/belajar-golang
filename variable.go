package main

import "fmt"

func main() {
	var nama string
	nama = "Mat"

	// go dapat mendeteksi otomatis type data dalam variabel
	var fullName = "Muhammad Muhdar"
	var age = 25

	// sorthand variable menggunakan ":="
	alamat := "Jakarta"
	// note: penggunaan ":=" hanya untuk deklarasi saja, untuk re assign nilai tidak perlu lagi

	// mutliple variable
	var (
		firstname = "Muhammad"
		lastname = "Muhdar"
	)

	fmt.Println(nama)
	fmt.Println(fullName)
	fmt.Println(age)
	fmt.Println(alamat)
	fmt.Println(firstname + " " + lastname)
}