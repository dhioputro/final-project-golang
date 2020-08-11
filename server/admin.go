package main

import (
	"bufio"
	"college-final-project/collegepb"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Pallinder/go-randomdata"
	"github.com/chilts/sid"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type Mahasiswa struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type Dosen struct {
	Id    string
	Nama  string
	Kelas string
}

type resumeMhs struct {
	idMhs   string
	namaMhs string
	kelas   [3]Kelas
}

type Kelas struct {
	idMhs   string  `json:"idMhs"`
	namaMhs string  `json:"namaMhs"`
	nilai   float32 `json:"nilai"`
	absen   int32   `json:"absen"`
}

var mhs = make(map[string]Mahasiswa)
var dosen = make(map[string]Dosen)
var kelas = make(map[string]Kelas)
var dataMhs = make(map[string]resumeMhs)
var kelasArr [3]Kelas

type Server struct {
}

func (s Server) GetAllDosen(e *empty.Empty, stream collegepb.AdminService_GetAllDosenServer) error {
	log.Printf("Get All Dosen Send to Client\n")
	for _, getDosen := range dosen {
		response := &collegepb.GetDosenResponse{
			Dosen: &collegepb.Dosen{
				Id:    getDosen.Id,
				Nama:  getDosen.Nama,
				Kelas: getDosen.Kelas,
			},
		}
		log.Printf("mengirim response : %v\n", response)
		_ = stream.Send(response)
	}
	return nil
}

func (s Server) CreateMhs(ctx context.Context, request *collegepb.CreateMhsRequest) (*collegepb.ResultResponse, error) {
	mhsCreate(request.Mahasiswa.Nama)

	response := &collegepb.ResultResponse{
		Result: "Mahasiswa Sukses Ditambahkan",
	}

	log.Printf("Success : %v\n", response)

	return response, nil
}

func (s Server) EditMhs(ctx context.Context, request *collegepb.EdiMhsRequest) (*collegepb.ResultResponse, error) {
	id := request.Id
	if _, found := mhs[id]; found {
		namaBaru := request.Nama

		mhs[id] = Mahasiswa{
			Id:   id,
			Nama: namaBaru,
		}

		response := &collegepb.ResultResponse{
			Result: "mahasiswa id =" + id + " berhasil ditambahkan",
		}
		return response, nil
	} else {
		response := &collegepb.ResultResponse{
			Result: "data tidak ditemukan",
		}
		return response, nil
	}
}

func (s Server) UpdateDosen(ctx context.Context, request *collegepb.EditDosenRequest) (*collegepb.ResultResponse, error) {
	id := request.Id
	if _, found := dosen[id]; found {
		namaBaru := request.Nama

		dosen[id] = Dosen{
			Id:   id,
			Nama: namaBaru,
		}
		response := &collegepb.ResultResponse{
			Result: "dosen id =" + id + " berhasil diubah",
		}
		return response, nil
	}
	response := &collegepb.ResultResponse{
		Result: "dosen id =" + id + " tidak ditemukan",
	}

	return response, nil
}

func (s Server) UpdateKelas(ctx context.Context, request *collegepb.EditKelasRequest) (*collegepb.ResultResponse, error) {
	id := request.Id
	if _, found := dosen[id]; found {
		kelasBaru := request.Kelas

		dosen[id] = Dosen{
			Id:    id,
			Kelas: kelasBaru,
		}
		response := &collegepb.ResultResponse{
			Result: "kelas id =" + id + " berhasil diubah",
		}
		return response, nil
	}
	return nil, nil
}

func (s Server) DeleteMhs(ctx context.Context, request *collegepb.DeleteMhsRequest) (*collegepb.ResultResponse, error) {
	id := request.Id
	if _, found := mhs[id]; found {
		delete(mhs, id)

		response := &collegepb.ResultResponse{
			Result: "mahasiswa id =" + id + " berhasil dihapus",
		}
		return response, nil
	} else {
		response := &collegepb.ResultResponse{
			Result: "data tidak ditemukan",
		}
		return response, nil
	}
}

func (s Server) GetAllMhs(empty *empty.Empty, stream collegepb.AdminService_GetAllMhsServer) error {
	log.Printf("Get All Mhs Send to Client\n")
	for _, getMahasiswa := range mhs {
		response := &collegepb.GetMhsResponse{
			Mhs: &collegepb.Mahasiswa{
				Id:   getMahasiswa.Id,
				Nama: getMahasiswa.Nama,
			},
		}
		log.Printf("mengirim response : %v\n", response)
		_ = stream.Send(response)
	}
	return nil
}

// ================================================================================== //

func (s Server) UpdateDataMhs(ctx context.Context, request *collegepb.InputDataReq) (*collegepb.ResultResponse, error) {
	id := request.Id

	if _, found := mhs[id]; found {
		dataMhs := findMahasiswa(id)

		response := &collegepb.ResultResponse{
			Result: inputNilaiMhs(request, dataMhs),
		}
		return response, nil
	}
	response := &collegepb.ResultResponse{
		Result: "dosen id =" + id + " tidak ditemukan",
	}

	return response, nil
}

// ================================================================================== //

func main() {
	//mhsGenerator(5)
	kelas := []string{"algoritma", "dasar pemograman", "aljabar"}
	dosenGenerator(kelas)

	fmt.Println("Go gRPC Beginners Tutorial!")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := Server{}
	grpcServer := grpc.NewServer()
	collegepb.RegisterAdminServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	// ========== TODO CRUD MAHASISWA ==========
	//fmt.Println("======== Crud Mahasiswa ========")
	//mhsGenerator(5)
	//mhsDemo()
	//
	//// ========== TODO UPDATE DOSEN ==========
	//fmt.Println("======== Update Dosen ========")
	//kelas := []string{"algoritma", "dasar pemograman", "aljabar"}
	//dosenGenerator(kelas)
	//dosenDemo()
	//
	//
	//// ========== TODO UPDATE KELAS ==========
	//fmt.Println("======== Update Kelas ========")
	//kelasEdit()
	//dosenRead()
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
	mhsCreate("andi")
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

func mhsCreate(namaBaru string) {
	log.Println("membuat data mahasiswa baru")
	idBaru := sid.Id()

	// simpan mahasiswa
	mhs[idBaru] = Mahasiswa{
		Id:   idBaru,
		Nama: namaBaru,
	}
}

func mhsGenerator(jumlahMhs int) {
	for i := 0; i < jumlahMhs; i++ {
		id := sid.Id()
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

//========================================================================//

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

func inputNilaiMhs(request *collegepb.InputDataReq, data resumeMhs) string {
	var result string
	id := data.idMhs
	nama := data.namaMhs
	subj := request.Kelas
	inputA := request.Absen
	inputN := request.Nilai

	switch subj {
	case "algoritma":
		kelas[subj] = Kelas{
			idMhs:   id,
			namaMhs: nama,
			absen:   inputA,
			nilai:   inputN,
		}

		// hitung nilai total
		// bobotAbsen := (float32(inputA) / 14) * 30
		// bobotNilai := inputN * 70 / 100
		// totalNilai := bobotAbsen + bobotNilai

		// fmt.Println(bobotAbsen)
		// fmt.Println(bobotNilai)
		// fmt.Println(totalNilai)

		kelasArr[0] = kelas[subj]
		data = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
			kelas:   kelasArr,
		}

		dataMhs[subj] = data
		fmt.Printf("Nilai Algoritma %s, telah di input dengan nilai %g\n", nama, inputN)
		result = "Nilai dan absen" + subj + " " + nama + ", sukses di input!"

	case "dasar pemrograman":
		kelas[subj] = Kelas{
			idMhs:   id,
			namaMhs: nama,
			absen:   inputA,
			nilai:   inputN,
		}

		// input ke data mahasiswa
		kelasArr[1] = kelas[subj]
		data = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
			kelas:   kelasArr,
		}

		// masukkan ke map dataMhs
		dataMhs[id] = data
		fmt.Printf("Nilai Algoritma %s, telah di input dengan nilai %g\n", nama, inputN)
		result = "Nilai dan absen" + subj + " " + nama + ", sukses di input!"

	case "aljabar":
		kelas[subj] = Kelas{
			idMhs:   id,
			namaMhs: nama,
			absen:   inputA,
			nilai:   inputN,
		}

		// input ke data mahasiswa
		kelasArr[2] = kelas[subj]
		data = resumeMhs{
			idMhs:   id,
			namaMhs: nama,
			kelas:   kelasArr,
		}

		// masukkan ke map dataMhs
		dataMhs[id] = data
		fmt.Printf("Nilai Algoritma %s, telah di input dengan nilai %g\n", nama, inputN)
		result = "Nilai dan absen" + subj + " " + nama + ", sukses di input!"
	default:
		fmt.Println("Pilihan yang anda masukkan tidak valid")
		fmt.Println()
		break
	}

	return result
}
