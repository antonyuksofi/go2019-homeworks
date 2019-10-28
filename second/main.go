package main

import (
	"fmt"
	"homeworks/second/aggregatefuncs"
	"homeworks/second/morse"
	"homeworks/second/panics"
)

func main() {
	testMinMax()

	testMorse()

	testPanics()
}

func testMinMax() {
	fmt.Println("min:", aggregatefuncs.Max(2, 3, 1, 7, 0, 1, 8))

	fmt.Println("max:", aggregatefuncs.Max(2, 3, 1, 7, 0, 1, 8))
}

func testMorse() {

	rawCode := "--. --- / .-- .- ... / -.. . ... .. --. -. . -.. / .- - / --. --- --- --. .-.. . / - --- / .. -- .--. .-. --- ...- . / .--. .-. --- --. .-. .- -- -- .. -. --. / .--. .-. --- -.. ..- -.-. - .. ...- .. - -.--"
	result, err := morse.Translate(morse.GetCode(rawCode))
	if err != nil {
		fmt.Println("Error!")
		return
	}
	fmt.Println(result)
}

func testPanics() {
	/*song := [...]string{
		"I feel something's wrong",
		"But the panic is on",
		"I feel something's wrong",
		"But the panic is on",
		"But the panic is on",
		"But the panic is on",

		"I feel something's wrong",
		"But the panic is on",
		"I feel something's wrong",
		"But the panic is on",
	}*/

	song := "I feel something's wrong But the panic is on I feel something's wrong But the panic is on But the panic is on But the panic is on I feel something's wrong But the panic is on I feel something's wrong But the panic is on"

	panics.ProcessPanicSong(song)

	// fmt.Println(song)
}
