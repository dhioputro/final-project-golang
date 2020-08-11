package main

import (
	"college-final-project/collegepb"
	pb "college-final-project/collegepb"
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func dosenMain() {
	// var login = false
	var opt int
	c := ConnectServer()

	// loginUser(login)

	for opt != 99 {
		menu(
			"Lihat Daftar Mahasiswa",
			"Input Data Mahasiswa",
			"Show Data Kelas",
		)
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			fmt.Println("=========== Data Mahasiswa =========")
			listMahasiswa(c)
			fmt.Println("=====================================")
		case 2:
			inputDataMhs(c)
		case 3:
			var subj string
			fmt.Println("Silahkan masukkan Kelas: ")
			fmt.Scanln(&subj)

		case 99:
			fmt.Println("Sampai jumpa!")
		default:
			fmt.Println("Pilihan yang anda masukkan tidak valid")
			fmt.Println()
			break
		}

	}
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

func listMahasiswa(c pb.AdminServiceClient) {
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

func inputDataMhs(c pb.AdminServiceClient) {
	var id, subj string
	var inputA int32
	var inputN float32
	fmt.Println("Silahkan masukkan ID mahasiswa:")
	fmt.Scanln(&id)
	fmt.Println("Silahkan pilih data yang ingin di input (algoritma/dasar pemograman/aljabar)")
	fmt.Scanln(&subj)
	fmt.Println("Silahkan input Absen Kelas", subj)
	fmt.Scanln(&inputA)
	fmt.Println("Silahkan input Nilai", subj)
	fmt.Scanln(&inputN)

	request := &collegepb.InputDataReq{
		Id:    id,
		Kelas: subj,
		Absen: inputA,
		Nilai: inputN,
	}
	response, err := c.UpdateDataMhs(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when Input Data Mahasiswa: %s", err)
	}
	log.Printf("Response from server: %v\n", response)
}

// connection setup
func ConnectServer() pb.AdminServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("it fails to connect: %v", err)
	}
	defer conn.Close()
	return pb.NewAdminServiceClient(conn)
}
