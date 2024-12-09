SOAL 1
Fungsi insertionSort : 
Mengurutkan elemen array dalam urutan menaik menggunakan Insertion Sort
Mulai dari elemen kedua (indeks 1),Bandingkan elemen tersebut (key) dengan elemen sebelumnya,Geser elemen-elemen yang lebih besar dari key ke kanan hingga posisi yang benar ditemukan untuk key
Fungsi hitungjarakArr:
Mengecek apakah elemen-elemen array memiliki jarak yang tetap antar elemen
Jika array memiliki kurang dari 2 elemen, dianggap jaraknya tetap dengan nilai 0, Hitung jarak antara elemen pertama dan kedua (spacing),
Periksa apakah jarak antara elemen-elemen berikutnya sama dengan spacing,Jika jarak tidak sama, kembalikan false dan 0.


SOAL 2
Struct Buku dan DaftarBuku:
Buku menyimpan informasi tentang buku, DaftarBuku menyimpan array buku (pustaka) dan jumlah buku (nPustaka)
Fungsi-Fungsi:
DaftarkanBuku: Meminta pengguna memasukkan data buku satu per satu, Data yang dimasukkan berupa: id, judul, penulis, penerbit, eksemplar, tahun, dan rating, Data tersebut disimpan ke dalam array pustaka
CetakTerFavorit: dengan rating buku pertama sebagai nilai maksimum, Loop melalui seluruh buku untuk menemukan buku dengan rating tertinggi
UrutkanBuku: Mengurutkan array buku berdasarkan rating dari yang tertinggi ke terendah dengan insertion sort
Cetak5Teratas: Iterasi hingga lima buku atau hingga jumlah buku yang ada, Menampilkan ID, judul, dan rating dari buku tersebut
CariBuku: dari indeks awal (low) dan akhir (high) array
Cek elemen tengah:
Jika rating cocok, tampilkan data buku
Jika rating lebih besar, ganti pada bagian kanan array
Jika rating lebih kecil, ganti pada bagian kiri array
