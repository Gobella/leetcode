package basic


func longestCommonPrefix(strs []string) string {
	if strs==nil||len(strs)==0{
		return ""
	}

	one:=[]byte(strs[0])
	index:=0

	for k,v:=range one{
		for j:=1;j<len(strs);j++{
			temp:=[]byte(strs[j])
			if len(strs[j])==k||temp[k]!=v{
				index=k
				return string(one[0:index])
			}
		}
	}
	return strs[0]
}
