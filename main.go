package main

import (
	"fmt"
	"perpusmini/library"
)

func main() {
    lib := library.Library{}
    var pilihan int

    for {
        fmt.Println("\n=== APLIKASI PERPUSTAKAAN MINI ===")
        fmt.Println("1. Tambah Buku (CREATE)")
        fmt.Println("2. Lihat Semua Buku (READ)")
        fmt.Println("3. Cari Buku (READ)")
        fmt.Println("4. Pinjam Buku (UPDATE)")
        fmt.Println("5. Kembalikan Buku (UPDATE)")
        fmt.Println("6. Hapus Buku (DELETE)")
        fmt.Println("0. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilihan)

        switch pilihan {
        case 1:
            var judul, penulis string
            var tahun int
            fmt.Print("Masukkan Judul: ")
            fmt.Scan(&judul)
            fmt.Print("Masukkan Penulis: ")
            fmt.Scan(&penulis)
            fmt.Print("Masukkan Tahun: ")
            fmt.Scan(&tahun)
            lib.TambahBuku(judul, penulis, tahun)

        case 2:
            lib.DaftarBuku()

        case 3:
            var judul string
            fmt.Print("Masukkan Judul Buku: ")
            fmt.Scan(&judul)
            lib.CariBuku(judul)

        case 4:
            var id int
            fmt.Print("Masukkan ID Buku: ")
            fmt.Scan(&id)
            lib.PinjamBuku(id)

        case 5:
            var id int
            fmt.Print("Masukkan ID Buku: ")
            fmt.Scan(&id)
            lib.KembalikanBuku(id)

        case 6:
            var id int
            fmt.Print("Masukkan ID Buku: ")
            fmt.Scan(&id)
            lib.HapusBuku(id)

        case 0:
            fmt.Println("Terima kasih! Keluar...")
            return

        default:
            fmt.Println("Pilihan tidak valid.")
        }
    }
}