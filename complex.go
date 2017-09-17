package flatten

func complex64ToBytes(d complex64) []byte {
	f32r := real(d)
	f32i := imag(d)

	return append(float32ToBytes(f32r), float32ToBytes(f32i)...)
}

func complex128ToBytes(d complex128) []byte {
	f64i := imag(d)
	f64r := real(d)

	return append(float64ToBytes(f64r), float64ToBytes(f64i)...)
}
