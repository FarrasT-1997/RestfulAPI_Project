<a href="https://github.com/FarrasT-1997/RestfulAPI_Project"><img height="80" src="https://image.flaticon.com/icons/png/512/3081/3081648.png"></a>

# ALTA SHOPPING
API E-commerce untuk Alta Store 

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)

# Table of Content

- [Introduction](#introduction)
- [Feature Overview](#feature-overview)
- [How to use](#how-to-use)
- [HTTP Request User](#httprequestUser)
- [HTTP Request Category](#httprequestCategory)
- [HTTP Request Mengakses User](#httprequestMengaksesUser)
- [HTTP Request Transaksi](#httprequestTransaksi)
- [HTTP Request Cart](#httprequestCart)
- [Understanding the Project](#memahami-projek)
- [Image source](#image-source)

# Introduction
Project ini merupakan API untuk aplikasi E-commerce ALTA Shop. Project ini ditulis menggunakan bahasa pemrograman Golang menggunakan Echo framework dan GORM.

# Feature Overview

- Kostumer dapat melihat list product berdasarkan kategori product
- Kostumer dapat memasukan produk ke chart mereka
- Kostumer dapat melihat isi dari chart mereka
- Kostumer dapat menghapus produk yang mereka masukan di chart mereka
- Kostumer dapat melakukan checkout dan melakukan transaksi pembayaran pada produk yang mereka masukan ke chart
- Kostumer dapat membuat akun dan login ke akun yang mereka buat.

# How to use

Pastikan projek anda menggunakan Modul Go(harus ada berkas `go.mod` pada folder projek anda):

``` sh
go mod init
```

Lalu, referensikan Alta-Shopping di dalam Program Go dengan menggunakan `import`:

``` go
import (
    "https://github.com/FarrasT-1997/RestfulAPI_Project"
)
```
Untuk menjalankan API maka:
```
$ ./start.sh
```

Untuk payment method, harus membuat database yang berisi tabel paymentmethod dan datanya bisa dimasukkan secara manual. Contoh: id_bank= 1, nama_bank= bank abc, dst.
# HTTP Request User
- Login ke User
```
curl -X POST localhost:{yourport}/login
   -H 'Content-Type: application/json'
   -d '{"email}":"{your email}","password":"{your password}"}'
```
- Signup/mendaftar User
```
curl -X POST localhost:{yourport}/signup
   -H 'Content-Type: application/json'-d '{"name}":"{your name}","gender":"{your gender}","username":"{your username}","email":"{your email}","password":"{your password}","address":"{your address}"}'
```

# HTTP Request Category
- Mengambil seluruh kategori barang yang tersedia
```
curl localhost:{your port}/category
```
- Mengambil barang yang ada pada kategory tertentu
```
curl localhost:{your port}/category/{category id}
```
- Mengambil barang berdasarkan id barang
```
curl localhost:{your port}/product/{product category}
```

# HTTP Request Mengakses User

Untuk mendapatkan token JWT maka harus melakukan http request login, jika login berhasil maka token akan didapatkan.

- Mengganti Profil
```
curl -X PUT localhost:{your port}/jwt/users/{user id}
	-H "Accept: application/json"
    -H "Authorization: Bearer {token}"
    -d '{"name":"{your name}","gender":"{your gender}","username":"{your username}","email":"{your email}","password":"{your password}","address":"{your address}"}'

```
- Melihat Profil
```
curl localhost:{your port}/jwt/users/{user id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}"
```
- Logout User
```
curl -X PUT localhost:{your port}/jwt/logout/{user id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}" 
```
# HTTP Request Transaksi
- Membuat transaksi
```
curl -X POST localhost:{your port}/jwt/transaction
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}" 
```
- Melihat semua transaksi
```
curl localhost:{your port}/jwt/transaction
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}" 
```
- Menghapus transaksi
```
curl -X DELETE localhost:{your port}/jwt/transaction/{transaction id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}" 
```
- Mengubah payment method
```
curl -X PUT localhost:{your port}/jwt/transaction/payment/{transaction id}/{payment id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}"
```
- Checkout transaksi
```
curl -X PUT localhost:{your port}/jwt/transaction/checkout/{transaction id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}"
```
- Mengubah status transaksi
```
curl -X PUT localhost:{your port}/jwt/transaction/paid/{transaction id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}"
```
# HTTP Request Cart
- Memasukkan barang ke cart
```
curl -X POST localhost:{your port}/jwt/transaction/{transaction id}/cart/{product id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}"
```
- Mengganti banyaknya barang yang dimasukkan
```
curl -X PUT localhost:{your port}/jwt/transaction/{transaction id}/cart/{cart id}/{quantity}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}"
```
- Menghapus cart
```
curl -X DELETE localhost:{your port}/jwt/transaction/{transaction id}/{cart id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}"
```
- Mendapatkan semua produk yang ada di dalam cart
```
curl -X localhost:{your port}/jwt/transaction/cart/{transaction id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}"
```
* Kemungkinan akan banyak curl method yang tidak berfungsi haha.
* Dapat juga menggunakan Postman untuk melakukan HTTP Request, tinggal memasukkan value yang ada pada tiap method.

# Understanding the Project

Project ini bertujuan sebagai API untuk Alta Shop, dari pemilihan produk, memasukkan produk ke dalam chart, pembayaran produk, serta pembuatan akun pembeli.

Struktur projek:
- models: Mvc, data utama yang ada pada Alta Shop
- controllers: merupakan kumpulan fungsi untuk permintaan dari user, mereka meminta servis untuk melakukan aksi yang diminta user langsung ke database
- jwt : merupakan fungsi untuk autentikasi pada http request tertentu.
- middleware: mengandung middleware yaitu fungsi echo.

# Image Source
- "https://www.flaticon.com/authors/iconixar"
