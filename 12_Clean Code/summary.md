# 12) Clean Code

## Overview

Dalam chapter ini, dapat dipelajari:

1. [Clean Code dan Karakteristiknya](#clean-code-dan-karakteristiknya);
2. [Prinsip Clean Code](#prinsip-clean-code);
3. [Refactoring](#refactoring).

### Clean Code dan Karakteristiknya

Clean code merupakan pola penulisan kode yang mudah dibaca, dipahami dan diubah oleh programmer dan maintainer kode.

Karakteristiknya antara lain:

- Mudah dipahami dan deskriptif, karena nama variabel menggambarkan data yang diisikan ke variabel tersebut;
- Mudah dieja dan dicari, penamaan yang deskriptif dan menghindari singkatan yang tidak perlu akan lebih mudah untuk dicari dan digunakan lagi;
- Konsisten, penamaan yang konsisten menolong maintainer untuk mencari dan me-reuse variabel tersebut;
- Nama variabel sederhana dan langsung ke topik utama;
- Komentar, jika kode sudah cukup deskriptif, tidak perlu untuk memberi komentar tambahan;
- Penggunaan konvensi koding, penggunaan ketetapan yang baku membuat koding makin konsisten sehingga mudah terbaca;
- Menggunakan format yang konsisten.

### Prinsip Clean Code

#### KISS (Keep It So Simple)

Prinsip ini menghindari god-like function, yaitu satu fungsi yang melakukan banyak prosedur sekaligus.

#### DRY (Don't Repeat Yourself)

Prinsip ini mencegah duplikasi dengan menggunakan function untuk membungkus prosedur yang sering dipakai.

### Refactoring

Refactoring adalah proses me-restruktur kode dengan prinsip dan karakterisik clean code, sehingga dapat menambah tingkat keterbacaan kode tanpa mengubah behavior dari kode itu sendiri. Perubahan behavior yang disebabkan karena refactor dapat disebut dengan regression.

## Tasks

### P1 - Analysis

```go
package main

type user struct {
  id       int
  username int
  password int
}

type userservice struct {
  t []user
}

func (u userservice) getallusers() []user {
  return u.t
}

func (u userservice) getuserbyid(id int) user {
  for _, r := range u.t {
    if id == r.id {
      return r
    }
  }
  return user{}
}
```

- Hasil Analisa:
  1. The name of structs and function is lack of readability. This can be improved by using the common naming conventions, like:
      - The structs name and (will be used as public) function must be named with pascal case;
  2. The userservice seems a redundant struct, since it was a same content like []user type. Alternatively, it can be replaced by using direct typing aliasing instead of creating a new struct. It can be written like: `type userservice []user`;
  3. Sub-variable name of userservice struct didn't describe this struct correctfully. the next-maintainer may can't directly recognize what data contained here when the `t`-named variable declared. Hence, the code that using this struct can't use this code cleanly. it can be improved by using more desciptive name like `userGroup`, or it can be removed if point 2 is applied;
  4. Struct-managed object isn't using a descriptive name, maybe it will confuse the next-code manager. It can be improved with naming the variable like `users`, and point 1 must be applied first to avoid duplication;
  5. The variable `r` in `getuserbyid` method, doesn't describe what data will be contained here. Alternatively, it can be renamed as `user` or `userData`. But, the point 1 must be applied first to avoid duplication;
  6. Struct `user` contains unnecessary data to define, like `username` and `password` it can be removed since they don't used in the entire project.

### P2 - Rewrite

```go
package main

type Kendaraan struct {
  kecepatanPerJam int
}

func (mobil *Kendaraan) berjalan() {
  mobil.tambahKecepatan(10)
}

func (mobil *Kendaraan) tambahKecepatan(kecepatanBaru int) {
  mobil.kecepatanPerJam += kecepatanBaru
}

func main() {
  mobilCepat := Kendaraan{}
  mobilCepat.berjalan()
  mobilCepat.berjalan()
  mobilCepat.berjalan()

  mobilLambat := Kendaraan{}
  mobilLambat.berjalan()
}

```
