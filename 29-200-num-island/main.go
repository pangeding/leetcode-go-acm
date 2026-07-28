package main


func numIslands(grid [][]byte) int {
	// if len(grid) == 0 return 0
	if len(grid) == 0 {
		return 0 
	}
	// rows cols
	rows, cols := len(grid), len(grid[0])

	// define the count 
	count := 0 

	// func dfs
	// the dfs function don't count the num of the islands.
	// it just clear the island
	var dfs func(r, c int)
	dfs = func(r, c int) {
		// corner case 

		// solve all the status
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
			return 
		}
		grid[r][c] = '0'
		dfs(r - 1, c)
		dfs(r + 1, c)
		dfs(r , c - 1)
		dfs(r , c + 1)
	}


	// 

	// 遍历 
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r,c)
			}
		}
	}

	return count 
}