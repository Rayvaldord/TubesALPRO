package main
import "fmt"

type PaketMCU struct {
	IDPaketMCU string
	NamaPaket  string
	HargaPaket int
}

type Pasien struct {
	IDPasien    string
	NamaLengkap string
	JenisKelamin string
	Alamat       string
	NoKontak     string
	TanggalMCU   string
	PaketDipilih string
}

type RekapHasilMCU struct {
	IDRekap    string
	IDPasien   string
	TanggalMCU string
	HasilMCU   string
}

const (
	MenuUtama         = "Pilih Menu: 1. Kelola Paket MCU 2. Kelola Pasien 3. Rekap Hasil MCU 4. Laporan Pemasukan 5. Keluar"
	MenuKelolaPaket   = "1. Tambah Paket MCU 2. Edit Paket MCU 3. Hapus Paket MCU 4. Tampilkan Daftar Paket MCU"
	MenuKelolaPasien  = "1. Tambah Pasien 2. Edit Pasien 3. Hapus Pasien 4. Tampilkan Daftar Pasien 5. Cari Pasien"
	MenuRekapHasil    = "1. Tambah Hasil MCU 2. Edit Hasil MCU 3. Hapus Hasil MCU 4. Tampilkan Rekap Hasil MCU"
	MenuLaporan       = "1. Laporan Harian 2. Laporan Bulanan 3. Laporan Tahunan"
)

var (
	paketMCUList  [100]PaketMCU
	pasienList    [100]Pasien
	rekapHasilList [100]RekapHasilMCU
	nPaketMCU     int
	nPasien       int
	nRekapHasil   int
)

func main() {
	for {
		fmt.Println(MenuUtama)
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			kelolaPaketMCU()
		case 2:
			kelolaPasien()
		case 3:
			rekapHasilMCU()
		case 4:
			tampilkanLaporan()
		case 5:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func kelolaPaketMCU() {
	fmt.Println(MenuKelolaPaket)
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tambahPaketMCU()
	case 2:
		editPaketMCU()
	case 3:
		hapusPaketMCU()
	case 4:
		tampilkanDaftarPaket()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func tambahPaketMCU() {
	if nPaketMCU >= 100 {
		fmt.Println("Kapasitas data paket MCU penuh.")
		return
	}
	var paket PaketMCU
	fmt.Println("Masukkan ID Paket MCU:")
	fmt.Scanln(&paket.IDPaketMCU)
	fmt.Println("Masukkan Nama Paket MCU:")
	fmt.Scanln(&paket.NamaPaket)
	fmt.Println("Masukkan Harga Paket MCU:")
	fmt.Scanln(&paket.HargaPaket)
	paketMCUList[nPaketMCU] = paket
	nPaketMCU++
	fmt.Println("Paket MCU berhasil ditambahkan.")
}

func editPaketMCU() {
	var id string
	fmt.Println("Masukkan ID Paket MCU yang akan diedit:")
	fmt.Scanln(&id)

	for i := 0; i < nPaketMCU; i++ {
		if paketMCUList[i].IDPaketMCU == id {
			fmt.Println("Masukkan Nama Paket MCU baru:")
			fmt.Scanln(&paketMCUList[i].NamaPaket)
			fmt.Println("Masukkan Harga Paket MCU baru:")
			fmt.Scanln(&paketMCUList[i].HargaPaket)
			fmt.Println("Paket MCU berhasil diubah.")
			return
		}
	}
	fmt.Println("Paket MCU tidak ditemukan.")
}

func hapusPaketMCU() {
	var id string
	fmt.Println("Masukkan ID Paket MCU yang akan dihapus:")
	fmt.Scanln(&id)

	for i := 0; i < nPaketMCU; i++ {
		if paketMCUList[i].IDPaketMCU == id {
			for j := i; j < nPaketMCU-1; j++ {
				paketMCUList[j] = paketMCUList[j+1]
			}
			nPaketMCU--
			fmt.Println("Paket MCU berhasil dihapus.")
			return
		}
	}
	fmt.Println("Paket MCU tidak ditemukan.")
}

func tampilkanDaftarPaket() {
	if nPaketMCU == 0 {
		fmt.Println("Belum ada paket MCU yang terdaftar.")
		return
	}

	fmt.Println("Daftar Paket MCU:")
	for i := 0; i < nPaketMCU; i++ {
		paket := paketMCUList[i]
		fmt.Printf("ID: %s, Nama: %s, Harga: %d\n", paket.IDPaketMCU, paket.NamaPaket, paket.HargaPaket)
	}
}

func kelolaPasien() {
	fmt.Println(MenuKelolaPasien)
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tambahPasien()
	case 2:
		editPasien()
	case 3:
		hapusPasien()
	case 4:
		tampilkanDaftarPasien()
	case 5:
		cariPasien()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func tambahPasien() {
	if nPasien >= 100 {
		fmt.Println("Jumlah maksimal pasien tercapai.")
		return
	}

	// Masukkan Nama Pasien
	fmt.Println("Masukkan Nama Pasien:")
	fmt.Scanln(&pasienList[nPasien].NamaLengkap)

	// Masukkan ID Pasien
	fmt.Println("Masukkan ID Pasien:")
	fmt.Scanln(&pasienList[nPasien].IDPasien)

	// Masukkan Jenis Kelamin
	fmt.Println("Masukkan Jenis Kelamin (L/P):")
	fmt.Scanln(&pasienList[nPasien].JenisKelamin)

	// Masukkan Alamat
	fmt.Println("Masukkan Alamat:")
	fmt.Scanln(&pasienList[nPasien].Alamat)

	// Masukkan NoKontak
	fmt.Println("Masukkan No. Kontak:")
	fmt.Scanln(&pasienList[nPasien].NoKontak)

	// Masukkan Tanggal MCU (dengan validasi format tanggal)
	fmt.Println("Masukkan Tanggal MCU (format: YYYY-MM-DD):")
	var tanggal string
	for {
		fmt.Scanln(&tanggal)
		// Cek apakah tanggal memiliki format yang valid, yaitu karakter ke-5 dan ke-8 adalah '-' dan panjang 10 karakter
		if tanggal[4] == '-' && tanggal[7] == '-' && tanggal[0] >= '0' && tanggal[1] >= '0' && tanggal[2] >= '0' && tanggal[3] >= '0' &&
			tanggal[5] >= '0' && tanggal[6] >= '0' && tanggal[8] >= '0' && tanggal[9] >= '0' {
			pasienList[nPasien].TanggalMCU = tanggal
			break
		} else {
			fmt.Println("Format tanggal tidak valid. Masukkan dalam format YYYY-MM-DD:")
		}
	}

	// Pilih Paket MCU
	fmt.Println("Pilih Paket MCU:")
	for i := 0; i < nPaketMCU; i++ {
		fmt.Printf("%d. %s (Rp %d)\n", i+1, paketMCUList[i].NamaPaket, paketMCUList[i].HargaPaket)
	}
	var pilihan int
	fmt.Scanln(&pilihan)
	if pilihan > 0 && pilihan <= nPaketMCU {
		pasienList[nPasien].PaketDipilih = paketMCUList[pilihan-1].NamaPaket
		fmt.Println("Paket berhasil dipilih.")
	} else {
		fmt.Println("Pilihan tidak valid. Data pasien tidak disimpan.")
		return
	}

	// Menambah pasien ke dalam daftar
	nPasien++
	fmt.Println("Pasien berhasil ditambahkan.")
}


func editPasien() {
	var id string
	fmt.Println("Masukkan ID Pasien yang akan diedit:")
	fmt.Scanln(&id)

	for i := 0; i < nPasien; i++ {
		if pasienList[i].IDPasien == id {
			fmt.Println("Masukkan Nama Pasien baru:")
			fmt.Scanln(&pasienList[i].NamaLengkap)
			fmt.Println("Masukkan Jenis Kelamin baru (L/P):")
			fmt.Scanln(&pasienList[i].JenisKelamin)
			fmt.Println("Masukkan Alamat baru:")
			fmt.Scanln(&pasienList[i].Alamat)
			fmt.Println("Masukkan No. Kontak baru:")
			fmt.Scanln(&pasienList[i].NoKontak)
			fmt.Println("Masukkan Tanggal MCU baru (YYYY-MM-DD):")
			var tanggal string
			for {
				fmt.Scanln(&tanggal)
				// Cek apakah tanggal memiliki format yang valid, yaitu karakter ke-5 dan ke-8 adalah '-' dan panjang 10 karakter
				if tanggal[4] == '-' && tanggal[7] == '-' && tanggal[0] >= '0' && tanggal[1] >= '0' && tanggal[2] >= '0' && tanggal[3] >= '0' &&
					tanggal[5] >= '0' && tanggal[6] >= '0' && tanggal[8] >= '0' && tanggal[9] >= '0' {
					pasienList[i].TanggalMCU = tanggal
					break
				} else {
					fmt.Println("Format tanggal tidak valid. Masukkan dalam format YYYY-MM-DD:")
				}
			}
			fmt.Println("Masukkan Paket MCU baru:")
				for i := 0; i < nPaketMCU; i++ {
					fmt.Printf("%d. %s (Rp %d)\n", i+1, paketMCUList[i].NamaPaket, paketMCUList[i].HargaPaket)
				}
				var pilihan int
				fmt.Scanln(&pilihan)
				if pilihan > 0 && pilihan <= nPaketMCU {
					pasienList[i].PaketDipilih = paketMCUList[pilihan-1].NamaPaket
					fmt.Println("Paket berhasil dipilih.")
				} else {
					fmt.Println("Pilihan tidak valid. Data pasien tidak disimpan.")
					return
				}
			fmt.Println("Data pasien berhasil diubah.")
			return
		}
	}
	fmt.Println("Pasien dengan ID tersebut tidak ditemukan.")
}

func hapusPasien() {
	var id string
	fmt.Println("Masukkan ID Pasien yang akan dihapus:")
	fmt.Scanln(&id)

	for i := 0; i < nPasien; i++ {
		if pasienList[i].IDPasien == id {
			for j := i; j < nPasien-1; j++ {
				pasienList[j] = pasienList[j+1]
			}
			nPasien--
			fmt.Println("Pasien berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pasien dengan ID tersebut tidak ditemukan.")
}

func tampilkanDaftarPasien() {
	if nPasien == 0 {
		fmt.Println("Belum ada data pasien yang terdaftar.")
		return
	}

	fmt.Println("Pilih kriteria pengurutan: ")
	fmt.Println("1. Waktu MCU (Ascending)")
	fmt.Println("2. Waktu MCU (Descending)")
	fmt.Println("3. Paket MCU (Ascending)")
	fmt.Println("4. Paket MCU (Descending)")
	var pilihan int
	fmt.Scanln(&pilihan)

	// Mengurutkan berdasarkan pilihan
	switch pilihan {
	case 1: // Waktu MCU Ascending
		for i := 0; i < nPasien-1; i++ {
			for j := 0; j < nPasien-i-1; j++ {
				if pasienList[j].TanggalMCU > pasienList[j+1].TanggalMCU {
					pasienList[j], pasienList[j+1] = pasienList[j+1], pasienList[j]
				}
			}
		}
	case 2: // Waktu MCU Descending
		for i := 0; i < nPasien-1; i++ {
			for j := 0; j < nPasien-i-1; j++ {
				if pasienList[j].TanggalMCU < pasienList[j+1].TanggalMCU {
					pasienList[j], pasienList[j+1] = pasienList[j+1], pasienList[j]
				}
			}
		}
	case 3: // Paket MCU Ascending
		for i := 0; i < nPasien-1; i++ {
			for j := 0; j < nPasien-i-1; j++ {
				if pasienList[j].PaketDipilih > pasienList[j+1].PaketDipilih {
					pasienList[j], pasienList[j+1] = pasienList[j+1], pasienList[j]
				}
			}
		}
	case 4: // Paket MCU Descending
		for i := 0; i < nPasien-1; i++ {
			for j := 0; j < nPasien-i-1; j++ {
				if pasienList[j].PaketDipilih < pasienList[j+1].PaketDipilih {
					pasienList[j], pasienList[j+1] = pasienList[j+1], pasienList[j]
				}
			}
		}
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	// Menampilkan daftar pasien setelah diurutkan
	fmt.Println("Daftar Pasien:")
	for i := 0; i < nPasien; i++ {
		pasien := pasienList[i]
		fmt.Printf("ID: %s, Nama: %s, Jenis Kelamin: %s, Alamat: %s, No. Kontak: %s, Tanggal MCU: %s, Paket Dipilih: %s\n",
			pasien.IDPasien, pasien.NamaLengkap, pasien.JenisKelamin, pasien.Alamat, pasien.NoKontak, pasien.TanggalMCU, pasien.PaketDipilih)
	}
}

func cariPasien() {
	fmt.Println("Pilih kriteria pencarian: 1. Nama Pasien 2. Paket MCU 3. Periode Waktu Tertentu")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		var nama string
		fmt.Println("Masukkan Nama Pasien:")
		fmt.Scanln(&nama)
		fmt.Println("Hasil Pencarian:")
		for i := 0; i < nPasien; i++ {
			if pasienList[i].NamaLengkap == nama {
				fmt.Printf("ID: %s, Nama: %s, Jenis Kelamin: %s, Alamat: %s, No. Kontak: %s, Tanggal MCU: %s, Paket Dipilih: %s\n",
					pasienList[i].IDPasien, pasienList[i].NamaLengkap, pasienList[i].JenisKelamin, pasienList[i].Alamat, pasienList[i].NoKontak, pasienList[i].TanggalMCU, pasienList[i].PaketDipilih)
			}
		}
	case 2:
		var paket int
		fmt.Println("Masukkan Paket MCU:")
		for i := 0; i < nPaketMCU; i++ {
			fmt.Printf("%d. %s (Rp %d)\n", i+1, paketMCUList[i].NamaPaket, paketMCUList[i].HargaPaket)
		}
		fmt.Scanln(&paket)
		if paket > 0 && paket <= nPaketMCU {
			paketDipilih := paketMCUList[paket-1].NamaPaket // Ambil ID paket yang dipilih
			fmt.Println("Hasil Pencarian:")
			found := false // Variabel untuk mengecek apakah ada pasien yang ditemukan
			for i := 0; i < nPasien; i++ {
				if pasienList[i].PaketDipilih == paketDipilih { // Cek apakah paket yang dipilih sama
					fmt.Printf("ID: %s, Nama: %s, Jenis Kelamin: %s, Alamat: %s, No. Kontak: %s, Tanggal MCU: %s, Paket Dipilih: %s\n",
						pasienList[i].IDPasien, pasienList[i].NamaLengkap, pasienList[i].JenisKelamin, pasienList[i].Alamat, pasienList[i].NoKontak, pasienList[i].TanggalMCU, pasienList[i].PaketDipilih)
					found = true // Tandai bahwa ada pasien yang ditemukan
				}
			}
			if !found {
				fmt.Println("Tidak ada pasien yang memilih paket ini.")
			}
		} else {
			fmt.Println("Paket Tidak Ada")
			return
		}

		
	case 3:
		var tanggalMulai, tanggalAkhir string
		fmt.Println("Masukkan Tanggal Mulai (YYYY-MM-DD):")
		fmt.Scanln(&tanggalMulai)
		fmt.Println("Masukkan Tanggal Akhir (YYYY-MM-DD):")
		fmt.Scanln(&tanggalAkhir)
		fmt.Println("Hasil Pencarian:")
		for i := 0; i < nPasien; i++ {
			if pasienList[i].TanggalMCU >= tanggalMulai && pasienList[i].TanggalMCU <= tanggalAkhir {
				fmt.Printf("ID: %s, Nama: %s, Jenis Kelamin: %s, Alamat: %s, No. Kontak: %s, Tanggal MCU: %s, Paket Dipilih: %s\n",
					pasienList[i].IDPasien, pasienList[i].NamaLengkap, pasienList[i].JenisKelamin, pasienList[i].Alamat, pasienList[i].NoKontak, pasienList[i].TanggalMCU, pasienList[i].PaketDipilih)
			}
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func rekapHasilMCU() {
	fmt.Println(MenuRekapHasil)
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tambahHasilMCU()
	case 2:
		editHasilMCU()
	case 3:
		hapusHasilMCU()
	case 4:
		tampilkanRekapHasilMCU()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func tambahHasilMCU() {
	if nPasien == 0 {
		fmt.Println("Tidak ada pasien untuk ditambahkan hasil MCU.")
		return
	}

	var idPasien, hasil string
	fmt.Println("Masukkan ID Pasien:")
	fmt.Scanln(&idPasien)

	for i := 0; i < nPasien; i++ {
		if pasienList[i].IDPasien == idPasien {
			fmt.Println("Masukkan Hasil MCU:")
			fmt.Scanln(&hasil)
			rekapHasilList[nRekapHasil] = RekapHasilMCU{
				IDRekap:    fmt.Sprintf("R%d", nRekapHasil+1),
				IDPasien:   idPasien,
				TanggalMCU: pasienList[i].TanggalMCU,
				HasilMCU:   hasil,
			}
			nRekapHasil++
			fmt.Println("Rekap hasil MCU berhasil ditambahkan.")
			return
		}
	}
	fmt.Println("ID Pasien tidak ditemukan.")
}

func editHasilMCU() {
	if nRekapHasil == 0 {
		fmt.Println("Tidak ada rekap hasil MCU yang dapat diedit.")
		return
	}

	var idRekap, hasilBaru string
	fmt.Println("Masukkan ID Rekap Hasil MCU yang ingin diedit:")
	fmt.Scanln(&idRekap)

	for i := 0; i < nRekapHasil; i++ {
		if rekapHasilList[i].IDRekap == idRekap {
			fmt.Println("Masukkan Hasil MCU baru:")
			fmt.Scanln(&hasilBaru)
			rekapHasilList[i].HasilMCU = hasilBaru
			fmt.Println("Hasil MCU berhasil diperbarui.")
			return
		}
	}
	fmt.Println("ID Rekap Hasil MCU tidak ditemukan.")
}

func hapusHasilMCU() {
	if nRekapHasil == 0 {
		fmt.Println("Tidak ada rekap hasil MCU yang dapat dihapus.")
		return
	}

	var idRekap string
	fmt.Println("Masukkan ID Rekap Hasil MCU yang ingin dihapus:")
	fmt.Scanln(&idRekap)

	for i := 0; i < nRekapHasil; i++ {
		if rekapHasilList[i].IDRekap == idRekap {
			for j := i; j < nRekapHasil-1; j++ {
				rekapHasilList[j] = rekapHasilList[j+1]
			}
			nRekapHasil--
			fmt.Println("Rekap hasil MCU berhasil dihapus.")
			return
		}
	}
	fmt.Println("ID Rekap Hasil MCU tidak ditemukan.")
}

func tampilkanRekapHasilMCU() {
	if nRekapHasil == 0 {
		fmt.Println("Tidak ada rekap hasil MCU.")
		return
	}

	fmt.Println("Rekap Hasil MCU:")
	for i := 0; i < nRekapHasil; i++ {
		fmt.Printf("ID Rekap: %s, ID Pasien: %s, Tanggal MCU: %s, Hasil: %s\n",
			rekapHasilList[i].IDRekap, rekapHasilList[i].IDPasien, rekapHasilList[i].TanggalMCU, rekapHasilList[i].HasilMCU)
	}
}

func tampilkanLaporan() {
	if nPasien == 0 {
		fmt.Println("Belum ada data pasien yang terdaftar.")
		return
	}

	fmt.Println("Pilih Periode Laporan: 1. Harian 2. Bulanan 3. Tahunan")
	var pilihan int
	fmt.Scanln(&pilihan)

	var tanggalMulai, tanggalAkhir string
	switch pilihan {
	case 1:
		fmt.Println("Masukkan Tanggal (YYYY-MM-DD):")
		fmt.Scanln(&tanggalMulai)
		tanggalAkhir = tanggalMulai
	case 2:
		fmt.Println("Masukkan Bulan dan Tahun (YYYY-MM):")
		var bulan string
		fmt.Scanln(&bulan)
		tanggalMulai = bulan + "-01"
		tanggalAkhir = bulan + "-31"
	case 3:
		fmt.Println("Masukkan Tahun (YYYY):")
		var tahun string
		fmt.Scanln(&tahun)
		tanggalMulai = tahun + "-01-01"
		tanggalAkhir = tahun + "-12-31"
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	totalPemasukan := 0
	var namaPaketTerjual [100]string
	var jumlahPaketTerjual [100]int
	jumlahPaketTercatat := 0

	// Hitung pemasukan dan catat paket terjual
	for i := 0; i < nPasien; i++ {
		tanggalMCU := pasienList[i].TanggalMCU
		if tanggalMCU >= tanggalMulai && tanggalMCU <= tanggalAkhir {
			paketDipilih := pasienList[i].PaketDipilih
			for j := 0; j < nPaketMCU; j++ {
				if paketMCUList[j].IDPaketMCU == paketDipilih {
					totalPemasukan += paketMCUList[j].HargaPaket

					// Periksa apakah paket sudah tercatat
					found := false
					for k := 0; k < jumlahPaketTercatat; k++ {
						if namaPaketTerjual[k] == paketMCUList[j].NamaPaket {
							jumlahPaketTerjual[k]++
							found = true
							break
						}
					}

					// Jika belum tercatat, tambahkan ke daftar
					if !found {
						namaPaketTerjual[jumlahPaketTercatat] = paketMCUList[j].NamaPaket
						jumlahPaketTerjual[jumlahPaketTercatat] = 1
						jumlahPaketTercatat++
					}
				}
			}
		}
	}

	// Tampilkan laporan
	fmt.Printf("Laporan Pemasukan Periode %s sampai %s:\n", tanggalMulai, tanggalAkhir)
	fmt.Println("Daftar Paket MCU Terjual:")
	for i := 0; i < jumlahPaketTercatat; i++ {
		fmt.Printf("- Paket: %s, Terjual: %d kali\n", namaPaketTerjual[i], jumlahPaketTerjual[i])
	}
	fmt.Printf("Total Pemasukan: Rp %d\n", totalPemasukan)
}
