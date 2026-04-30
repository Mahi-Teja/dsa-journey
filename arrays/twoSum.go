package arrays
 
func TwoSum(nums []int, target int) []int { 
    // indecise sum = target : brute force
    // tc O(n^2)
    //  sc O(n)
    for  i :=0;i<len(nums);i++{ 
        for j:=0;j<len(nums);j++{
			  if(i==j) {continue}
            if(nums[i]+nums[j]==target){
				return []int{i,j}
            }
        }
    } 
	return  []int{}

}