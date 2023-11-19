package main

import (
	"fmt"
)

// struct dari pertanyaan yang ada
type soal struct {
	pertanyaan   string
	pilihan      [4]string
	kunciJawaban int
}

// fungsi untuk menampilkan soal kepada user
func tampilkanSoal(q soal) {
	fmt.Println(q.pertanyaan)
	for i, opsi := range q.pilihan {
		fmt.Printf("%d. %s\n", i, opsi)
	}
}

// fungsi untuk mengecek jawaban user dengan kunci jawaban dari soal yang ada
func cekJawaban(jawabanUser int, kunciJawaban int) bool {
	return jawabanUser == kunciJawaban
}

func main() {

	// mendeklarasikan variabel yang diperlukan
	var nama string      //variabel untuk menyimpan nama user
	var score int        //variabel untuk menyimpan jumlah score user
	var jawabanBenar int //variabel untuk menyimpan jumlah jawaban benar dari user
	var jawabanSalah int //variabel untuk menyimpan jumlah jawaban salah dari user
	var jawabanUser int  //variabel untuk menyimpan jawaban yang diberikan oleh user

	//variabel tipe slice agar dapat menambah pertanyaan sebanyak mungkin
	//slice yang berisi daftar soal-soal dengan format sesuai struct soal di line ke 8
	var daftarPertanyaan = []soal{
		{ //pertanyaan 1
			pertanyaan:   "nama presiden pertama indonesia adalah...",
			pilihan:      [4]string{"Soeharto", "Megawati", "Soekarno", "Jokowi"},
			kunciJawaban: 2,
		},

		{ //pertanyaan 2
			pertanyaan:   "Manusia bernapas dengan...",
			pilihan:      [4]string{"paru-paru", "insang", "trakea", "ijin Tuhan"},
			kunciJawaban: 0,
		},

		{ //pertanyaan 3
			pertanyaan:   "a = 1, b = -2 dan c = -1\nPilihlah jawaban yang tepat sesuai pernyataan diatas!",
			pilihan:      [4]string{"b > c", "c < a", "a < c", "c < b"},
			kunciJawaban: 1,
		},
	}

	//meminta user untuk menginput nama nya dan menyimpan nama tersebut ke variabel "nama" yang sudah dideklarasi sebelumnya
	fmt.Print("Input nama: ")
	fmt.Scanln(&nama)

	//melakukan looping sebanyak jumlah soal yang ada lalu mengeksekusi fungsi-fungsi yang dituliskan dibawahnya
	for _, soal := range daftarPertanyaan {
		//menampilkan soal
		tampilkanSoal(soal)

		//menampilkan prosedur jawaban yang harus dipilih kepada user dan menyimpan jawaban dari user ke variabel "jawabanUser" yang sudah dideklarasi sebelumnya
		fmt.Print("Jawaban (0,1,2,3) : ")
		fmt.Scanln(&jawabanUser)

		//melakukan pengecekan jawaban user dengan kunci jawaban dari soal yang ada
		//jika jawaban user benar, maka value dari variabel "jawabanBenar" dan "score" bertambah 1
		//jika jawaban user salah, maka value dari variabel "jawabanSalah" bertambah 1
		if cekJawaban(jawabanUser, soal.kunciJawaban) {
			jawabanBenar++
			score++
		} else {
			jawabanSalah++
		}
		fmt.Println("\n")
	}
	//menampilkan statistic user, seperti nama, score dari user, jawaban benar dan jawaban salah dari user
	fmt.Printf("Statistic Kuis\n")
	fmt.Printf("Nama         : %s\n", nama)
	fmt.Printf("Score        : %d\n", score)
	fmt.Printf("Jawaban benar: %d\n", jawabanBenar)
	fmt.Printf("Jawaban salah: %d\n", jawabanSalah)
}
