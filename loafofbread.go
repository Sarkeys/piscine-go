package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	ans := ""
	count := 0
	for i, char := range str {
		if char != ' ' && count != 5 {
			ans += string(char)
			count++
		} else if count == 5 {
			ans += " "
			count = 0
		}
		if i == len(str)-1 && len(ans) > 0 && ans[len(ans)-1] == ' ' {
			ans = ans[:len(ans)-1]
		}
	}
	ans += "\n"
	return ans
}
