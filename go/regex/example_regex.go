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

	tmpURL := "/ui/user"
	if uiRegex.MatchString(tmpURL) {
		fmt.Printf("Mached 'ui' string in regex : %s\n", tmpURL)
	} else {
		fmt.Printf("Unmached 'ui' string in regex : %s\n", tmpURL)
	}
	if uiRegex.Match([]byte(tmpURL)) {
		fmt.Printf("Mached 'ui' pattern in regex : %s\n", tmpURL)
	} else {
		fmt.Printf("Unmached 'ui' pattern in regex : %s\n", tmpURL)
	}
	tmpURL1 := "/abcd/ui/user"
	if uiRegex.MatchString(tmpURL1) {
		fmt.Printf("Mached 'ui' string in regex : %s\n", tmpURL1)
	} else {
		fmt.Printf("Unmached 'ui' string in regex : %s\n", tmpURL1)
	}
	if uiRegex.Match([]byte(tmpURL1)) {
		fmt.Printf("Mached 'ui' pattern in regex : %s\n", tmpURL1)
	} else {
		fmt.Printf("Unmached 'ui' pattern in regex : %s\n", tmpURL1)
	}

	tmpDownload := "/downloads/user"
	if dnRegex.MatchString(tmpDownload) {
		fmt.Printf("Mached 'downloads' string in regex : %s\n", tmpDownload)
	} else {
		fmt.Printf("Unmached 'downloads' string in regex : %s\n", tmpDownload)
	}
	if uiRegex.Match([]byte(tmpDownload)) {
		fmt.Printf("Mached 'downloads' pattern in regex : %s\n", tmpDownload)
	} else {
		fmt.Printf("Unmached 'downloads' pattern in regex : %s\n", tmpDownload)
	}

	tmpDownload1 := "/abcd/downloads/user"
	if dnRegex.MatchString(tmpDownload) {
		fmt.Printf("Mached 'downloads' string in regex : %s\n", tmpDownload1)
	} else {
		fmt.Printf("Unmached 'downloads' string in regex : %s\n", tmpDownload1)
	}
	if uiRegex.Match([]byte(tmpDownload)) {
		fmt.Printf("Mached 'downloads' pattern in regex : %s\n", tmpDownload1)
	} else {
		fmt.Printf("Unmached 'downloads' pattern in regex : %s\n", tmpDownload1)
	}
}
