# BOILERPLATE GOLANG

## Installation

- clone [Repository](https://github.com/iqbaloke/template_go.git)
- go run main.go

## Module Package
- [Gofiber](https://gofiber.io/)
- [Gorm](https://gorm.io/)
- [Go playground](https://github.com/go-playground/validator)

## Penjelasan
### App
- Controllers
```
Untuk menghendel error seperti validasi
```
- Exception
```
Controller digunakan untuk mengolah data
```
- Models
```
Models digunakan untuk mengatur struct (table database)
```
- Providers
```
Untuk membuat global response seperti (success, notfound)
```
- Request

    - Validation
    ```
    Untuk mengatur validasi format, validasi message dan validasi request
    ```
- Resource
```
Mengolah response / mengubah response data
```

### config
- database (connection database)

### databases
- migrations
- seeders

### routes
- router digunakan sebagai definisi endpoint (url)