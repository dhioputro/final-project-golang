package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	_ "github.com/Pallinder/go-randomdata"
	"github.com/chilts/sid"
)

type Mahasiswa struct {
	Id   string
	Nama string
}

type Dosen struct {
	Id    string
	Nama  string
	Kelas string
}

var mhs = make(map[string]Mahasiswa)
var dosen = make(map[string]Dosen)

func main() {
	// ========== TODO CRUD MAHASISWA ==========
	mhsGenerator(5)
	mhsDemo()

	// ========== TODO UPDATE DOSEN ==========
	// kelas := []string{"algoritma", "dasar pemograman", "aljabar"}
	// dosenGenerator(kelas)
	// dosenDemo()

	// ========== TODO UPDATE KELAS ==========
	// kelasEdit()
	// dosenRead()
}

func kelasEdit() {
	cariId := ScanString("input berdasarkan id dosen diedit")
	if _, found := dosen[cariId]; found {
		kelasBaru := ScanString("input nama kelas baru")

		// save to map
		dosen[cariId] = Dosen{
			Id:    dosen[cariId].Id,
			Nama:  dosen[cariId].Nama,
			Kelas: kelasBaru,
		}
	} else {
		fmt.Println("id dosen tidak ditemukan")
	}
}

func dosenDemo() {
	dosenRead()

	dosenEdit()
	dosenRead()
}

func dosenEdit() {
	cariId := ScanString("input id dosen yang diedit")
	if _, found := dosen[cariId]; found {
		namaBaru := ScanString("input nama dosen baru")

		// save to map
		dosen[cariId] = Dosen{
			Id:    dosen[cariId].Id,
			Nama:  namaBaru,
			Kelas: dosen[cariId].Kelas,
		}
	} else {
		fmt.Println("id dosen tidak ditemukan")
	}
}

func dosenRead() {
	for _, getDosen := range dosen {
		fmt.Printf("id: %s --- nama: %s --- kelas: %s\n", getDosen.Id, getDosen.Nama, getDosen.Kelas)
	}
}

func mhsDemo() {
	mhsCreate()
	mhsRead()

	mhsEdit()
	mhsRead()

	mhsDelete()
	mhsRead()
}

func dosenGenerator(kelas []string) {
	for i := range kelas {
		id := sid.Id()
		nama := randomdata.FullName(randomdata.RandomGender)
		dosen[id] = Dosen{
			Id:    id,
			Nama:  nama,
			Kelas: kelas[i],
		}
	}
}

func mhsRead() {
	for _, getMahasiswa := range mhs {
		fmt.Printf("id: %s --- nama: %s\n", getMahasiswa.Id, getMahasiswa.Nama)
	}
}

func mhsDelete() {
	idDelete := ScanString("input id mahasiwa yang dihapus")
	if _, found := mhs[idDelete]; found {
		delete(mhs, idDelete)
	} else {
		fmt.Println("id mahasiswa tidak ditemukan")
	}
}

func mhsEdit() {
	cariId := ScanString("input id yang diedit")
	if _, found := mhs[cariId]; found {
		namaBaru := ScanString("input nama baru")

		// save to map
		mhs[cariId] = Mahasiswa{
			Id:   cariId,
			Nama: namaBaru,
		}
	} else {
		fmt.Println("id mahasiswa tidak ditemukan")
	}
}

func mhsCreate() {
	fmt.Println("membuat data mahasiswa baru")
	idBaru := sid.Id()
	namaBaru := ScanString("input nama baru")

	// simpan mahasiswa
	mhs[idBaru] = Mahasiswa{
		Id:   idBaru,
		Nama: namaBaru,
	}
}

func mhsGenerator(jumlahMhs int) {
	for i := 0; i < jumlahMhs; i++ {
		id := strconv.Itoa(i)
		nama := randomdata.FullName(randomdata.RandomGender)
		mhs[id] = Mahasiswa{
			Id:   id,
			Nama: nama,
		}
	}
}

func ScanPilihan() int {
	var pilihan int
	fmt.Print("Pilihan Anda : ")
	_, err := fmt.Scanf("%d", &pilihan)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return pilihan
}

func ScanString(text string) string {
	fmt.Printf("%s : ", text)
	scanner := bufio.NewScanner(os.Stdin)
	line := scanner.Text()

	return line
}
