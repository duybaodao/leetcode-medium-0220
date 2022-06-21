package leetcode

import (
	"math"
	"sync"
)

/*
	Given an integer array nums and two integers k and t, return true if there are
	two distinct indices i and j in the array such that:
	--- abs(nums[i] - nums[j]) <= t
	--- abs(i - j) <= k

	Ref: https://leetcode.com/problems/contains-duplicate-iii/
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// return containsNearbyAlmostDuplicate_bruceForce(nums, k, t)
	return containsNearbyAlmostDuplicate_multiThreads(nums, k, t)
}

/*
	containsNearbyAlmostDuplicate_bruceForce is a solution for containsNearbyAlmostDuplicate
	Solution:
	- Iterate through nums, and for each iteration, sub iterate one more to check
	- Time complexity: O( nums.length * k )
*/
func containsNearbyAlmostDuplicate_bruceForce(nums []int, k int, t int) bool {
	// Iterate through nums
	for i := 0; i < len(nums)-1; i++ {
		// And for each iteration, sub iterate one more to check
		// Add loop ending condition ( j-i <= k ) to clear the condition of:
		// --- abs(i - j) <= k
		for j := i + 1; (j < len(nums)) && (j-i <= k); j++ {
			// Check for condition:
			// --- abs(nums[i] - nums[j]) <= t
			diff := math.Abs(float64(nums[i] - nums[j]))
			if diff <= float64(t) {
				// All conditions cleared, return true
				return true
			}
		}
	}

	// All iteration done, but no valid result found -> return false
	return false
}

// Number of concurrent workers for containsNearbyAlmostDuplicate_multiThreads
const CON_COUNT = 8

/*
	containsNearbyAlmostDuplicate_multiThreads uses the same base algorithm as
	containsNearbyAlmostDuplicate_bruceForce to solve. But it will splits the tasks across sub-routines.
*/
func containsNearbyAlmostDuplicate_multiThreads(nums []int, k int, t int) bool {
	// Create channel to hold cancel signal
	cancelChan := make(chan bool, 1)
	// And upon existing, send cancel signal to close all child routine
	defer func() {
		cancelChan <- true
	}()

	// Create channel of tasks
	taskChan := make(chan int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		taskChan <- i
	}

	// Create channel for worker to send result to
	resultChan := make(chan bool, 1)

	// Use WaitGroup to check when all worker done
	var wg sync.WaitGroup
	go func() {
		// When all done, it means no valid result found, return false
		wg.Wait()
		resultChan <- false
	}()

	// Define task of worker
	workerTask := func(index int) bool {
		for j := index + 1; (j < len(nums)) && (j-index <= k); j++ {
			diff := math.Abs(float64(nums[index] - nums[j]))
			if diff <= float64(t) {
				return true
			}
		}
		return false
	}

	// Start workers
	for i := 0; i < CON_COUNT; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case index := <-taskChan:
					result := workerTask(index)
					if result {
						resultChan <- result
						return
					}
				case cancel := <-cancelChan:
					if cancel {
						return
					}
				default:
					return
				}
			}
		}()
	}

	// Wait for result
	result := <-resultChan
	return result
}
