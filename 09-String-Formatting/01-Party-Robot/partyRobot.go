package partyrobot
import (
    "fmt"
    "math"
)
// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name+"!"
}
// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}
// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	formattedTable := fmt.Sprintf("%03d", table)
  // Round distance to one decimal place
  roundedDistance := math.Round(distance*10) / 10
    return fmt.Sprintf(Welcome(name) + "\n" + (fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %.1f meters from here.", formattedTable, direction, roundedDistance))+"\n"+ (fmt.Sprintf("You will be sitting next to %s.", neighbor)))
}