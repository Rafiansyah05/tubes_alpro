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

func main() {
	// var dataCrypto tabCrypto
	// var dataMining tabMining
	// var cariAs, urutkan int
	var pilih int

	opsi(pilih)


}


func opsi(pilihan int){

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
			cariAset(&search)
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

func tambahAset(A *tabCrypto){

	fmt.Println("\n=============== Tambah Aset =================")
	var i int

	fmt.Print("\nNama Aset: ")
	fmt.Scan(&A[i].Nama)

	fmt.Print("Tingkat Kesulitan: ")
	fmt.Scan(&A[i].Kesulitan)

	fmt.Print("\nReward: ")
	fmt.Scan(&A[i].Reward)
	
	fmt.Println(" ")

	fmt.Print("Berikut daftar algoritma yang biasa di gunakan:  \n")
	
	fmt.Printf("%-12s %-35s\n", "ALGORITMA", "CONTOH KOIN")
	fmt.Printf("%-12s %-35s\n", "SHA-256", "Bitcoin(BTC), Bitcoin Cash(BCH)")
	fmt.Printf("%-12s %-35s\n", "Scrypt", "Litecoin(LTC)")
	fmt.Printf("%-12s %-35s\n", "X11", "Dash(DASH)")
	fmt.Printf("%-12s %-35s\n", "Ethash", "Ethereum Classic(ETC)")

	fmt.Print("\nAlgoritma: ")
	fmt.Scan(&A[i].Algoritma)

	i++
	
	fmt.Println("Aset berhasil ditambahkan.")

}

func ubahAset(A *tabCrypto){
	var i int
	var nama string

	fmt.Println("\n=============== Ubah Aset =================")

	fmt.Print("Masukkan nama aset yang ingin diubah: ")
	fmt.Scan(&nama)

	for i = 0; i < len(A); i++{
		if A[i].Nama == nama{
			fmt.Print("\nNama Aset Baru: ")
			fmt.Scan(&A[i].Nama)

			fmt.Print("\nTingkat Kesulitan Baru: ")
			fmt.Scan(&A[i].Kesulitan)

			fmt.Print("\nReward Baru: ")
			fmt.Scan(&A[i].Reward)

			fmt.Print("\nAlgoritma Baru: ")
			fmt.Scan(&A[i].Algoritma)

			fmt.Println("Data Berhasil Diubah")
		}else{
			fmt.Println("Aset tidak ditemukan\n")
		}
	}
}

func hapusAset(A *tabCrypto){
	var i int
	var nama string
	var selesai bool = false

	fmt.Println("\n=============== Hapus Aset =================")

	fmt.Print("Masukkan nama aset yang ingin dihapus: ")
	fmt.Scan(&nama)

	i = 0
	for i < len(A) && !selesai {
		if A[i].Nama == nama{
			A[i] = A[i+1]
			fmt.Println("Aset berhasil dihapus.")
			selesai = true
		}else{
			fmt.Println("Aset tidak ditemukan\n")
		}
	}
}

func lihatAset(A *tabCrypto){
	fmt.Println("\n=============== Lihat Aset =================")
	var i int

	if len(A) == 0{
		fmt.Println("Tidak ada aset yang tersedia")
	}else{
		fmt.Printf("%-12s %-10s %-10s %-10s\n", "Nama Aset", "Kesulitan", "Reward", "Algoritma")

		for i = 0; i < len(A); i++{
			fmt.Printf("%-12s %-10s %-10.2f %-10.2f\n", A[i].Nama, A[i].Kesulitan, A[i].Reward, A[i].Algoritma)
		}
	}

	fmt.Println(" ")
}

func simulasiMining(A *tabCrypto, M *tabMining){
	var nama string
	var power, estimasiWaktu, totalReward float64
	var hari, i int	
	var ditemukan bool = false

	fmt.Println("\n=============== Simulasi Mining =================")
	fmt.Print("Nama aset: ")
	fmt.Scan(&nama)

	
	if len(A) == 0 {
		fmt.Println("Tidak ada aset yang tersedia")
	}else{
		i = 0
		for i < len(A) && !ditemukan {
			if A[i].Nama == nama{

				fmt.Print("Daya Komputasi (hash/s): ")
				fmt.Scan(&power)
				fmt.Print("Jumlah Hari: ")
				fmt.Scan(&hari)

				estimasiWaktu = A[i].Kesulitan / power
				totalReward = float64(hari) / estimasiWaktu * A[i].Reward

				M[i].Nama = nama
				M[i].Reward = totalReward
				M[i].Hari = hari

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

func cariAset(cari *int) {
	var nm string
	var tbCry tabCrypto

	fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
	fmt.Println("1. Binary Search")
	fmt.Println("2. Sequential Search")
	fmt.Print(">")
	fmt.Scan(&cari)

	if *cari == 1 {
		fmt.Scan(&nm)
		binarySearch(tbCry, nm)

	}else if *cari == 2 {
		fmt.Scan(&nm)
		sequentialSearch(tbCry, nm)
	}else{
		fmt.Println("Pilihan tidak valid")
	}

	fmt.Println(" ")
}


func binarySearch(A tabCrypto, nama string) {

	fmt.Println("\n=============== Cari Aset Menggunakan Binary Search =================")

	var low int = 0
	var high int = len(A) - 1
	var ditemukan bool = false

	for low <= high && !ditemukan {
		mid := (low + high) / 2
		if A[mid].Nama == nama {
			fmt.Printf("\nDitemukan Aset bernama: \n", A[mid].Nama) 

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

	if !ditemukan{
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
			fmt.Printf("\nDitemukan Aset bernama: \n", A[i].Nama)

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
	fmt.Scan(&urut)

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
		}else if pilihUrut == 2 {
			fmt.Println("\n=============== Urut Menurun Menggunakan Selection =================")

			urutMenurunSelect(&tbCry)
			lihatAset(&tbCry)

			fmt.Println(" ")
		}else{
			fmt.Println("Pilihan tidak ada")
		}

	}else if *urut == 2{
		fmt.Println("\n=============== Urut Menggunakan Insertion Sort =================")
		fmt.Println("\nsilahkan pilih (1)/(2) mau menggunakan apa: ")
		fmt.Println("1. Urut Menaik")
		fmt.Println("2. Urut Menurun")
		fmt.Print(">")
		fmt.Scan(&pilihUrut)

		if pilihUrut == 1 {
			fmt.Println("\n=============== Urut Menaik Menggunakan Insertion Sort =================")

			urutMenaikInsert(&tbCry)
			lihatAset(&tbCry)
			
			fmt.Println(" ")
		}else if pilihUrut == 2 {
			fmt.Println("\n=============== Urut Menurun Menggunakan Insertion Sort =================")

			urutMenurunInsert(&tbCry)
			lihatAset(&tbCry)

			fmt.Println(" ")
		}else{
			fmt.Println("Pilihan tidak ada")
		}
	}else {
		fmt.Println("Pilihan tidak ada")
	}

	fmt.Println(" ")
}

func urutMenaikSelect(A *tabCrypto) {
	var i, pass, idx int
	var temp CryptoAsset

	for i = 0; i < len(A); i++{
		idx = i

		for pass = i + 1; pass < len(A); pass++{
			if A[pass].Kesulitan < A[idx].Kesulitan {
				idx = pass
			}
		}

		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}
}

func urutMenurunSelect(A *tabCrypto) {
	var i, pass, idx int
	var temp CryptoAsset

	for i = 0; i < len(A); i++{
		idx = i

		for pass = i + 1; pass < len(A); pass++{
			if A[pass].Kesulitan > A[idx].Kesulitan {
				idx = pass
			}
		}

		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}
}



func urutMenaikInsert(A *tabCrypto) {
	var  i, pass int
	var key CryptoAsset

	for i = 1; i < len(A); i++{

		key = A[i]
		pass = i - 1

		for pass >= 0 && A[pass].Kesulitan > key.Kesulitan {
			A[pass + 1] = A[pass]
			pass--
		}

		A[pass+1] = key
	}

}

func urutMenurunInsert(A *tabCrypto) {
	var  i, pass int
	var key CryptoAsset

	for i = 1; i < len(A); i++{

		key = A[i]
		pass = i - 1

		for pass >= 0 && A[pass].Kesulitan < key.Kesulitan {
			A[pass + 1] = A[pass]
			pass--
		}

		A[pass+1] = key
	}

}


func laporanMining(M *tabMining) {
	var total float64
	var i int

	fmt.Println("\n=============== Laporan Mining =================")

	if len(M) == 0 {
		fmt.Println("Tidak ada data mining yang tersedia.")

		fmt.Println(" ")
	}else{
		for i = 0; i < len(M); i++{
			fmt.Println("Nama Aset\t: ", M[i].Nama)
			fmt.Println("Reward\t: ", M[i].Reward)
			fmt.Println("Durasi\t: %d hari", M[i].Hari)
			total += M[i].Reward

			fmt.Println(" ")
		}

		fmt.Printf("Total Semua Reward: %.4f\n", total)
	}
}