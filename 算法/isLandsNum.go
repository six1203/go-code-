package main

func dfs(row, col, rows, cols int, grid [][]byte, visited [][]bool) {
	if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] == '0' || visited[row][cols] == true {
		return
	}
	visited[row][col] = true
	dfs(row-1, col, rows, cols, grid, visited)
	dfs(row+1, col, rows, cols, grid, visited)
	dfs(row, col-1, rows, cols, grid, visited)
	dfs(row, col+1, rows, cols, grid, visited)

}

func numIslands(grid [][]byte) int {
	var rows, cols int
	rows = len(grid)
	cols = len(grid[0])

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		res := make([]bool, cols)
		for j := 0; j < cols; j++ {
			res[j] = false
		}
		visited[i] = res
	}

	var islandNum int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && visited[i][j] == false {
				islandNum++
				dfs(i, j, rows, cols, grid, visited)
			}
		}
	}
	return islandNum
}
