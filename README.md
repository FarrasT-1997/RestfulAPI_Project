# ![ALTA SHOPPING](https://ibb.co/dbMMBSd)

API E-commerce untuk Alta Store 

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)

# Table of Content

- [Introduction](#introduction)
- [Feature Overview](#Feature/Overview)
- [How to use](#How/to/use)
- [HTTP Request User Menggunakan JSON](#HTTPRequestUserMenggunakanJSON)
- [HTTP Request Category Menggunakan JSON](#HTTPRequestCategoryMenggunakanJSON)
- [Understanding the Project](#Memahami/Projek)

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

# HTTP Request User Menggunakan JSON
- Login ke User
```
curl -X POST localhost:{yourport}/login
   -H 'Content-Type: application/json'
   -d '{"email":"{your email}","password":"{your password}"}'
```
- Signup/mendaftar User
```
curl -X POST localhost:{yourport}/signup
   -H 'Content-Type: application/json'-d '{"name":"{your name}","gender":"{your gender}","username":"{your username","email":"{your email}","password":"{your password}","address":"{your address}"}'
```

# HTTP Request Category Menggunakan JSON
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
- Mengganti Profil
```
curl -X PUT localhost:{your port}/users/{user id}
	-H "Accept: application/json"
    -H "Authorization: Bearer {token}"
    -d '{"name":"{your name}","gender":"{your gender}","username":"{your username","email":"{your email}","password":"{your password}","address":"{your address}"}'

```
- Melihat Profil
```
curl localhost:{your port}/users/{user id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}"
```
- Logout User
```
curl -X PUT localhost:{your port}/logout/{user id}
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}" 
```
# HTTP Request Transaksi Menggunakan JSON
- Membuat transaksi
```
curl -X POST localhost:{your port}/transaction
 -H 'Accept: application/json' 
 -H "Authorization: Bearer {token}" 
 
```
# Understanding the Project

Project ini bertujuan sebagai API untuk Alta Shop, dari pemilihan produk, memasukkan produk ke dalam chart, pembayaran produk, serta pembuatan akun pembeli.

Struktur projek:
- models: Mvc, data utama yang ada pada Alta Shop
- controllers: merupakan kumpulan fungsi untuk permintaan dari user, mereka meminta servis untuk melakukan aksi yang diminta user langsung ke database
- servis: mengandung logika bisnis untuk tiap model, dan juga untuk otorisasi
- middleware: mengandung middleware yaitu fungsi echo.
