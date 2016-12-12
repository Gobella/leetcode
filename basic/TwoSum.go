package basic
/*Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution.*/
import (
	"sort"
)

type pair struct {
	index int
	value int
}

type pairList []*pair

func (list pairList) Len() int{
	return len(list)
}

func (list pairList) Swap(i,j int)  {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}
func (list pairList) Less(i,j int) bool  {
	if list[i].value < list[j].value {
		return true
	}else {
		return false
	}

}

func twoSum(nums []int, target int) []int {
	length := len(nums)

	var mynums pairList
	for i:=0;i<length ;i++  {
		mynums = append(mynums,&pair{i,nums[i]})
	}
	var result []int
	sort.Sort(mynums)
	i := 0
	j := length - 1
	for i < j{
		if  mynums[i].value + mynums[j] == target{
			result = append(result,mynums[i].index)
			result = append(result,mynums[j].index)
			break
		}else if mynums[i].value + mynums[j] < target{
			i++
		}else{
			j--;
		}


	}
	return result

}
