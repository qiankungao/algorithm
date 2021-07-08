package leetCode

import (
	"reflect"
	"testing"
)

func TestFindMinDifference(t *testing.T) {
	timePoints := []string{"23:59", "00:00"}
	FindMinDifference(timePoints)

}

func TestMinDeletionSize(t *testing.T) {
	strs := []string{"abc", "bce", "cae"}
	if res := MinDeletionSize(strs); res != 1 {
		t.Errorf("res is 1,not %d", res)
	}

	strs = []string{"cba", "daf", "ghi"}
	if res := MinDeletionSize(strs); res != 1 {
		t.Errorf("res is 1,not %d", res)
	}

	strs = []string{"a", "b"}
	if res := MinDeletionSize(strs); res != 0 {
		t.Errorf("res is 0,not %d", res)
	}

	strs = []string{"zyx", "wvu", "tsr"}
	if res := MinDeletionSize(strs); res != 3 {
		t.Errorf("res is 3,not %d", res)
	}

}

func TestFindWords(t *testing.T) {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	if res := FindWords(words); !reflect.DeepEqual(res, []string{"Alaska", "Dad"}) {
		t.Errorf("res is Alaska Dad,not %s", res)
	}

	words = []string{"omk"}
	if res := FindWords(words); !reflect.DeepEqual(res, []string{}) {
		t.Errorf("res is [],not %s", res)
	}
}

func TestShortestToChar(t *testing.T) {
	s := "loveleetcode"
	c := s[3]
	if res := ShortestToChar(s, c); !reflect.DeepEqual(res, []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}) {
		t.Errorf("res is [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0],not %d", res)
	}

	s = "aaab"
	c = s[3]
	if res := ShortestToChar(s, c); !reflect.DeepEqual(res, []int{3, 2, 1, 0}) {
		t.Errorf("res is [3,2,1,0],not %d", res)
	}

}

func TestSubdomainVisits(t *testing.T) {
	//cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	cpdomains := []string{"9001 discuss.leetcode.com"}
	SubdomainVisits(cpdomains)


}
