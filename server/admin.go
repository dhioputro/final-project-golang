package main

import (
	"bufio"
	"college-final-project/collegepb"
	"context"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/chilts/sid"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
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

var mhs = make(map[string]Mahasiswa)
var dosen = make(map[string]Dosen)

type Server struct {
}

func (s Server) DeleteMhs(ctx context.Context, request *collegepb.EdiMhsRequest) (*collegepb.ResultMhsResponse, error) {
	id := request.Id
	if _, found := mhs[id]; found {
		delete(mhs, id)

		response := &collegepb.ResultMhsResponse{
			Result: "mahasiswa id =" + id + " berhasil dihapus",
		}
		return response, nil
	} else {
		response := &collegepb.ResultMhsResponse{
			Result: "data tidak ditemukan",
		}
		return response, nil
	}
}

func (s Server) EditMhs(ctx context.Context, request *collegepb.EdiMhsRequest) (*collegepb.ResultMhsResponse, error) {
	id := request.Id
	if _, found := mhs[id]; found {
		namaBaru := request.Nama

		mhs[id] = Mahasiswa{
			Id:   id,
			Nama: namaBaru,
		}

		response := &collegepb.ResultMhsResponse{
			Result: "mahasiswa id =" + id + " berhasil ditambahkan",
		}
		return response, nil
	} else {
		response := &collegepb.ResultMhsResponse{
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

func (s Server) CreateMhs(ctx context.Context, request *collegepb.CreateMhsRequest) (*collegepb.ResultMhsResponse, error) {
	mhsCreate(request.Mahasiswa.Nama)

	response := &collegepb.ResultMhsResponse{
		Result: "Mahasiswa Sukses Ditambahkan",
	}

	log.Printf("Success : %v\n", response)

	return response, nil
}

func main() {
	mhsGenerator(5)

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
