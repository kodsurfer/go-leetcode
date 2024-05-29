package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{[]byte{'1', '1', '1', '1', '0'}, []byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '0', '0', '0'}}))
	fmt.Println(numIslands([][]byte{[]byte{'1', '1', '0', '0', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '1', '0', '0'}, []byte{'0', '0', '0', '1', '1'}}))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				cnt++
				markIsland(grid, i, j)
			}
		}
	}
	return cnt
}

func markIsland(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	markIsland(grid, i-1, j)
	markIsland(grid, i+1, j)
	markIsland(grid, i, j-1)
	markIsland(grid, i, j+1)
}
