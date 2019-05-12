package main

import iso "./isopakke"

func main() {

	stringliteral := iso.GetExtASCIIStringLiteral()
	iso.IterateOverExtendedASCIIStringLiteral(stringliteral)
	iso.GreetingExtendedASCII()

}
