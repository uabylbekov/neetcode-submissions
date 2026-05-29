type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, str := range strs{
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)

	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	decoded := make([]string, 0)
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		end := start + length

		decoded = append(decoded, encoded[start:end])

		i = end
	}

	return decoded
}
