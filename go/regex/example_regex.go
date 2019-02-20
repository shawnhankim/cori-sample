package main

import (
	"fmt"
	"regexp"
)

func main() {
	uiRegex, uiErr := regexp.Compile(`^\/ui\/.*`)
	if uiErr != nil {
		fmt.Printf("Unable to compile a temporary regex\n")
	}
	dnRegex, dnErr := regexp.Compile(`/downloads/*`)
	if dnErr != nil {
		fmt.Printf("Unable to compile a temporary regex\n")
	}

	fmt.Printf("\n1. Test Case : %s \n", `^\/ui\/.*`)
	tmpURL := "/ui/user"
	if uiRegex.MatchString(tmpURL) {
		fmt.Printf("MatchString() : Mached   'ui' string  in regex : %s\n", tmpURL)
	} else {
		fmt.Printf("MatchString() : Unmached 'ui' string  in regex : %s\n", tmpURL)
	}
	if uiRegex.Match([]byte(tmpURL)) {
		fmt.Printf("Match()       : Mached   'ui' pattern in regex : %s\n", tmpURL)
	} else {
		fmt.Printf("Match()       : Unmached 'ui' pattern in regex : %s\n", tmpURL)
	}

	tmpURL1 := "/abcd/ui/user"
	if uiRegex.MatchString(tmpURL1) {
		fmt.Printf("MatchString() : Mached   'ui' string  in regex : %s\n", tmpURL1)
	} else {
		fmt.Printf("MatchString() : Unmached 'ui' string  in regex : %s\n", tmpURL1)
	}
	if uiRegex.Match([]byte(tmpURL1)) {
		fmt.Printf("Match()       : Mached   'ui' pattern in regex : %s\n", tmpURL1)
	} else {
		fmt.Printf("Match()       : Unmached 'ui' pattern in regex : %s\n", tmpURL1)
	}

	fmt.Printf("\n2. Test Case : %s \n", `/downloads/*`)
	tmpDownload := "/downloads/user"
	if dnRegex.MatchString(tmpDownload) {
		fmt.Printf("MatchString() : Mached   'downloads' string  in regex : %s\n", tmpDownload)
	} else {
		fmt.Printf("MatchString() : Unmached 'downloads' string  in regex : %s\n", tmpDownload)
	}
	if uiRegex.Match([]byte(tmpDownload)) {
		fmt.Printf("Match()       : Mached   'downloads' pattern in regex : %s\n", tmpDownload)
	} else {
		fmt.Printf("Match()       : Unmached 'downloads' pattern in regex : %s\n", tmpDownload)
	}

	tmpDownload1 := "/abcd/downloads/user"
	if dnRegex.MatchString(tmpDownload) {
		fmt.Printf("MatchString() : Mached   'downloads' string  in regex : %s\n", tmpDownload1)
	} else {
		fmt.Printf("MatchString() : Unmached 'downloads' string  in regex : %s\n", tmpDownload1)
	}
	if uiRegex.Match([]byte(tmpDownload)) {
		fmt.Printf("Match()       : Mached   'downloads' pattern in regex : %s\n", tmpDownload1)
	} else {
		fmt.Printf("Match()       : Unmached 'downloads' pattern in regex : %s\n", tmpDownload1)
	}
}
