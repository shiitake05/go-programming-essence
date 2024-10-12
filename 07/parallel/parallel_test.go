package parallel

import "testing"

// このテストは9秒かかる
// func TestAdd(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		lhs  int
// 		rhs  int
// 		want int
// 	}{
// 		{name: "test1", lhs: 0, rhs: 1, want: 1},
// 		{name: "test2", lhs: 1, rhs: -1, want: 0},
// 		{name: "test3", lhs: 2, rhs: 1, want: 3},
// 	}

// 	for _, test := range tests {
// 		got := Add(test.lhs, test.rhs)
// 		if got != test.want {
// 			t.Errorf("%v:eant %v, but %v:", test.name, test.want, got)
// 		}
// 	}
// }

// 上のようにテストを追加するたびに遅くなると困るので、テストを並列実行する
// このテストは3秒で終わる
// ※注意点: テストにユニークな名前をつけること。
// このテストは並列実行されるため、テスト名が同じだと結果が混ざってしまう
// ※注意点: Table Driven Testsを並行で実行する場合、ループの中でtest変数を束縛する必要がある
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		lhs  int
		rhs  int
		want int
	}{
		{name: "test1", lhs: 0, rhs: 1, want: 1},
		{name: "test2", lhs: 1, rhs: -1, want: 0},
		{name: "test3", lhs: 2, rhs: 1, want: 3},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Add(test.lhs, test.rhs)
			if got != test.want {
				t.Errorf("%v:eant %v, but %v:", test.name, test.want, got)
			}
		})
	}
}
