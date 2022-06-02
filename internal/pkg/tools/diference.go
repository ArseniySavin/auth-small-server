package tools

// GetDifference -
func GetDifference(c1, c2 []string) []string {
	var res []string
	for _, item := range c1 {
		doAdd := false
		for _, gItem := range c2 {
			if gItem == item {
				doAdd = false
				break
			}
			doAdd = true
		}
		if doAdd {
			res = append(res, item)
		}
	}
	return res
}
