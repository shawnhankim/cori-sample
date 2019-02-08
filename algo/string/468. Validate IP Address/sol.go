func isIP4(str string) bool {
    for _, t := range str {
        if (t >= '0' && t <= '9') {
            continue
        } else {
            return false
        }
    }
    val, _ := strconv.Atoi(str)
    if (val > 255 || val < 0 || (len(str) > 1 && str[0] == '0')) {
        return false
    }
    return true
}

func isIP6(str string) bool {
    if (len(str) > 4) {
        return false
    }
    return true
}

func validIPAddress(IP string) string {
    dcnt := 0
    ccnt := 0
    prev := ' '
    //s    := ' '
    //i    := 0
    
    var str string
    
    for _, s := range IP {
        switch s {
        case '.': 
            dcnt++
            if !isIP4(str) || ccnt > 0 || prev == s || dcnt > 3 {
                return "Neither"
            }
            str = ""
            prev = s
            continue

        case ':':
            ccnt++;
            if !isIP6(str) || dcnt > 0 || prev == s || ccnt > 7 {
                return "Neither"
            }
            str = ""
            prev = s
            continue
        }
        str += string(s)
        prev = s
    }

    strLen := len(str)
    if (ccnt == 7) {
        if strLen == 0 || (strLen > 0 && !isIP6(str)) {
            return "Neither" 
        }
        return "IPv6";
    }
    if (dcnt == 3) {
        if strLen == 0 || (strLen > 0 && !isIP4(str)) {
            return "Neither" 
        }
        return "IPv4";
    }
    return "Neither";
}
