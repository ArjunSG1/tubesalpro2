package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	MAX_LAPANGAN = 20
	MAX_PENYEWA  = 50
	MAX_TRANS    = 500
	MAX_JAM      = 14 // 08:00 - 22:00
)

type Lapangan struct {
	ID        int
	Nama      string
	HargaSewa int
	Status    string
}

type Penyewa struct {
	ID      int
	Nama    string
	Telepon string
}

type Transaksi struct {
	IDTransaksi int
	IDPenyewa   int
	IDLapangan  int
	Tanggal     string
	JamMulai    int
	JamSelesai  int
	TotalBayar  int
	Status      string
}

type SlotJadwal struct {
	Tanggal     string
	Status      string
	IDTransaksi int
}

type JamFreq struct {
	Jam    int
	Jumlah int
}

var (
	dataLapangan [MAX_LAPANGAN]Lapangan
	dataPenyewa  [MAX_PENYEWA]Penyewa
	dataTrans    [MAX_TRANS]Transaksi
	dataJadwal   [MAX_LAPANGAN][MAX_JAM]SlotJadwal

	jmlLapangan int
	jmlPenyewa  int
	jmlTrans    int

	reader = bufio.NewReader(os.Stdin)
)

func main() {
	inisialisasiJadwal()
	menuUtama()
}

func inputString(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func inputStringWajib(prompt string, pesan string) string {
	for {
		s := inputString(prompt)
		if s == "0" {
			return ""
		}
		if s == "" {
			fmt.Println(pesan)
			continue
		}
		return s
	}
}

func inputInt(prompt string) int {
	for {
		s := inputString(prompt)
		if s == "0" {
			return 0
		}
		n, err := strconv.Atoi(s)
		if err == nil {
			return n
		}
		fmt.Println("Input harus angka.")
	}
}

func inisialisasiJadwal() {
	for i := 0; i < MAX_LAPANGAN; i++ {
		for j := 0; j < MAX_JAM; j++ {
			dataJadwal[i][j].Tanggal = ""
			dataJadwal[i][j].Status = "Kosong"
			dataJadwal[i][j].IDTransaksi = 0
		}
	}
}

func inputStatusLapangan() string {
	for {
		fmt.Println("Pilih Status Lapangan:")
		fmt.Println("1. Aktif")
		fmt.Println("2. Perbaikan")
		fmt.Println("3. Nonaktif")
		pilihan := inputInt("Pilih (0 untuk batal): ")

		switch pilihan {
		case 1:
			return "Aktif"
		case 2:
			return "Perbaikan"
		case 3:
			return "Nonaktif"
		case 0:
			return ""
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuUtama() {
	for {
		fmt.Println("\n+++ FUTSAL-BOOK +++")
		fmt.Println("1. Kelola Data Lapangan")
		fmt.Println("2. Kelola Data Penyewa")
		fmt.Println("3. Transaksi Penyewaan")
		fmt.Println("4. Pencarian Penyewa")
		fmt.Println("5. Pengurutan")
		fmt.Println("6. Statistik")
		fmt.Println("7. Keluar")

		pilihan := inputInt("Pilih: ")
		switch pilihan {
		case 1:
			menuLapangan()
		case 2:
			menuPenyewa()
		case 3:
			menuTransaksi()
		case 4:
			menuPencarian()
		case 5:
			menuPengurutan()
		case 6:
			menuStatistik()
		case 7:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuLapangan() {
	for {
		fmt.Println("\n--- Kelola Data Lapangan ---")
		fmt.Println("1. Tambah")
		fmt.Println("2. Ubah")
		fmt.Println("3. Hapus")
		fmt.Println("4. Tampil")
		fmt.Println("5. Kembali")

		pilihan := inputInt("Pilih: ")
		switch pilihan {
		case 1:
			tambahLapangan()
		case 2:
			ubahLapangan()
		case 3:
			hapusLapangan()
		case 4:
			tampilLapangan()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuPenyewa() {
	for {
		fmt.Println("\n--- Kelola Data Penyewa ---")
		fmt.Println("1. Tambah")
		fmt.Println("2. Ubah")
		fmt.Println("3. Hapus")
		fmt.Println("4. Tampil")
		fmt.Println("5. Kembali")

		pilihan := inputInt("Pilih: ")
		switch pilihan {
		case 1:
			tambahPenyewa()
		case 2:
			ubahPenyewa()
		case 3:
			hapusPenyewa()
		case 4:
			tampilPenyewa()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuTransaksi() {
	for {
		fmt.Println("\n--- Transaksi Penyewaan ---")
		fmt.Println("1. Tambah Transaksi")
		fmt.Println("2. Tampil Jadwal")
		fmt.Println("3. Tampil Jadwal Kosong")
		fmt.Println("4. Kembali")

		pilihan := inputInt("Pilih: ")
		switch pilihan {
		case 1:
			tambahTransaksi()
		case 2:
			tampilJadwal()
		case 3:
			tampilJadwalKosong()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuPencarian() {
	for {
		fmt.Println("\n--- Pencarian Penyewa ---")
		fmt.Println("1. Sequential Search Nama")
		fmt.Println("2. Sequential Search Telepon")
		fmt.Println("3. Binary Search Nama")
		fmt.Println("4. Binary Search Telepon")
		fmt.Println("5. Kembali")

		pilihan := inputInt("Pilih: ")
		switch pilihan {
		case 1:
			nama := inputString("Masukkan nama (0 untuk batal): ")
			if nama == "0" {
				fmt.Println("Batal.")
				continue
			}
			idx := sequentialSearchNama(nama)
			tampilHasilPenyewa(idx)
		case 2:
			telp := inputString("Masukkan telepon (0 untuk batal): ")
			if telp == "0" {
				fmt.Println("Batal.")
				continue
			}
			idx := sequentialSearchTelepon(telp)
			tampilHasilPenyewa(idx)
		case 3:
			sortPenyewaByNamaAscending()
			nama := inputString("Masukkan nama (0 untuk batal): ")
			if nama == "0" {
				fmt.Println("Batal.")
				continue
			}
			idx := binarySearchNama(nama)
			tampilHasilPenyewa(idx)
		case 4:
			sortPenyewaByTeleponAscending()
			telp := inputString("Masukkan telepon (0 untuk batal): ")
			if telp == "0" {
				fmt.Println("Batal.")
				continue
			}
			idx := binarySearchTelepon(telp)
			tampilHasilPenyewa(idx)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuPengurutan() {
	for {
		fmt.Println("\n--- Pengurutan ---")
		fmt.Println("1. Lapangan Harga Ascending")
		fmt.Println("2. Lapangan Harga Descending")
		fmt.Println("3. Penyewa Nama Ascending")
		fmt.Println("4. Penyewa Nama Descending")
		fmt.Println("5. Kembali")

		pilihan := inputInt("Pilih: ")
		switch pilihan {
		case 1:
			sortLapanganByHargaAscending()
			tampilLapangan()
		case 2:
			sortLapanganByHargaDescending()
			tampilLapangan()
		case 3:
			sortPenyewaByNamaAscending()
			tampilPenyewa()
		case 4:
			sortPenyewaByNamaDescending()
			tampilPenyewa()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuStatistik() {
	for {
		fmt.Println("\n--- Statistik ---")
		fmt.Println("1. Total Pendapatan Bulanan")
		fmt.Println("2. Jam Paling Sering Dipesan")
		fmt.Println("3. Kembali")

		pilihan := inputInt("Pilih: ")
		switch pilihan {
		case 1:
			bulan := inputInt("Masukkan bulan (1-12): ")
			if bulan == 0 {
				fmt.Println("Batal.")
				continue
			}
			tahun := inputInt("Masukkan tahun: ")
			if tahun == 0 {
				fmt.Println("Batal.")
				continue
			}
			fmt.Printf("Total pendapatan bulan %02d tahun %d: %d\n", bulan, tahun, statistikPendapatanBulanan(bulan, tahun))
		case 2:
			statistikJamPalingSering()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func nextIDLapangan() int  { return jmlLapangan + 1 }
func nextIDPenyewa() int   { return jmlPenyewa + 1 }
func nextIDTransaksi() int { return jmlTrans + 1 }

func tambahLapangan() {
	if jmlLapangan >= MAX_LAPANGAN {
		fmt.Println("Data lapangan penuh.")
		return
	}

	nama := inputStringWajib("Nama lapangan (0 untuk batal): ", "Masukkan nama dengan benar.")
	if nama == "" {
		fmt.Println("Batal.")
		return
	}

	harga := inputInt("Harga sewa per jam (0 untuk batal): ")
	if harga == 0 {
		fmt.Println("Batal.")
		return
	}

	status := inputStatusLapangan()
	if status == "" {
		fmt.Println("Batal.")
		return
	}

	dataLapangan[jmlLapangan] = Lapangan{
		ID:        nextIDLapangan(),
		Nama:      nama,
		HargaSewa: harga,
		Status:    status,
	}
	jmlLapangan++
	fmt.Println("Lapangan berhasil ditambah.")
}

func ubahLapangan() {
	tampilLapangan()
	id := inputInt("Masukkan ID lapangan yang diubah (0 untuk batal): ")
	if id == 0 {
		fmt.Println("Batal.")
		return
	}

	idx := cariLapanganByID(id)
	if idx < 0 {
		fmt.Println("Lapangan tidak ditemukan.")
		return
	}

	nama := inputString("Nama baru (Enter jika tidak diubah, 0 untuk batal): ")
	if nama == "0" {
		fmt.Println("Batal ubah data.")
		return
	}
	if nama != "" {
		dataLapangan[idx].Nama = nama
	}

	hargaStr := inputString("Harga sewa baru (Enter jika tidak diubah, 0 untuk batal): ")
	if hargaStr == "0" {
		fmt.Println("Batal ubah data.")
		return
	}
	if hargaStr != "" {
		harga, err := strconv.Atoi(hargaStr)
		if err != nil {
			fmt.Println("Harga harus angka.")
			return
		}
		dataLapangan[idx].HargaSewa = harga
	}

	status := inputStatusLapangan()
	if status != "" {
		dataLapangan[idx].Status = status
	}

	fmt.Println("Data lapangan berhasil diubah.")
}

func hapusLapangan() {
	tampilLapangan()
	id := inputInt("Masukkan ID lapangan yang dihapus (0 untuk batal): ")
	if id == 0 {
		fmt.Println("Batal.")
		return
	}

	idx := cariLapanganByID(id)
	if idx < 0 {
		fmt.Println("Lapangan tidak ditemukan.")
		return
	}

	for i := idx; i < jmlLapangan-1; i++ {
		dataLapangan[i] = dataLapangan[i+1]
	}
	jmlLapangan--
	fmt.Println("Lapangan berhasil dihapus.")
}

func tampilLapangan() {
	if jmlLapangan == 0 {
		fmt.Println("Belum ada data lapangan.")
		return
	}

	fmt.Println("+----+----------------------+------------+------------+")
	fmt.Println("| ID | Nama                 | Harga/Jam  | Status     |")
	fmt.Println("+----+----------------------+------------+------------+")
	for i := 0; i < jmlLapangan; i++ {
		fmt.Printf("| %-2d | %-20s | %-10d | %-10s |\n",
			dataLapangan[i].ID,
			dataLapangan[i].Nama,
			dataLapangan[i].HargaSewa,
			dataLapangan[i].Status)
	}
	fmt.Println("+----+----------------------+------------+------------+")
}

func tambahPenyewa() {
	if jmlPenyewa >= MAX_PENYEWA {
		fmt.Println("Data penyewa penuh.")
		return
	}

	nama := inputStringWajib("Nama penyewa (0 untuk batal): ", "Masukkan nama dengan benar.")
	if nama == "" {
		fmt.Println("Batal.")
		return
	}

	telp := inputStringWajib("Telepon (0 untuk batal): ", "Masukkan telepon dengan benar.")
	if telp == "" {
		fmt.Println("Batal.")
		return
	}

	dataPenyewa[jmlPenyewa] = Penyewa{
		ID:      nextIDPenyewa(),
		Nama:    nama,
		Telepon: telp,
	}
	jmlPenyewa++
	fmt.Println("Penyewa berhasil ditambah.")
}

func ubahPenyewa() {
	tampilPenyewa()
	id := inputInt("Masukkan ID penyewa yang diubah (0 untuk batal): ")
	if id == 0 {
		fmt.Println("Batal.")
		return
	}

	idx := cariPenyewaByID(id)
	if idx < 0 {
		fmt.Println("Penyewa tidak ditemukan.")
		return
	}

	nama := inputString("Nama baru (Enter jika tidak diubah, 0 untuk batal): ")
	if nama == "0" {
		fmt.Println("Batal ubah data.")
		return
	}
	if nama != "" {
		dataPenyewa[idx].Nama = nama
	}

	telp := inputString("Telepon baru (Enter jika tidak diubah, 0 untuk batal): ")
	if telp == "0" {
		fmt.Println("Batal ubah data.")
		return
	}
	if telp != "" {
		dataPenyewa[idx].Telepon = telp
	}

	fmt.Println("Data penyewa berhasil diubah.")
}

func hapusPenyewa() {
	tampilPenyewa()
	id := inputInt("Masukkan ID penyewa yang dihapus (0 untuk batal): ")
	if id == 0 {
		fmt.Println("Batal.")
		return
	}

	idx := cariPenyewaByID(id)
	if idx < 0 {
		fmt.Println("Penyewa tidak ditemukan.")
		return
	}

	for i := idx; i < jmlPenyewa-1; i++ {
		dataPenyewa[i] = dataPenyewa[i+1]
	}
	jmlPenyewa--
	fmt.Println("Penyewa berhasil dihapus.")
}

func tampilPenyewa() {
	if jmlPenyewa == 0 {
		fmt.Println("Belum ada data penyewa.")
		return
	}

	fmt.Println("+----+----------------------+----------------+")
	fmt.Println("| ID | Nama                 | Telepon        |")
	fmt.Println("+----+----------------------+----------------+")
	for i := 0; i < jmlPenyewa; i++ {
		fmt.Printf("| %-2d | %-20s | %-14s |\n",
			dataPenyewa[i].ID,
			dataPenyewa[i].Nama,
			dataPenyewa[i].Telepon)
	}
	fmt.Println("+----+----------------------+----------------+")
}

func tambahTransaksi() {
	if jmlLapangan == 0 || jmlPenyewa == 0 {
		fmt.Println("Data lapangan atau penyewa masih kosong.")
		return
	}

	tampilLapangan()
	idLap := inputInt("Pilih ID lapangan (0 untuk batal): ")
	if idLap == 0 {
		fmt.Println("Batal.")
		return
	}

	idxLap := cariLapanganByID(idLap)
	if idxLap < 0 {
		fmt.Println("Lapangan tidak ditemukan.")
		return
	}

	if strings.ToLower(dataLapangan[idxLap].Status) != "aktif" {
		fmt.Println("Lapangan tidak aktif.")
		return
	}

	tampilPenyewa()
	idPen := inputInt("Pilih ID penyewa (0 untuk batal): ")
	if idPen == 0 {
		fmt.Println("Batal.")
		return
	}

	idxPen := cariPenyewaByID(idPen)
	if idxPen < 0 {
		fmt.Println("Penyewa tidak ditemukan.")
		return
	}

	hariIni := time.Now().Format("2006-01-02")
	fmt.Println("Tanggal transaksi:", hariIni)

	jamMulai := inputInt("Masukkan jam mulai (8-21, 0 untuk batal): ")
	if jamMulai == 0 {
		fmt.Println("Batal.")
		return
	}

	jamSelesai := inputInt("Masukkan jam selesai (9-22, 0 untuk batal): ")
	if jamSelesai == 0 {
		fmt.Println("Batal.")
		return
	}

	if jamMulai < 8 || jamMulai > 21 {
		fmt.Println("Jam mulai tidak valid.")
		return
	}
	if jamSelesai <= jamMulai || jamSelesai > 22 {
		fmt.Println("Jam selesai tidak valid.")
		return
	}

	if !slotTersedia(hariIni, idxLap, jamMulai, jamSelesai) {
		fmt.Println("Slot sudah dipesan.")
		return
	}

	total := dataLapangan[idxLap].HargaSewa * (jamSelesai - jamMulai)

	dataTrans[jmlTrans] = Transaksi{
		IDTransaksi: nextIDTransaksi(),
		IDPenyewa:   idPen,
		IDLapangan:  idLap,
		Tanggal:     hariIni,
		JamMulai:    jamMulai,
		JamSelesai:  jamSelesai,
		TotalBayar:  total,
		Status:      "Dipesan",
	}
	idTrans := dataTrans[jmlTrans].IDTransaksi
	jmlTrans++

	sinkronJadwal(hariIni, idxLap, jamMulai, jamSelesai, idTrans)

	fmt.Println("Transaksi berhasil.")
	fmt.Println("Total pembayaran:", total)
}

func tampilJadwal() {
	if jmlLapangan == 0 {
		fmt.Println("Belum ada data lapangan.")
		return
	}

	tampilLapangan()
	idLap := inputInt("Pilih ID lapangan (0 untuk batal): ")
	if idLap == 0 {
		fmt.Println("Batal.")
		return
	}

	idxLap := cariLapanganByID(idLap)
	if idxLap < 0 {
		fmt.Println("Lapangan tidak ditemukan.")
		return
	}

	hariIni := time.Now().Format("2006-01-02")
	tampilJadwalLapangan(hariIni, idxLap)
}

func tampilJadwalKosong() {
	if jmlLapangan == 0 {
		fmt.Println("Belum ada data lapangan.")
		return
	}

	tampilLapangan()
	idLap := inputInt("Pilih ID lapangan (0 untuk batal): ")
	if idLap == 0 {
		fmt.Println("Batal.")
		return
	}

	idxLap := cariLapanganByID(idLap)
	if idxLap < 0 {
		fmt.Println("Lapangan tidak ditemukan.")
		return
	}

	hariIni := time.Now().Format("2006-01-02")
	fmt.Printf("\n=== Jadwal Kosong Lapangan %s (%s) ===\n", dataLapangan[idxLap].Nama, hariIni)
	fmt.Println("+------------+----------+")
	fmt.Println("| Jam        | Status   |")
	fmt.Println("+------------+----------+")
	for j := 0; j < MAX_JAM; j++ {
		jam := 8 + j
		if dataJadwal[idxLap][j].Tanggal != hariIni || dataJadwal[idxLap][j].Status == "Kosong" {
			fmt.Printf("| %02d:00-%02d:00 | %-8s |\n", jam, jam+1, "Kosong")
		}
	}
	fmt.Println("+------------+----------+")
}

func tampilJadwalLapangan(tanggal string, idxLap int) {
	fmt.Printf("\n=== Jadwal Lapangan %s (%s) ===\n", dataLapangan[idxLap].Nama, tanggal)
	fmt.Println("+------------+----------+")
	fmt.Println("| Jam        | Status   |")
	fmt.Println("+------------+----------+")
	for j := 0; j < MAX_JAM; j++ {
		jam := 8 + j
		status := "Kosong"
		if dataJadwal[idxLap][j].Tanggal == tanggal {
			status = dataJadwal[idxLap][j].Status
		}
		fmt.Printf("| %02d:00-%02d:00 | %-8s |\n", jam, jam+1, status)
	}
	fmt.Println("+------------+----------+")
}

func slotTersedia(tanggal string, idxLap, jamMulai, jamSelesai int) bool {
	for j := jamMulai; j < jamSelesai; j++ {
		idxJam := j - 8
		if idxJam < 0 || idxJam >= MAX_JAM {
			return false
		}
		if dataJadwal[idxLap][idxJam].Tanggal == tanggal && dataJadwal[idxLap][idxJam].Status == "Dipesan" {
			return false
		}
	}
	return true
}

func sinkronJadwal(tanggal string, idxLap, jamMulai, jamSelesai, idTrans int) {
	for j := jamMulai; j < jamSelesai; j++ {
		idxJam := j - 8
		dataJadwal[idxLap][idxJam].Tanggal = tanggal
		dataJadwal[idxLap][idxJam].Status = "Dipesan"
		dataJadwal[idxLap][idxJam].IDTransaksi = idTrans
	}
}

func sequentialSearchNama(nama string) int {
	nama = strings.ToLower(nama)
	for i := 0; i < jmlPenyewa; i++ {
		if strings.ToLower(dataPenyewa[i].Nama) == nama {
			return i
		}
	}
	return -1
}

func sequentialSearchTelepon(telp string) int {
	for i := 0; i < jmlPenyewa; i++ {
		if dataPenyewa[i].Telepon == telp {
			return i
		}
	}
	return -1
}

func sortPenyewaByNamaAscending() {
	for i := 0; i < jmlPenyewa-1; i++ {
		minIdx := i
		for j := i + 1; j < jmlPenyewa; j++ {
			if strings.ToLower(dataPenyewa[j].Nama) < strings.ToLower(dataPenyewa[minIdx].Nama) {
				minIdx = j
			}
		}
		dataPenyewa[i], dataPenyewa[minIdx] = dataPenyewa[minIdx], dataPenyewa[i]
	}
}

func sortPenyewaByNamaDescending() {
	for i := 0; i < jmlPenyewa-1; i++ {
		maxIdx := i
		for j := i + 1; j < jmlPenyewa; j++ {
			if strings.ToLower(dataPenyewa[j].Nama) > strings.ToLower(dataPenyewa[maxIdx].Nama) {
				maxIdx = j
			}
		}
		dataPenyewa[i], dataPenyewa[maxIdx] = dataPenyewa[maxIdx], dataPenyewa[i]
	}
}

func sortPenyewaByTeleponAscending() {
	for i := 1; i < jmlPenyewa; i++ {
		key := dataPenyewa[i]
		j := i - 1
		for j >= 0 && dataPenyewa[j].Telepon > key.Telepon {
			dataPenyewa[j+1] = dataPenyewa[j]
			j--
		}
		dataPenyewa[j+1] = key
	}
}

func binarySearchNama(nama string) int {
	low, high := 0, jmlPenyewa-1
	nama = strings.ToLower(nama)
	for low <= high {
		mid := (low + high) / 2
		midNama := strings.ToLower(dataPenyewa[mid].Nama)
		if midNama == nama {
			return mid
		} else if midNama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func binarySearchTelepon(telp string) int {
	low, high := 0, jmlPenyewa-1
	for low <= high {
		mid := (low + high) / 2
		if dataPenyewa[mid].Telepon == telp {
			return mid
		} else if dataPenyewa[mid].Telepon < telp {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func tampilHasilPenyewa(idx int) {
	if idx >= 0 {
		fmt.Printf("Ditemukan: %d | %s | %s\n", dataPenyewa[idx].ID, dataPenyewa[idx].Nama, dataPenyewa[idx].Telepon)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func sortLapanganByHargaAscending() {
	for i := 1; i < jmlLapangan; i++ {
		key := dataLapangan[i]
		j := i - 1
		for j >= 0 && dataLapangan[j].HargaSewa > key.HargaSewa {
			dataLapangan[j+1] = dataLapangan[j]
			j--
		}
		dataLapangan[j+1] = key
	}
	fmt.Println("Lapangan diurutkan ascending berdasarkan harga.")
}

func sortLapanganByHargaDescending() {
	for i := 1; i < jmlLapangan; i++ {
		key := dataLapangan[i]
		j := i - 1
		for j >= 0 && dataLapangan[j].HargaSewa < key.HargaSewa {
			dataLapangan[j+1] = dataLapangan[j]
			j--
		}
		dataLapangan[j+1] = key
	}
	fmt.Println("Lapangan diurutkan descending berdasarkan harga.")
}

func statistikPendapatanBulanan(bulan, tahun int) int {
	total := 0
	for i := 0; i < jmlTrans; i++ {
		t, err := time.Parse("2006-01-02", dataTrans[i].Tanggal)
		if err != nil {
			continue
		}
		if int(t.Month()) == bulan && t.Year() == tahun {
			total += dataTrans[i].TotalBayar
		}
	}
	return total
}

func statistikJamPalingSering() {
	var frek [24]int

	for i := 0; i < jmlTrans; i++ {
		for jam := dataTrans[i].JamMulai; jam < dataTrans[i].JamSelesai; jam++ {
			frek[jam]++
		}
	}

	max := 0
	for i := 8; i <= 22; i++ {
		if frek[i] > max {
			max = frek[i]
		}
	}

	if max == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}

	fmt.Println("\n+------------+--------+")
	fmt.Println("| Jam        | Kali   |")
	fmt.Println("+------------+--------+")

	for i := 8; i <= 22; i++ {
		if frek[i] > 0 {
			fmt.Printf("| %02d:00      | %-6d |\n", i, frek[i])
		}
	}

	fmt.Println("+------------+--------+")

	fmt.Println("\nJam paling sering dipesan:")
	fmt.Println("+------------+--------+")
	fmt.Println("| Jam        | Kali   |")
	fmt.Println("+------------+--------+")

	for i := 8; i <= 22; i++ {
		if frek[i] == max {
			fmt.Printf("| %02d:00      | %-6d |\n", i, frek[i])
		}
	}

	fmt.Println("+------------+--------+")
}
func cariLapanganByID(id int) int {
	for i := 0; i < jmlLapangan; i++ {
		if dataLapangan[i].ID == id {
			return i
		}
	}
	return -1
}

func cariPenyewaByID(id int) int {
	for i := 0; i < jmlPenyewa; i++ {
		if dataPenyewa[i].ID == id {
			return i
		}
	}
	return -1
}