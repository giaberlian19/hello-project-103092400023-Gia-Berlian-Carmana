package main

import "fmt"

const nmax = 100

type lagu struct {
	judul  string
	artis  string
	genre  string
	rating float64
}

type tabint [nmax]lagu

func main() {
	var musik tabint
	var jumlah int
	var pilih int

	dashboard()
	fmt.Printf("masukan pilihan : ")
	fmt.Scan(&pilih)
	for pilih != 6 {
		if pilih == 1 {
			TambahMusik(&musik, &jumlah)
		} else if pilih == 2 {
			CariMusik(&musik, &jumlah)
		} else if pilih == 3 {
			UbahMusik(&musik, &jumlah)
		} else if pilih == 4 {
			HapusMusik(&musik, &jumlah)
		} else if pilih == 5 {
			UrutanMusik(&musik, &jumlah)
		} else if pilih == 6 {
			fmt.Println("Keluar dari program.")
		} else {
			fmt.Println("Pilihan tidak valid")
		}

		fmt.Println("masukan pilihan :")
		fmt.Scan(&pilih)
	}
}
func dashboard() {
	fmt.Println("~~~~~~~~||~~~~~~~~~||~~~~~~~")
	fmt.Println(" Selamat Datang di MyMusik  ")
	fmt.Println("============================")
	fmt.Println("1. Tambah Musik             ")
	fmt.Println("2. Cari Musik               ")
	fmt.Println("3. Ubah Musik               ")
	fmt.Println("4. Hapus Musik              ")
	fmt.Println("5. list musik anda          ")
	fmt.Println("6. Keluar                   ")
	fmt.Println("============================")
	fmt.Println("~~  Let's Play your music ~~")
	fmt.Println("~~~~~~~||~~~~~~~~~~||~~~~~~~")

}

func TambahMusik(musik *tabint, jumlah *int) {
	var lagi int
	var tamjud, tamart, tamgen, tamrat lagu

	for *jumlah < nmax {
		fmt.Println("Masukkan judul lagu: ")
		fmt.Scan(&tamjud.judul)
		fmt.Println("Masukkan nama artis: ")
		fmt.Scan(&tamart.artis)
		fmt.Println("Masukkan genre: ")
		fmt.Scan(&tamgen.genre)

		tamrat.rating = 0
		for tamrat.rating < 1.00 || tamrat.rating > 5.00 {
			fmt.Println("Masukkan rating lagu kamu (1 - 5): ")
			fmt.Scan(&tamrat.rating)
			if tamrat.rating < 1.00 || tamrat.rating > 5.0 {
				fmt.Println("Rating tidak valid")
			}
		}
		musik[*jumlah].judul = tamjud.judul
		musik[*jumlah].artis = tamart.artis
		musik[*jumlah].genre = tamgen.genre
		musik[*jumlah].rating = tamrat.rating

		fmt.Println("====== Lagu berhasil ditambah ========")
		fmt.Printf("%s - %s - %s - %.2f\n", musik[*jumlah].judul, musik[*jumlah].artis, musik[*jumlah].genre, musik[*jumlah].rating)
		fmt.Println("======================================")

		*jumlah = *jumlah + 1

		fmt.Println("===============")
		fmt.Println("1. Tambah lagi")
		fmt.Println("2. Kembali ke menu")
		fmt.Println("===============")
		fmt.Scan(&lagi)

		if lagi == 2 {
			dashboard()
			return
		}
	}
}
func CariMusik(musik *tabint, jumlah *int) {
	var carjud string
	var cakem int

	fmt.Print("Masukkan judul lagu yang dicari: ")
	fmt.Scan(&carjud)

	SelectionSortByJudul(musik, *jumlah)
	if BinarySearch(musik, *jumlah, carjud) {
		index := CariIndex(musik, *jumlah, carjud)
		fmt.Println("=== Lagu ditemukan cuyyy ===")
		fmt.Printf("%s - %s - %s - %.2f\n", musik[index].judul, musik[index].artis, musik[index].genre, musik[index].rating)
	} else {
		fmt.Println("Lagu tidak ada di list")
	}

	fmt.Println("===================")
	fmt.Println("1.cari lagi        ")
	fmt.Println("2.kembali ke menu  ")
	fmt.Scan(&cakem)

	if cakem == 2 {
		dashboard()
		return
	}
}
func UbahMusik(musik *tabint, jumlah *int) {
	var cariJudul string
	var ubakem int
	var ubah int

	for *jumlah < nmax {
		fmt.Printf("%-5s  %-20s  %-20s  %-20s  %-20s\n", "NO", "JUDUL", "ARTIS", "GENRE", "RATING")
		for i := 0; i < *jumlah; i++ {
			fmt.Printf("%-5d. %-20s  %-20s  %-20s  %-20.2f\n", i+1, musik[i].judul, musik[i].artis, musik[i].genre, musik[i].rating)
		}

		fmt.Print("Masukkan judul lagu yang ingin diubah: ")
		fmt.Scan(&cariJudul)

		if !SequentialSearch(musik, *jumlah, cariJudul) {
			fmt.Println("Lagu tidak ditemukan.")
			dashboard()
			return
		}

		index := CariIndex(musik, *jumlah, cariJudul)
		if index == -1 {
			fmt.Println("Lagu tidak ditemukan.")
			return
		}

		fmt.Println("==============================")
		fmt.Printf("%s - %s - %s - %.2f\n", musik[index].judul, musik[index].artis, musik[index].genre, musik[index].rating)
		fmt.Println("Bagian mana yang mau diubah:")
		fmt.Println("1. Judul   3. Genre")
		fmt.Println("2. Artis   4. Rating")
		fmt.Println("==============================")
		fmt.Scan(&ubah)

		if ubah == 1 {
			fmt.Print("Judul baru: ")
			fmt.Scan(&musik[index].judul)
		} else if ubah == 2 {
			fmt.Print("Artis baru: ")
			fmt.Scan(&musik[index].artis)
		} else if ubah == 3 {
			fmt.Print("Genre baru: ")
			fmt.Scan(&musik[index].genre)
		} else if ubah == 4 {
			fmt.Print("Rating baru: ")
			fmt.Scan(&musik[index].rating)
		} else {
			fmt.Println("Pilihan tidak ada.")
		}

		fmt.Println("================================================")
		fmt.Println("============ Data berhasil diubah ==============")
		fmt.Printf("%s - %s - %s - %.2f\n", musik[index].judul, musik[index].artis, musik[index].genre, musik[index].rating)
		fmt.Println("================================================")

		fmt.Println("=================")
		fmt.Println("1. Ubah lagi")
		fmt.Println("2. Kembali ke menu")
		fmt.Printf("Masukan pilihan: ")
		fmt.Scan(&ubakem)

		if ubakem == 2 {
			dashboard()
			return
		}
	}
}

func HapusMusik(musik *tabint, jumlah *int) {
	var cariJudul string
	var hapkem int

	fmt.Printf("%-5s  %-20s  %-20s  %-20s  %-20s\n", "NO", "JUDUL", "ARTIS", "GENRE", "RATING")
	for i := 0; i < *jumlah; i++ {
		fmt.Printf("%-5d. %-20s  %-20s  %-20s  %-20.2f\n", i+1, musik[i].judul, musik[i].artis, musik[i].genre, musik[i].rating)
	}
	fmt.Print("Masukkan judul lagu yang ingin dihapus: ")
	fmt.Scan(&cariJudul)

	if !SequentialSearch(musik, *jumlah, cariJudul) {
		fmt.Println("Lagu tidak ditemukan.")
		dashboard()
		return
	}

	index := CariIndex(musik, *jumlah, cariJudul)

	for i := index; i < *jumlah-1; i++ {
		musik[i] = musik[i+1]
	}
	*jumlah = *jumlah - 1
	fmt.Println("Data berhasil dihapus.")
	for i := 0; i < *jumlah; i++ {
		fmt.Printf("%-5d %-20s  %-20s  %-20s  %-20.2f\n", i+1, musik[i].judul, musik[i].artis, musik[i].genre, musik[i].rating)
	}
	fmt.Println("=================")
	fmt.Println("1.hapus lagi     ")
	fmt.Println("2.kembali ke menu")
	fmt.Scan(&hapkem)

	if hapkem == 2 {
		dashboard()
		return
	}
}
func UrutanMusik(musik *tabint, jumlah *int) {
	var urut int
	fmt.Println("urutan musik")
	fmt.Println("by genre")

	SelectionSortBygenre(musik, jumlah)
	for i := 0; i < *jumlah; i++ {
		fmt.Printf("%-5d %-20s  %-20s  %-20s  %-20.2f\n", i+1, musik[i].judul, musik[i].artis, musik[i].genre, musik[i].rating)
	}
	fmt.Println("pencet 2 bila ingin kembali")
	if urut == 2 {
		dashboard()
		return
	}

}

// sortingnya ya guys wkwkwwk
func BinarySearch(musik *tabint, jumlah int, target string) bool {
	kiri := 0
	kanan := jumlah - 1
	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		if musik[mid].judul == target {
			return true
		} else if musik[mid].judul < target {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}
	return false
}
func SequentialSearch(musik *tabint, jumlah int, target string) bool {
	for i := 0; i < jumlah; i++ {
		if musik[i].judul == target {
			return true
		}
	}
	return false
}
func CariIndex(musik *tabint, jumlah int, target string) int {
	for i := 0; i < jumlah; i++ {
		if musik[i].judul == target {
			return i
		}
	}
	return -1
}
func SelectionSortByJudul(musik *tabint, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		min := i
		for j := i + 1; j < jumlah; j++ {
			if musik[j].judul < musik[min].judul {
				min = j
			}
		}
		temp := musik[i]
		musik[i] = musik[min]
		musik[min] = temp
	}
}
func InsertionSortByRating(musik *tabint, jumlah int) {
	for i := 1; i < jumlah; i++ {
		temp := musik[i]
		j := i - 1
		for j >= 0 && musik[j].rating < temp.rating {
			musik[j+1] = musik[j]
			j--
		}
		musik[j+1] = temp
	}
}
func SelectionSortBygenre(musik *tabint, jumlah *int) {
	for i := 0; i < *jumlah; i++ {
		min := i
		for j := i + 1; j < *jumlah; j++ {
			if musik[j].genre < musik[min].genre {
				min = j
			}
		}
		temp := musik[i]
		musik[i] = musik[min]
		musik[min] = temp
	}
}
