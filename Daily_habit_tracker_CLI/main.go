package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const dataFile = "data/habits.json"

type Habit struct {
	Name      string   `json:"name"`
	Completed []string `json:"completed"`
}

type Data struct {
	Habits []Habit `json:"habits"`
}

func loadData() Data {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return Data{Habits: []Habit{}}
	}

	file, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("Error membaca data:", err)
		return Data{}
	}

	var data Data
	json.Unmarshal(file, &data)
	return data
}

func saveData(data Data) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error menyimpan data:", err)
		return
	}
	os.WriteFile(dataFile, file, 0644)
}

func addHabit() {
	var habitName string
	fmt.Print("Nama kebiasaan baru: ")
	fmt.Scanln(&habitName)

	data := loadData()
	data.Habits = append(data.Habits, Habit{Name: habitName, Completed: []string{}})

	saveData(data)
	fmt.Println("✔ Habit ditambahkan:", habitName)
}

func viewHabits() {
	data := loadData()

	if len(data.Habits) == 0 {
		fmt.Println("Belum ada habit.")
		return
	}

	fmt.Println("\n📋 Daftar Habit:")
	for i, h := range data.Habits {
		fmt.Printf("%d. %s (%d kali selesai)\n", i+1, h.Name, len(h.Completed))
	}
}

func markDone() {
	data := loadData()
	if len(data.Habits) == 0 {
		fmt.Println("Belum ada habit.")
		return
	}

	viewHabits()
	var index int
	fmt.Print("\nPilih nomor habit: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(data.Habits) {
		fmt.Println("❌ Index tidak valid.")
		return
	}

	today := time.Now().Format("2006-01-02")
	habit := &data.Habits[index-1]

	// Cek apakah sudah ditandai hari ini
	for _, d := range habit.Completed {
		if d == today {
			fmt.Println("⚠ Habit sudah completed hari ini.")
			return
		}
	}

	habit.Completed = append(habit.Completed, today)
	saveData(data)
	fmt.Println("🎉 Good job! Habit diselesaikan hari ini.")
}

func viewProgress() {
	data := loadData()
	if len(data.Habits) == 0 {
		fmt.Println("Belum ada habit.")
		return
	}

	fmt.Println("\n📈 Progress Habit:")
	for _, h := range data.Habits {
		fmt.Printf("- %s: %d kali selesai\n", h.Name, len(h.Completed))
	}
}

func menu() {
	for {
		fmt.Println("\n===== Habit Tracker =====")
		fmt.Println("1. Tambah Habit")
		fmt.Println("2. Lihat Habit")
		fmt.Println("3. Tandai Selesai Hari Ini")
		fmt.Println("4. Lihat Progress")
		fmt.Println("5. Keluar")
		fmt.Print("> ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			addHabit()
		case "2":
			viewHabits()
		case "3":
			markDone()
		case "4":
			viewProgress()
		case "5":
			fmt.Println("Bye 👋")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func main() {
	menu()
}
