package basic

/*
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container.
*/

func maxArea(height []int) int {
	left:=0;
	right:=len(height)-1
	area:=0
	for ;left<right;{
		temp:=(right-left)*min(height[left],height[right])
		if height[left]>=height[right]{
			right--
		}else{
			left++
		}
		if temp>area{
			area=temp
		}
	}
	return area

}
func min(a int,b int) int{
	if a>b{
		return b
	}
	return a
}