## How to
`masih error gunain $GOROOT dan package module dari github langsung`
>> dikarenakan bukan main root di githubnya, remember?

```
> go mod init folder
> go run main.go
> go build
```
===

# MATERI

## 9. Introduction to Hexagonal Architecture
![separating](https://user-images.githubusercontent.com/20918862/107870682-0ff50900-6ecd-11eb-9b9b-9974ef540077.JPG)

![HECAGON](https://user-images.githubusercontent.com/20918862/107870734-9a3d6d00-6ecd-11eb-8e80-c781eeb5cd97.png)

1. Perbedaan MOCK Adapter dan STUB
2. Mock digunain tuntuk test our expectations when we weite behavior based test? whatever that means

## 10. Implementating Hexagonal Architecture
1. Nama komponen harus selalu jelas menggambarkan INTENT (gunain I di depan nama interface)
2. begini **rules**nya
```
(a attached to it) | NAMA_FUNC() | (Return, error)
```
---
1. Membuat folder domain dengan isi customer.go

2. Membuat adapter di domain dengan isi customerRepositoryStub.go
![Mock customerService.go](https://user-images.githubusercontent.com/20918862/107871416-dbd11680-6ed3-11eb-9366-d08cc15e6643.png)

3. lingkaran ungu untuk file customerService.go
![Mock customerService.go](https://user-images.githubusercontent.com/20918862/107871416-dbd11680-6ed3-11eb-9366-d08cc15e6643.png)
dimana dilakukan koneksi antara port 1 dan 2

4. lingkaran ungu untuk file handler.go (CustomerHandlers)
![handler.go)](https://user-images.githubusercontent.com/20918862/107871861-01601f00-6ed8-11eb-91c5-1a2df43f8024.png)

5. lingkaran ungu untuk file app.go (**ALL THE WIRING**)
![app.go](https://user-images.githubusercontent.com/20918862/107871858-fe652e80-6ed7-11eb-9aa0-30b46cb841d9.png)


## 11. Database adapter
