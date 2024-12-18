package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total int
	for i := range birdsPerDay {
		total += birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var res int
	for i := (week - 1) * 7; i < week*7; i++ {
		res = birdsPerDay[i]
	}
	return res
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i/2 != 0 {
			birdsPerDay[i] = birdsPerDay[i]+1 
		}
	}
	return birdsPerDay
}
