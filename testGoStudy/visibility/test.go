package visibility

import "fmt"

const PI = 3.145
const pi = 3.14
const _PI = 3.14

var P int = 1
var p int = 1

func private_function() {
	fmt.Println("only used in this package!")
}

func Public_fuction() {
	fmt.Println("used in anywhere!")
}
