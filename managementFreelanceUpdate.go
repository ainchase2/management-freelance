package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Project struct {
	Nama     string
	Client   string
	Status   int
	Deadline time.Time
	DaysLeft int
}

var projects []Project

// Func untuk menampilkan pilihan menu
func PilihanMenu() {
	for {
		fmt.Println("\nPilihan Menu:")
		fmt.Println("1. Tambah Project")
		fmt.Println("2. Ubah Project")
		fmt.Println("3. Hapus Project")
		fmt.Println("4. Perbarui Status")
		fmt.Println("5. Search Project")
		fmt.Println("6. Sorting Project")
		fmt.Println("7. Tampilkan Hasil Project")
		fmt.Println("8. Exit")

		var pilih int
		fmt.Print("Tentukan Pilihanmu: ")
		fmt.Scan(&pilih)

		switch {
		case pilih == 1:
			tambahProject()
		case pilih == 2:
			ubahProject()
		case pilih == 3:
			hapusProject()
		case pilih == 4:
			perbaruiStatus()
		case pilih == 5:
			searchProject()
		case pilih == 6:
			sortingProject()
		case pilih == 7:
			tampilkanProject()
		case pilih == 8:
			fmt.Println("Program Selesai!")
			os.Exit(0)
		default:
			fmt.Println("Invalid! Pilih Menu yang tersedia")

		}
	}
}

/*
	Func untuk membaca input dari pengguna

menggunakan bufio untuk menangani spasi dan newline
*/
func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

/*
	Func untuk input deadline project

menggunakan format DD MM YYYY
*/
func inputDeadline() time.Time {
	var day, month, year int

	fmt.Println("Masukkan tanggal deadline (format: DD MM YYYY):")
	fmt.Print("Tanggal (DD): ")
	fmt.Scan(&day)
	fmt.Print("Bulan (MM): ")
	fmt.Scan(&month)
	fmt.Print("Tahun (YYYY): ")
	fmt.Scan(&year)

	if day < 1 || day > 31 || month < 1 || month > 12 || year < 2024 {
		fmt.Println("Format tanggal tidak valid! Menggunakan tanggal hari ini.")
		return time.Now()
	}

	deadline := time.Date(year, time.Month(month), day, 23, 59, 59, 0, time.Local)
	return deadline
}

/*
	Func untuk menghitung sisa hari dari deadline

Menggunakan time.Time untuk menghitung selisih waktu
*/
func hitungSisaHari(deadline time.Time) int {
	today := time.Now()
	duration := deadline.Sub(today)
	days := int(duration.Hours() / 24)

	return days
}

/*
Func untuk menkonversi status project dari int ke string
dan untuk mendapatkan status project sebagai string
*/
func getStatus(status int) string {
	switch status {
	case 1:
		return "Pending"
	case 2:
		return "OnProgress"
	case 3:
		return "Selesai"
	default:
		return "Unknown"
	}
}

// Func untuk input status project
func inputStatus() int {
	var status int
	for {
		fmt.Print("Status Project:\n1.Pending\n2.OnProgress\n3.Selesai\n")
		fmt.Print("Input Status: ")
		fmt.Scan(&status)
		if status >= 1 && status <= 3 {
			fmt.Printf("Status Project: %s\n", getStatus(status))
			return status
		} else {
			fmt.Println("Invalid! Pilih Status yang tersedia")
		}

	}
}

/*
	Func untuk membaca input setelah scan

untuk menghindari masalah newline yang tersisa
*/
func newLine(prompt string) string {
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

/*
	Func untuk menambahkan project baru

meminta input nama project, client, status, dan deadline
*/
func tambahProject() {
	var newProject Project
	fmt.Println("\n--- Tambah Project Baru ---")

	// Input nama project
	newProject.Nama = newLine("Nama Project: ")

	// Input nama client
	newProject.Client = readInput("Nama Client: ")

	// Input status project
	newProject.Status = inputStatus()

	// Input deadline
	newProject.Deadline = inputDeadline()

	// Hitung sisa hari
	newProject.DaysLeft = hitungSisaHari(newProject.Deadline)

	projects = append(projects, newProject)

	fmt.Println("\nSukses menambahkan project!")
}

/*
	Func untuk mengubah project yang sudah ada

memungkinkan pengguna untuk mengubah nama, client, status, dan deadline project
*/
func ubahProject() {
	if len(projects) == 0 {
		fmt.Println("Belum ada project yang diinput!")
		return
	}

	fmt.Println("\n--- Ubah Project ---")
	tampilkanProject()

	// input project yang akan diubah

	var index int
	fmt.Print("Pilih project yang akan diubah: ")
	fmt.Scan(&index)
	if index < 1 || index > len(projects) {
		fmt.Println("Nomor project tidak valid!")
		return
	}

	index--

	fmt.Println("\n--- Masukkan Data Baru ---")

	// Input nama project baru
	projects[index].Nama = newLine("Nama Project Baru: ")
	// Input nama client baru
	projects[index].Client = readInput("Nama Client Baru: ")
	/* Input status project baru
	menggunakan fungsi inputStatus untuk mendapatkan status baru*/
	fmt.Print("Status Project Baru: ")
	projects[index].Status = inputStatus()

	// Update deadline
	fmt.Println("Update deadline? (yes/no): ")
	var updateDeadline string
	fmt.Scan(&updateDeadline)

	if strings.ToLower(updateDeadline) == "yes" {
		projects[index].Deadline = inputDeadline()
		projects[index].DaysLeft = hitungSisaHari(projects[index].Deadline)
	}

	fmt.Println("Project sukses diperbarui!")
}

/*
	Func untuk menghapus project yang sudah ada

memungkinkan pengguna untuk menghapus project berdasarkan nomor yang ditampilkan
*/
func hapusProject() {
	if len(projects) == 0 {
		fmt.Println("Belum ada project yang diinput!")
		return
	}

	fmt.Println("\n--- Hapus Project ---")
	tampilkanProject()
	fmt.Println("--------------------------------------- ")
	// input project yang akan dihapus
	var index int
	fmt.Print("\nPilih project yang ingin dihapus: ")
	fmt.Scan(&index)
	if index < 1 || index > len(projects) {
		fmt.Println("Nomor project tidak valid!")
		return
	}

	index--
	// project akan dihapus sesuai dengan input
	projects = append(projects[:index], projects[index+1:]...)

	fmt.Println("Project berhasil dihapus!")
}

/*
	Func untuk memperbarui status project

memungkinkan pengguna untuk memperbarui status project yang sudah ada
*/
func perbaruiStatus() {
	if len(projects) == 0 {
		fmt.Println("Tidak ada project untuk diperbarui!")
		return
	}

	fmt.Println("\n--- Perbarui Status ---")
	tampilkanProject()
	// input perbarui status
	var index int
	fmt.Print("Masukkan no project untuk perbarui status: ")
	fmt.Scan(&index)
	if index < 1 || index > len(projects) {
		fmt.Println("Nomor project tidak valid!")
		return
	}
	// perbarui status yang baru di input
	projects[index-1].Status = inputStatus()

	fmt.Println("Status berhasil diperbarui!")

	tampilkanProject()

}

// Func untuk mencari project berdasarkan nama project atau client
func searchProject() {
	if len(projects) == 0 {
		fmt.Println("Belum ada project yang diinput!")
		return
	}
	// input keyword untuk mencari project
	var keyword string
	keyword = newLine("\nMasukkan nama project atau client yang dicari: ")

	fmt.Println("\n--- Hasil Pencarian ---")
	found := false
	for i, project := range projects {
		if strings.Contains(strings.ToLower(project.Nama), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(project.Client), strings.ToLower(keyword)) {
			fmt.Printf("%d. Nama Project: %s\n   Client: %s\n   Status: %s\n   Deadline: %s\n   Sisa Hari: %d\n\n",
				i+1, project.Nama, project.Client, getStatus(project.Status),
				project.Deadline.Format("02-01-2006"), project.DaysLeft)
			found = true
			fmt.Println("--------------------------------------- ")
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan project dengan kata kunci tersebut.")
	}
}

// Func untuk sorting project berdasarkan deadline terdekat (ascending order)
func sortingProject() {
	if len(projects) == 0 {
		fmt.Println("Belum ada project yang diinput!")
		return
	}

	fmt.Println("\n--- Sorting Project Berdasarkan Deadline Terdekat ---")

	// Sort berdasarkan deadline terdekat
	for i := 0; i < len(projects)-1; i++ {
		for j := i + 1; j < len(projects); j++ {
			if projects[i].Deadline.After(projects[j].Deadline) {
				projects[i], projects[j] = projects[j], projects[i]
			}
		}
	}

	fmt.Println("Berhasil sorting berdasarkan deadline terdekat!")
	tampilkanProject()
}

/*
	Func untuk menampilkan daftar project yang sudah diinput

menampilkan nama project, client, status, deadline, dan sisa hari
*/
func tampilkanProject() {
	if len(projects) == 0 {
		fmt.Println("Belum ada project yang diinput!")
		return
	}

	fmt.Println("\n--- Daftar Project ---")
	for i, project := range projects {
		// Update sisa hari setiap kali menampilkan
		projects[i].DaysLeft = hitungSisaHari(project.Deadline)

		var statusDisplay string
		switch project.Status {
		case 1:
			statusDisplay = "🟠 Pending"
		case 2:
			statusDisplay = "🔵 OnProgress"
		case 3:
			statusDisplay = "✅ Selesai"
		default:
			statusDisplay = "❓ Unknown"
		}
		// memunculkan deadline dari project
		var deadlineDisplay string
		if project.DaysLeft < 0 {
			deadlineDisplay = fmt.Sprintf("TERLAMBAT %d hari!", -project.DaysLeft)
		} else {
			deadlineDisplay = fmt.Sprintf("%d hari lagi", project.DaysLeft)
		}

		fmt.Printf("%d. Nama Project: %s\n   Client: %s\n   Status: %s\n   Deadline: %s (%s)\n\n",
			i+1, project.Nama, project.Client, statusDisplay,
			project.Deadline.Format("02-01-2006"), deadlineDisplay)
		fmt.Println("--------------------------------------- ")
	}
}

func main() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("  Selamat datang di aplikasi manajemen freelance!  ")
	fmt.Println("---------------------------------------------------")

	fmt.Println("1. Main Menu")
	fmt.Println("0. Exit")
	fmt.Print("Input Pilihan: ")
	var pilih int
	fmt.Scan(&pilih)
	if pilih == 1 {
		PilihanMenu()
	} else if pilih == 0 {
		fmt.Println("Terima kasih telah menggunakan aplikasi manajemen freelance!")
		return
	} else {
		fmt.Println("Invalid! Pilih menu yang tersedia.")
		fmt.Print("Input Pilihan: ")
		fmt.Scan(&pilih)
		for {
			if pilih != 0 && pilih != 1 {
				fmt.Println("Invalid! Pilih menu yang tersedia.")
				fmt.Print("Input Pilihan: ")
				fmt.Scan(&pilih)
			}
			if pilih == 1 {
				PilihanMenu()
			} else if pilih == 0 {
				fmt.Println("Terima kasih telah menggunakan aplikasi manajemen freelance!")
				return
			}
		}
	}
}
