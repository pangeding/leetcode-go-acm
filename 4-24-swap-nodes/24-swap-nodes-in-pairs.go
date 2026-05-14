package swapNode

import {

}

func main() {

	var nums int[]
	
	head := buildList(nums)
	newHead := swapPairs(head)
	printList(newHead)
}