func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)
    for i:=0;i<len(nums);i=i+1 {
        m[target - nums[i]]=i
    }
    //fmt.Println("map:", m)
    for i:=0;i<len(nums); i=i+1 {
        v, prs := m[nums[i]]
        //fmt.Println("prs", prs, "v", v)
        if prs && i!=v {
            res := []int{i,v}
            return res
        }
    }
    return []int{}
}
