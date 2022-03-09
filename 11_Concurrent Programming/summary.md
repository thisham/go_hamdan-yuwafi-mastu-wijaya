# 11) Concurrent Programming

## Overview

Hal yang dapat dipelajari dari chapter 11 ini antara lain:

1. [Typical of Running Procedures](#typical-of-running-procedures);
2. [Go-Routines, Channels and Select](#go-routines-channels-and-select);
3. [Race Condition](#race-condition).

### Typical of Running Procedures

- Sequential
  
  Program dieksekusi bergantian setiap prosedurnya dan dilakukan oleh satu core CPU.
  
- Parallel

  Setiap prosedur yang tidak berkaitan dieksekusi secara bersamaan dalam core CPU yang berbeda.

- Concurrent
  
  Setiap prosedur dieksekusi bergantian, namun prosedur itu dipecah dalam bagian-bagian kecil. Setiap bagian kecil-kecil tersebut dapat saling berkomunikasi untuk menyelesaikan maslah melalui channel yang didefinisikan sebelumnya.

### Go-Routines, Channels and Select

**Go-Routines** merupakan unsur utama dalam concurrent programming untuk menyelesaikan sub-masalah. Setiap unsur concurrent dapat dihubungkan melalui **channel** untuk memproses keterkaitan masalah yang ada pada setiap routine. Jika terdapat lebih dari satu channel, channel dapat dipilih menggunakan **select** untuk mengunakannya.

### Race Condition

Suatu kondisi pengiriman data dari variabel di go-routine pengirim melalui channel yang prosesnya belum selesai atau belum terkirim ke variabel penerima di go-routine target, namun prosedur dari fungsi utama sudah mengalami return value atau sudah selesai. Kondisi ini dapat dicegah dengan **waitgroup** (menunggu proses di go-routine selesai), **blocking with channel** (go-routine utama akan di-blok hingga child go-routine membalikkan nilai), **returning the channel** (menjadikan channel sebagai nilai return), **using mutex** (menyelesaikan salah satu child go-routine selesai, kemudian menjalankan child go-routine lainnya hingga variabel yang dijadikan return value oleh go-routine utama terisi).

## Tasks

### P1 - Count Letter Concurrently

- solution [p1-letter-frequencies.go](praktikum/p1-letter-frequencies.go)
- command to run:

  ```bash
    go run .\p1-letter-frequencies.go "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam reiciendis qui in eius consectetur magni sint beatae a neque laborum exercitationem ab cum, ipsa suscipit quam quasi voluptate debitis repellendus."
  ```

- screenshot: ![p1-letter-frequencies.go](screenshots/p1-letter-frequencies.png)
