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
	idMhs      string
	namaMhs    string
	algo       float64
	js         float64
	golang     float64
	nilaiTotal float64
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
	// var login = false
	var opt int

	// demo isi mahasiswa
	mhs["1"] = Mahasiswa{Id: "1", Nama: "Dhio"}
	mhs["2"] = Mahasiswa{Id: "2", Nama: "Andi"}
	mhs["3"] = Mahasiswa{Id: "3", Nama: "Pawit"}

	//demo isi dosen
	dosen["wawan01"] = Dosen{Id: "1", Nama: "Wawan", Kelas: "Golang", Password: "dosenku123"}
	dosen["harry02"] = Dosen{Id: "2", Nama: "Harry", Kelas: "JS", Password: "dosenku124"}
	dosen["ferry03"] = Dosen{Id: "3", Nama: "Ferry", Kelas: "Algo", Password: "dosenku125"}

	// loginUser(login)

	for opt != 99 {
		menu(
			"Lihat Daftar Mahasiswa",
			"Input Absen",
			"Input Nilai",
			"Delete Nilai",
		)
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			fmt.Println("=========== Data Mahasiswa =========")
			listMahasiswa(mhs)

		case 2:
			var id string
			fmt.Println("Silahkan masukkan ID mahasiswa:")
			fmt.Scanln(&id)
			found := findMahasiswa(id)
			inputAbsenMhs(found)

		case 3:
			var id string
			fmt.Println("Silahkan masukkan ID mahasiswa:")
			fmt.Scanln(&id)
			found := findMahasiswa(id)
			inputNilaiMhs(found)

		case 4:
			fmt.Println(dataMhs)
			fmt.Println(algo)

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

func inputAbsenMhs(dataMhs resumeMhs) resumeMhs {
	var subj string
	var input int
	fmt.Println("Silahkan masukkan absen kelas yang ingin di tambahkan (algoritma/golang/js)")
	fmt.Scanln(&subj)

	id := dataMhs.idMhs
	nama := dataMhs.namaMhs

	switch subj {
	case "algoritma":
		fmt.Println("Silahkan input Absen kelas Algoritma")
		fmt.Scanln(&input)
		absen := algo[id].absenAlgo + input
		algo[id] = kelasAlgo{
			namaMhs:   nama,
			absenAlgo: absen,
			nilaiAlgo: algo[id].nilaiAlgo,
		}

		fmt.Printf("Absen kelas Algoritma %s, telah di update dengan jumlah %d\n", nama, algo[id].absenAlgo)
		return dataMhs

	case "golang":
		fmt.Println("Silahkan input Absen kelas Golang")
		fmt.Scanln(&input)
		absen := golang[id].absenGo + input
		golang[id] = kelasGo{
			namaMhs: nama,
			absenGo: absen,
			nilaiGo: golang[id].nilaiGo,
		}

		fmt.Printf("Absen kelas Golang %s, telah di update dengan jumlah %d\n", nama, golang[id].absenGo)
		return dataMhs

	case "js":
		fmt.Println("Silahkan input Nilai JS")
		fmt.Scanln(&input)
		absen := js[id].absenJS + input
		js[id] = kelasJS{
			namaMhs: nama,
			absenJS: absen,
			nilaiJS: js[id].nilaiJS,
		}

		fmt.Printf("Nilai Algoritma %s, telah di input dengan nilai %d\n", nama, js[id].absenJS)
		return dataMhs

	default:
		fmt.Println("Pilihan yang anda masukkan tidak valid")
		fmt.Println()
		break
	}
	return dataMhs
}

func inputNilaiMhs(dataMhs resumeMhs) resumeMhs {
	var subj string
	var input float64
	fmt.Println("Silahkan masukkan nilai yang ingin di input (algoritma/golang/js)")
	fmt.Scanln(&subj)

	id := dataMhs.idMhs
	nama := dataMhs.namaMhs

	switch subj {
	case "algoritma":
		fmt.Println("Silahkan input Nilai Algoritma")
		fmt.Scanln(&input)
		algo[id] = kelasAlgo{
			namaMhs:   nama,
			absenAlgo: algo[id].absenAlgo,
			nilaiAlgo: input,
		}

		// input ke data mahasiswa
		algo := algo[id].nilaiAlgo
		dataMhs = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
			algo:    algo,
		}

		fmt.Printf("Nilai Algoritma %s, telah di input dengan nilai %g\n", nama, algo)
		return dataMhs

	case "golang":
		fmt.Println("Silahkan input Nilai Golang")
		fmt.Scanln(&input)
		golang[id] = kelasGo{
			namaMhs: nama,
			absenGo: golang[id].absenGo,
			nilaiGo: input,
		}
		// input ke data mahasiswa
		golang := golang[id].nilaiGo
		dataMhs = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
			golang:  golang,
		}
		fmt.Printf("Nilai Golang %s, telah di input dengan nilai %g\n", nama, golang)
		return dataMhs

	case "js":
		fmt.Println("Silahkan input Nilai JS")
		fmt.Scanln(&input)
		js[id] = kelasJS{
			namaMhs: nama,
			absenJS: js[id].absenJS,
			nilaiJS: input,
		}

		// input ke data mahasiswa
		js := js[id].nilaiJS
		dataMhs = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
			js:      js,
		}
		fmt.Printf("Nilai Algoritma %s, telah di input dengan nilai %g\n", nama, js)
		return dataMhs

	default:
		fmt.Println("Pilihan yang anda masukkan tidak valid")
		fmt.Println()
		break
	}
	return dataMhs
}
