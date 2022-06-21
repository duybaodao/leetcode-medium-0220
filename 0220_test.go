package leetcode

import "testing"

type test struct {
	nums []int
	k    int
	t    int
	want bool
}

var tests = []test{
	{[]int{1, 2, 3, 1}, 3, 0, true},
	{[]int{1, 0, 1, 1}, 1, 2, true},
	{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
	{[]int{2147483646, 2147483647}, 3, 3, true},
	{[]int{1, 2, 1, 1}, 1, 0, true},
}

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
