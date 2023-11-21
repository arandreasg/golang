package main

import (
	"fmt"
	"os"
	"sesi-baru/method"
	"sesi-baru/models"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage : main <nama>")
		return
	}

	nama := args[1]

	peserta := method.CustomPeserta{
		Peserta: models.Peserta{
			Nama: nama,
		},
	}

	peserta.FindPeserta()

}
