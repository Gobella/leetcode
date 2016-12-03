package basic

import (
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	global_min:=1<<62
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
			temp:=sum-target

			if abs(temp)<abs(global_min){
				global_min=temp
			}
			if temp>0{
				k3--
				continue
			}else if temp<0{
				k2++
				continue
			}else{
				k2++
				k3--
				global_min=0
				return target
			}

		}
	}
	return global_min+target
}

func abs(i int) int{
	if i<0{
		return -i
	}
	return i
}