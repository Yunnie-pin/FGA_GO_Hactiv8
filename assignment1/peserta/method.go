package peserta

func (p Peserta) Cetakan() string {
	if p.Id == 0 {
		return "Peserta tidak ditemukan"
	}
	return "Nama: " + p.Nama + "\nAlamat: " + p.Alamat + "\nPekerjaan: " + p.Pekerjaan + "\nAlasan: " + p.Alasan
}
