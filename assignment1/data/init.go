package data

import "assignment1/peserta"

func InitData() []peserta.Peserta {
	return []peserta.Peserta{
		{
			Id:        1,
			Nama:      "Arifin Yunianta",
			Alamat:    "Klaten, Jawa Tengah",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Menambah Portofolio",
		},
		{
			Id:        2,
			Nama:      "Andi",
			Alamat:    "Jl. Jendral Sudirman No. 2",
			Pekerjaan: "PNS",
			Alasan:    "Mau belajar Go",
		},
	}
}
