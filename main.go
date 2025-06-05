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
			lihatAset(crypAss)
			kondisi = true
		case 5:
			simulasiMining(&crypAss, &mining)
			kondisi = true
		case 6:
			cariAset(&search, &crypAss)
			kondisi = true
		case 7:
			urutkanAset(mengurut)
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
	var keluar bool = false

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
            keluar = true
        }
    }

	if keluar == false {
		fmt.Println("Aset tidak ditemukan")
	}
    
}


func hapusAset(A *tabCrypto) {
    var i, j int
    var nama string
    var ditemukan bool = false

    fmt.Println("\n=============== Hapus Aset =================")

    fmt.Print("Masukkan nama aset yang ingin dihapus: ")
    fmt.Scan(&nama)

    for i = 0; i < jumlahAset; i++ {
        if (*A)[i].Nama == nama {
            for j = i; j < jumlahAset-1; j++ {
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

func lihatAset(A tabCrypto) {
    fmt.Println("\n=============== Lihat Aset =================")
    var i int

    if jumlahAset == 0 {
        fmt.Println("Tidak ada aset yang tersedia")
    } else {
        fmt.Printf("%-12s %-10s %-10s %-10s\n", "Nama Aset", "Kesulitan", "Reward", "Algoritma")

        for i = 0; i < jumlahAset; i++ {
            fmt.Printf("%-12s %-10.2f %-10.2f %-10s\n", (A)[i].Nama, (A)[i].Kesulitan, (A)[i].Reward, (A)[i].Algoritma)
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
		urutMenaikSelect(A, )
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
	var mid int
	var high int = jumlahAset - 1 
	var ditemukan bool = false

	if jumlahAset == 0 {
		fmt.Println("Tidak ada aset untuk dicari (data kosong).") 
		
	}

	for low <= high && !ditemukan {
		mid = low + (high-low)/2 
		if A[mid].Nama == nama {
			fmt.Print("\nDitemukan Aset bernama: ", A[mid].Nama)
			fmt.Println(" ")
			fmt.Println("DETAIL ASET")
			fmt.Println(" ")
			fmt.Printf("Nama Aset : %-12s \n", A[mid].Nama)
			fmt.Printf("Algoritma : %-12s\n", A[mid].Algoritma)
			fmt.Printf("Kesulitan : %-12.2f \n", A[mid].Kesulitan)
			fmt.Printf("Reward : %-12.2f  \n", A[mid].Reward)
			ditemukan = true
		} else if A[mid].Nama < nama {
			low = mid + 1
		} else {
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
	var n int = jumlahAset

	i = 0
	for i < n && !ditemukan {
		if A[i].Nama == nama {
			fmt.Print("\nDitemukan Aset bernama: ", A[i].Nama)

			fmt.Println(" ")
			fmt.Println("DETAIL ASET")
			fmt.Println(" ")

			fmt.Printf("Nama Aset : %-12s \n", A[i].Nama)
			fmt.Printf("Algoritma : %-12s\n", A[i].Algoritma)
			fmt.Printf("Kesulitan : %-12.2f \n", A[i].Kesulitan)
			fmt.Printf("Reward : %-12.2f  \n", A[i].Reward)
			
			ditemukan = true
		}
		i++
	}

	if !ditemukan {
		fmt.Println("Aset tidak ditemukan")
	}

	fmt.Println(" ")
}




func urutkanAset(urut int){
	var tbCry tabCrypto
	var pilihUrut, ururtBerdasararkan int

	fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
	fmt.Println("1. Selection")
	fmt.Println("2. Insertion Sort")
	fmt.Print("> ")
	fmt.Scan(&urut)

	if urut == 1 {
		fmt.Println("\n=============== Urut Menggunakan Selection =================")
		fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
		fmt.Println("1. Urut Menaik")
		fmt.Println("2. Urut Menurun")
		fmt.Print("> ")
		fmt.Scan(&pilihUrut)

		if pilihUrut == 1 {
			fmt.Println("\n=============== Urut Menaik Menggunakan Selection =================")
			fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
			fmt.Println("1. Berdasarkan Reward")
			fmt.Println("2. Berdasarkan Kesulitan")
			fmt.Print("> ")
			fmt.Scan(&ururtBerdasararkan)

			urutMenaikSelectPilih(&tbCry, ururtBerdasararkan)
			lihatAset(tbCry)
			fmt.Println(" ")
		} else if pilihUrut == 2 {
			fmt.Println("\n=============== Urut Menurun Menggunakan Selection =================")
			fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
			fmt.Println("1. Berdasarkan Reward")
			fmt.Println("2. Berdasarkan Kesulitan")
			fmt.Print("> ")
			fmt.Scan(&ururtBerdasararkan)

			urutMenurunSelect(&tbCry, ururtBerdasararkan)
			lihatAset(tbCry)
			fmt.Println(" ")
		} else {
			fmt.Println("Pilihan tidak ada")
		}
	} else if urut == 2 {
		fmt.Println("\n=============== Urut Menggunakan Insertion Sort =================")
		fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
		fmt.Println("1. Urut Menaik")
		fmt.Println("2. Urut Menurun")
		fmt.Print("> ")
		fmt.Scan(&pilihUrut)

		if pilihUrut == 1 {
			fmt.Println("\n=============== Urut Menaik Menggunakan Insertion Sort =================")

			fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
			fmt.Println("1. Berdasarkan Reward")
			fmt.Println("2. Berdasarkan Kesulitan")
			fmt.Print("> ")
			fmt.Scan(&ururtBerdasararkan)

			urutMenaikInsertion(&tbCry, ururtBerdasararkan)
			lihatAset(tbCry)
			fmt.Println(" ")
		} else if pilihUrut == 2 {
			fmt.Println("\n=============== Urut Menurun Menggunakan Insertion Sort =================")
			fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
			fmt.Println("1. Berdasarkan Reward")
			fmt.Println("2. Berdasarkan Kesulitan")
			fmt.Print("> ")
			fmt.Scan(&ururtBerdasararkan)

			urutMenurunInsertion(&tbCry, ururtBerdasararkan)
			lihatAset(tbCry)
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
	var temp CryptoAsset


	for i = 0; i < jumlahAset ; i++ {
		minIdx = i

		for j = i + 1; j < jumlahAset; j++ {
			if ((*A)[j].Reward > (*A)[minIdx].Reward) {
				minIdx = j
			}
		}

		temp = (*A)[i]
		(*A)[i] = (*A)[minIdx]
		(*A)[minIdx] = temp

		// (*A)[i], (*A)[minIdx] = (*A)[minIdx], (*A)[i]
	}
}


func urutMenaikSelectPilih(A *tabCrypto, nUrut int) {
	var j, i, minIdx int
	var temp CryptoAsset

	for i = 0; i < jumlahAset; i++ {
		minIdx = i

		for j = i + 1; j < jumlahAset; j++ {
			if nUrut == 1 {
				if ((*A)[j].Reward < (*A)[minIdx].Reward) {
					minIdx = j
				}
			}

			if nUrut == 2 {
				if ((*A)[j].Kesulitan < (*A)[minIdx].Kesulitan) {
					minIdx = j
				}
			}

		}

		temp = (*A)[i]
		(*A)[i] = (*A)[minIdx]
		(*A)[minIdx] = temp

	}
}


func urutMenurunSelect(A *tabCrypto, nUrut int) {
	var j, i, maxIdx int
	// var n int
	var temp CryptoAsset

	// n = jumlahAset

	for i = 0; i < jumlahAset; i++ {
		for j = i + 1; j < jumlahAset; j++ {
			if nUrut == 1 {
				if ((*A)[j].Reward > (*A)[maxIdx].Reward) {
					maxIdx = j
				}
			}

			if nUrut == 2 {
				if ((*A)[j].Kesulitan > (*A)[maxIdx].Kesulitan) {
					maxIdx = j
				}
			}

		}
		temp = (*A)[i]
		(*A)[i] = (*A)[maxIdx]
		(*A)[maxIdx] = temp

	}
}

func urutMenaikInsertion(A *tabCrypto, nUrut int) {
	var i, j int
	var key CryptoAsset

	// n = jumlahAset
	for i = 1; i < jumlahAset; i++ {
		key = (*A)[i]
		j = i - 1

		if nUrut == 1 {
			for j >= 0 && ((*A)[j].Reward > key.Reward){
				(*A)[j+1] = (*A)[j]
				j--
			}
		}

		if nUrut == 2 {
			for j >= 0 && ((*A)[j].Kesulitan > key.Kesulitan){
				(*A)[j+1] = (*A)[j]
				j--
			}
		}
		
		(*A)[j+1] = key
	}
}

func urutMenurunInsertion(A *tabCrypto, nUrut int) {
	var i, j int
	var key CryptoAsset

	for i = 1; i < jumlahAset; i++ {
		key = (*A)[i]
		j = i - 1
		if nUrut == 1 {
			for j >= 0 && ((*A)[j].Reward < key.Reward){
				(*A)[j+1] = (*A)[j]
				j--
			}
		}

		if nUrut == 2 {
			for j >= 0 && ((*A)[j].Kesulitan < key.Kesulitan){
				(*A)[j+1] = (*A)[j]
				j--
			}
		}
		(*A)[j+1] = key
	}
}

func laporanMining(M *tabMining) {
	var total float64
	var i int
	var n int = jumlahMining

	fmt.Println("\n=============== Laporan Mining =================")

	if n == 0 {
		fmt.Println("Tidak ada data mining yang tersedia.")

		fmt.Println(" ")
	}else{
		for i = 0; i < n; i++{
			fmt.Printf("Nama Aset %-12s : ", (*M)[i].Nama)
			fmt.Printf("Reward %-12.2f : ", (*M)[i].Reward)
			fmt.Printf("Durasi %-12d : hari", (*M)[i].Hari)
			fmt.Println(" ")

			total += (*M)[i].Reward

			fmt.Println(" ")
		}

		fmt.Printf("Total Semua Reward: %.4f\n", total)
	}
}