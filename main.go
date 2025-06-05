package main

import "fmt"

const NMAX int = 1000000

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

	opsi(pilih)
}

func opsi(pilihan int) {

	var crypAss tabCrypto
	var mining tabMining
	var search, mengurut int
	// var nData int

	var kondisi bool = true

	for kondisi == true {
		fmt.Println("\n=============== SIMULASI CRYPTO =================")
		fmt.Println("1. Tambah Aset Kripto")
		fmt.Println("2. Ubah Aset Kripto")
		fmt.Println("3. Hapus Aset Kripto")
		fmt.Println("4. Lihat Semua Aset")
		fmt.Println("5. Simulasi Mining")
		fmt.Println("6. Cari Aset (Sequential/Binary)")
		fmt.Println("7. Urutkan Aset (Selection/Insertion Sort)")
		fmt.Println("8. Laporan Mining")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: \n")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahAset(&crypAss)
			kondisi = true
		case 2:
			ubahAset(&crypAss)
			kondisi = true
		case 3:
			hapusAset(&crypAss)
			kondisi = true
		case 4:
			lihatAset(&crypAss)
			kondisi = true
		case 5:
			simulasiMining(&crypAss, &mining)
			kondisi = true
		case 6:
			cariAset(&search, &crypAss)
			kondisi = true
		case 7:
			urutkanAset(&mengurut)
			kondisi = true
		case 8:
			laporanMining(&mining)
			kondisi = true
		case 9:
			kondisi = false
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}

}


func tambahAset(A *tabCrypto) {
    fmt.Println("\n=============== Tambah Aset =================")

    fmt.Print("\nNama Aset: ")
    fmt.Scan(&A[jumlahAset].Nama)

    fmt.Print("Tingkat Kesulitan: ")
    fmt.Scan(&A[jumlahAset].Kesulitan)

    fmt.Print("Reward: ")
    fmt.Scan(&A[jumlahAset].Reward)

    fmt.Println(" ")

    fmt.Print("Berikut daftar algoritma yang biasa di gunakan:  \n")

    fmt.Printf("%-12s %-35s\n", "ALGORITMA", "CONTOH KOIN")
    fmt.Printf("%-12s %-35s\n", "SHA-256", "Bitcoin(BTC), Bitcoin Cash(BCH)")
    fmt.Printf("%-12s %-35s\n", "Scrypt", "Litecoin(LTC)")
    fmt.Printf("%-12s %-35s\n", "X11", "Dash(DASH)")
    fmt.Printf("%-12s %-35s\n", "Ethash", "Ethereum Classic(ETC)")

    fmt.Print("\nAlgoritma: ")
    fmt.Scan(&A[jumlahAset].Algoritma)

    jumlahAset++

    fmt.Println("Aset berhasil ditambahkan.")
}

func ubahAset(A *tabCrypto) {
    var i int
    var nama string

    fmt.Println("\n=============== Ubah Aset =================")

    fmt.Print("Masukkan nama aset yang ingin diubah: ")
    fmt.Scan(&nama)

    for i = 0; i < jumlahAset; i++ {
        if (*A)[i].Nama == nama {
            fmt.Print("\nNama Aset Baru: ")
            fmt.Scan(&(*A)[i].Nama)

            fmt.Print("Tingkat Kesulitan Baru: ")
            fmt.Scan(&(*A)[i].Kesulitan)

            fmt.Print("Reward Baru: ")
            fmt.Scan(&(*A)[i].Reward)

            fmt.Print("Algoritma Baru: ")
            fmt.Scan(&(*A)[i].Algoritma)

            fmt.Println("Data Berhasil Diubah")
            return
        }
    }
    fmt.Println("Aset tidak ditemukan")
}


func hapusAset(A *tabCrypto) {
    var i int
    var nama string
    var ditemukan bool = false

    fmt.Println("\n=============== Hapus Aset =================")

    fmt.Print("Masukkan nama aset yang ingin dihapus: ")
    fmt.Scan(&nama)

    for i = 0; i < jumlahAset; i++ {
        if (*A)[i].Nama == nama {
            // Geser ke kiri semua elemen setelah i
            for j := i; j < jumlahAset-1; j++ {
                (*A)[j] = (*A)[j+1]
            }
            jumlahAset--
            fmt.Println("Aset berhasil dihapus.")
            ditemukan = true
            break
        }
    }
    if !ditemukan {
        fmt.Println("Aset tidak ditemukan")
    }
}

func lihatAset(A *tabCrypto) {
    fmt.Println("\n=============== Lihat Aset =================")
    var i int

    if jumlahAset == 0 {
        fmt.Println("Tidak ada aset yang tersedia")
    } else {
        fmt.Printf("%-12s %-10s %-10s %-10s\n", "Nama Aset", "Kesulitan", "Reward", "Algoritma")

        for i = 0; i < jumlahAset; i++ {
            fmt.Printf("%-12s %-10.2f %-10.2f %-10s\n", (*A)[i].Nama, (*A)[i].Kesulitan, (*A)[i].Reward, (*A)[i].Algoritma)
        }
    }

    fmt.Println(" ")
}

func simulasiMining(A *tabCrypto, M *tabMining) {
    var nama string
    var power, estimasiWaktu, totalReward float64
    var hari, i int
    var ditemukan bool = false

    fmt.Println("\n=============== Simulasi Mining =================")
    fmt.Print("Nama aset: ")
    fmt.Scan(&nama)

    if jumlahAset == 0 {
        fmt.Println("Tidak ada aset yang tersedia")
    } else {
        for i = 0; i < jumlahAset; i++ {
            if (*A)[i].Nama == nama {
                fmt.Print("Daya Komputasi (hash/s): ")
                fmt.Scan(&power)
                fmt.Print("Jumlah Hari: ")
                fmt.Scan(&hari)

                estimasiWaktu = (*A)[i].Kesulitan / power
                totalReward = float64(hari) / estimasiWaktu * (*A)[i].Reward

                (*M)[jumlahMining].Nama = nama
                (*M)[jumlahMining].Reward = totalReward
                (*M)[jumlahMining].Hari = hari
                jumlahMining++

                fmt.Printf("Estimasi waktu per blok: %.2f hari\n", estimasiWaktu)
                fmt.Printf("Estimasi reward total: %.4f\n", totalReward)

                ditemukan = true
                break
            }
        }
        if !ditemukan {
            fmt.Println("Aset tidak ditemukan")
        }
    }
    fmt.Println(" ")
}


func cariAset(cari *int, A *tabCrypto) {
    var nm string

    fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
    fmt.Println("1. Binary Search")
    fmt.Println("2. Sequential Search")
    fmt.Print(">")
    fmt.Scan(cari)

    fmt.Print("Masukkan nama aset yang dicari: ")
    fmt.Scan(&nm)

    if *cari == 1 {
		urutMenaikSelect(A)
        binarySearch(*A, nm)
    } else if *cari == 2 {
        sequentialSearch(*A, nm)
    } else {
        fmt.Println("Pilihan tidak valid")
    }

    fmt.Println(" ")
}

func binarySearch(A tabCrypto, nama string) {

	fmt.Println("\n=============== Cari Aset Menggunakan Binary Search =================")

	var low int = 0
	var high int = len(A) - 1
	var ditemukan bool = false

	for low <= high {
		mid := (low + high) / 2
		if A[mid].Nama == nama {
			fmt.Print("\nDitemukan Aset bernama: ", A[mid].Nama) 

			fmt.Println(" ")
			fmt.Println("DETAIL ASET")
			fmt.Println(" ")

			fmt.Println("Nama Aset: ", A[mid].Nama)
			fmt.Println("Algoritma: ", A[mid].Algoritma)
			fmt.Println("Kesulitan: ", A[mid].Kesulitan)
			fmt.Println("Reward: ", A[mid].Reward)
			
			ditemukan = true

		}else if A[mid].Nama < nama {
			low = mid + 1
		}else{
			high = mid - 1
		}
	}

	if ditemukan == false{
		fmt.Println("Aset tidak ditemukan") 
	}

	fmt.Println(" ")
	
}

func sequentialSearch(A tabCrypto, nama string){
	fmt.Println("\n=============== Cari Aset Menggunakan Sequential Search =================")
	var i int
	var ditemukan bool = false

	i = 0
	for i < len(A) && !ditemukan {
		if A[i].Nama == nama {
			fmt.Print("\nDitemukan Aset bernama: ", A[i].Nama)

			fmt.Println(" ")
			fmt.Println("DETAIL ASET")
			fmt.Println(" ")

			fmt.Println("Nama Aset: ", A[i].Nama)
			fmt.Println("Algoritma: ", A[i].Algoritma)
			fmt.Println("Kesulitan: ", A[i].Kesulitan)
			fmt.Println("Reward: ", A[i].Reward)
			
			ditemukan = true
		}
		i++
	}

	if !ditemukan {
		fmt.Println("Aset tidak ditemukan")
	}

	fmt.Println(" ")
}




func urutkanAset(urut *int){
	var tbCry tabCrypto
	var pilihUrut int

	fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
	fmt.Println("1. Selection")
	fmt.Println("2. Insertion Sort")
	fmt.Print(">")
	fmt.Scan(urut)

	if *urut == 1 {
		fmt.Println("\n=============== Urut Menggunakan Selection =================")
		fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
		fmt.Println("1. Urut Menaik")
		fmt.Println("2. Urut Menurun")
		fmt.Print(">")
		fmt.Scan(&pilihUrut)

		if pilihUrut == 1 {
			fmt.Println("\n=============== Urut Menaik Menggunakan Selection =================")
			urutMenaikSelect(&tbCry)
			lihatAset(&tbCry)
			fmt.Println(" ")
		} else if pilihUrut == 2 {
			fmt.Println("\n=============== Urut Menurun Menggunakan Selection =================")
			urutMenurunSelect(&tbCry)
			lihatAset(&tbCry)
			fmt.Println(" ")
		} else {
			fmt.Println("Pilihan tidak ada")
		}
	} else if *urut == 2 {
		fmt.Println("\n=============== Urut Menggunakan Insertion Sort =================")
		fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
		fmt.Println("1. Urut Menaik")
		fmt.Println("2. Urut Menurun")
		fmt.Print(">")
		fmt.Scan(&pilihUrut)

		if pilihUrut == 1 {
			fmt.Println("\n=============== Urut Menaik Menggunakan Insertion Sort =================")
			urutMenaikInsertion(&tbCry)
			lihatAset(&tbCry)
			fmt.Println(" ")
		} else if pilihUrut == 2 {
			fmt.Println("\n=============== Urut Menurun Menggunakan Insertion Sort =================")
			urutMenurunInsertion(&tbCry)
			lihatAset(&tbCry)
			fmt.Println(" ")
		} else {
			fmt.Println("Pilihan tidak ada")
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}


func urutMenaikSelect(A *tabCrypto) {
	var j, i, minIdx int
	var n int
	var temp CryptoAsset

	n = len(*A)

	for i = 0; i < n-1; i++ {
		minIdx = i

		for j = i + 1; j < n; j++ {
			if (*A)[j].Nama < (*A)[minIdx].Nama {
				minIdx = j
			}
		}

		temp = (*A)[i]
		(*A)[i] = (*A)[minIdx]
		(*A)[minIdx] = temp
		// (*A)[i], (*A)[minIdx] = (*A)[minIdx], (*A)[i]
	}
}

func urutMenurunSelect(A *tabCrypto) {
	var j, i, minIdx int
	var n int
	var temp CryptoAsset

	n = len(*A)

	for i = 0; i < n-1; i++ {
		minIdx = i

		for j = i + 1; j < n; j++ {
			if (*A)[j].Nama > (*A)[minIdx].Nama {
				minIdx = j
			}
		}

		temp = (*A)[i]
		(*A)[i] = (*A)[minIdx]
		(*A)[minIdx] = temp
		// (*A)[i], (*A)[maxIdx] = (*A)[maxIdx], (*A)[i]
	}
}

func urutMenaikInsertion(A *tabCrypto) {
	var i, j, n int
	var key CryptoAsset

	n = len(*A)
	for i = 1; i < n; i++ {
		key = (*A)[i]
		j = i - 1
		for j >= 0 && (*A)[j].Nama > key.Nama {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = key
	}
}

func urutMenurunInsertion(A *tabCrypto) {
	var i, j, n int
	var key CryptoAsset

	n = len(*A)
	for i = 1; i < n; i++ {
		key = (*A)[i]
		j = i - 1
		for j >= 0 && (*A)[j].Nama < key.Nama {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = key
	}
}

func laporanMining(M *tabMining) {
	var total float64
	var i int

	fmt.Println("\n=============== Laporan Mining =================")

	if len(*M) == 0 {
		fmt.Println("Tidak ada data mining yang tersedia.")

		fmt.Println(" ")
	}else{
		for i = 0; i < len(*M); i++{
			fmt.Println("Nama Aset\t: ", (*M)[i].Nama)
			fmt.Println("Reward\t: ", (*M)[i].Reward)
			fmt.Printf("Durasi\t: %d hari", (*M)[i].Hari)
			total += (*M)[i].Reward

			fmt.Println(" ")
		}

		fmt.Printf("Total Semua Reward: %.4f\n", total)
	}
}