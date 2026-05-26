import (
    "slices"
    "maps"
)
func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[[26]int][]string)

    for _, str := range strs {
        countStr := [26]int{}

        for _, char := range str {
            countStr[char - 'a']++
        }

        anagrams[countStr] = append(anagrams[countStr], str)
    }

    return slices.Collect(maps.Values(anagrams))
}
