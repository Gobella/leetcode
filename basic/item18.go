package basic

import (
	"sort"
)

//http://blog.csdn.net/hcbbt/article/details/44063375 hash
func FourSum(nums []int, target int) [][]int {
	result:=make([][]int,0)
	length:=len(nums)
	sort.Ints(nums)
	for k,v:=range nums{
		num1:=v-target
		if k>0&&nums[k]==nums[k-1]{
			continue
		}
		for i:=k+1;i<length-2;i++{
			num2:=num1+nums[i]
			num3Index:=i+1
			num4Index:=length-1
			if i>k+1&&nums[i-1]==nums[i]{
				continue
			}
			for ;num3Index<num4Index;{
				if num3Index>i+1&&nums[num3Index-1]==nums[num3Index]{
					num3Index++
					continue
				}
				if num4Index<length-1&&nums[num4Index+1]==nums[num3Index]{
					num4Index--
					continue
				}
				sum:=num2+nums[num4Index]+nums[num3Index]
				if sum>0{
					num4Index--
				}else if sum<0{
					num3Index++
				}else{
					result=append(result,[]int{v,nums[i],nums[num3Index],nums[num4Index]})
					num3Index++
					num4Index--
				}
			}
		}
	}
	return result

}
