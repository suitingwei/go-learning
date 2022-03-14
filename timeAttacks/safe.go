package timeAttacks

func CheckPwd(input string) bool {

	if len(input) != len(pwd) {
		return false
	}
	for i := 0; i < len(pwd); i++ {
		if pwd[i] != input[i] {
			return false
		}
	}
	return true
}
