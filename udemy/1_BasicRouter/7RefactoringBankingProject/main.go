package main

// rapikan karena banyak clutter, mulai bikin mvc
/*

1. semua logic ada di folder app
	- semua penyerahan ditangani oleh handler.go
	- app.go sebagai main
2. jangan lupa kalau semua yan di export, awalan harus kapital

*/

import (
	"7RefactoringBankingProject/app"
)

func main() {
	app.Start()
}
