package main

import (
	"fmt"

	"github.com/jaden-song/learngo/accounts/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	base := "hello"
	dictionary.Add(base, "first")
	err := dictionary.Update(base, "second")
	if err != nil {
		fmt.Println(err)
	}
	err = dictionary.Delete("xx")
	fmt.Println(err)
}
