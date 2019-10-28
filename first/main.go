package main

import (
	"fmt"
	"homeworks/first/mylib"
)

func main() {
	fmt.Println("hello")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers) // [1 2 3 4]
	invNumbers := mylib.Invert(numbers)
	fmt.Println(invNumbers) // [4 3 2 1]

	fmt.Println(mylib.IsAnagram("abc", "cba")) // true
	fmt.Println(mylib.IsAnagram("qaz", "zse")) // false

	names := []string{} // cap=0, len=0
	fmt.Println(names, cap(names), len(names))
	names = mylib.AppendString(names, "Bob") // cap=1, len=1
	fmt.Println(names, cap(names), len(names))
	names = mylib.AppendString(names, "Charlie") // cap=2, len=2
	fmt.Println(names, cap(names), len(names))
	names = mylib.AppendString(names, "Jhon") // cap=4, len=3
	fmt.Println(names, cap(names), len(names))
	names = mylib.AppendString(names, "Roger") // cap=4, len=4
	fmt.Println(names, cap(names), len(names))
	names = mylib.AppendString(names, "Ric") // cap=8, len=5
	fmt.Println(names, cap(names), len(names))
	names = mylib.AppendString(names, "Zenny") // cap=8, len=6
	fmt.Println(names, cap(names), len(names))
	names = mylib.AppendString(names, "Abba") // cap=8, len=7
	fmt.Println(names, cap(names), len(names))
	names = mylib.AppendString(names, "Bobby") // cap=8, len=8
	fmt.Println(names, cap(names), len(names))
	names = mylib.AppendString(names, "Stewart") // cap=16, len=9
	fmt.Println(names, cap(names), len(names))
}
