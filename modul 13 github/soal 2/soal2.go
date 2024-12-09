package main

import "fmt"

const nMax = 7919

type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku struct {
	pustaka  [nMax]Buku
	nPustaka int
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	fmt.Println("Masukkan data buku (id, judul, penulis, penerbit, eksemplar, tahun, rating):")
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka.pustaka[i].id, &pustaka.pustaka[i].judul, &pustaka.pustaka[i].penulis, &pustaka.pustaka[i].penerbit, &pustaka.pustaka[i].eksemplar, &pustaka.pustaka[i].tahun, &pustaka.pustaka[i].rating)
	}
	pustaka.nPustaka = n
}

func CetakTerFavorit(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}

	maxRating := pustaka.pustaka[0].rating
	index := 0
	for i := 1; i < pustaka.nPustaka; i++ {
		if pustaka.pustaka[i].rating > maxRating {
			maxRating = pustaka.pustaka[i].rating
			index = i
		}
	}

	b := pustaka.pustaka[index]
	fmt.Println("Buku Terfavorit:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n", b.id, b.judul, b.penulis, b.penerbit, b.tahun, b.rating)
}

func UrutkanBuku(pustaka *DaftarBuku) {
	for i := 1; i < pustaka.nPustaka; i++ {
		key := pustaka.pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.pustaka[j].rating < key.rating {
			pustaka.pustaka[j+1] = pustaka.pustaka[j]
			j--
		}
		pustaka.pustaka[j+1] = key
	}
}

func Cetak5Teratas(pustaka DaftarBuku) {
	fmt.Println("Lima Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < pustaka.nPustaka; i++ {
		b := pustaka.pustaka[i]
		fmt.Printf("ID: %d, Judul: %s, Rating: %d\n", b.id, b.judul, b.rating)
	}
}

func CariBuku(pustaka DaftarBuku, r int) {
	low, high := 0, pustaka.nPustaka-1
	found := false
	for low <= high {
		mid := (low + high) / 2
		if pustaka.pustaka[mid].rating == r {
			b := pustaka.pustaka[mid]
			fmt.Printf("Buku Ditemukan: ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n", b.id, b.judul, b.penulis, b.penerbit, b.tahun, b.rating)
			found = true
			break
		} else if pustaka.pustaka[mid].rating < r {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating yang dicari.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)
	CetakTerFavorit(pustaka)
	UrutkanBuku(&pustaka)
	Cetak5Teratas(pustaka)

	var r int
	fmt.Print("Masukkan rating untuk mencari buku: ")
	fmt.Scan(&r)
	CariBuku(pustaka, r)
}
