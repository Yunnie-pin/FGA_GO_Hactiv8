package peserta

func CariPeserta(peserta []Peserta, absensi int) Peserta {
	for _, p := range peserta {
		if p.Id == absensi {
			return p
		}
	}
	return Peserta{}
}
