package leetCode

/*
	有效的括号
*/

func IsValid(s string) bool {

	tmp := make([]string, 0)
	flag := false
	for _, v := range s {
		if flag {
			break
		}
		if (string(v) == ")" || string(v) == "]" || string(v) == "}") && len(tmp) == 0 {
			flag = true
			break
		}
		vv := string(v)
		switch vv {
		case "(", "[", "{":
			tmp = append(tmp, string(v))
		case ")":
			if tmp[len(tmp)-1] == "(" && len(tmp) > 0 {
				tmp = tmp[:len(tmp)-1]
			} else {
				flag = true
			}
		case "]":
			if tmp[len(tmp)-1] == "[" {
				tmp = tmp[:len(tmp)-1]
			} else {
				flag = true
			}

		case "}":

			if tmp[len(tmp)-1] == "{" && len(tmp) > 0 {
				tmp = tmp[:len(tmp)-1]
			} else {
				flag = true
			}
		}

	}
	if len(tmp) > 0 || flag {
		return false
	}
	return true
}
