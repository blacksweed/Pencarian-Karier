func rekomendasiKarier() {
	if len(daftarKeahlian) == 0 {
		fmt.Println("Belum ada data keahlian yang Anda masukkan. Silakan tambahkan dulu.")
		return
	}

	fmt.Println("\n====== REKOMENDASI KARIER BERDASARKAN KECOCOKAN ======")
	fmt.Println("(Perbandingan bersifat case-sensitive)")

	totalCocokGlobal := 0
	for _, karier := range daftarKarier {
		cocokLokal := 0
		for _, keahlianUser := range daftarKeahlian {
			// Iterasi melalui setiap keahlian yang dibutuhkan oleh karier
			for _, keahlianKarier := range karier.KeahlianDibutuhkan {
				// Perbandingan langsung (case-sensitive) karena tidak ada strings.ToLower()
				if keahlianUser.nama == keahlianKarier {
					cocokLokal++
					break // Jika satu keahlian user cocok, lanjut ke keahlian user berikutnya
				}
			}
		}

		if cocokLokal > 0 && len(karier.KeahlianDibutuhkan) > 0 {
			totalCocokGlobal++
			// Hitung persentase kecocokan
			persen := (float64(cocokLokal) / float64(len(karier.KeahlianDibutuhkan))) * 100
			fmt.Printf("Karier: %-20s | Industri: %-15s | Gaji: Rp%d | Kecocokan: %.0f%%\n",
				karier.Nama, karier.Industri, karier.GajiRataRata, persen)
		}
	}

	if totalCocokGlobal == 0 {
		fmt.Println("Belum ada karier yang cocok dengan keahlian Anda saat ini.")
	}
}
