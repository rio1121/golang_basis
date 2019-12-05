package testlib

// ConcatStringSlice スライスsの文字列を全て連結する
func ConcatStringSlice(s []string, ch chan string) {
	result := ""
	for _, element := range s {
		result += element
		ch <- result
	}
	close(ch)
}
