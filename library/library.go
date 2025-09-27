package library

import (
	"fmt"
	"perpusmini/book"
)

type Library struct {
    Books []book.Book
}

// CREATE → Tambah data buku baru ke dalam slice
func (l *Library) TambahBuku(judul, penulis string, tahun int) {
    id := len(l.Books) + 1 // ID otomatis naik
    buku := book.Book{ID: id, Judul: judul, Penulis: penulis, Tahun: tahun, Dipinjam: false}
    l.Books = append(l.Books, buku)
    fmt.Println("Buku berhasil ditambahkan!")
}

// READ → Menampilkan semua buku yang ada
func (l *Library) DaftarBuku() {
    if len(l.Books) == 0 {
        fmt.Println("Belum ada buku.")
        return
    }
    fmt.Println("Daftar Buku:")
    for _, b := range l.Books {
        status := "Tersedia"
        if b.Dipinjam {
            status = "Dipinjam"
        }
        fmt.Printf("[%d] %s - %s (%d) [%s]\n", b.ID, b.Judul, b.Penulis, b.Tahun, status)
    }
}

// READ (khusus) → Mencari buku berdasarkan judul
func (l *Library) CariBuku(judul string) {
    ketemu := false
    for _, b := range l.Books {
        if b.Judul == judul {
            status := "Tersedia"
            if b.Dipinjam {
                status = "Dipinjam"
            }
            fmt.Printf("Ditemukan: %s - %s (%d) [%s]\n", b.Judul, b.Penulis, b.Tahun, status)
            ketemu = true
        }
    }
    if !ketemu {
        fmt.Println("Buku tidak ditemukan.")
    }
}

// UPDATE → Mengubah status buku menjadi "dipinjam"
func (l *Library) PinjamBuku(id int) {
    for i, b := range l.Books {
        if b.ID == id {
            if b.Dipinjam {
                fmt.Println("Buku sudah dipinjam.")
                return
            }
            l.Books[i].Dipinjam = true
            fmt.Println("Buku berhasil dipinjam!")
            return
        }
    }
    fmt.Println("ID buku tidak ditemukan.")
}

// UPDATE → Mengubah status buku menjadi "tersedia" (dikembalikan)
func (l *Library) KembalikanBuku(id int) {
    for i, b := range l.Books {
        if b.ID == id {
            if !b.Dipinjam {
                fmt.Println("Buku ini belum dipinjam.")
                return
            }
            l.Books[i].Dipinjam = false
            fmt.Println("Buku berhasil dikembalikan!")
            return
        }
    }
    fmt.Println("ID buku tidak ditemukan.")
}

// DELETE (opsional) → Menghapus buku dari daftar
// Tidak wajib, tapi bisa ditambah agar CRUD lengkap
func (l *Library) HapusBuku(id int) {
    for i, b := range l.Books {
        if b.ID == id {
            l.Books = append(l.Books[:i], l.Books[i+1:]...) // hapus item slice
            fmt.Println("Buku berhasil dihapus!")
            return
        }
    }
    fmt.Println("ID buku tidak ditemukan.")
}
