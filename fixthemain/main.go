package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(Door *Door) {
	PrintStr("Door Opening...")
	Door.state = OPEN
}

func CloseDoor(Door *Door) {
	PrintStr("Door Closing...")
	Door.state = CLOSE
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	return Door.state
}

func IsDoorClose(Door *Door) bool {
	PrintStr("is the Door closed ?")
	return Door.state
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
}

const OPEN, CLOSE = true, false
