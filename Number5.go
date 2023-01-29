package main

import "fmt"

type Kendaraan struct {
	Suara string
}

func (k *Kendaraan) Akselerasi() {
	fmt.Println(k.Suara)
}

type Sepeda struct {
	Kendaraan
	Rantai string
}

func (s *Sepeda) Akselerasi() {
	s.Kendaraan.Akselerasi()
	s.Rantai = "Perlu perbaikan"
}

type Mobil struct {
	Kendaraan
	Bensin string
}

func (m *Mobil) Akselerasi() {
	m.Kendaraan.Akselerasi()
	m.Bensin = "Kosong"
}

func main() {
	mobil := &Mobil{
		Kendaraan: Kendaraan{Suara: "Vroom"},
		Bensin:    "Penuh",
	}

	sepeda := &Sepeda{
		Kendaraan: Kendaraan{Suara: "Swoosh"},
		Rantai:    "Normal",
	}

	sepeda.Akselerasi()
	fmt.Println(sepeda.Rantai)

	fmt.Println("")
	mobil.Akselerasi()
	fmt.Println(mobil.Bensin)
}
