func longestPalindrome(s string) string {
	flags := [1000][1000]int{}
	length := len(s)
	var startPostion int
	var resultLen int
	for k,_:= range s  {
		flags[k][k] = 1
		startPostion = k
		resultLen = 1
	}
	for i:=0;i<length-1;i++{
		if s[i] == s[i+1] {
			flags[i][i+1] = 1
			startPostion = i
			resultLen = 2
		}
	}
	for len:=3;len<=length;len++{
		for i:=0;i<= length - len;i++{
			j:= len-1+i
			if s[j] == s[i] && flags[i+1][j-1] == 1 {
				flags[i][j] = 1
				startPostion = i
				resultLen = len
			}
		}
	}
    return s[startPostion:startPostion+resultLen]
}
