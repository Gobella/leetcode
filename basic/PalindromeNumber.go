func isPalindrome(x int) bool {
	var temp int64
	var result int64
	temp = int64(x)
	if x < 0{
	    return false
	}
	for temp != 0{
		result = 10*result + temp%10
		temp = temp / 10
	}
	if int(result) == x{
		return true
	}else {
		return false
	}
}
