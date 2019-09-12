package main

import (
	"fmt"
	"regexp"
)

func main() {
	var partern, s string
	partern = `^[\+0-9]+$`
	s = "+0123112312"
	isMatch, err := regexp.MatchString(partern, s)
	fmt.Println("isMatch", isMatch, err)

	partern = `^v\.[A-Za-z0-9]+@vinid.net$`
	for _, e := range []string{
		"v.tungdt11@vinid.net",
        "v.tungdt11@vinid.nethaha",
        "1v.tungdt11@vinid.nethaha",
        "v.5@vinid.net",
        "v.@vinid.net",
	} {
		isMatch, err := regexp.MatchString(partern, e)
		fmt.Println("zzz", isMatch, e, err)
	}

	partern = "^(male|female)$"
	for _, e := range []string{
		"male1",
		"male",
		"female",
		"femalefemale",
	} {
		isMatch, err := regexp.MatchString(partern, e)
		fmt.Println("zzz", isMatch, e, err)
	}
}
