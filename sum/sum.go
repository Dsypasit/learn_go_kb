package sum

func sum(args ...int) int {
	ans := 0
	for _, v := range args {
		ans += v
	}
	return ans
}
