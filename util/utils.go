package util

// Find takes a slice and looks for an element in it. If found it will
// return true else false.
func Find(slice []string, val string) (bool, int) {
    for index, item := range slice {
        if item == val {
            return true, index
        }
    }
    return false, -1
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}