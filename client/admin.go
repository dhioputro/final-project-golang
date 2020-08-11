package main

import (
	"college-final-project/collegepb"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := collegepb.NewAdminServiceClient(conn)

	//demoMhs(c)
	getAllDosen(c)
	editDosen(c)
	editKelas(c)
}

func getAllDosen(c collegepb.AdminServiceClient) {
	log.Println("Get All Data Dosen")
	resStream, err := c.GetAllDosen(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatalf("Error when calling MhsGetAll: %s", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// untuk mengakhiri end of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while getting message: %s", err)
		}
		log.Printf("log msg.GetMhs() :  %v", msg.GetDosen())
	}
}

func editKelas(c collegepb.AdminServiceClient) {
	log.Println("Edit Data Kelas")
	id := ScanString("masukan id yang diubah")
	kelas := ScanString("masukan nama kelas baru")
	request := &collegepb.EditKelasRequest{
		Id:    id,
		Kelas: kelas,
	}
	response, err := c.UpdateKelas(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling editDosen: %s", err)
	}
	log.Printf("Response from server: %v\n", response)
}

func editDosen(c collegepb.AdminServiceClient) {
	log.Println("Edit Data Dosen")
	id := ScanString("masukan id yang diubah")
	nama := ScanString("masukan nama baru dosen")
	request := &collegepb.EditDosenRequest{
		Id:   id,
		Nama: nama,
	}
	response, err := c.UpdateDosen(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling editDosen: %s", err)
	}
	log.Printf("Response from server: %v\n", response)

}

func demoMhs(c collegepb.AdminServiceClient) {
	MhsCreate(c)
	log.Println()
	MhsGetAll(c)
	log.Println()
	MhsEdit(c)
	log.Println()
	MhsGetAll(c)
	log.Println()
	MhsDelete(c)
	log.Println()
	MhsGetAll(c)
}

func MhsDelete(c collegepb.AdminServiceClient) {
	log.Println("Delete Data Mahasiswa")
	idBaru := ScanString("masukan id yang dihapus")
	request := &collegepb.DeleteMhsRequest{
		Id: idBaru,
	}
	response, err := c.DeleteMhs(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling MhsEdit: %s", err)
	}
	log.Printf("Response from server: %v\n", response)
}

func MhsEdit(c collegepb.AdminServiceClient) {
	log.Println("Edit Data Mahasiswa")
	idBaru := ScanString("masukan id yang diedit")
	namaBaru := ScanString("masukan id yang diedit")

	request := &collegepb.EdiMhsRequest{
		Id:   idBaru,
		Nama: namaBaru,
	}
	response, err := c.EditMhs(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling MhsEdit: %s", err)
	}
	log.Printf("Response from server: %v\n", response)
}

func MhsGetAll(c collegepb.AdminServiceClient) {
	log.Println("Get All Data Mahasiswa")
	resStream, err := c.GetAllMhs(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatalf("Error when calling MhsGetAll: %s", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// untuk mengakhiri end of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while getting message: %s", err)
		}
		log.Printf("log msg.GetMhs() :  %v", msg.GetMhs())
	}
}

func MhsCreate(c collegepb.AdminServiceClient) {
	log.Println("Membuat Data Mahasiswa Baru")
	var mhs = collegepb.Mahasiswa{
		Nama: "udin",
	}
	request := &collegepb.CreateMhsRequest{
		Mahasiswa: &mhs,
	}
	response, err := c.CreateMhs(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling Create Mahasiswa: %s", err)
	}
	// response bisa dengan response.Success
	log.Printf("Response from server: %v\n", response)
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
	var pilihan string
	fmt.Printf("%s : ", text)
	_, err := fmt.Scanf("%s\n", &pilihan)

	if err != nil {
		fmt.Println(err)
		return "inputan kosong" // debug mode
	}

	return pilihan
}
