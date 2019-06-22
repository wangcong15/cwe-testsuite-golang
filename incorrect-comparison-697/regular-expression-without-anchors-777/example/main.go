package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path"
	"regexp"

	"github.com/wangcong15/goassert"
)

func getPathFromInput() string {
	// suppose we input `../../etc/passwd`
	return "../../etc/passwd"
}

func main() {
	pat := "/[A-Za-z0-9]+/"
	u, _ := user.Current()
	home := u.HomeDir
	p := getPathFromInput()
	if flag, _ := regexp.MatchString(pat, p); flag == true {
		goassert.AssertStrNotIn("..", p)
		// pass regular expression check to read security information
		filePath := path.Join(home, p)
		if b, err := ioutil.ReadFile(filePath); err == nil {
			fmt.Println(string(b))
		}
	}
}
