## How to
`masih error gunain $GOROOT dan package module dari github langsung`
>> dikarenakan bukan main root di githubnya, remember?

```
> go mod init folder
> go run main.go
> go build
```

---

## MATERI

### 9. Introduction to Hexagonal Architecture
![separating](https://user-images.githubusercontent.com/20918862/107870682-0ff50900-6ecd-11eb-9b9b-9974ef540077.JPG)

![HECAGON](https://user-images.githubusercontent.com/20918862/107870734-9a3d6d00-6ecd-11eb-8e80-c781eeb5cd97.png)

1. Perbedaan Mock Adapter dan STUB? or struct?
2. Mock digunain tuntuk test our expectations when we weite behavior based test? whatever that means

### 10. Implementating Hexagonal Architecture
1. Nama komponen harus selalu jelas menggambarkan INTENT (gunain I di depan nama interface)

### 11. Database adapter
