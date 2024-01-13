package partyrobot

import (
	"strconv"
    "fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	var str string = "Welcome to my party, " + name + "!"
	return str
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	var str string = "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
	return str
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
     return fmt.Sprintf("%s\nYou have been assigned to table %03d. Your table is %s, exactly %0.1f meters from here.\nYou will be sitting next to %s.",Welcome(name),table,direction,distance,neighbor)
}
