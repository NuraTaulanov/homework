package main

type Room struct {
	name        string
	description string
	status      string
	exits       []string
	items       map[string][]string
}

type Player struct {
	currentRoom string
	Exits       map[string]Room
	DoorIsOpen  bool
	HasBackPack bool
	inventory   []string
	Task        string
}
