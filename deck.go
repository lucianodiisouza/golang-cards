package main

import "fmt"

type deck []string

func (cards deck) print() {
	for index, card := range cards {
		fmt.Println(index, card)
	}
}
