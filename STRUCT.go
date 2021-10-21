package main

import "fmt"

//Membuat struct dengan nama identitas
//type digunakan untuk deklarasi
type Identitas struct {
	Nik int
	Nama, Tempat_Tgl_Lahir, Jenis_kelamin,Gol_Darah, Alamat, RT_RW, Kel_Desa, Kecamatan, Agama, Status_perkawinan, Pekerjaan, Kewarganegaraan, Berlaku_hingga string	
	
}

func main() {
	//inisialisasi
	I := Identitas {
	Nik : 7306117110010000 ,
	Nama : "NUR AYU ASRHI",
	Tempat_Tgl_Lahir : "JANNAYA, 30-10-2001",
	Jenis_kelamin : "PEREMPUAN",
	Gol_Darah : "-",
	Alamat : "BONTONOMPO",
	RT_RW : "005/004" ,
	Kel_Desa : "MANJAPAI" ,
	Kecamatan : "BONTOMPO" ,
	Agama : "ISLAM" ,
	Status_perkawinan : "BELUM KAWIN" ,
	Pekerjaan : "PELAJAR/MAHASISWA" ,
	Kewarganegaraan : "WNI" ,
	Berlaku_hingga : "SEUMUR HIDUP" ,
}
fmt.Println("----------------------------------------")
fmt.Println("PROVINSI SULAWESI SELATAN KABUPATEN GOWA")
fmt.Println("----------------------------------------")
fmt.Println("Nik               :", I.Nik)
fmt.Println("Nama              :", I.Nama)
fmt.Println("Tempat_Tgl_Lahir  :", I.Tempat_Tgl_Lahir)
fmt.Println("Jenis_kelamin     :", I.Jenis_kelamin)
fmt.Println("Gol_Darah         :", I.Gol_Darah)
fmt.Println("Alamat            :", I.Alamat)
fmt.Println("RT/RW             :", I.RT_RW)
fmt.Println("Kel/Desa          :", I.Kel_Desa)
fmt.Println("Kecamatan         :", I.Kecamatan)
fmt.Println("Agama             :", I.Agama)
fmt.Println("Status_perkawinan :", I.Status_perkawinan)
fmt.Println("Pekerjaan         :", I.Pekerjaan)
fmt.Println("Kewarganegaraan   :", I.Kewarganegaraan)
fmt.Println("Berlaku_hingga    :", I.Berlaku_hingga)

}