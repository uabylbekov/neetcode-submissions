func numIslands(grid [][]byte) int {
    res := 0
    rows := len(grid)
    cols := len(grid[0])

    var dfs func(r, c int)

    dfs = func(r, c int) {
        if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == '1' {
            grid[r][c] = '0'
            dfs(r + 1, c)
            dfs(r - 1, c)
            dfs(r, c + 1)
            dfs(r, c - 1)
        }
    }

    for r := range rows {
        for c := range cols {
            if grid[r][c] == '1' {
                res += 1
                dfs(r, c)
            }
        }
    }
    return res
}
