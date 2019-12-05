package nextcode

// NextCode 文字列の各文字を次の文字コードの文字に置換した結果を返す.
func NextCode(s string) string {
	sRuned := []rune(s)
	for i := 0; i < len(sRuned); i++ {
		sRuned[i] = sRuned[i] + 1
	}
	return string(sRuned)
}
