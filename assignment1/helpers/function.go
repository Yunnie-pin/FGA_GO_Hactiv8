package helpers

import (
	"fmt"
	"os"
	"strconv"
)

func HandleArgs(args []string) int {
	arg := os.Args[1]
	// ubah argumen dari string ke int
	var absensi, err = strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Argument harus berupa angka")
		return 0
	}

	return absensi
}
