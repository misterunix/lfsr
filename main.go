package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	p4 := lfsr4()
	fmt.Println("lsfr4(): period:", p4, "seconds:", time.Since(t).Seconds())

	t = time.Now()
	p8 := lfsr8()
	fmt.Println("lsfr8(): period:", p8, "seconds:", time.Since(t).Seconds())

	t = time.Now()
	p16 := lfsr16()
	fmt.Println("lsfr16(): period:", p16, "seconds:", time.Since(t).Seconds())

	t = time.Now()
	p24 := lfsr24()
	fmt.Println("lsfr24(): period:", p24, "seconds:", time.Since(t).Seconds())

	t = time.Now()
	p32 := lfsr32()
	fmt.Println("lsfr32(): period:", p32, "seconds:", time.Since(t).Seconds())

	t = time.Now()
	p63 := lfsr63()
	fmt.Println("lsfr63(): period:", p63, "seconds:", time.Since(t).Seconds())

	t = time.Now()
	p64 := lfsr64()
	fmt.Println("lsfr64(): period:", p64, "seconds:", time.Since(t).Seconds())

}

func lfsr64() uint64 {

	var start_state uint64
	var lfsr uint64
	var bit uint64

	start_state = 0xACE1
	lfsr = start_state

	var period uint64 = 0

	/* taps: 32,22,2,1 + 11*/
	for {
		//	fmt.Print(lfsr & 1)
		//         64,             63,            61,            60      + 1
		bit = ((lfsr >> 0) ^ (lfsr >> 1) ^ (lfsr >> 3) ^ (lfsr >> 4)) & 1

		lfsr = (lfsr >> 1) | (bit << 63)

		if lfsr == start_state {
			break
		}

		period++
	}

	//fmt.Print("\nPeriod: ", period)
	return period
}

func lfsr63() uint64 {

	var start_state uint64
	var lfsr uint64
	var bit uint64

	start_state = 0xACE1
	lfsr = start_state

	var period uint64 = 0

	/* taps: 63,62+ 1*/
	for {
		//	fmt.Print(lfsr & 1)
		//         63,             62      + 1
		bit = ((lfsr >> 0) ^ (lfsr >> 1)) & 1

		lfsr = (lfsr >> 1) | (bit << 62)

		if lfsr == start_state {
			break
		}

		period++
	}

	//fmt.Print("\nPeriod: ", period)
	return period
}

func lfsr32() uint64 {
	var start_state uint32
	var lfsr uint32
	var bit uint32

	start_state = 0xACE1
	lfsr = start_state

	var period uint64 = 0

	/* taps: 32,22,2,1 + 11*/
	for {
		//	fmt.Print(lfsr & 1)
		//         32             22             2             1         + 1
		bit = ((lfsr >> 0) ^ (lfsr >> 10) ^ (lfsr >> 30) ^ (lfsr >> 31)) & 1

		lfsr = (lfsr >> 1) | (bit << 31)

		if lfsr == start_state {
			break
		}

		period++
	}
	return period
	//fmt.Print("\nPeriod: ", period)
}

func lfsr24() uint {
	var start_state uint32
	var lfsr uint32
	var bit uint32

	start_state = 1 << 23
	lfsr = start_state

	var period uint = 0

	/* taps: 32,22,2,1 + 11*/
	for {
		//	fmt.Print(lfsr & 1)
		//           24,          23,             22,          17        + 1
		bit = ((lfsr >> 0) ^ (lfsr >> 1) ^ (lfsr >> 2) ^ (lfsr >> 7)) & 1

		lfsr = (lfsr >> 1) | (bit << 23)

		if lfsr == start_state {
			break
		}

		period++
	}
	return period
	//fmt.Print("\nPeriod: ", period)
}

func lfsr16() int {
	var start_state uint16
	var lfsr uint16
	var bit uint16

	start_state = 0xACE1
	lfsr = start_state

	var period int = 0

	/* taps: 16 14 13 11; feedback polynomial: x^16 + x^14 + x^13 + x^11 + 1 */
	for {
		//fmt.Print(lfsr & 1)
		//         16             14             13             11    + 1
		bit = ((lfsr >> 0) ^ (lfsr >> 2) ^ (lfsr >> 3) ^ (lfsr >> 5)) & 1

		lfsr = (lfsr >> 1) | (bit << 15)

		if lfsr == start_state {
			break
		}

		period++
	}
	return period
	//fmt.Print("\nPeriod: ", period)
}

func lfsr8() int {
	var start_state uint8
	var lfsr uint8
	var bit uint8

	start_state = 0xAC
	lfsr = start_state

	var period int = 0

	/* taps: 8,6,5,4; */
	for {
		//	fmt.Print(lfsr & 1)
		//         8             6             5             4    + 1
		bit = ((lfsr >> 0) ^ (lfsr >> 2) ^ (lfsr >> 3) ^ (lfsr >> 4)) & 1

		lfsr = (lfsr >> 1) | (bit << 7)

		if lfsr == start_state {
			break
		}

		period++
	}
	return period
	//fmt.Print("\nPeriod: ", period)
}
func lfsr4() int {
	var start_state uint8
	var lfsr uint8
	var bit uint8

	start_state = 0b1001
	lfsr = start_state

	var period int = 0

	/* taps: 4,3 */
	for {
		//fmt.Print(lfsr & 1)
		//         4             3         + 1
		bit = ((lfsr >> 0) ^ (lfsr >> 1)) & 1

		lfsr = (lfsr >> 1) | (bit << 3)

		if lfsr == start_state {
			break
		}

		period++
	}
	return period
	//fmt.Print("\nPeriod: ", period)
}
