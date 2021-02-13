## How to
`masih error gunain $GOROOT dan package module dari github langsung`
>> dikarenakan bukan main root di githubnya, remember?

```
> go mod init folder
> go run main.go
> go build
```

## MATERI

### 7. refactoring
> fixing the clutter dan diubah ke file `tree`

📦7RefactoringBankingProject
 ┣ 📂app
 ┃ ┣ 📜app.go
 ┃ ┗ 📜handler.go 
 ┣ 📜7RefactoringBankingProject.exe
 ┣ 📜go.mod
 ┣ 📜main.go
 ┗ 📜README.md

> folder **app** berisi seluruh `logic`
> **handlers** berisi seluruh `fungsi` yang bertugas menyerahkan tugas dari DATABASE -> MAIN

### 8. gorilla mux
```
"github.com/gorilla/mux"
```
jadinya gunain mux (multiplexer) instead of ~~http (default multiplexer)~~
> `gorilla mux` hanya **library** yang cuma `nyederhanain route definitions` eg. json dan xml encoder
>> lebih populer lain, ada `gin gonic` sebagai **framework** features lebih canggih eg. *validation* dan *customized response writing*