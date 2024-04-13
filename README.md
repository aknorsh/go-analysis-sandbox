# Go Analysis Sandbox

静的解析のサンドボックス。

## サードパーティの linter を差し込みたいとき

TODO。analysis.Analyzer 提供してたらそれを main.go に差す？

## 独自 linter を追加したいとき

- internal/{hoge}
  - hoge.go <- Analyzer を定義。テンプレ化できる
  - hoge_test.go <- テスト書く。テンプレ固定。
  - testdata/src/a/a.go <- テストで使うやつを書く。実質テスト定義箇所はこっち。テンプレ化できなくもない
- cmd/mylint/main.go に hoge.Analyzer 追加
