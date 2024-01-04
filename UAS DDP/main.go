package main

import ( //ini import data
	"fmt"
	"bufio"
	"os"
	"strings"
)

// Struct Recepi untuk menyimpan detail resep
type Recepi struct {
	Nama string
	Bahan []Bahan
	Langkah []string
}

// Struct Bahan untuk menyimpan detail bahan
type Bahan struct {
	Nama string
	Satuan string
	Jumlah string
}

var resep []Recepi

func main(){
	loadData() // Memuat resep yang ada di file penyimpanan

	for {
		tampilkanMenu()
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahResep()
		case 2:
			cariResep()
		case 3:
			hapusResep()
		case 4:
			simpanData() // Menyimpan Resep sebelum keluar
			os.Exit(0)
		}
	}
}

func tampilkanMenu() {
	fmt.Println("\nProgram Resep Makanan")
	fmt.Println("=====================")
	fmt.Println("1. Tambah Resep")
	fmt.Println("2. Cari Resep")
	fmt.Println("3. Hapus Resep")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih Menu [1-4]: ")
}

func tambahResep(){
	clearScreen()

	var namaResep string
	var bahan []Bahan
	var langkah []string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nama Resep: ")
	scanner.Scan()
	namaResep = scanner.Text()

	// Pemeriksaan apakah nama resep tersebut sudah ada
	for resepSudahAda(namaResep) {
		fmt.Println("Resep dengan nama tersebut sudah ada. Silakan gunakan nama yang berbeda.")
		fmt.Print("Nama Resep: ")
		scanner.Scan()
		namaResep = scanner.Text()
	}
	
	// Memasukkan bahan-bahan
	fmt.Println("Masukan bahan-bahan (tekan underscore untuk menghentikan penambahan bahan):")
	bahanIndex := 1
	for {
		fmt.Printf("Bahan ke - %d: ", bahanIndex)
		scanner.Scan()
		namaBahan := scanner.Text()
		if namaBahan == "_"{
			break
		}

		fmt.Printf("Satuan bahan ke - %d: ", bahanIndex)
		scanner.Scan()
		satuan := scanner.Text()

		fmt.Printf("Jumlah bahan ke - %d: ", bahanIndex)
		scanner.Scan()
		jumlah := scanner.Text()

		bahanBaru := Bahan{Nama: namaBahan, Satuan: satuan, Jumlah: jumlah}
		bahan = append(bahan, bahanBaru)

		bahanIndex++
	}

	// Memasukkan langkah-langkah
	fmt.Println("Masukan langkah-langkah (tekan underscore untuk menghentikan penambahan langkah):")
	langkahIndex := 1
	for {
		fmt.Printf("Langkah ke - %d: ", langkahIndex)
		scanner.Scan()
		langkahBaru := scanner.Text()
		if langkahBaru == "_" {
			break
		}
		langkah = append(langkah, langkahBaru)
		langkahIndex++
	}

	resepBaru := Recepi{Nama: namaResep, Bahan: bahan, Langkah: langkah}
	resep = append(resep, resepBaru)
	fmt.Println("Resep berhasil ditambahkan!")
	simpanData()

}


// Fungsi untuk memeriksa apakah resep dengan nama yang sama sudah ada atau belum
func resepSudahAda(namaResep string) bool {
	for _, r := range resep {
		if strings.EqualFold(r.Nama, namaResep) {
			return true
		}
	}
	return false
}

func cariResep() {
	clearScreen()

	var cariNamaResep string

	if len(resep) == 0 {
		fmt.Println("Belum ada resep yang ditambahkan")
		return
	}

	listResep()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkkan nama resep: ")
	scanner.Scan()
	cariNamaResep = scanner.Text()

	resepDitemukan := false

	for _, resep := range resep {
		if strings.Contains(strings.ToLower(resep.Nama),strings.ToLower(cariNamaResep)) {
			resepDitemukan = true
			fmt.Println("\nResep ditemukan!")
			fmt.Println("Nama Resep:", resep.Nama)
			fmt.Println("Bahan-bahan:")
			for i, bahan := range resep.Bahan {
				fmt.Printf("%d. %s - %s %s\n", i+1, bahan.Nama, bahan.Jumlah, bahan.Satuan)
			}
			fmt.Println("Langkah-langkah:")
			for i, langkah := range resep.Langkah {
				fmt.Printf("%d. %s\n", i+1, langkah)
			}
		}
	}

	if !resepDitemukan {
		fmt.Println("Resep tidak ditemukan.")
	}
}

func listResep() {
	fmt.Println("Resep-resep:")
	for i, resep := range resep {
		fmt.Printf("%d. %s\n", i+1, resep.Nama)
	}
}

func hapusResep() {
	clearScreen()

	fmt.Println("Pilih Resep yang akan dihapus:")
	listResep()

	var pilihanHapus int
	fmt.Print("Pilih Resep [1-", len(resep), "]: ")
    fmt.Scanln(&pilihanHapus)
	if pilihanHapus < 1 || pilihanHapus > len(resep) {
        fmt.Println("Pilihan tidak valid")
    }
	resep = append(resep[:pilihanHapus-1], resep[pilihanHapus:]...)
    fmt.Println("Resep berhasil dihapus!")
}

func simpanData() {
	file, err := os.Create("resep.txt")
	if err != nil {
		fmt.Println("Error menyimpan data:", err)
		return
	}
	defer file.Close()

	for _, resep := range resep {
		// Menyimpan Nama Resep
		fmt.Fprintf(file, "Nama Resep: %s\n", resep.Nama)
		
		// Menyimpan Nama Bahan
		fmt.Fprint(file, "Nama Bahan: ")
		for i, bahan := range resep.Bahan {
			fmt.Fprintf(file, "%s", bahan.Nama)
			if i != len(resep.Bahan)-1 {
				fmt.Fprint(file, ". ")
			}
		}
		fmt.Fprintln(file) //tambah baris baru di file

		// Menyimpan Satuan Bahan
		fmt.Fprint(file, "Satuan Bahan: ")
		for i, bahan := range resep.Bahan {
			fmt.Fprintf(file, "%s", bahan.Satuan)
			if i != len(resep.Bahan)-1 {
				fmt.Fprint(file, ". ")
			}
		}
		fmt.Fprintln(file)

		// Menyimpan Jumlah Bahan
		fmt.Fprint(file, "Jumlah Bahan: ")
		for i, bahan := range resep.Bahan {
			fmt.Fprintf(file, "%s", bahan.Jumlah)
			if i != len(resep.Bahan)-1 {
				fmt.Fprint(file, ". ")
			}
		}
		fmt.Fprintln(file)

		// Menyimpan Langkah-Langkah
		fmt.Fprintf(file, "Langkah-Langkah: %s\n\n", strings.Join(resep.Langkah, ". "))
	}

	fmt.Println("Data berhasil disimpan!")
}

func loadData() {
	file, err := os.Open("resep.txt")
	if err != nil {
		fmt.Println("File tidak ditemukan. Memulai dengan daftar resep kosong.")
        return
	}
	defer file.Close()

	var resepTemp Recepi
	var bahanTemp Bahan
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Nama Resep:") { //pengecekan
			resepTemp.Nama = strings.TrimSpace(strings.TrimPrefix(line, "Nama Resep:")) // hapus spasi dan hapus awalan disimpan
		} else if strings.HasPrefix(line, "Nama Bahan:") {
            // Membaca Nama Bahan
            bahanLine := strings.TrimPrefix(line, "Nama Bahan: ")
            namaBahan := strings.Split(bahanLine, ". ")
            for _, nama := range namaBahan {
                bahanTemp = Bahan{Nama: nama, Satuan: "", Jumlah: ""}
                resepTemp.Bahan = append(resepTemp.Bahan, bahanTemp)
            }
		} else if strings.HasPrefix(line, "Satuan Bahan:") {
            // Membaca Satuan Bahan
            satuanLine := strings.TrimPrefix(line, "Satuan Bahan: ")
            satuanBahan := strings.Split(satuanLine, ". ")
            for i, satuan := range satuanBahan {
                // Memastikan panjang array resepTemp.Bahan sesuai dengan satuanBahan
                if i < len(resepTemp.Bahan) {
                    resepTemp.Bahan[i].Satuan = satuan
                }
            }
		} else if strings.HasPrefix(line, "Jumlah Bahan:") {
			// Membaca Jumlah Bahan
			jumlahLine := strings.TrimPrefix(line, "Jumlah Bahan: ")
			jumlahBahan := strings.Split(jumlahLine, ". ")
			for i, jumlah := range jumlahBahan {
                // Memastikan panjang array resepTemp.Bahan sesuai dengan jumlahBahan
                if i < len(resepTemp.Bahan) {
                    resepTemp.Bahan[i].Jumlah = jumlah
                }
            }
		 }else if strings.HasPrefix(line, "Langkah-Langkah:") {
            // Membaca Langkah-Langkah
            langkahLine := strings.TrimSpace(strings.TrimPrefix(line, "Langkah-Langkah:"))
            resepTemp.Langkah = strings.Split(langkahLine, ". ")

            // Menambahkan resep ke dalam slice resep
            resep = append(resep, resepTemp)

            // Me-reset resepTemp untuk resep berikutnya
            resepTemp = Recepi{}
        }
	}
	if err := scanner.Err(); err != nil {
        fmt.Println("Error membaca file:", err)
        return
    }
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
	// /033 untuk bawa kursor ke baris 1
	// bersihin layar konsul
}