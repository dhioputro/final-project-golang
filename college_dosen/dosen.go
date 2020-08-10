package main

import (
	"fmt"
	"sort"
	"strconv"

	// "college-final-project"

	"gopkg.in/validator.v2" //>> validator untuk regex
)

type dosenLogin struct {
	Username string `validate:"min=3,max=10,regexp=^[a-zA-Z0-9._-]*$"`
	Password string `validate:"min=8,max=15,regexp=^[a-zA-Z0-9._-]*$"`
}

type Mahasiswa struct {
	Id   string
	Nama string
}

type Dosen struct {
	Id       string
	Nama     string
	Kelas    string
	Username string
	Password string
}

type resumeMhs struct {
	idMhs   string
	namaMhs string
	algo    float64
	js      float64
	golang  float64
}

type kelasAlgo struct {
	namaMhs   string
	nilaiAlgo float64
	absenAlgo int
}
type kelasJS struct {
	namaMhs string
	nilaiJS float64
	absenJS int
}
type kelasGo struct {
	namaMhs string
	nilaiGo float64
	absenGo int
}

var mhs = make(map[string]Mahasiswa)
var dosen = make(map[string]Dosen)
var algo = make(map[string]kelasAlgo)
var js = make(map[string]kelasJS)
var golang = make(map[string]kelasGo)
var dataMhs = make(map[string]resumeMhs)

func main() {
	var login = false
	var opt int

	// demo isi mahasiswa
	mhs["1"] = Mahasiswa{Id: "1", Nama: "Dhio"}
	mhs["2"] = Mahasiswa{Id: "2", Nama: "Andi"}
	mhs["3"] = Mahasiswa{Id: "3", Nama: "Pawit"}

	// demo isi dosen
	dosen["wawan01"] = Dosen{Id: "1", Nama: "Wawan", Kelas: "Golang", Password: "dosenku123"}
	dosen["harry02"] = Dosen{Id: "2", Nama: "Harry", Kelas: "JS", Password: "dosenku124"}
	dosen["ferry03"] = Dosen{Id: "3", Nama: "Ferry", Kelas: "Algo", Password: "dosenku125"}

	loginUser(login)

	for opt != 99 {
		menu(
			"Lihat Daftar Mahasiswa",
			"Input Data Mahasiswa",
			"Show Data",
		)
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			fmt.Println("=========== Data Mahasiswa =========")
			listMahasiswa(mhs)
			fmt.Println("=====================================")
		case 2:
			var id string
			fmt.Println("Silahkan masukkan ID mahasiswa:")
			fmt.Scanln(&id)
			found := findMahasiswa(id)
			fmt.Printf("ID %s ditemukan dengan nama %s\n", id, found.namaMhs)
			inputNilaiMhs(found)

		case 3:
			fmt.Println("=====================================")
			fmt.Println("ID\t Nama\t Algoritma\t JS\t Golang\t")

			for _, val := range dataMhs {
				fmt.Printf("%s\t %s\t %.2f\t\t %.2f\t %.2f\t\n", val.idMhs, val.namaMhs, val.algo, val.js, val.golang)
			}
			fmt.Println("=====================================")
			fmt.Println()

		case 99:
			fmt.Println("Sampai jumpa!")
		default:
			fmt.Println("Pilihan yang anda masukkan tidak valid")
			fmt.Println()
			break
		}

	}
}

func loginUser(login bool) bool {
	for login != true {
		var user, pass string
		fmt.Println("Username :")
		fmt.Scanln(&user)
		fmt.Println("Password :")
		fmt.Scanln(&pass)

		log := dosenLogin{Username: user, Password: pass}
		if err := validator.Validate(log); err != nil {
			fmt.Println("Username harus berisi 3-10 character, terdiri dari alphanumeric")
			fmt.Println("Password harus berisi 8-15 character, terdiri dari alphanumeric")
			fmt.Println(err)

		} else {
			for getUsername, getDosen := range dosen {
				// fmt.Println(getUsername)
				// fmt.Println(getDosen)
				if getUsername == user && getDosen.Password == pass {
					login = true
					fmt.Println("==============================")
					fmt.Println("Hi", getDosen.Nama, ", Selamat Datang!")
					fmt.Println("==============================")
					break
				} else {
					login = false
				}
			}
			if login == false {
				fmt.Println("==============================")
				fmt.Println("Maaf, username tidak terdaftar atau password salah")
				fmt.Println("==============================")
			}
		}
	}
	return login
}

func menu(menu ...string) {
	fmt.Println("===== Menu Dosen =====")
	for opt, val := range menu {
		a := strconv.Itoa(opt + 1)
		fmt.Println(a + ". " + val)
	}
	fmt.Println("99. Selesai")
	fmt.Println("Masukkan pilihan anda :")
}

func listMahasiswa(listMhs map[string]Mahasiswa) {
	keys := make([]string, 0, len(listMhs))
	for k := range listMhs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("NIM \t| Nama")
	for _, k := range keys {
		fmt.Println(k, "\t ", listMhs[k].Nama)
	}
	fmt.Println()
}

func findMahasiswa(id string) resumeMhs {
	if _, getMhs := mhs[id]; getMhs {
		nama := mhs[id].Nama
		dataMhs[id] = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
		}
	}
	return dataMhs[id]
}

func inputNilaiMhs(data resumeMhs) resumeMhs {
	var subj string
	var inputN float64
	var inputA int
	fmt.Println("Silahkan pilih data yang ingin di input (algoritma/golang/js)")
	fmt.Scanln(&subj)

	id := data.idMhs
	nama := data.namaMhs

	fmt.Println("Silahkan input Absen Kelas", subj)
	fmt.Scanln(&inputA)
	fmt.Println("Silahkan input Nilai", subj)
	fmt.Scanln(&inputN)

	switch subj {
	case "algoritma":
		algo[id] = kelasAlgo{
			namaMhs:   nama,
			absenAlgo: inputA,
			nilaiAlgo: inputN,
		}

		// hitung nilai total
		bobotAbsen := (float64(inputA) / 14) * 30
		bobotNilai := inputN * 70 / 100
		totalNilai := bobotAbsen + bobotNilai

		fmt.Println(bobotAbsen)
		fmt.Println(bobotNilai)
		fmt.Println(totalNilai)

		// input ke data mahasiswa
		data = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
			algo:    totalNilai,
			js:      js[id].nilaiJS,
			golang:  golang[id].nilaiGo,
		}

		// masukkan ke map dataMhs
		algo := algo[id].nilaiAlgo
		dataMhs[id] = data
		fmt.Printf("Nilai Algoritma %s, telah di input dengan nilai %g\n", nama, algo)

	case "golang":
		golang[id] = kelasGo{
			namaMhs: nama,
			absenGo: inputA,
			nilaiGo: inputN,
		}

		// hitung nilai total
		bobotAbsen := (float64(inputA) / 14) * 30
		bobotNilai := inputN * 70 / 100
		totalNilai := bobotAbsen + bobotNilai

		// input ke data mahasiswa

		fmt.Println(algo[id])
		fmt.Println(golang[id])
		fmt.Println(dataMhs[id])
		data = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
			algo:    algo[id].nilaiAlgo,
			js:      js[id].nilaiJS,
			golang:  totalNilai,
		}

		// masukkan ke map dataMhs
		golang := golang[id].nilaiGo
		dataMhs[id] = data
		fmt.Printf("Nilai Golang %s, telah di input dengan nilai %g\n", nama, golang)

	case "js":
		js[id] = kelasJS{
			namaMhs: nama,
			absenJS: inputA,
			nilaiJS: inputN,
		}

		// hitung nilai total
		bobotAbsen := (float64(inputA) / 14) * 30
		bobotNilai := inputN * 70 / 100
		totalNilai := bobotAbsen + bobotNilai

		// input ke data mahasiswa
		data = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
			algo:    algo[id].nilaiAlgo,
			js:      totalNilai,
			golang:  golang[id].nilaiGo,
		}

		// masukkan ke map dataMhs
		js := js[id].nilaiJS
		dataMhs[id] = data
		fmt.Printf("Nilai Algoritma %s, telah di input dengan nilai %g\n", nama, js)

	default:
		fmt.Println("Pilihan yang anda masukkan tidak valid")
		fmt.Println()
		break
	}
	return data
}
