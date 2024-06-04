package main

import "fmt"

const NMAX int = 100

type Mahasiswa struct {
	NIM       string
	Nama      string
	SKS       int
	Matkul    [NMAX]Matakuliah
	Transkrip [NMAX]Nilai
	JmlMatkul int
}
type Matakuliah struct {
	Kode string
	Nama string
	SKS  int
}
type Nilai struct {
	NamaMK string
	Kode   string
	UTS    float64
	UAS    float64
	Quiz   float64
	Total  float64
	Grade  string
}
type daftarMahasiswa [NMAX]Mahasiswa
type daftarMatakuliah [NMAX]Matakuliah

func main() {
	var mahasiswa daftarMahasiswa
	var matakuliah daftarMatakuliah
	var nMahasiswa, nMatakuliah int
	mahasiswa[nMahasiswa].NIM = "123456"
	mahasiswa[nMahasiswa].Nama = "Jundi"
	mahasiswa[nMahasiswa].SKS = 20
	mahasiswa[nMahasiswa].JmlMatkul = 2
	mahasiswa[nMahasiswa].Matkul[0] = Matakuliah{Kode: "MK001", Nama: "Matematika", SKS: 3}
	mahasiswa[nMahasiswa].Matkul[1] = Matakuliah{Kode: "MK002", Nama: "Fisika", SKS: 4}
	mahasiswa[nMahasiswa].Transkrip[0] = Nilai{NamaMK: "Matematika", Kode: "MK001", UTS: 80, UAS: 85, Quiz: 90, Total: 85, Grade: "A"}
	mahasiswa[nMahasiswa].Transkrip[1] = Nilai{NamaMK: "Fisika", Kode: "MK002", UTS: 75, UAS: 80, Quiz: 85, Total: 80, Grade: "B"}
	nMahasiswa++

	mahasiswa[nMahasiswa].NIM = "654321"
	mahasiswa[nMahasiswa].Nama = "Rizki"
	mahasiswa[nMahasiswa].SKS = 18
	mahasiswa[nMahasiswa].JmlMatkul = 2
	mahasiswa[nMahasiswa].Matkul[0] = Matakuliah{Kode: "MK003", Nama: "Kimia", SKS: 3}
	mahasiswa[nMahasiswa].Matkul[1] = Matakuliah{Kode: "MK004", Nama: "Biologi", SKS: 4}
	mahasiswa[nMahasiswa].Transkrip[0] = Nilai{NamaMK: "Kimia", Kode: "MK003", UTS: 85, UAS: 90, Quiz: 95, Total: 90, Grade: "A"}
	mahasiswa[nMahasiswa].Transkrip[1] = Nilai{NamaMK: "Biologi", Kode: "MK004", UTS: 70, UAS: 75, Quiz: 80, Total: 75, Grade: "B"}
	nMahasiswa++

	mahasiswa[nMahasiswa].NIM = "789012"
	mahasiswa[nMahasiswa].Nama = "Faisal"
	mahasiswa[nMahasiswa].SKS = 22
	mahasiswa[nMahasiswa].JmlMatkul = 3
	mahasiswa[nMahasiswa].Matkul[0] = Matakuliah{Kode: "MK005", Nama: "Statistika", SKS: 3}
	mahasiswa[nMahasiswa].Matkul[1] = Matakuliah{Kode: "MK006", Nama: "Algoritma", SKS: 4}
	mahasiswa[nMahasiswa].Matkul[2] = Matakuliah{Kode: "MK007", Nama: "Pemrograman", SKS: 4}
	mahasiswa[nMahasiswa].Transkrip[0] = Nilai{NamaMK: "Statistika", Kode: "MK005", UTS: 88, UAS: 92, Quiz: 85, Total: 88.33, Grade: "A"}
	mahasiswa[nMahasiswa].Transkrip[1] = Nilai{NamaMK: "Algoritma", Kode: "MK006", UTS: 78, UAS: 80, Quiz: 82, Total: 80, Grade: "B"}
	mahasiswa[nMahasiswa].Transkrip[2] = Nilai{NamaMK: "Pemrograman", Kode: "MK007", UTS: 90, UAS: 95, Quiz: 92, Total: 92.33, Grade: "A"}
	nMahasiswa++

	matakuliah[nMatakuliah] = Matakuliah{Kode: "MK001", Nama: "Matematika", SKS: 3}
	nMatakuliah++
	matakuliah[nMatakuliah] = Matakuliah{Kode: "MK002", Nama: "Fisika", SKS: 4}
	nMatakuliah++
	matakuliah[nMatakuliah] = Matakuliah{Kode: "MK003", Nama: "Kimia", SKS: 3}
	nMatakuliah++
	matakuliah[nMatakuliah] = Matakuliah{Kode: "MK004", Nama: "Biologi", SKS: 4}
	nMatakuliah++
	matakuliah[nMatakuliah] = Matakuliah{Kode: "MK005", Nama: "Statistika", SKS: 3}
	nMatakuliah++
	matakuliah[nMatakuliah] = Matakuliah{Kode: "MK006", Nama: "Algoritma", SKS: 4}
	nMatakuliah++
	matakuliah[nMatakuliah] = Matakuliah{Kode: "MK007", Nama: "Pemrograman", SKS: 4}
	nMatakuliah++
	main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
}

func main_menu(mahasiswa daftarMahasiswa, matakuliah daftarMatakuliah, nMahasiswa, nMatakuliah int) {
	var pilihan int
	fmt.Println("===============================================")
	fmt.Println("==========  APLIKASI NILAI MAHASISWA ==========")
	fmt.Println("===============================================")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Data Matakuliah")
	fmt.Println("3. Menambahkan Matakuliah ke Mahasiswa")
	fmt.Println("4. Edit Matakuliah Mahasiswa")
	fmt.Println("5. Hapus Matakuliah dari Mahasiswa")
	fmt.Println("6. Tampilkan Transkrip Nilai Mahasiswa")
	fmt.Println("7. Tampilkan Daftar mahasiswa yang mengambil suatu Mata Kuliah")
	fmt.Println("8. Tampilkan Daftar Mata Kuliah yang diambil oleh mahasiswa")
	fmt.Println("9. Keluar")
	fmt.Println("===============================================")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihan)
	fmt.Println("===============================================")
	switch pilihan {
	case 1:
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 2:
		menu_data_matakuliah(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 3:
		tambahMatakuliahMahasiswa(&mahasiswa, &matakuliah, nMatakuliah, nMahasiswa)
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 4:
		editMatakuliahMahasiswa(&mahasiswa, &matakuliah, nMahasiswa, nMatakuliah)
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 5:
		hapusMatakuliahMahasiswa(&mahasiswa, nMahasiswa)
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 6:
		TampilkanTranskripNilai(&mahasiswa, nMahasiswa)
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 7:
		TampilkanMahasiswaBerdasarkanMatakuliah(&mahasiswa, nMahasiswa)
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 8:
		tampilkanMatakuliahBerdasarkanMahasiswa(&mahasiswa, nMahasiswa)
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 9:
		return
	default:
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	}
}

func menu_data_mahasiswa(mahasiswa daftarMahasiswa, matakuliah daftarMatakuliah, nMahasiswa, nMatakuliah int) {
	var pilihan int
	fmt.Println("===============================================")
	fmt.Println("1. Input Data Mahasiswa")
	fmt.Println("2. Edit Data Mahasiswa")
	fmt.Println("3. Hapus Data Mahasiswa")
	fmt.Println("4. Edit Nilai Data Mahasiswa")
	fmt.Println("5. Tampil Data Mahasiswa")
	fmt.Println("6. Tampil Daftar Mahasiswa secara Terurut NIM Terkecil-Terbesar")
	fmt.Println("7. Tampil Daftar Mahasiswa secara Terurut NIM Terbesar-Terkecil")
	fmt.Println("8. Tampil Daftar Mahasiswa secara Terurut Jumlah SKS Terkecil-Terbesar")
	fmt.Println("9. Tampil Daftar Mahasiswa secara Terurut Jumlah SKS Terbesar-Terkecil")
	fmt.Println("10. Kembali")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihan)
	fmt.Println("===============================================")
	switch pilihan {
	case 1:
		TambahDataMahasiswa(&mahasiswa, &nMahasiswa)
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 2:
		EditDataMahasiswa(&mahasiswa, nMahasiswa)
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 3:
		HapusDataMahasiswa(&mahasiswa, &nMahasiswa)
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 4:
		editMatakuliahMahasiswa(&mahasiswa, &matakuliah, nMahasiswa, nMatakuliah)
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 5:
		TampilDataMahasiswa(mahasiswa, nMahasiswa)
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 6:
		TampilkanMahasiswaTerurut_NIM_Ascending(&mahasiswa, nMahasiswa)
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 7:
		TampilkanMahasiswaTerurut_NIM_Descending(&mahasiswa, nMahasiswa)
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 8:
		TampilkanMahasiswaTerurut_SKS_Ascending(&mahasiswa, nMahasiswa)
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 9:
		TampilkanMahasiswaTerurut_SKS_Descending(&mahasiswa, nMahasiswa)
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 10:
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	default:
		menu_data_mahasiswa(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	}
}

func menu_data_matakuliah(mahasiswa daftarMahasiswa, matakuliah daftarMatakuliah, nMahasiswa, nMatakuliah int) {
	fmt.Println("===============================================")
	fmt.Println("1. Input Data Matakuliah")
	fmt.Println("2. Edit Data Matakuliah")
	fmt.Println("3. Hapus Data Matakuliah")
	fmt.Println("4. Tampil Data Matakuliah")
	fmt.Println("5. Kembali")
	var pilihan int
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihan)
	fmt.Println("===============================================")
	switch pilihan {
	case 1:
		TambahDataMatakuliah(&matakuliah, &nMatakuliah)
		menu_data_matakuliah(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 2:
		EditDataMatakuliah(&matakuliah, nMatakuliah)
		menu_data_matakuliah(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 3:
		HapusDataMatakuliah(&matakuliah, &nMatakuliah)
		menu_data_matakuliah(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 4:
		TampildataMatakuliah(matakuliah, nMatakuliah)
		menu_data_matakuliah(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	case 5:
		main_menu(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	default:
		menu_data_matakuliah(mahasiswa, matakuliah, nMahasiswa, nMatakuliah)
	}
}

func TambahDataMahasiswa(MH *daftarMahasiswa, N *int) {
	var jumlah int
	fmt.Print("Masukkan jumlah mahasiswa yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)
	fmt.Println("===============================================")

	if *N+jumlah > NMAX {
		jumlah = NMAX - *N
		fmt.Printf("Jumlah mahasiswa melebihi kapasitas maksimum. Hanya %d mahasiswa yang akan ditambahkan.\n", jumlah)
	}

	for i := 0; i < jumlah; i++ {
		fmt.Printf("Data Mahasiswa ke-%d \n", *N+1)
		fmt.Print("Masukkan Nama Mahasiswa Baru : ")
		fmt.Scan(&MH[*N].Nama)
		fmt.Print("Masukkan NIM Mahasiswa Baru : ")
		fmt.Scan(&MH[*N].NIM)
		fmt.Print("Masukkan SKS Mahasiswa Baru : ")
		fmt.Scan(&MH[*N].SKS)
		fmt.Println("===============================================")
		*N++
	}
}

func TampilDataMahasiswa(MH daftarMahasiswa, N int) {
	if N == 0 {
		fmt.Println("Data mahasiswa masih kosong.")
	} else {
		for i := 0; i < N; i++ {
			fmt.Println("===============================================")
			fmt.Println("Data Mahasiswa ke", i+1)
			fmt.Println("Nama Mahasiswa :", MH[i].Nama)
			fmt.Println("NIM Mahasiswa :", MH[i].NIM)
			fmt.Println("Jumlah SKS :", MH[i].SKS)
		}
	}
}

func EditDataMahasiswa(E *daftarMahasiswa, N int) {
	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang ingin diedit: ")
	fmt.Scan(&nim)
	nampung := sequentialSearchMahasiswa(E, N, nim)
	if nampung == -1 {
		fmt.Println("===============================================")
		fmt.Println("NIM yang anda cari tidak ada")
	} else {
		pilihanEditMahasiswa(E, nampung)
	}
}

func HapusDataMahasiswa(E *daftarMahasiswa, N *int) {
	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dihapus: ")
	fmt.Scan(&nim)

	nampung := sequentialSearchMahasiswa(E, *N, nim)
	if nampung == -1 {
		fmt.Println("NIM yang anda cari tidak ada")
	} else {
		fmt.Println("===============================================")
		fmt.Println("Data mahasiswa dengan nama", E[nampung].Nama, "berhasil dihapus.")
		for i := nampung; i < *N-1; i++ {
			E[i] = E[i+1]
		}
		*N--
	}
}

func pilihanEditMahasiswa(P *daftarMahasiswa, index int) {
	for {
		var pilihan int
		fmt.Println("===============================================")
		fmt.Println("DATA MAHASISWA")
		fmt.Println("Nama Mahasiswa :", P[index].Nama)
		fmt.Println("NIM Mahasiswa :", P[index].NIM)
		fmt.Println("Jumlah SKS :", P[index].SKS)
		fmt.Println("===============================================")
		fmt.Println("1. NIM")
		fmt.Println("2. Nama")
		fmt.Println("3. SKS")
		fmt.Println("4. Selesai Mengedit")
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Input NIM baru: ")
			fmt.Scan(&P[index].NIM)
			fmt.Println("NIM Mahasiswa", P[index].Nama, "Berhasil di ubah!")
		case 2:
			fmt.Print("Input Nama baru: ")
			fmt.Scan(&P[index].Nama)
			fmt.Println("NAMA Mahasiswa dengan NIM", P[index].NIM, "Berhasil di ubah!")
		case 3:
			fmt.Print("Input SKS baru: ")
			fmt.Scan(&P[index].SKS)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func sequentialSearchMahasiswa(A *daftarMahasiswa, N int, nim string) int {
	idx := -1
	for i := 0; i < N; i++ {
		if A[i].NIM == nim {
			idx = i
			break
		}
	}
	return idx
}

func binarySearchMahasiswa(B *daftarMahasiswa, N int, nim string) int {
	var left, right int
	left = 0
	right = N - 1
	for left <= right {
		mid := (left + right) / 2
		if B[mid].NIM == nim {
			return mid
		} else if B[mid].NIM < nim {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func TambahDataMatakuliah(MK *daftarMatakuliah, N *int) {
	var jumlah int
	fmt.Print("Masukkan jumlah matakuliah yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)

	if *N+jumlah > NMAX {
		jumlah = NMAX - *N
		fmt.Printf("Jumlah mahasiswa melebihi kapasitas maksimum. Hanya %d mahasiswa yang akan ditambahkan.\n", jumlah)
	}

	for i := 0; i < jumlah; i++ {
		fmt.Printf("Data Mata Kuliah ke-%d \n", *N+1)
		fmt.Print("Masukkan Nama Mata Kuliah Baru : ")
		fmt.Scan(&MK[*N].Nama)
		fmt.Print("Masukkan Kode Mata Kuliah Baru : ")
		fmt.Scan(&MK[*N].Kode)
		fmt.Print("Masukkan SKS Mata Kuliah Baru : ")
		fmt.Scan(&MK[*N].SKS)
		*N++
	}
}

func TampildataMatakuliah(MK daftarMatakuliah, N int) {
	if N == 0 {
		fmt.Println("Data mahasiswa masih kosong.")
	} else {
		for i := 0; i < N; i++ {
			fmt.Println("===============================================")
			fmt.Println("Data Mata Kuliah ke", i+1)
			fmt.Println("Nama Mata Kuliah :", MK[i].Nama)
			fmt.Println("Kode Mata Kuliah :", MK[i].Kode)
			fmt.Println("Jumlah SKS :", MK[i].SKS)
		}
	}
}

func EditDataMatakuliah(E *daftarMatakuliah, N int) {
	var kode string
	fmt.Print("Masukkan Kode Matakuliah yang ingin diedit: ")
	fmt.Scan(&kode)
	nampung := sequentialSearchMatakuliah(E, N, kode)
	if nampung == -1 {
		fmt.Println("Kode yang anda cari tidak ada")
	} else {
		pilihanEditMatakuliah(E, nampung)
	}
}

func HapusDataMatakuliah(E *daftarMatakuliah, N *int) {
	var kode string
	fmt.Print("Masukkan Kode matakuliah yang ingin dihapus: ")
	fmt.Scan(&kode)

	nampung := sequentialSearchMatakuliah(E, *N, kode)
	if nampung == -1 {
		fmt.Println("Kode yang anda cari tidak ada")
	} else {
		for i := nampung; i < *N-1; i++ {
			E[i] = E[i+1]
		}
		*N--
		fmt.Println("Data matakuliah berhasil dihapus")
	}
}

func pilihanEditMatakuliah(P *daftarMatakuliah, index int) {
	for {
		var pilihan int
		fmt.Println("===============================================")
		fmt.Println("DATA MATAKULIAH")
		fmt.Println("Kode Matakuliah :", P[index].Kode)
		fmt.Println("Nama Matakuliah :", P[index].Nama)
		fmt.Println("Jumlah SKS :", P[index].SKS)
		fmt.Println("===============================================")
		fmt.Println("1. Kode")
		fmt.Println("2. Nama")
		fmt.Println("3. SKS")
		fmt.Println("4. Selesai Mengedit")
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Input Kode baru: ")
			fmt.Scan(&P[index].Kode)
		case 2:
			fmt.Print("Input Nama baru: ")
			fmt.Scan(&P[index].Nama)
		case 3:
			fmt.Print("Input SKS baru: ")
			fmt.Scan(&P[index].SKS)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func sequentialSearchMatakuliah(A *daftarMatakuliah, N int, kode string) int {
	idx := -1
	for i := 0; i < N; i++ {
		if A[i].Kode == kode {
			idx = i
			break
		}
	}
	return idx
}

func AssignMatakuliahkeMahasiswa(MH *daftarMahasiswa, MK daftarMatakuliah, Nmahasiswa int, Nmatakuliah int) {
	var nim, kode string
	i := sequentialSearchMahasiswa(MH, Nmahasiswa, nim)
	j := sequentialSearchMatakuliah(&MK, Nmatakuliah, kode)
	if i == -1 || j == -1 {
		fmt.Println("Assign gagal")
		if i == -1 {
			fmt.Println("NIM Mahasiswa tidak ditemukan!")
		}
		if j == -1 {
			fmt.Println("Kode Mata Kuliah tidak ditemukan!")
		}
	} else {
		MH[i].Matkul[MH[i].JmlMatkul] = MK[j]
		MH[i].JmlMatkul++
		fmt.Printf("%s mengambil mata kuliah %s\n", MH[i].Nama, MK[j].Nama)
	}
}

func selectionSortNIM_Ascending(M *daftarMahasiswa, N int) {
	for i := 0; i < N-1; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if (*M)[j].NIM < (*M)[min].NIM {
				min = j
			}
		}
		(*M)[i], (*M)[min] = (*M)[min], (*M)[i]
	}
}

func selectionSortNIM_Descending(M *daftarMahasiswa, N int) {
	for i := 0; i < N-1; i++ {
		max := i
		for j := i + 1; j < N; j++ {
			if (*M)[j].NIM > (*M)[max].NIM {
				max = j
			}
		}
		(*M)[i], (*M)[max] = (*M)[max], (*M)[i]
	}
}

func TampilkanMahasiswaTerurut_NIM_Ascending(M *daftarMahasiswa, N int) {
	selectionSortNIM_Ascending(M, N)
	TampilDataMahasiswa(*M, N)
}

func TampilkanMahasiswaTerurut_NIM_Descending(M *daftarMahasiswa, N int) {
	selectionSortNIM_Descending(M, N)
	TampilDataMahasiswa(*M, N)
}

func insertionSortSKS_Ascending(M *daftarMahasiswa, N int) {
	for i := 1; i < N; i++ {
		key := (*M)[i]
		j := i - 1
		for j >= 0 && (*M)[j].SKS > key.SKS {
			(*M)[j+1] = (*M)[j]
			j--
		}
		(*M)[j+1] = key
	}
}

func insertionSortSKS_Descending(M *daftarMahasiswa, N int) {
	for i := 1; i < N; i++ {
		key := (*M)[i]
		j := i - 1
		for j >= 0 && (*M)[j].SKS < key.SKS {
			(*M)[j+1] = (*M)[j]
			j--
		}
		(*M)[j+1] = key
	}
}

func TampilkanMahasiswaTerurut_SKS_Ascending(M *daftarMahasiswa, N int) {
	insertionSortSKS_Ascending(M, N)
	TampilDataMahasiswa(*M, N)
}

func TampilkanMahasiswaTerurut_SKS_Descending(M *daftarMahasiswa, N int) {
	insertionSortSKS_Descending(M, N)
	TampilDataMahasiswa(*M, N)
}

func TampilkanTranskripNilai(T *daftarMahasiswa, N int) {
	var nim string
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scan(&nim)
	fmt.Println("===============================================")
	idx := sequentialSearchMahasiswa(T, N, nim)
	if idx != -1 {
		fmt.Printf("=============== TRANSKRIP NILAI ===============\nNama Mahasiswa : %s\nNIM Mahasiswa : %s\n", T[idx].Nama, T[idx].NIM)
		fmt.Println("===============================================")
		for i := 0; i < T[idx].JmlMatkul; i++ {
			nilai := T[idx].Transkrip[i]
			fmt.Printf("Nama Matakuliah : %s\nKode mata Kuliah: %s\nUTS: %.2f\nUAS: %.2f\nQuiz: %.2f\nTotal: %.2f\nGrade: %s\n", nilai.NamaMK, nilai.Kode, nilai.UTS, nilai.UAS, nilai.Quiz, nilai.Total, nilai.Grade)
			fmt.Println("===============================================")
		}
	} else {
		fmt.Println("Mahasiswa tidak ditemukan.")
	}

}

func TampilkanMahasiswaBerdasarkanMatakuliah(M *daftarMahasiswa, N int) {
	var kode string
	var count int
	fmt.Print("Masukkan kode matakuliah: ")
	fmt.Scan(&kode)
	fmt.Println("===============================================")
	fmt.Printf("===== DAFTAR MAHASISWA MATA KULIAH %s ======\n", kode)
	for i := 0; i < N; i++ {
		for j := 0; j < M[i].JmlMatkul; j++ {
			if M[i].Matkul[j].Kode == kode {
				count++
				fmt.Println("===============================================")
				fmt.Println("Data Mahasiswa ke", count)
				fmt.Println("Nama Mahasiswa :", M[i].Nama)
				fmt.Println("NIM Mahasiswa :", M[i].NIM)
			}
		}
	}
}

func tampilkanMatakuliahBerdasarkanMahasiswa(M *daftarMahasiswa, N int) {
	var nim string
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scan(&nim)
	fmt.Println("===============================================")
	selectionSortNIM_Ascending(M, N)
	idx := binarySearchMahasiswa(M, N, nim)
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
	} else {
		fmt.Printf("========== MATA KULIAH YANG DIAMBIL ===========\nNama Mahasiswa : %s\nNIM Mahasiswa : %s\n", M[idx].Nama, M[idx].NIM)
		fmt.Println("===============================================")
		for i := 0; i < M[idx].JmlMatkul; i++ {
			nilai := M[idx].Transkrip[i]
			fmt.Printf("Nama Matakuliah : %s\nKode mata Kuliah: %s\n", nilai.NamaMK, nilai.Kode)
			fmt.Println("===============================================")
		}
	}
}

func InputNilaiMahasiswa(M *daftarMahasiswa, N int) {
	var nim, kode string
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scan(&nim)
	i := sequentialSearchMahasiswa(M, N, nim)
	if i == -1 {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan kode matakuliah: ")
	fmt.Scan(&kode)
	m := -1
	for j := 0; j < M[i].JmlMatkul; j++ {
		if M[i].Matkul[j].Kode == kode {
			m = j
		}
	}
	if m == -1 {
		fmt.Println("Matakuliah dengan kode tersebut tidak ditemukan pada mahasiswa ini.")
		return
	}

	fmt.Print("Masukkan nilai UTS: ")
	fmt.Scan(&M[i].Transkrip[m].UTS)
	fmt.Print("Masukkan nilai UAS: ")
	fmt.Scan(&M[i].Transkrip[m].UAS)
	fmt.Print("Masukkan nilai Quiz: ")
	fmt.Scan(&M[i].Transkrip[m].Quiz)
	M[i].Transkrip[m].Total = (M[i].Transkrip[m].UTS * 0.3) + (M[i].Transkrip[m].UAS * 0.4) + (M[i].Transkrip[m].Quiz * 0.3)
	M[i].Transkrip[m].Grade = hitungGrade(M[i].Transkrip[m].Total)

	fmt.Printf("Nilai mahasiswa %s di matakuliah %s telah diinput.\n", M[i].Nama, M[i].Matkul[m].Nama)
}

func hitungGrade(total float64) string {
	if total >= 85 {
		return "A"
	} else if total >= 70 {
		return "B"
	} else if total >= 55 {
		return "C"
	} else if total >= 40 {
		return "D"
	} else {
		return "E"
	}
}

func tambahMatakuliahMahasiswa(MH *daftarMahasiswa, MK *daftarMatakuliah, nMatakuliah, nMahasiswa int) {
	var nim, kode string
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scan(&nim)
	idx := sequentialSearchMahasiswa(MH, nMahasiswa, nim)
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	} else {
		fmt.Print("Masukkan Kode Matakuliah yang ingin ditambahkan: ")
		fmt.Scan(&kode)
		mkIdx := sequentialSearchMatakuliah(MK, nMatakuliah, kode)
		if mkIdx == -1 {
			fmt.Println("Matakuliah tidak ditemukan.")
			return
		}

		MH[idx].Matkul[MH[idx].JmlMatkul] = (*MK)[mkIdx]
		MH[idx].Transkrip[MH[idx].JmlMatkul] = Nilai{
			NamaMK: (*MK)[mkIdx].Nama,
			Kode:   (*MK)[mkIdx].Kode,
		}
		MH[idx].JmlMatkul++
		fmt.Printf("Matakuliah %s berhasil ditambahkan ke mahasiswa %s\n", (*MK)[mkIdx].Nama, MH[idx].Nama)
	}
}

func editMatakuliahMahasiswa(MH *daftarMahasiswa, MK *daftarMatakuliah, nMahasiswa int, nMatakuliah int) {
	var nim, kode string
	var nilai Nilai
	var MKIDX int
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan kode matakuliah: ")
	fmt.Scan(&kode)
	idx := sequentialSearchMahasiswa(MH, nMahasiswa, nim)
	if idx != -1 {
		matkulIdx := sequentialSearchMatkulMahasiswa(MH, idx, kode)
		if matkulIdx != -1 {
			for x := 0; x < nMatakuliah; x++ {
				if MK[x].Kode == kode {
					MKIDX = x
					break
				}
			}
			nilai.NamaMK = MK[MKIDX].Nama
			nilai.Kode = MK[MKIDX].Kode
			fmt.Print("Masukkan nilai UTS baru: ")
			fmt.Scan(&nilai.UTS)
			fmt.Print("Masukkan nilai UAS baru: ")
			fmt.Scan(&nilai.UAS)
			fmt.Print("Masukkan nilai Quiz baru: ")
			fmt.Scan(&nilai.Quiz)
			nilai.Total = (nilai.UTS*0.3 + nilai.UAS*0.4 + nilai.Quiz*0.3)
			nilai.Grade = hitungGrade(nilai.Total)
			MH[idx].Transkrip[matkulIdx] = nilai
			fmt.Println("Nilai matakuliah berhasil diubah.")
		} else {
			fmt.Println("Mahasiswa belum mengambil matakuliah tersebut.")
		}
	} else {
		fmt.Println("Mahasiswa tidak ditemukan.")
	}
}

func hapusMatakuliahMahasiswa(MH *daftarMahasiswa, nMahasiswa int) {
	var nim, kode string
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan kode matakuliah: ")
	fmt.Scan(&kode)
	idx := sequentialSearchMahasiswa(MH, nMahasiswa, nim)
	if idx != -1 {
		matkulIdx := sequentialSearchMatkulMahasiswa(MH, idx, kode)
		if matkulIdx != -1 {
			MH[idx].SKS -= MH[idx].Matkul[matkulIdx].SKS
			MH[idx].Transkrip[matkulIdx] = MH[idx].Transkrip[MH[idx].JmlMatkul-1]
			MH[idx].Matkul[matkulIdx] = MH[idx].Matkul[MH[idx].JmlMatkul-1]
			MH[idx].JmlMatkul--
			fmt.Println("Matakuliah berhasil dihapus.")
		} else {
			fmt.Println("Mahasiswa belum mengambil matakuliah tersebut.")
		}
	} else {
		fmt.Println("Mahasiswa tidak ditemukan.")
	}
}

func sequentialSearchMatkulMahasiswa(M *daftarMahasiswa, N int, kode string) int {
	idx := -1
	for i := 0; i < M[N].JmlMatkul; i++ {
		if M[N].Transkrip[i].Kode == kode {
			idx = i
			break
		}
	}
	return idx
}
