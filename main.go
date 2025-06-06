package main

import "fmt"

const NMAX int = 100

type CryptoAsset struct {
	Nama      string
	Kesulitan float64
	Reward    float64
	Algoritma string
}

type hasilMining struct {
	Nama   string
	Reward float64
	Hari   int
}

type tabCrypto [NMAX]CryptoAsset
type tabMining [NMAX]hasilMining

var jumlahAset int = 0
var jumlahMining int = 0

func main() {
	var pilih int

	fmt.Println(`
  ███████╗██╗███╗   ███╗██╗   ██╗██╗      █████╗ ███████╗  ██╗
  ██╔════╝██║████╗ ████║██║   ██║██║     ██╔══██╗██╔════╝  ██║
  ███████╗██║██╔████╔██║██║   ██║██║     ███████║███████╗  ██║
  ╚════██║██║██║╚██╔╝██║██║   ██║██║     ██╔══██║╚════██║  ██║
  ███████║██║██║ ╚═╝ ██║╚██████╔╝███████╗██║  ██║███████║  ██║
  ╚══════╝╚═╝╚═╝     ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝╚══════╝  ╚═╝

    ██████╗██████╗ ██╗   ██╗██████╗ ████████╗ ██████╗
   ██╔════╝██╔══██╗╚██╗ ██╔╝██╔══██╗╚══██╔══╝██╔═══██╗
   ██║     ██████╔╝ ╚████╔╝ ██████╔╝   ██║   ██║   ██║
   ██║     ██╔══██╗  ╚██╔╝  ██╔═══╝    ██║   ██║   ██║
   ╚██████╗██║  ██║   ██║   ██║        ██║   ╚██████╔╝
    ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝        ╚═╝    ╚═════╝
	`)

	opsi(pilih)
}

func opsi(pilihan int) {
	var crypAss tabCrypto
	var mining tabMining
	var search, mengurut int
	var kondisi bool = true

	for kondisi {

		
		fmt.Println(" ")
		fmt.Println("=================================================")
		fmt.Println("=============== SIMULASI CRYPTO =================")
		fmt.Println("=================================================")
		
		fmt.Println("1. Tambah Aset Kripto")
		fmt.Println("2. Ubah Aset Kripto")
		fmt.Println("3. Hapus Aset Kripto")
		fmt.Println("4. Lihat Semua Aset")
		fmt.Println("5. Simulasi Mining")
		fmt.Println("6. Cari Aset (Sequential/Binary)")
		fmt.Println("7. Urutkan Aset (Selection/Insertion Sort)")
		fmt.Println("8. Laporan Mining")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahAset(&crypAss)
		case 2:
			ubahAset(&crypAss)
		case 3:
			hapusAset(&crypAss)
		case 4:
			lihatAset(crypAss)
		case 5:
			simulasiMining(&crypAss, &mining)
		case 6:
			cariAset(&search, &crypAss)
		case 7:
			urutkanAset(&crypAss, &mengurut)
		case 8:
			laporanMining(&mining)
		case 9:
			kondisi = false
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahAset(A *tabCrypto) {
	fmt.Println("\n=============================================")
	fmt.Println("=============== Tambah Aset =================")
	fmt.Println("=============================================")
	fmt.Print("Nama Aset: ")
	fmt.Scan(&A[jumlahAset].Nama)
	fmt.Print("Tingkat Kesulitan: ")
	fmt.Scan(&A[jumlahAset].Kesulitan)
	fmt.Print("Reward: ")
	fmt.Scan(&A[jumlahAset].Reward)

	fmt.Println("\nBerikut daftar algoritma yang biasa digunakan:")
	fmt.Printf("%-12s %-35s\n", "ALGORITMA", "CONTOH KOIN")
	fmt.Printf("%-12s %-35s\n", "SHA-256", "Bitcoin(BTC), Bitcoin Cash(BCH)")
	fmt.Printf("%-12s %-35s\n", "Scrypt", "Litecoin(LTC)")
	fmt.Printf("%-12s %-35s\n", "X11", "Dash(DASH)")
	fmt.Printf("%-12s %-35s\n", "Ethash", "Ethereum Classic(ETC)")

	fmt.Print("Algoritma: ")
	fmt.Scan(&A[jumlahAset].Algoritma)

	jumlahAset++
	fmt.Println("Aset berhasil ditambahkan.")
}

func ubahAset(A *tabCrypto) {
	var i int
	var nama string
	var keluar bool = false

	fmt.Println("\n===========================================")
	fmt.Println("=============== Ubah Aset =================")
	fmt.Println("===========================================")
	fmt.Print("Masukkan nama aset yang ingin diubah: ")
	fmt.Scan(&nama)

	i = 0
	for i < jumlahAset && !keluar {
		if (*A)[i].Nama == nama {
			fmt.Print("Nama Aset Baru: ")
			fmt.Scan(&(*A)[i].Nama)
			fmt.Print("Tingkat Kesulitan Baru: ")
			fmt.Scan(&(*A)[i].Kesulitan)
			fmt.Print("Reward Baru: ")
			fmt.Scan(&(*A)[i].Reward)
			fmt.Print("Algoritma Baru: ")
			fmt.Scan(&(*A)[i].Algoritma)

			fmt.Println("Data Berhasil Diubah")
			keluar = true
		
		}
		i++
	}

	if !keluar {
		fmt.Println("Aset tidak ditemukan")
	}
}

func hapusAset(A *tabCrypto) {
	var i, j int
	var nama string
	var ditemukan bool = false

	fmt.Println("\n============================================")
	fmt.Println("=============== Hapus Aset =================")
	fmt.Println("============================================")
	fmt.Print("Masukkan nama aset yang ingin dihapus: ")
	fmt.Scan(&nama)

	i = 0
	for i < jumlahAset && !ditemukan {
		if (*A)[i].Nama == nama {
			for j = i; j < jumlahAset-1; j++ {
				(*A)[j] = (*A)[j+1]
			}
			jumlahAset--
			fmt.Println("Aset berhasil dihapus.")
			ditemukan = true
		}
		i++
	}
	if !ditemukan {
		fmt.Println("Aset tidak ditemukan")
	}
}

func lihatAset(A tabCrypto) {
	fmt.Println("\n============================================")
	fmt.Println("=============== Lihat Aset =================")
	fmt.Println("============================================")

	var i int

	if jumlahAset == 0 {
		fmt.Println("Tidak ada aset yang tersedia")
		return
	}
	fmt.Printf("%-12s %-10s %-10s %-10s\n", "Nama Aset", "Kesulitan", "Reward", "Algoritma")
	for i = 0; i < jumlahAset; i++ {
		fmt.Printf("%-12s %-10.2f %-10.2f %-10s\n", A[i].Nama, A[i].Kesulitan, A[i].Reward, A[i].Algoritma)
	}
}

func simulasiMining(A *tabCrypto, M *tabMining) {
	var nama string
	var power float64
	var hari, i int
	var estimasiWaktu, totalReward float64
	var ditemukan bool = false

	fmt.Println("\n=================================================")
	fmt.Println("=============== Simulasi Mining =================")
	fmt.Println("=================================================")
	fmt.Print("Nama aset: ")
	fmt.Scan(&nama)

	for i = 0; i < jumlahAset; i++ {
		if (*A)[i].Nama == nama {
			fmt.Print("Daya Komputasi (hash/s): ")
			fmt.Scan(&power)
			fmt.Print("Jumlah Hari: ")
			fmt.Scan(&hari)

			estimasiWaktu = (*A)[i].Kesulitan / power
			totalReward = float64(hari) / estimasiWaktu * (*A)[i].Reward

			(*M)[jumlahMining] = hasilMining{nama, totalReward, hari}
			jumlahMining++

			fmt.Printf("Estimasi waktu per blok: %.2f hari\n", estimasiWaktu)
			fmt.Printf("Estimasi reward total: %.2f\n", totalReward)

			ditemukan = true
			
		}
	}
	if !ditemukan {
		fmt.Println("Aset tidak ditemukan")
	}
}

func cariAset(cari *int, A *tabCrypto) {
	var nm string

	fmt.Println("\nPilih metode pencarian:")
	fmt.Println("1. Binary Search")
	fmt.Println("2. Sequential Search")
	fmt.Print(">")
	fmt.Scan(cari)

	fmt.Print("Masukkan nama aset yang dicari: ")
	fmt.Scan(&nm)

	if *cari == 1 {
		urutMenaikNama(A)
		binarySearch(*A, nm)
	} else if *cari == 2 {
		sequentialSearch(*A, nm)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func binarySearch(A tabCrypto, nama string) {
	fmt.Println("\n===============================================")
	fmt.Println("=============== Binary Search =================")
	fmt.Println("===============================================")

	var low int = 0
	var high int = jumlahAset-1
	var mid int
	var ditemukan bool = false

	for low <= high && !ditemukan {
		mid = (low + high) / 2
		if A[mid].Nama == nama {
			tampilDetailAset(A[mid])
			ditemukan = true
		} else if A[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !ditemukan {
		fmt.Println("Aset tidak ditemukan")
	}
}

func sequentialSearch(A tabCrypto, nama string) {
	fmt.Println("\n===================================================")
	fmt.Println("=============== Sequential Search =================")
	fmt.Println("===================================================")

	var i int
	var ditemukan bool = false
	for i = 0; i < jumlahAset; i++ {
		if A[i].Nama == nama {
			tampilDetailAset(A[i])
			ditemukan = true
	
		}
	}
	if !ditemukan {
		fmt.Println("Aset tidak ditemukan")
	}
}

func tampilDetailAset(c CryptoAsset) {
	fmt.Println("\nDitemukan Aset:")
	fmt.Printf("Nama Aset : %s\n", c.Nama)
	fmt.Printf("Algoritma : %s\n", c.Algoritma)
	fmt.Printf("Kesulitan : %.2f\n", c.Kesulitan)
	fmt.Printf("Reward    : %.2f\n", c.Reward)
}

func urutkanAset(A *tabCrypto, pilih *int) {
	var kriteria, urutan int

	fmt.Println("\nPilih metode pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print(">")
	fmt.Scan(pilih)

	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Kesulitan")
	fmt.Println("2. Reward")
	fmt.Print("> ")
	fmt.Scan(&kriteria)

	fmt.Println("\nPilih urutan:")
	fmt.Println("1. Menaik (Ascending)")
	fmt.Println("2. Menurun (Descending)")
	fmt.Print("Urutan: ")
	fmt.Scan(&urutan)

	if *pilih == 1 {
		urutSelection(A, kriteria, urutan)
		lihatAset(*A)
		fmt.Println("Aset berhasil diurutkan.")
	} else if *pilih == 2 {
		urutInsertion(A, kriteria, urutan)
		lihatAset(*A)
		fmt.Println("Aset berhasil diurutkan.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}

	
}

func urutSelection(A *tabCrypto, kriteria, urutan int) {

	var i, j, min int
	for i = 0; i < jumlahAset-1; i++ {
		min = i
		for j = i + 1; j < jumlahAset; j++ {
			if bandingkan((*A)[j], (*A)[min], kriteria, urutan) {
				min = j
			}
		}

		(*A)[i] = (*A)[min]
		(*A)[min] = (*A)[i]
	}
}

func urutInsertion(A *tabCrypto, kriteria, urutan int) {
	var i, j int
	var temp CryptoAsset

	for i = 1; i < jumlahAset; i++ {
		temp = (*A)[i]
		j = i - 1
		for j >= 0 && bandingkan(temp, (*A)[j], kriteria, urutan) {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = temp
	}
}

func bandingkan(a, b CryptoAsset, kriteria, urutan int) bool {
	switch kriteria {
	case 1: 
		if urutan == 1 {
			return a.Kesulitan < b.Kesulitan
		}
		return a.Kesulitan > b.Kesulitan
	case 2: 
		if urutan == 1 {
			return a.Reward < b.Reward
		}
		return a.Reward > b.Reward
	default:
		return false
	}
}


func urutMenaikNama(A *tabCrypto) {
	var i, j, min int

	for i = 0; i < jumlahAset-1; i++ {
		min = i
		for j = i + 1; j < jumlahAset; j++ {
			if (*A)[j].Nama < (*A)[min].Nama {
				min = j
			}
		}

		(*A)[i] = (*A)[min]
		(*A)[min] = (*A)[i]

	}
}

func laporanMining(M *tabMining) {
	var total float64
	var i int
	var n int = jumlahMining

	fmt.Println("\n================================================")
	fmt.Println("=============== Laporan Mining =================")
	fmt.Println("================================================")

	if n == 0 {
		fmt.Println("Tidak ada data mining yang tersedia.")

		fmt.Println(" ")
	}else{
		for i = 0; i < n; i++{
			fmt.Printf("Nama Aset  : %s \n", (*M)[i].Nama)
			fmt.Printf("Reward     : %.2f \n", (*M)[i].Reward)
			fmt.Printf("Durasi     : %d hari", (*M)[i].Hari)
			fmt.Println(" ")

			total += (*M)[i].Reward

			fmt.Println(" ")
		}

		fmt.Printf("Total Semua Reward: %.4f\n", total)
	}
}