package main

import (
	"fmt"
)

type keahlian struct {
	nama, jenis string
}


type Karier struct {
	Nama               string
	Industri           string
	KeahlianDibutuhkan []string
	GajiRataRata       int
}


var daftarKeahlian []keahlian


var daftarKarier = []Karier{
	{"Data Scientist", "Teknologi", []string{"statistik", "machine learning", "python"}, 15000000},
	{"UI/UX Designer", "Teknologi", []string{"desain", "figma", "user research"}, 12000000},
	{"Digital Marketer", "Pemasaran", []string{"seo", "content creation", "sosial media"}, 10000000},
	{"Akuntan", "Keuangan", []string{"akuntansi", "excel", "laporan keuangan"}, 9000000},
	{"Dokter", "Kesehatan", []string{"biologi", "komunikasi pasien"}, 20000000},
	{"Guru", "Pendidikan", []string{"komunikasi", "manajemen kelas"}, 8000000},
	{"Pengacara", "Hukum", []string{"analisis kasus", "negosiasi"}, 18000000},
	{"Wirausahawan", "Bisnis", []string{"manajemen", "kreativitas"}, 10000000},
	{"Musisi", "Seni", []string{"instrumen musik", "kreativitas"}, 7000000},
	{"Penyanyi", "Hiburan", []string{"vokal", "penampilan panggung"}, 8000000},
	{"Stand-up Comedian", "Hiburan", []string{"public speaking", "humor"}, 6000000},
	{"Pedagang", "Perdagangan", []string{"negosiasi", "manajemen stok"}, 5000000},
	{"Chef", "Kuliner", []string{"memasak", "kreativitas resep"}, 9000000},
	{"Fotografer", "Kreatif", []string{"fotografi", "editing foto"}, 7000000},
	{"Content Creator", "Media Sosial", []string{"video editing", "kreativitas"}, 10000000},
	{"Psikolog", "Kesehatan Mental", []string{"empati", "analisis psikologi"}, 12000000},
	{"Aktor/Aktris", "Hiburan", []string{"akting", "ekspresi emosi"}, 10000000},
	{"Pelukis", "Seni", []string{"melukis", "kreativitas visual"}, 6000000},
	{"Software Engineer", "Teknologi", []string{"go", "java", "git", "sql"}, 14000000},
	{"Network Engineer", "Teknologi", []string{"jaringan", "cisco", "tcp/ip"}, 11000000},
}


func showMenu() {
	fmt.Println("\n============ MENU UTAMA ============ ")
	fmt.Println("1. Tambah minat / keahlian")
	fmt.Println("2. Lihat Daftar Keahlian")
	fmt.Println("3. Edit Data Keahlian")
	fmt.Println("4. Hapus Data Keahlian")
	fmt.Println("5. Rekomendasi Karier")
	fmt.Println("6. Pencarian Karier (Linear Search)")
	fmt.Println("7. Urutkan Karier Berdasarkan Gaji (Terendah ke Tertinggi)")
	fmt.Println("8. Filter Karier berdasarkan Industri (Exact Match)")
	fmt.Println("9. Tampilkan Karier Gaji Tertinggi") // Teks menu disesuaikan
	fmt.Println("10. Keluar")
	fmt.Print("Pilih menu (1-10): ")
}

func tambahKeahlian() {
	var namaInput, jenisInput string

	fmt.Print("Masukkan nama minat/keahlian: ")
	fmt.Scanln(&namaInput)

	fmt.Print("Masukkan jenis keahlian (contoh: teknis, interpersonal, kreatif): ")
	fmt.Scanln(&jenisInput)

	if namaInput == "" || jenisInput == "" {
		fmt.Println("Input tidak boleh kosong. Harap masukkan nama dan jenis keahlian.")
		return
	}

	dataBaru := keahlian{nama: namaInput, jenis: jenisInput}
	daftarKeahlian = append(daftarKeahlian, dataBaru)
	fmt.Println("Data minat/keahlian berhasil ditambahkan.")
}


func tampilKeahlian() {
	if len(daftarKeahlian) == 0 {
		fmt.Println("Belum ada data minat/keahlian yang dimasukkan.")
		return
	}
	fmt.Println("\n====== DAFTAR MINAT dan KEAHLIAN ======")
	for i, k := range daftarKeahlian {
		fmt.Printf("%d. Nama: %s | Jenis: %s\n", i+1, k.nama, k.jenis)
	}
}


func editKeahlian() {
	tampilKeahlian()
	if len(daftarKeahlian) == 0 {
		return
	}

	var index int
	fmt.Print("Masukkan nomor data yang ingin diubah: ")
	_, err := fmt.Scanln(&index)
	if err != nil || index < 1 || index > len(daftarKeahlian) {
		fmt.Println("Nomor tidak valid. Pastikan nomor sesuai dengan daftar.")
		return
	}

	var namaBaru, jenisBaru string
	fmt.Printf("Mengedit data ke-%d:\n", index)
	fmt.Print("Masukkan nama minat/keahlian baru: ")
	fmt.Scanln(&namaBaru)

	fmt.Print("Masukkan jenis keahlian baru: ")
	fmt.Scanln(&jenisBaru)

	if namaBaru == "" || jenisBaru == "" {
		fmt.Println("Input nama atau jenis baru tidak boleh kosong.")
		return
	}

	daftarKeahlian[index-1] = keahlian{nama: namaBaru, jenis: jenisBaru}
	fmt.Println("Data minat/keahlian berhasil diperbarui.")
}


func hapusKeahlian() {
	tampilKeahlian()
	if len(daftarKeahlian) == 0 {
		return
	}

	var index int
	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	_, err := fmt.Scanln(&index)
	if err != nil || index < 1 || index > len(daftarKeahlian) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	daftarKeahlian = append(daftarKeahlian[:index-1], daftarKeahlian[index:]...)
	fmt.Println("Data minat/keahlian berhasil dihapus.")
}


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
			for _, keahlianKarier := range karier.KeahlianDibutuhkan {
				if keahlianUser.nama == keahlianKarier {
					cocokLokal++
					break
				}
			}
		}

		if cocokLokal > 0 && len(karier.KeahlianDibutuhkan) > 0 {
			totalCocokGlobal++
			persen := (float64(cocokLokal) / float64(len(karier.KeahlianDibutuhkan))) * 100
			fmt.Printf("Karier: %-20s | Industri: %-15s | Gaji: Rp%d | Kecocokan: %.0f%%\n",
				karier.Nama, karier.Industri, karier.GajiRataRata, persen)
		}
	}

	if totalCocokGlobal == 0 {
		fmt.Println("Belum ada karier yang cocok dengan keahlian Anda saat ini.")
	}
}


func urutkanKarierBerdasarkanGaji() {
	karierUntukDiurut := make([]Karier, len(daftarKarier))
	copy(karierUntukDiurut, daftarKarier)

	n := len(karierUntukDiurut)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if karierUntukDiurut[j].GajiRataRata < karierUntukDiurut[minIndex].GajiRataRata {
				minIndex = j
			}
		}
		karierUntukDiurut[i], karierUntukDiurut[minIndex] = karierUntukDiurut[minIndex], karierUntukDiurut[i]
	}

	fmt.Println("\n=== DAFTAR KARIER BERDASARKAN GAJI (TERENDAH KE TERTINGGI) ===")
	for _, k := range karierUntukDiurut {
		fmt.Printf("Karier: %-20s | Industri: %-15s | Gaji: Rp%d\n", k.Nama, k.Industri, k.GajiRataRata)
	}
}


func linearSearchKarier() {
	var targetNamaKarier string
	fmt.Print("Masukkan nama karier yang ingin dicari (case-sensitive): ")
	fmt.Scanln(&targetNamaKarier)

	if targetNamaKarier == "" {
		fmt.Println("Nama karier tidak boleh kosong.")
		return
	}

	fmt.Println("\n--- HASIL PENCARIAN ---")
	found := false
	for _, k := range daftarKarier {
		if k.Nama == targetNamaKarier {
			fmt.Printf("Ditemukan: Karier: %-20s | Industri: %-15s | Gaji: Rp%d\n",
				k.Nama, k.Industri, k.GajiRataRata)
			found = true
		}
	}

	if !found {
		fmt.Println("Karier dengan nama '"+targetNamaKarier+"' tidak ditemukan.")
	}
}


func filterIndustri() {
	var targetIndustri string
	fmt.Print("Masukkan nama industri yang ingin difilter (exact match, case-sensitive): ")
	fmt.Scanln(&targetIndustri)

	if targetIndustri == "" {
		fmt.Println("Nama industri tidak boleh kosong.")
		return
	}

	fmt.Println("\n--- HASIL FILTER BERDASARKAN INDUSTRI:", targetIndustri, "---")
	found := false
	for _, k := range daftarKarier {
		if k.Industri == targetIndustri {
			fmt.Printf("Karier: %-20s | Gaji: Rp%d\n", k.Nama, k.GajiRataRata)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada karier yang ditemukan pada industri '"+targetIndustri+"'.")
	}
}


func tampilkanNilaiEkstrimGaji() {
	if len(daftarKarier) == 0 {
		fmt.Println("Daftar karier kosong, tidak ada gaji tertinggi untuk ditampilkan.")
		return
	}

	var karierGajiMax []Karier

	
	maxGaji := daftarKarier[0].GajiRataRata
	karierGajiMax = append(karierGajiMax, daftarKarier[0]) 

	
		karierSaatIni := daftarKarier[i]

		
		if karierSaatIni.GajiRataRata > maxGaji {
			maxGaji = karierSaatIni.GajiRataRata
			karierGajiMax = []Karier{karierSaatIni} 
		} else if karierSaatIni.GajiRataRata == maxGaji {
			karierGajiMax = append(karierGajiMax, karierSaatIni) 
		}
	}

	fmt.Println("\n========== KARIER DENGAN GAJI TERTINGGI ==========")
	fmt.Printf("\nKARIER DENGAN GAJI TERTINGGI (Rp%d):\n", maxGaji)
	for _, k := range karierGajiMax {
		fmt.Printf("- %s (Industri: %s)\n", k.Nama, k.Industri)
	}
}


func main() {
	var pilihan string
	for {
		showMenu()
		_, err := fmt.Scanln(&pilihan)
		if err != nil {
			fmt.Println("Input tidak valid, coba lagi.")
			continue
		}

		switch pilihan {
		case "1":
			tambahKeahlian()
		case "2":
			tampilKeahlian()
		case "3":
			editKeahlian()
		case "4":
			hapusKeahlian()
		case "5":
			rekomendasiKarier()
		case "6":
			linearSearchKarier()
		case "7":
			urutkanKarierBerdasarkanGaji()
		case "8":
			filterIndustri()
		case "9":
			tampilkanNilaiEkstrimGaji() 
		case "10":
			fmt.Println("Program selesai. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan masukkan angka antara 1-10.")
		}
	}
}
