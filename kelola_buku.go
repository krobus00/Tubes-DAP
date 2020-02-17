/*
	Kelompok 1
	- Michael Putera Wardana (1301194056)
	- Rizky Ahmad Saputra (1301194207)
	Program Aplikasi Kelola Stok Buku
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ID, judul, penerbit, tahun, stok.
type identitas struct {
	id       int
	judul    string
	penerbit string
	tahun    int
	stok     int
}

// Array untuk indentitas buku
var buku []identitas

// procedure untuk tambah buku
func add(id int, judul, penerbit string, tahun int, stok int) {
	a := identitas{id: id, judul: judul, penerbit: penerbit, tahun: tahun, stok: stok}
	buku = append(buku, a)
}

// function untuk sorting
func sorting(arr []identitas) []identitas {
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1].stok < arr[j].stok {
				t := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = t
			}
		}
	}
	return arr
}

// func untuk mencari data buku dengan id
func findbyID(id int) int {
	for i := 0; i < len(buku); i++ {
		if id == buku[i].id {
			return i
		}
	}
	return -1
}

// func untuk mencari data buku dengan judul
func findbyTitle(title string) int {
	for i := 0; i < len(buku); i++ {
		if title == buku[i].judul {
			return i
		}
	}
	return -1
}

// func untuk mencari data buku dengan penerbit
func findbyPublisher(publisher string) []int {
	var listID []int
	for i := 0; i < len(buku); i++ {
		if publisher == buku[i].penerbit {
			listID = append(listID, i)
		}
	}
	return listID
}

// validation untuk string
func inputStr(text string, input *string) {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print(text)
		scanner.Scan()
		*input = scanner.Text()
		if *input == "" {
			fmt.Println("Tidak boleh kosong!")
		} else {
			break
		}
	}
}

// validation untuk pertanyaan yes/no (y/n)
func validationQuestion(text string) bool {
	var answer string
	inputStr(text, &answer)
	for answer != "y" && answer != "Y" && answer != "n" && answer != "N" {
		inputStr(text, &answer)
	}
	if answer == "y" || answer == "Y" {
		return true
	}
	return false
}

// validation untuk Integer
func inputInt(text string, number *int) {
	var input string
	var err error
	for true {
		fmt.Print(text)
		fmt.Scanln(&input)
		*number, err = strconv.Atoi(input)
		if err == nil {
			break
		} else {
			fmt.Println("Silahkan masukan angka dengan benar!")
		}
	}
}

// Menampilkan data buku
func specificRead(kurang bool) []identitas {
	var newBuku []identitas
	for i := 0; i < len(buku); i++ {
		a := identitas{id: buku[i].id, judul: buku[i].judul, penerbit: buku[i].penerbit, tahun: buku[i].tahun, stok: buku[i].stok}
		if kurang {
			if buku[i].stok >= 7 && buku[i].stok <= 9 {
				newBuku = append(newBuku, a)
			}
		} else {
			newBuku = append(newBuku, a)
		}
	}
	if kurang {
		newBuku = sorting(newBuku)
	}
	return newBuku
}

// Menampilkan data buku
func readBuku(kurang bool) {
	no := 0
	dataBuku := specificRead(kurang)
	fmt.Printf("------------------------------------------------------------------------------------------------\n")
	fmt.Printf("| No | ID Buku |Judul Buku\t\t\t| Penerbit\t\t| Tahun Terbit\t| Stok |\n")
	fmt.Printf("------------------------------------------------------------------------------------------------\n")
	if len(dataBuku) > 0 {
		for i := 0; i < len(dataBuku); i++ {
			no++
			fmt.Printf("| %-2d |", no)
			fmt.Printf(" %-7d |", dataBuku[i].id)
			fmt.Printf(" %-30s |", dataBuku[i].judul)
			fmt.Printf(" %-21s |", dataBuku[i].penerbit)
			fmt.Printf(" %-4d\t\t|", dataBuku[i].tahun)
			fmt.Printf(" %-4d |\n", dataBuku[i].stok)
		}
	} else {
		fmt.Printf("%60s\n", "DATA TIDAK DITEMUKAN")
	}
	fmt.Printf("------------------------------------------------------------------------------------------------\n")
	if len(dataBuku) > 0 {
		if validationQuestion("(y/n) Ingin menambah stok ? ") {
			cariBuku()
		}
	} else {
		if validationQuestion("(y/n) Ingin menambah tambah buku ? ") {
			tambahBuku()
		}
	}
}

// menambah buku
func tambahBuku() {
	var id, tahun, stok int
	var judul, penerbit string
	for true {
		inputInt("Masukan ID buku : ", &id)
		for findbyID(id) >= 0 {
			fmt.Println("ID telah digunakan!")
			inputInt("Masukan ID buku : ", &id)
		}
		inputStr("Masukan Judul buku : ", &judul)
		for findbyTitle(judul) >= 0 {
			fmt.Println("Judul telah digunakan!")
			inputStr("Masukan Judul buku : ", &judul)
		}
		inputStr("Masukan penerbit : ", &penerbit)
		inputInt("Masukan tahun terbit : ", &tahun)
		inputInt("Masukan stok : ", &stok)
		fmt.Println("Detail buku :")
		fmt.Printf("ID : %v\n", id)
		fmt.Printf("Judul : %v\n", judul)
		fmt.Printf("Penerbit : %v\n", penerbit)
		fmt.Printf("Tahun terbit : %v\n", tahun)
		fmt.Printf("Stok : %v\n", stok)
		if validationQuestion("(y/n) Yakin menambahkan buku tersebut ? ") {
			add(id, judul, penerbit, tahun, stok)
		}
		if !validationQuestion("(y/n) Tambah buku lagi ? ") {
			break
		}

	}

}

// menambah stok buku
func tambahStok(id int) {
	var max int
	var tambahan int
	var tahun int = buku[id].tahun
	var stok int = buku[id].stok
	if tahun >= 2010 && tahun <= 2018 {
		if stok >= 3 && stok <= 6 {
			max = 7
		} else if stok >= 7 && stok <= 9 {
			max = 2
		} else if stok == 15 {
			max = -1
		} else {
			max = 0
		}
	} else if tahun > 2010 {
		if stok >= 3 && stok <= 6 {
			max = 5
		} else if stok >= 7 && stok <= 9 {
			max = 1
		} else if stok == 10 {
			max = -1
		} else {
			max = 0
		}
	} else {
		max = 0
	}
	if max > 0 {
		fmt.Printf("Hanya dapat menambahkan maksimal %d buah", max)
		inputInt("Tambah Sebanyak : ", &tambahan)
		for tambahan > max {
			inputInt("Tambah Sebanyak : ", &tambahan)
		}
		buku[id].stok += tambahan
	} else if max == 0 {
		inputInt("Tambah Sebanyak : ", &tambahan)
		for tambahan < 0 {
			inputInt("Tambah Sebanyak : ", &tambahan)
		}
		buku[id].stok += tambahan
	} else {
		fmt.Println("Tidak dapat menambah buku!")
	}
}

// mencari dan menampilkan identitas buku dengan menggunakan id buku
func cariBuku() {
	var id int
	inputInt("Masukan ID buku : ", &id)
	id = findbyID(id)
	if id >= 0 {
		fmt.Println("OK")
		fmt.Println("Detail buku :")
		fmt.Printf("ID : %v\n", buku[id].id)
		fmt.Printf("Judul : %v\n", buku[id].judul)
		fmt.Printf("Penerbit : %v\n", buku[id].penerbit)
		fmt.Printf("Tahun terbit : %v\n", buku[id].tahun)
		fmt.Printf("Stok : %v\n", buku[id].stok)
		if validationQuestion("(y/n) Ingin menambah stok buku ? ") {
			tambahStok(id)
		}
	} else {
		fmt.Println("ID tidak ditemukan!")
		if validationQuestion("(y/n) Ingin menambah buku ? ") {
			tambahBuku()
		}
	}
}

// mengganti semua penerbit yang sama (a menjadi b)
func gantiPenerbit() {
	var penerbit, penerbitBaru string
	var listID []int
	inputStr("Masukan nama penerbit : ", &penerbit)
	listID = findbyPublisher(penerbit)
	if len(listID) > 0 {
		fmt.Printf("Ditemukan %v buku dengan nama penerbit %v\n", len(findbyPublisher(penerbit)), penerbit)
		inputStr("Masukan nama baru dari pernerbit : ", &penerbitBaru)
		if validationQuestion("(y/n) Ganti nama penerbit ? ") {
			for i := 0; i < len(listID); i++ {
				buku[listID[i]].penerbit = penerbitBaru
			}
			fmt.Printf("Penerbit %s telah diganti menjadi %s", penerbit, penerbitBaru)
		}
	} else {
		fmt.Println("Buku dengan penerbit %s tidak ditemukan", penerbit)
	}

}

// menghapus buku dengan id tertentu
func hapusBuku() {
	var id int
	for true {
		var temp []identitas
		inputInt("Masukan ID buku : ", &id)
		id = findbyID(id)
		if id >= 0 {
			fmt.Println("Detail buku :")
			fmt.Printf("ID : %v\n", buku[id].id)
			fmt.Printf("Judul : %v\n", buku[id].judul)
			fmt.Printf("Penerbit : %v\n", buku[id].penerbit)
			fmt.Printf("Tahun terbit : %v\n", buku[id].tahun)
			fmt.Printf("Stok : %v\n", buku[id].stok)
			if validationQuestion("(y/n) Yakin menghapus buku tersebut ? ") {
				for i := 0; i < len(buku); i++ {
					if i != id {
						temp = append(temp, buku[i])
					}
				}
				buku = temp
			}
		} else {
			fmt.Println("Buku tidak")
		}
		if !validationQuestion("(y/n) Ingin hapus buku lagi ? ") {
			break
		}

	}
}

// menampilkan menu
func mainMenu() {
	var pilihan int
	var exit bool
	for true {
		fmt.Println("==============================================")
		fmt.Println("=                                            =")
		fmt.Println("= Selamat datang di aplikasi management buku =")
		fmt.Println("=                                            =")
		fmt.Println("==============================================")
		fmt.Println("Menu Utama")
		fmt.Println("1. Lihat semua daftar buku")
		fmt.Println("2. Lihat daftar buku yang kurang")
		fmt.Println("3. Tambah stok buku")
		fmt.Println("4. Tambah buku")
		fmt.Println("5. Ganti Penerbit")
		fmt.Println("6. Hapus buku")
		fmt.Println("7. Keluar")
		fmt.Println("")
		inputInt("Pilihan anda : ", &pilihan)
		switch pilihan {
		case 1:
			readBuku(false)
		case 2:
			readBuku(true)
		case 3:
			cariBuku()
		case 4:
			tambahBuku()
		case 5:
			gantiPenerbit()
		case 6:
			hapusBuku()
		case 7:
			exit = true
			break
		default:
			fmt.Println("Silahkan pilih menu dari 1-6")
		}
		if exit {
			fmt.Println("Terimakasih~")
			break
		}
	}
}

func main() {
	mainMenu()
}
