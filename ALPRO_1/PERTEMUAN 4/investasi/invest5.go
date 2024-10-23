package main

import "fmt"

type nilai struct {
	hargaBeli, hargaJual                                float64
	jumlahSaham                                         int
	totalInvestasiAwal, totalPenjualan, keuntunganKotor float64
	pajakKeuntungan                                     float64
	biayaTransaksi, keuntunganBersih                    float64
}

func main() {
	var t nilai
	fmt.Println("Masukan Harrga beli")
	fmt.Scan(&t.hargaBeli)

	fmt.Println("Masukan Harga jual")
	fmt.Scan(&t.hargaJual)

	fmt.Println("Masukan Jumlah saham")
	fmt.Scan(&t.jumlahSaham)

	t.totalInvestasiAwal = t.hargaBeli * float64(t.jumlahSaham)
	t.totalPenjualan = t.hargaJual * float64(t.jumlahSaham)
	t.keuntunganKotor = t.totalPenjualan - t.totalInvestasiAwal
	t.biayaTransaksi = (0.2 / 100) * t.totalPenjualan
	t.pajakKeuntungan = 0.10 * t.keuntunganKotor

	if t.pajakKeuntungan < 0 {
		t.pajakKeuntungan = 0
	}

	t.keuntunganBersih = t.keuntunganKotor - (t.biayaTransaksi + t.pajakKeuntungan)

	fmt.Printf("biaya transaksi %.2f ", t.biayaTransaksi)
	fmt.Printf("pajak keuntungan %.2f ", t.pajakKeuntungan)
	fmt.Printf("hasilnya adalah %.2f ", t.keuntunganBersih)


}
