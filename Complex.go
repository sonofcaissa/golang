package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	x := complex(2.5, 3.1)
	//y := complex(10.2, 2)
	//fmt.Println(x + y)
	//fmt.Println(x - y)
	//
	//fmt.Println(x * y)
	//fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

	//  Модуль комлексного числа вычисляется по формуле - |z| = sqrt( a*a + b*b ) где b с мнимой частью
	var xx = math.Sqrt(math.Pow(real(x), 2) + math.Pow(imag(x), 2))
	fmt.Println(xx)
}
