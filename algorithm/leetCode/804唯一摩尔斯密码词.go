package leetCode

func UniqueMorseRepresentations(words []string) int {
	temple := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]interface{}, len(words))
	for _, w := range words {
		s := ""
		for v := range w {
			s += temple[w[v]-97]
		}
		m[s] = struct{}{}
	}
	return len(m)
}
