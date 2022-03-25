package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// var m1 map[int]string   这样写无法给字典赋值 会报错
	hashTable := map[int]int{}

	for index,value := range nums{
		targetIndex,hasExisted := hashTable[target - value]
		if hasExisted {
			return []int{targetIndex,index}
		}else{
			hashTable[value] = index
		}
	}
	return []int{}
}

func isValid(s string) bool {
     n :=len(s)
     if n %2 ==1{
     	return false
	 }
	 pairs := map[byte]byte{
		 ')': '(',
		 ']': '[',
		 '}': '{',
	 }
	 stack := []byte{}
	 for i:=0;i<n;i++{
	 	value, existed := pairs[s[i]]
	 	if existed {
	 		if len(stack)!=0 && value == stack[len(stack)-1]{
	 			stack = stack[:len(stack)-1]
			}else{
				return false
			}
		}else {
			stack = append(stack, s[i])
		}
	 }
	 return len(stack)==0
}


func main(){
	nums := []int{2,7,11,15}
	target := 9
	index := twoSum(nums,target)
	fmt.Println(index)

	s := "()[]{}"
	result := isValid(s)
	fmt.Println(result)
}
