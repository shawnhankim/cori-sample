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
	uxRegexStr = `^\/api\/sample\/[0-9a-z]+\/ux`
	//uxRegexStr = `^\/api\/sample\/[0-9a-z]+\/ux\/?.*`
	//uxRegexStr = `/api/sample/v1alpha1/ux`
	//uxRegexStr = `^\/api\/sample\/v1alpha1\/ux\/?.*`
)

var (
	regexList = make(map[string]*regexp.Regexp)
)

func sampleRegexMatching() {
	uxReg, uxErr := regexp.Compile(uxRegexStr)
	if uxErr != nil {
		fmt.Printf("Unable to compile a temporary regex\n")
	}

	fmt.Printf("\n1. Test Case : %s \n", uxRegexStr)
	tmpURL := "/ux/user"
	if uxReg.Match([]byte(tmpURL)) {
		fmt.Printf("Mached   'ux' pattern in regex : %s\n", tmpURL)
	} else {
		fmt.Printf("Unmached 'ux' pattern in regex : %s\n", tmpURL)
	}

	fmt.Printf("\n2. Test Case : %s \n", uxRegexStr)
	tmpURL1 := "/api/sample/v1alpha1/ux/user"
	if uxReg.Match([]byte(tmpURL1)) {
		fmt.Printf("Mached   'ux' pattern in regex : %s\n", tmpURL1)
	} else {
		fmt.Printf("Unmached 'ux' pattern in regex : %s\n", tmpURL1)
	}

	fmt.Printf("\n3. Test Case : %s \n", uxRegexStr)
	tmpURL2 := "/api/sample/v1alpha1/ux"
	if uxReg.Match([]byte(tmpURL2)) {
		fmt.Printf("Mached   'ux' pattern in regex : %s\n", tmpURL2)
	} else {
		fmt.Printf("Unmached 'ux' pattern in regex : %s\n", tmpURL2)
	}

	fmt.Printf("\n4. Test Case : %s \n", uxRegexStr)
	tmpURL3 := "/api/sample/v1alpha2/ux/"
	if uxReg.Match([]byte(tmpURL3)) {
		fmt.Printf("Mached   'ux' pattern in regex : %s\n", tmpURL3)
	} else {
		fmt.Printf("Unmached 'ux' pattern in regex : %s\n", tmpURL3)
	}

}

func sampleRegexMapping() {
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

func main() {
	sampleRegexMatching()
}
