package leetcode

import (
	"log"
	"path"
	"testing"

	"github.com/emirpasic/gods/trees/avltree"
	"github.com/kr/pretty"
)

func Test93(t *testing.T) {
	Solve93("102535011")
}

func Test1143(t *testing.T) {
	// r := leetcode.Solve1143("ace", "abcde")
	r := Solve1143("oxcpqrsvwf", "shmtulqrypy")
	t.Log(r)
}

func Test152(t *testing.T) {
	nums := []int{-1, -1, 0, -1, 2, 3, -5}
	r := MaxProduct(nums)
	t.Log(r)
}

func Test213(t *testing.T) {
	nums := []int{2, 2, 4, 3, 2, 5}
	r := RobII(nums)
	t.Log(r)
}

func Test263(t *testing.T) {
	n := 16
	t.Log(IsUgly(n))
}

func Test19(t *testing.T) {
	t.Setenv("GOEXPERIMENT", "loopvar")
	list := ListNodeFromSlice(1)
	RemoveNthFromEnd(list, 1)
	pretty.Println(list)
}

func bin_search(arr []int, target int, start int, end int) int {
	if start > end {
		return -1
	}

	if start == end {
		if arr[start] == target {
			return start
		}
		return -1
	}

	if (target <= arr[end] && target >= arr[start]) ||
		(arr[start] >= arr[end]) {
		e := int((end-start)/2) + start
		l := bin_search(arr, target, start, e)
		if l != -1 {
			return l
		}
		return bin_search(arr, target, e+1, end)
	}

	return -1
}

func search(arr []int, target int) int {
	return bin_search(arr, target, 0, len(arr)-1)
}

func TestSearch(t *testing.T) {
	nums := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	r := search(nums, 5)
	t.Log(r)
}

func Test264(t *testing.T) {
	r := NthUglyNumber(10)
	t.Log(r)
}

func Test279(t *testing.T) {
	r := PerfectSquare(13)
	t.Log(r)
}

func Test322(t *testing.T) {
	coins := []int{2}
	amount := 3
	r := CoinChange322(coins, amount)
	t.Log(r)
}

func Test309(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5, 6, 6, 9, 2, 3, 4, 5, 33, 4, 2, 13, 45, 23, 4512, 132, 12, 234, 24, 24, 324, 24, 4, 6, 77, 8, 3}
	r := MaxProfix_309(prices)
	t.Log(r)
}

func Test368(t *testing.T) {
	nums := []int{4, 8, 10, 240}
	r := LargestDivisibleSubset(nums)
	pretty.Logln(r)
}

func Test376(t *testing.T) {
	nums := []int{1, 1, 7, 4, 9, 2, 5}
	r := WiggleMaxLength(nums)
	pretty.Logln(r)
}

func Test413(t *testing.T) {
	nums := []int{1, 2, 3, 4, 6, 8, 10, 12}
	r := NumberOfArithmeticSlices(nums)
	pretty.Logln(r)
}

func Test416(t *testing.T) {
	nums := []int{3, 3, 3, 4, 5}
	r := CanPartition(nums)
	pretty.Logln(r)
}

func Test1(t *testing.T) {
	nums := []int{3, 2, 4}
	r := TwoSum(nums, 6)
	pretty.Logln(r)
}

func Test5(t *testing.T) {
	s := "a"
	r := LongestPalindrome(s)
	pretty.Logln(r)
}

func shortestBeautifulSubstring(s string, k int) string {
	res := ""
	dp := make([]int, len(s)+1)
	dp[0] = -k
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			dp[i+1] = dp[i]
		} else {
			l := dp[i] + 1
			if l >= 0 {
				for s[l] == '0' {
					l++
				}
			}
			dp[i+1] = l
		}

		if dp[i] >= 0 {
			sub := s[dp[i] : i+1]
			if res >= sub {
				res = sub
			}
		}
	}

	return res
}

func TestCm(t *testing.T) {
	pretty.Logln(shortestBeautifulSubstring("100011001", 3))
	// pretty.Logln(s2i("11"))
}

func Test996(t *testing.T) {
	nums := []int{1, 1, 8, 1, 8}
	r := numSquarefulPerms(nums)
	pretty.Logln(r)
}

func Test907(t *testing.T) {
	arr := []int{81, 55, 2}
	r := sumSubarrayMins(arr)
	pretty.Logln(r)
}

func Test1670(t *testing.T) {
	q := NewFrontMiddleBackQueue()
	q.PushMiddle(803226)
	print_queue(&q)
	q.PushMiddle(720129)
	print_queue(&q)
	q.PushMiddle(626048)
	print_queue(&q)
	q.PopMiddle()
	print_queue(&q)
	q.PushMiddle(860306)
	print_queue(&q)
	q.PopBack()
	print_queue(&q)
	q.PushMiddle(630886)
	print_queue(&q)
	q.PopMiddle()
	print_queue(&q)
	q.PopMiddle()
	print_queue(&q)
	q.PopMiddle()
	print_queue(&q)
}

func Test256(t *testing.T) {
	costs := [][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}
	r := minCost(costs)
	pretty.Logln(r)
}

func Test2366(t *testing.T) {
	tree := avltree.NewWithIntComparator()
	tree.Put(1, nil)
	tree.Put(1, nil)
	tree.Put(2, nil)
	pretty.Println(tree.Keys()...)
}

func Test15(t *testing.T) {
	nums := []int{-5853, -5787, -5595, -5586, -5412, -5372, -5220, -5171, -4988, -4970, -4933, -4895, -4794, -4781, -4779, -4641, -4621, -4589, -4392, -4239, -3902, -3680, -3541, -3464, -3455, -3454, -3255, -3196, -2986, -2940, -2916, -2898, -2625, -2406, -2402, -2397, -2162, -2110, -1900, -1182, -1108, -1080, -1076, -893, -869, -754, -637, -487, -387, -363, -302, -33, 340, 1965, 2149, 2441, 4024, 5764, 5813, 5825, 5895, 6141, 6335, 6476, 7082, 7186, 7209, 7558, 8712, 8883, 8967, 9328, 9371, 9540, 9664, 9914, 10517, 10693, 10779, 11063, 11129, 11239, 11683, 13330, 14573, 14685, 14783, 15157, 15758, 16172, 16245, 17215, 18575, 18807, 18882, 19191, 19322, 19908, 20106, 20456, 20572, 21249, 21725, 21796, 21916, 22589, 23257, 23842, 24736, 25187, 25195, 25413, 25667}
	res := threeSum(nums)
	pretty.Println(res)
}

func Test46(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	r := permute(nums)
	pretty.Logln(r)
}

func TestSliceAssign(t *testing.T) {
	a := []int{1, 2, 3}
	as := [][]int{}
	as = append(as, a)
	a = nil
	log.Println(as)
}

func Test2661(t *testing.T) {
	arr := []int{1, 3, 4, 2}
	mat := [][]int{
		{1, 4},
		{2, 3},
	}

	r := firstCompleteIndex(arr, mat)
	pretty.Logln(r)
}

func Test32(t *testing.T) {
	s := ""
	r := longestValidParentheses(s)
	pretty.Logln(r)
}

func TestTestCaseReader(t *testing.T) {
	reader := open_testcase[string](32, false)
	if reader == nil {
		return
	}
	defer reader.close()
	log.Println(reader.next())
}

func Test2952(t *testing.T) {
	coins := []int{1, 4, 10}
	r := minimumAddedCoins(coins, 19)
	pretty.Logln(r)
}

func Test122(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	pretty.Logln(solve122(prices))
}

func Test55(t *testing.T) {
	nums := []int{1, 1, 1, 0}
	pretty.Logln(canJump(nums))
}

func Test2772(t *testing.T) {
	nums := []int{2, 2, 3, 1, 1, 0}
	pretty.Logln(checkArray(nums, 3))
}

func Test20(t *testing.T) {
	s := "{{{{{{{((((((([[[[[[]]]]]])))))))}}}}}}}"
	pretty.Logln(isValid(s))
}

func Test51(t *testing.T) {
	pretty.Logln(solveNQueens(4))
}

func Test277(t *testing.T) {
	pretty.Logln(numWays(3, 2))
}

func Test56(t *testing.T) {
	t.Setenv("GOEXPERIMENT", "loopvar")
	intervals := [][]int{
		{2, 6}, {8, 10}, {15, 18}, {1, 3},
	}
	pretty.Logln(merge_interval(intervals))
}

func BenchmarkMergeIntervals(b *testing.B) {
	intervals := [][]int{{1, 3}, {2, 6}, {15, 18}, {8, 10}}

	for i := 0; i < b.N; i++ {
		merge_interval(intervals)
	}
}

func TestPlay(t *testing.T) {
}

func Test65(t *testing.T) {
	numstr := "-0.1"
	pretty.Logln(isNumber(numstr))
}

func Test7(t *testing.T) {
	reverse(1)
}

func Test71(t *testing.T) {
	p := "/home/../../.."
	pretty.Logln(simplifyPath(p))
	pretty.Logln(path.Clean(p))
	pretty.Logf("%x", 0.1)
}

func Test2976(t *testing.T) {
	source := "nepafdoolxkrvapckfrxxjxjvbreguxoxqtfjqrvsthvkcgllyfismhvgcyhgclzqqojegtxwrdvywpzkcqptvvrufmwvombolbp"
	target := "mdrwpvrexhwksjmgazaqbqbxpqffskkjxgzgeoqelewriqhfcipthoupassfwamiiivmhmeidcgekwjomkzkozfkjnofqthkuhwq"
	origin := []byte{'y', 'k', 'j', 'n', 'e', 'w', 'u', 'e', 'u', 'v', 'h', 'v', 'l', 'g', 'v', 'o', 'o', 'f', 'q', 'g', 'o', 'x', 'x', 'q', 'q', 'k', 's', 'x', 'j', 'w', 'g', 'g', 'f', 'x', 'l', 'u', 'g', 'n', 'n', 'm', 'u', 'b', 'r', 'v', 'r', 'q', 'e', 'd', 'y', 'u', 'u', 'b', 'j', 'x', 'q', 'w', 'a', 'y', 'm', 'a', 'l', 'g', 't', 'q', 'l', 't', 'e', 'k', 'q', 'e', 'o', 'c', 'j', 'a', 'm', 'h', 'j', 'j', 'v', 'w', 'e', 'r', 'a', 's', 'e', 'z', 'e', 'q', 'h', 'l', 'd', 'j', 'd', 'm', 'y', 'f', 'o', 'o', 'j', 'z', 'o', 'l', 'a', 'r', 'f', 'g', 'q', 'p', 'q', 'm', 'l', 'j', 'z', 'y', 'u', 'j', 'u', 'p', 'l', 'k', 'p', 'v', 'z', 's', 'l', 'e', 't', 's', 'p', 'n', 'f', 'l', 'h', 'i', 'a', 's', 'q', 'o', 'c', 'j', 't', 'u', 't', 'w', 'u', 'd', 'z', 't', 'k', 'y', 'h', 'a', 'o', 'y', 'b', 'r', 'z', 'k', 'r', 'w', 'z', 'l', 'o', 'c', 'w', 's', 'q', 'b', 'z', 'b', 'v', 'n', 'v', 'w', 'l', 't', 'n', 'v', 'r', 'e', 'n', 'k', 'o', 'u', 'e', 'i', 'r', 'c', 'r', 'u', 'u', 'w', 'm', 'r', 'b', 'b', 'j', 'c', 'g'}
	changed := []byte{'b', 'c', 'k', 'd', 't', 'e', 'v', 'f', 'n', 'y', 'x', 'h', 'q', 'k', 'o', 'i', 'z', 'd', 'm', 'p', 'l', 'k', 'z', 'r', 'a', 'f', 'k', 'w', 'n', 't', 'c', 'r', 'p', 'y', 'b', 'm', 'h', 'e', 'e', 'k', 'h', 'z', 'w', 'g', 'n', 'h', 's', 'y', 'w', 'p', 'e', 'w', 'h', 'r', 'd', 'j', 'i', 'p', 'k', 'e', 'b', 'm', 'r', 'b', 'b', 'c', 'n', 'y', 'x', 'k', 'a', 'z', 'r', 'f', 'j', 'f', 'a', 'm', 'q', 'g', 'z', 'j', 'c', 'j', 's', 'd', 'q', 'r', 'c', 'n', 'g', 'e', 'i', 'c', 'c', 'u', 'e', 'g', 'b', 'a', 'l', 'm', 's', 'c', 'r', 'i', 'u', 'f', 'o', 'h', 'p', 'f', 'd', 't', 'e', 'h', 's', 'a', 'i', 'd', 'j', 'm', 'u', 'n', 'v', 'a', 'g', 'q', 't', 'l', 'j', 'p', 'x', 'u', 'p', 't', 'n', 'e', 'r', 'y', 'p', 'j', 'y', 'a', 'e', 'q', 'r', 'z', 'x', 'f', 'l', 'r', 't', 'u', 'e', 'c', 'd', 'l', 'c', 'g', 'a', 'n', 't', 'w', 'd', 'n', 'r', 'k', 'm', 'i', 'd', 'y', 'o', 'c', 'z', 'y', 'i', 'z', 'i', 'l', 'd', 'y', 'p', 'x', 'w', 'g', 'j', 'r', 'y', 'q', 'h', 's', 'n', 'w', 't', 't', 'n', 'h', 'q'}
	costs := []int{629065, 837528, 28529, 72965, 650814, 212281, 287583, 608114, 291306, 183186, 734563, 403541, 316695, 82179, 810783, 11684, 446775, 316126, 975696, 86358, 344552, 42370, 11758, 299573, 86610, 736799, 584454, 176212, 25040, 679691, 393472, 300055, 401833, 173256, 753799, 50605, 936075, 250413, 340281, 256615, 60321, 253297, 605289, 576661, 245253, 186626, 249213, 242831, 773902, 931797, 155216, 614351, 524811, 985381, 542048, 62535, 158912, 436868, 934416, 844287, 740498, 554204, 837753, 455931, 671712, 941329, 195621, 532917, 809506, 33385, 64203, 951223, 86295, 519540, 78242, 102907, 58209, 93307, 262994, 103937, 584265, 578059, 888526, 320232, 24057, 674175, 701243, 57386, 508695, 333835, 459953, 428457, 226047, 559658, 216323, 798162, 25252, 708080, 119403, 872221, 928133, 211156, 647255, 803872, 267010, 774828, 724217, 726002, 775742, 765983, 583685, 720772, 273830, 242887, 255702, 108877, 948252, 401809, 159519, 535491, 789561, 114305, 275432, 622567, 864025, 152946, 345180, 76613, 437950, 703224, 794441, 952785, 69720, 814150, 517526, 775377, 514593, 621439, 585335, 464088, 85553, 952427, 826263, 133692, 96242, 167328, 778139, 459354, 708687, 996080, 8228, 936458, 117512, 462849, 104798, 150450, 953421, 206327, 687457, 30852, 782960, 377599, 673360, 260847, 418148, 555601, 513892, 510514, 276632, 916410, 600709, 63321, 3569, 605924, 177080, 515938, 110514, 567324, 462880, 782306, 793660, 986928, 680062, 351091, 914517, 332905, 102209, 135049, 191902, 899324, 529747, 771840, 293508, 434929, 145271, 943594, 151743, 813040, 51116}

	minimumCost(source, target, origin, changed, costs)
}

func Test97(t *testing.T) {
	isInterleave("aabcc", "dbbca", "aadbbcbcac")
}
