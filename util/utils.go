package util

// Find takes a slice and looks for an element in it. If found it will
// return true else false.
func Find(slice []string, val string) (bool) {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}