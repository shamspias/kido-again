package main

// BadStruct : Wastes memory! (24 bytes total)
// CPU reads bool (1 byte), has to pad 7 empty bytes to read the big int64, then reads bool.
type BadStruct struct {
	IsAdmin bool  // 1 byte + 7 bytes of wasted padding
	Score   int64 // 8 bytes
	IsAlive bool  // 1 byte + 7 bytes of wasted padding
}

// GoodStruct : Perfectly packed! (16 bytes total)
// By putting the biggest items first, the small items pack together at the end!
type GoodStruct struct {
	Score   int64 // 8 bytes
	IsAdmin bool  // 1 byte
	IsAlive bool  // 1 byte
}
