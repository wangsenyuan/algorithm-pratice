package main

import "fmt"

func main() {
	nums := []int{345, 234, 123}
	radixSort(nums, 3)
	fmt.Printf("%v\n", nums)
}

type node struct {
	val  int
	next *node
}

type bucket struct {
	head, tail *node
}

func radixSort(nums []int, numOfDigits int) {
	master := initMasterBucket(nums)

	buckets := make([]*bucket, 10)

	for i := range buckets {
		buckets[i] = &bucket{}
	}

	putIntoBucket := func(j int, val int) {
		buc := buckets[j]
		ptr := &node{val: val}
		if buc.tail == nil {
			buc.head = ptr
			buc.tail = ptr
		} else {
			buc.tail.next = ptr
			buc.tail = buc.tail.next
		}
	}
	distribute := func(i int) {
		for _, bucket := range buckets {
			bucket.head = nil
			bucket.tail = nil
		}
		base := 10
		for j := 0; j < i; j++ {
			base *= 10
		}
		ptr := master.head
		for ptr != nil {
			j := (ptr.val % base) % 10
			putIntoBucket(j, ptr.val)
			ptr = ptr.next
		}
	}

	coalesce := func() {
		master.head = nil
		master.tail = nil
		for _, buc := range buckets {
			for ptr := buc.head; ptr != nil; ptr = ptr.next {
				if master.head == nil {
					master.head = ptr
					master.tail = ptr
				} else {
					master.tail.next = ptr
					master.tail = master.tail.next
				}
			}
		}
	}
	for i := 1; i <= numOfDigits; i++ {
		distribute(i)
		coalesce()
	}

	i := 0
	for ptr := master.head; ptr != nil; ptr = ptr.next {
		nums[i] = ptr.val
		i++
	}

}

func initMasterBucket(nums []int) *bucket {
	master := &bucket{}
	var ptr *node
	for _, num := range nums {
		tmp := &node{val: num}
		if ptr == nil {
			ptr = tmp
			master.head = ptr
		} else {
			ptr.next = tmp
			ptr = tmp
		}
		master.tail = ptr
	}
	return master
}
