package pkg

func GenerateNeighbours(n int) [][]int {
	size := n - 1
	size = -size

	var res [][]int
	for i := size; i < n; i++ {
		for j := size; j < n; j++ {
			res = append(res, []int{i, j})
		}
	}
	return res
}
