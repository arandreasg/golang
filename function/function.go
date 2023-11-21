package function

import "sesi-baru/models"

func DataPeserta() []models.Peserta {
	pesertaList := []models.Peserta{
		{
			Nama:      "Gaveriel",
			Alamat:    "Jakarta",
			Pekerjaan: "Developer",
			Alasan:    "Tidak Ada",
		},
		{
			Nama:      "Sean",
			Alamat:    "Pematangsiantar",
			Pekerjaan: "UI/UX",
			Alasan:    "Ingin belajar Golang",
		},
		{
			Nama:      "Girsang",
			Alamat:    "Medan",
			Pekerjaan: "Manager",
			Alasan:    "Ingin belajar Golang",
		},
	}
	return pesertaList
}
