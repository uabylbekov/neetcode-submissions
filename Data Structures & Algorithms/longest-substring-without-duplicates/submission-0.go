func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]struct{})
	l := 0
	longest := 0

	for r := 0; r < len(s); r++ {
		for {
			if _, ok := seen[s[r]]; !ok {
				break
			}
			delete(seen, s[l])
			l++
		}
		seen[s[r]] = struct{}{}
		longest = max(longest, len(seen))
	}

    return longest
}
