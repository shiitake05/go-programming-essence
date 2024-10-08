// Code generater by stringer -type Fruit.go; DO NOT EDIT

package main

import "fmt"

const _Fruit_name = "AppleOrangeBanana"

type Fruit int

var _Fruit_index = [...]uint8{5, 11, 17}

func (i Fruit) String() string {
	if i < 0 || i >= Fruit(len(_Fruit_index)) {
		return fmt.Sprintf("Fruit(%d)", i)
	}
	hi := _Fruit_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _Fruit_index[i-1]
	}
	return _Fruit_name[lo:hi]
}
