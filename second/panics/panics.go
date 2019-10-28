package panics

import (
	"fmt"
	"strings"
)

type MyPanicParams struct {
	position int
}

func ProcessPanicSong(song string) {
	word := "panic"

	position := 0
	for {
		foundIndex := strings.Index(song[position:], word)
		if foundIndex == -1 {
			return
		}
		position += foundIndex
		panicAndRecover(position)
		position++
	}
}

func panicAndRecover(wordIndex int) {
	defer func() {
		p := recover()

		switch p {
		case nil:
			fmt.Println("it's ok")
		case MyPanicParams{position: wordIndex}:
			fmt.Println("caught my panic param", p)
		default:
			fmt.Println("DID NOT MATCH")
			panic(p)
		}
	}()

	myPanicParams := MyPanicParams{position: wordIndex}

	panic(myPanicParams)
}
