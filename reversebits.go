package piscine

func reverseBits(num byte) byte {
	var res byte
	for i := 0; i < 8; i++ {
		res = res<<1 + num&1
		num = num >> 1
	}
	return res
}

// func ReverseBits(oct byte) byte  {
// 	b1:= oct << 7
// 	b2:= oct << 6  >> 7 << 6
// 	b3:= oct << 5 >> 7 << 5
// 	b4:= oct << 4 >> 7 << 4
// 	b5:= oct << 3 >> 7 << 3
// 	b6:= oct << 2 >> 7 << 2
// 	b7:= oct << 1 >> 7 << 1
// 	b8:= oct << 0 >> 7 << 0

// 	return b1 + b2 + b3 + b4 + b5 + b6 + b7 + b8
// }
