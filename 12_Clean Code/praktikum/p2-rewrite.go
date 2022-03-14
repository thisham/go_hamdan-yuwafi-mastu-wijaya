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
