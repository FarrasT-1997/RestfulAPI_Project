# ALTA SHOPPING

API E-commerce untuk Alta Store

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)
- [Introduction](#introduction)

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

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference stripe-go in a Go program with `import`:

``` go
import (
    "https://github.com/FarrasT-1997/RestfulAPI_Project"
)
```
# Memahami Projek

Project ini bertujuan sebagai API untuk Alta Shop, dari pemilihan produk, memasukkan produk ke dalam chart, pembayaran produk, serta pembuatan akun pembeli.

Struktur projek:
- models: Mvc, data utama yang ada pada Alta Shop
- controllers: merupakan kumpulan fungsi untuk permintaan dari user, mereka meminta servis untuk melakukan aksi yang diminta user langsung ke database
- servis: mengandung logika bisnis untuk tiap model, dan juga untuk otorisasi
- middleware: mengandung middleware yaitu fungsi echo.
