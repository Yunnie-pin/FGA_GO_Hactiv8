package lib

import (
	"assignment1/helpers"
	"assignment1/peserta"
	"fmt"
	"os"
)

func Command(data []peserta.Peserta) {
	arg_len := len(os.Args[1:])

	switch arg_len {
	case 0:
		fmt.Println("==============================================")
		fmt.Printf("Jumlah Peserta yang terdaftar : %d\n", len(data))
		fmt.Println("Tidak ada argument yang diberikan")
		fmt.Println("Gunakan command `go run main.go [nomer peserta]`")
		fmt.Println("==============================================")
	case 1:
		// mendapatkan argument dari command line dan harus int
		var absensi = helpers.HandleArgs(os.Args[1:])

		var q = peserta.CariPeserta(data, absensi)
		fmt.Println("||==============================================||")
		fmt.Println(q.Cetakan())
		fmt.Println("||==============================================||")
	default:
		fmt.Println("Hanya menerima 1 argument")
	}
}
