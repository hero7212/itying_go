package main

import (
	"os"
)

func main() {

	// os.Mkdir("./abc", 0666)

	os.MkdirAll("./dir1/dir2/dir3", 0666)

}
