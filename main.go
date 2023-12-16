package main

import "fmt"

/*
	new		green
	worn	yellow
	broken	red
*/

type Wear string

const (
	New    Wear = "green"
	Worn   Wear = "yellow"
	Broken Wear = "red"
)

type Items struct {
	backpack []ItemType
}

type ItemType struct {
	name string
	wear Wear
}

func main() {
	fmt.Print("initial commit")
}
