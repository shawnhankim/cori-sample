package main

import (
	"fmt"
	"regexp"
)

/*

Regex

\/         ------- will match for a "/"
^\/        ------- will match for a "/" at the beginning of the line.
[^\/]      ------- ^ inside a [ ] will be for negation(opposite of). it will match for anything except a "/"
[^\/]+     ------- it will match for, one or more characters, anything except a "/"
([^\/]+)   ------- the rex matching should be put inside the flower bracket "( )".
([^\/]+)\/ ------- the "\/" tells that, match till this "/". (it will match only AAAA or whatever before a first "/".)

*/
const (
	uiRegexStr = `^\/ui\/.*`
	dnRegex    = `/downloads/*`
)

var (
	regexList = make(map[string]*regexp.Regexp)
)

func main() {
	var uiErr, dnErr error
	regexList[uiRegexStr], uiErr = regexp.Compile(uiRegexStr)
	if uiErr != nil {
		fmt.Printf("Unable to compile a temporary regex\n")
	}
	regexList[dnRegex], dnErr = regexp.Compile(dnRegex)
	if dnErr != nil {
		fmt.Printf("Unable to compile a temporary regex\n")
	}

	fmt.Printf("\n1. Test Case : %s \n", uiRegexStr)
	tmpURL := "/ui/user"
	if regexList[uiRegexStr].MatchString(tmpURL) {
		fmt.Printf("MatchString() : Mached   'ui' string  in regex : %s\n", tmpURL)
	} else {
		fmt.Printf("MatchString() : Unmached 'ui' string  in regex : %s\n", tmpURL)
	}
	if regexList[uiRegexStr].Match([]byte(tmpURL)) {
		fmt.Printf("Match()       : Mached   'ui' pattern in regex : %s\n", tmpURL)
	} else {
		fmt.Printf("Match()       : Unmached 'ui' pattern in regex : %s\n", tmpURL)
	}

	tmpURL1 := "/abcd/ui/user"
	if regexList[uiRegexStr].MatchString(tmpURL1) {
		fmt.Printf("MatchString() : Mached   'ui' string  in regex : %s\n", tmpURL1)
	} else {
		fmt.Printf("MatchString() : Unmached 'ui' string  in regex : %s\n", tmpURL1)
	}
	if regexList[uiRegexStr].Match([]byte(tmpURL1)) {
		fmt.Printf("Match()       : Mached   'ui' pattern in regex : %s\n", tmpURL1)
	} else {
		fmt.Printf("Match()       : Unmached 'ui' pattern in regex : %s\n", tmpURL1)
	}

	fmt.Printf("\n2. Test Case : %s \n", dnRegex)
	tmpDownload := "/downloads/user"
	if regexList[dnRegex].MatchString(tmpDownload) {
		fmt.Printf("MatchString() : Mached   'downloads' string  in regex : %s\n", tmpDownload)
	} else {
		fmt.Printf("MatchString() : Unmached 'downloads' string  in regex : %s\n", tmpDownload)
	}
	if regexList[dnRegex].Match([]byte(tmpDownload)) {
		fmt.Printf("Match()       : Mached   'downloads' pattern in regex : %s\n", tmpDownload)
	} else {
		fmt.Printf("Match()       : Unmached 'downloads' pattern in regex : %s\n", tmpDownload)
	}

	tmpDownload1 := "/abcd/downloads/user"
	if regexList[dnRegex].MatchString(tmpDownload) {
		fmt.Printf("MatchString() : Mached   'downloads' string  in regex : %s\n", tmpDownload1)
	} else {
		fmt.Printf("MatchString() : Unmached 'downloads' string  in regex : %s\n", tmpDownload1)
	}
	if regexList[dnRegex].Match([]byte(tmpDownload)) {
		fmt.Printf("Match()       : Mached   'downloads' pattern in regex : %s\n", tmpDownload1)
	} else {
		fmt.Printf("Match()       : Unmached 'downloads' pattern in regex : %s\n", tmpDownload1)
	}
}
