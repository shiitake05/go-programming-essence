// テスト対象のコード
package short

func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
