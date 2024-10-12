// テストをする際に外部ファイルを使うことがある。このようなときにはtestdataディレクトリを作成し、その中にテストデータを配置することが多い。
// しかし、testdataディレクトリは無視することが多く、.や_で始まるファイルやディレクトリも無視するようになっている。
// 実際にtestdataディレクトリでテストする際はfilepath.Globを使うと便利
// これはfilepath.Globの使用例
package testdata

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp" // go-cmpを使用すると比較が行えるだけでなく、比較した結果がdiff形式で得られる
)

func doSomething(s string) string {
	return s
}

func TestDoSomething(t *testing.T) {
	fns, err := filepath.Glob("testdata/*.dat")
	if err != nil {
		t.Fatal(err)
	}

	for _, fn := range fns {
		t.Log(fn)
		b, err := os.ReadFile(fn)
		if err != nil {
			t.Fatal(err)
		}

		got := doSomething(string(b))

		//.datを.outに入れ替えて結果データを読み込む
		b, err = os.ReadFile(fn[:len(fn)-4] + "out")
		if err != nil {
			t.Fatal(err)
		}
		want := string(b)

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Diff: %s", diff)
		}
	}
}
