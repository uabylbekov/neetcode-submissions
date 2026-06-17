func minimumRecolors(blocks string, k int) int {
	countOfW := 0

	for i := 0; i < k; i++ {
        if blocks[i] == 'W' {
            countOfW++
        }
    }

	minimum := countOfW

	for i := k; i < len(blocks); i++ {
        if blocks[i] == 'W' {
            countOfW++
        }
        if blocks[i-k] == 'W' {
            countOfW--
        }
        
        if countOfW < minimum {
            minimum = countOfW
        }
    }

    return minimum
}
