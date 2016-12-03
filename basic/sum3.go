package basic

import (
	"sort"
)
//忽略了交叉点！
func ThreeSum(nums []int) [][]int {
	result:=make([][]int,0)
	sort.Ints(nums)
	for k,v:=range nums{
		k2:=k+1
		k3:=len(nums)-1
		if k>0&&v==nums[k-1]{
			continue
		}
		for ;k2<k3;{
			if k2>k+1&&nums[k2]==nums[k2-1]{
				k2++
				continue
			}
			if k3<len(nums)-1&&nums[k3]==nums[k3+1]{
				k3--
				continue
			}
			sum:=nums[k2]+nums[k3]+v
			if sum>0{
				k3--
				continue
			}else if sum<0{
				k2++
				continue
			}else{
				temp:=[]int{v,nums[k2],nums[k3]}
				result=append(result,temp)
				k2++
				k3--
			}
		}
	}
	return result
}