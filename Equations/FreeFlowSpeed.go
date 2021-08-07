package main

import (
	"fmt"
)

//type Ffs interface {
//	multilane_ffs() float64
//	freeway_ffs() float64
//	basic_ffs() float64
//}

const sp_less50 int8 = 7
const sp_greater50 = 5

//func (b bffs ) mph() {
//	if Sp <= 50{
//		sp_adjust = sp_less50
//
//	}else{
//		sp_adjust = sp_greater50
//	}
//
//}

type Ffs struct {
	bffs, flw, ftlc, fm, fa float64
}

func (f Ffs) mph() float64 {

	a := float64(f.bffs) - (f.flw) - (f.ftlc) - (f.fm) - (f.fa)

	return a

}

func main() {

	multilaneFfs := Ffs{
		bffs: 90,
		flw:  3,
		ftlc: 1.3,
		fm:   2,
		fa:   1.8}

	fmt.Println(multilaneFfs.mph())

}
