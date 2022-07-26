package leetcode

func seeAreaOfOne(grid [][]int, indexRow int, indexCol int) int {

	if indexRow < 0 ||
		indexRow >= len(grid) ||
		indexCol < 0 ||
		indexCol >= len(grid[0]) ||
		grid[indexRow][indexCol] == 0 {

		return 0
	}

	grid[indexRow][indexCol] = 0

	return (1 + seeAreaOfOne(grid, indexRow-1, indexCol) +
		seeAreaOfOne(grid, indexRow, indexCol+1) +
		seeAreaOfOne(grid, indexRow, indexCol-1) +
		seeAreaOfOne(grid, indexRow+1, indexCol))
}

func MaxAreaOfIsland(grid [][]int) int {

	var bigArea int = 0

	for indexRow, row := range grid {
		for indexCol, col := range row {
			if col == 1 {
				if sizeArea := seeAreaOfOne(grid, indexRow, indexCol); bigArea < sizeArea {
					bigArea = sizeArea
				}
			}
		}
	}

	return bigArea
}
