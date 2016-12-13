func reverse(x int) int{
	var result int64
	for x!=0{
		result = 10*result +int64(x%10)
		x = x / 10
		fmt.Println(result)
	}
	if result<math.MinInt32 {
		return 0
	}
	if result> math.MaxInt32{
		return 0
	}
	return int(result)
}
