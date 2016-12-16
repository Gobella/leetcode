package basic

func If(b bool, t, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}


func lengthOfLongestSubstring(s string) int {
	var myMap = make(map[string] int)
	length := len(s)
	lastPosition := -1
	maxLen := -1
	i:=0
	for i=0; i<length ; i++ {
		if index,ok := myMap[string(s[i])];ok&& index > lastPosition {
			if maxLen < i-lastPosition -1 {
				maxLen = i-lastPosition -1
			}
			lastPosition = index;
		}
		myMap[string(s[i])] = i;
	}
	if maxLen < i-lastPosition -1 {
		maxLen = i-lastPosition -1
	}
	return maxLen
}