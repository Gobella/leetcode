func convert(s string, numRows int) string {
	resultArr := make([]string,numRows )
	flag :=1
	row:=0
	if numRows == 1{
	    return s
	}
	for i:=0;i<len(s);i++{
		resultArr[row]+= string(s[i])
		if row == numRows-1 {
			flag = -1
		}else if row == 0{
			flag = 1
		}
		row += flag
	}
	var result string
	for _,v := range resultArr{
		result += v
	}
	return result
}
