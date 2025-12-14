package page

import "fmt"

func Pagination(offset, size, count int) (string, string) {
	var previousLink string
	var nextLink string

	if offset == 0 {
		previousLink = ""
	} else {
		previousLink = fmt.Sprintf("page=%d&size=%d", offset-size, size)
	}

	if count-offset < size {
		nextLink = ""
	} else {
		nextLink = fmt.Sprintf("page=%d&size=%d", offset+size, size)
	}

	return previousLink, nextLink
}
