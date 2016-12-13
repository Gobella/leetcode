func myAtoi(str string) int {
	//1 先去除 空格
	var result int64
	flag := 1
	if len(str)==0{
	    return 0
	}
	for i,value := range str{
		if value != ' ' {
			str = str[i:]
			break
		}
	}
	if str[0] == '-'  {
		flag = -1
		str = str[1:]
	}else if(str[0] == '+'){
		str = str[1:]
	}
	
	for _,v := range str{
		if v<'0' || v>'9'{
			break;
		}
		result = 10*result + int64(v - '0')
		if result >math.MaxInt32 && flag == 1 {
			return math.MaxInt32
		}else if(result >math.MaxInt32+1 && flag==-1){
			return math.MinInt32
		}
		
	}
	if flag == -1{
		result = result * -1
	}
	return int(result)

}
