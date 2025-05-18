package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Project struct {
	Nama     string
	Client   string
	Status   string
	Deadline time.Time
	DaysLeft int
}

var projects []Project

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

func hitungSisaHari(deadline time.Time) int {
	today := time.Now()
	duration := deadline.Sub(today)
	days := int(duration.Hours() / 24)

	return days
}

func tambahProject() {
	var newProject Project
	fmt.Println("\n--- Tambah Project Baru ---")

	fmt.Print("Nama Project: ")
	fmt.Scan(&newProject.Nama)

	fmt.Print("Nama Client: ")
	fmt.Scan(&newProject.Client)

	fmt.Print("Status Project (Pending/OnProgress/Selesai): ")
	fmt.Scan(&newProject.Status)
	for {
		if newProject.Status != "Pending" && newProject.Status != "OnProgress" && newProject.Status != "Selesai" {
			fmt.Println("Status tidak valid! Pilih salah satu: Pending, OnProgress, Selesai")
			fmt.Print("Input Status: ")
			fmt.Scan(&newProject.Status)
		}
		if newProject.Status == "Pending" || newProject.Status == "OnProgress" || newProject.Status == "Selesai" {
			break
		}
	}

	// Input deadline
	newProject.Deadline = inputDeadline()

	// Hitung sisa hari
	newProject.DaysLeft = hitungSisaHari(newProject.Deadline)

	projects = append(projects, newProject)

	fmt.Println("\nSukses menambahkan project!")
}

func ubahProject() {
	if len(projects) == 0 {
		fmt.Println("Belum ada project yang diinput!")
		return
	}

	fmt.Println("\n--- Ubah Project ---")
	tampilkanProject()
	fmt.Println("--------------------------------------- ")
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
	fmt.Print("Nama Project Baru: ")
	fmt.Scan(&projects[index].Nama)

	fmt.Print("Nama Client Baru: ")
	fmt.Scan(&projects[index].Client)

	fmt.Print("Status Project Baru: ")
	fmt.Scan(&projects[index].Status)

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

func perbaruiStatus() {
	if len(projects) == 0 {
		fmt.Println("Tidak ada project untuk diperbarui!")
		return
	}

	fmt.Println("\n--- Perbarui Status ---")
	tampilkanProject()
	fmt.Println("--------------------------------------- ")
	// input perbarui status
	var index int
	fmt.Print("Masukkan no project untuk perbarui status: ")
	fmt.Scan(&index)
	if index < 1 || index > len(projects) {
		fmt.Println("Nomor project tidak valid!")
		return
	}
	// perbarui status yang baru di input
	fmt.Print("Masukkan status baru (Pending/OnProgress/Selesai): ")
	fmt.Scan(&projects[index-1].Status)

	fmt.Println("Status berhasil diperbarui!")
}

func searchProject() {
	if len(projects) == 0 {
		fmt.Println("Belum ada project yang diinput!")
		return
	}
	// input project atau client yang ingin dicari
	var keyword string
	fmt.Print("\nMasukkan nama project atau client yang dicari: ")
	fmt.Scan(&keyword)
	// memunculkan project yang dicari
	fmt.Println("\n--- Hasil Pencarian ---")
	found := false
	for i, project := range projects {
		if strings.Contains(strings.ToLower(project.Nama), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(project.Client), strings.ToLower(keyword)) {
			fmt.Printf("%d. Nama Project: %s\n   Client: %s\n   Status: %s\n   Deadline: %s\n   Sisa Hari: %d\n\n",
				i+1, project.Nama, project.Client, project.Status,
				project.Deadline.Format("02-01-2006"), project.DaysLeft)
			found = true
			fmt.Println("--------------------------------------- ")
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan project dengan kata kunci tersebut.")
	}
}

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
		case "Pending":
			statusDisplay = "🟠 Pending"
		case "OnProgress":
			statusDisplay = "🔵 OnProgress"
		case "Selesai":
			statusDisplay = "✅ Selesai"
		default:
			statusDisplay = project.Status
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
