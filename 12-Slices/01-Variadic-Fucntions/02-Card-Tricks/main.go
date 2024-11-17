package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	var favCards []int
	favCards = []int{2, 6, 9}
	return favCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	item := slice[index]
	if index >= len(slice) {
		return -1
	}
	return item
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if len(slice)-1 < index {
		slice = append(slice, value)
		return slice
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)

}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var temp []int
	temp = slice[0 : index-1]
	var temp2 []int
	temp2 = slice[index:]
	return append(temp, temp2...)
}
