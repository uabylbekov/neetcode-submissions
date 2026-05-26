func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    countS := [26]int{}
    countT := [26]int{}

    for _, char := range s {
        countS[char - 'a']++
    }

    for _, char := range t {
        countT[char - 'a']++
    }

    return countS == countT
}
