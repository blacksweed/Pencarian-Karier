package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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

func showMenu() {
	fmt.Println("\n============ MENU UTAMA ============ ")
	fmt.Println("1. Tambah minat / keahlian")
	fmt.Println("2. Lihat Daftar")
	fmt.Println("3. Edit Data")
	fmt.Println("4. Hapus Data")
	fmt.Println("5. Rekomendasi Karier")
	fmt.Println("6. Pencarian Karier (Binary Search)")
	fmt.Println("7. Urutkan Karier Berdasarkan Gaji")
	fmt.Println("8. Filter Karier berdasarkan Industri")
	fmt.Println("9. Keluar")
	fmt.Print("Pilih menu (1-9): ")
}

func tambahKeahlian() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan minat: ")
	namaInput, _ := reader.ReadString('\n')
	namaInput = strings.ToLower(strings.TrimSpace(namaInput))

	fmt.Print("Masukkan jenis keahlian: ")
	jenisInput, _ := reader.ReadString('\n')
	jenisInput = strings.ToLower(strings.TrimSpace(jenisInput))

	if namaInput == "" || jenisInput == "" {
		fmt.Println("Input tidak boleh kosong.")
		return
	}

	dataBaru := keahlian{nama: namaInput, jenis: jenisInput}
	daftarKeahlian = append(daftarKeahlian, dataBaru)
	fmt.Println("Data berhasil ditambahkan.")
}

func tampilKeahlian() {
	if len(daftarKeahlian) == 0 {
		fmt.Println("Belum ada data minat/keahlian.")
		return
	}
	fmt.Println("====== DAFTAR MINAT dan KEAHLIAN ======")
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
	fmt.Scanln(&index)
	if index < 1 || index > len(daftarKeahlian) {
		fmt.Println("Index tidak valid.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama baru: ")
	namaBaru, _ := reader.ReadString('\n')
	namaBaru = strings.ToLower(strings.TrimSpace(namaBaru))

	fmt.Print("Masukkan jenis keahlian baru: ")
	jenisBaru, _ := reader.ReadString('\n')
	jenisBaru = strings.ToLower(strings.TrimSpace(jenisBaru))

	if namaBaru == "" || jenisBaru == "" {
		fmt.Println("Input tidak boleh kosong.")
		return
	}

	daftarKeahlian[index-1] = keahlian{nama: namaBaru, jenis: jenisBaru}
	fmt.Println("Data berhasil diperbarui.")
}

func hapusKeahlian() {
	tampilKeahlian()
	if len(daftarKeahlian) == 0 {
		return
	}
	var index int
	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scanln(&index)
	if index < 1 || index > len(daftarKeahlian) {
		fmt.Println("Index tidak valid.")
		return
	}
	daftarKeahlian = append(daftarKeahlian[:index-1], daftarKeahlian[index:]...)
	fmt.Println("Data berhasil dihapus.")
}

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

func rekomendasiKarier() {
	if len(daftarKeahlian) == 0 {
		fmt.Println("Belum ada data keahlian.")
		return
	}

	fmt.Println("\nRekomendasi Karier Berdasarkan Kecocokan:")

	totalCocok := 0
	for _, karier := range daftarKarier {
		cocok := 0
		for _, keahlianUser := range daftarKeahlian {
			for _, keahlianKarier := range karier.KeahlianDibutuhkan {
				if strings.ToLower(keahlianUser.nama) == strings.ToLower(keahlianKarier) {
					cocok++
					break
				}
			}
		}
		if cocok > 0 {
			totalCocok++
			persen := (float64(cocok) / float64(len(karier.KeahlianDibutuhkan))) * 100
			fmt.Printf("Karier: %-20s | Industri: %-15s | Gaji: Rp%d | Kecocokan: %.0f%%\n",
				karier.Nama, karier.Industri, karier.GajiRataRata, persen)
		}
	}
	if totalCocok == 0 {
		fmt.Println("Belum ada karier yang cocok dengan keahlian Anda.")
	}
}

func urutkanKarierBerdasarkanGaji() {
	for i := 0; i < len(daftarKarier)-1; i++ {
		min := i
		for j := i + 1; j < len(daftarKarier); j++ {
			if daftarKarier[j].GajiRataRata < daftarKarier[min].GajiRataRata {
				min = j
			}
		}
		daftarKarier[i], daftarKarier[min] = daftarKarier[min], daftarKarier[i]
	}
	fmt.Println("Karier berhasil diurutkan berdasarkan gaji.\n")

	fmt.Println("=== Daftar Karier Berdasarkan Gaji Terendah ke Tertinggi ===")
	for _, k := range daftarKarier {
		fmt.Printf("Karier: %-20s | Industri: %-15s | Gaji: Rp%d\n", k.Nama, k.Industri, k.GajiRataRata)
	}
}


func binarySearchKarier() {
	sort.Slice(daftarKarier, func(i, j int) bool {
		return strings.ToLower(daftarKarier[i].Nama) < strings.ToLower(daftarKarier[j].Nama)
	})

	var target string
	fmt.Print("Masukkan nama karier: ")
	fmt.Scanln(&target)

	low, high := 0, len(daftarKarier)-1
	for low <= high {
		mid := (low + high) / 2
		comp := strings.ToLower(daftarKarier[mid].Nama)
		if comp == strings.ToLower(target) {
			k := daftarKarier[mid]
			fmt.Printf("Ditemukan: %s | Industri: %s | Gaji: Rp%d\n", k.Nama, k.Industri, k.GajiRataRata)
			return
		} else if comp < strings.ToLower(target) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Karier tidak ditemukan.")
}


func bacaDariFile() {
	file, err := os.Open("keahlian.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			daftarKeahlian = append(daftarKeahlian, keahlian{strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])})
		}
	}
}

func filterIndustri() {
	fmt.Print("Masukkan nama industri (atau bagian dari industri): ")
	var industri string
	fmt.Scanln(&industri)
	found := false
	for _, k := range daftarKarier {
		if strings.Contains(strings.ToLower(k.Industri), strings.ToLower(industri)) {
			fmt.Printf("Karier: %-20s | Gaji: Rp%d\n", k.Nama, k.GajiRataRata)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada karier pada industri tersebut.")
	}
}

func main() {
	bacaDariFile()
	for {
		showMenu()
		var pilihan string
		fmt.Scanln(&pilihan)

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
			binarySearchKarier()
		case "7":
			urutkanKarierBerdasarkanGaji()
		case "8":
			filterIndustri()
		case "9":
			fmt.Println("Program selesai. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
