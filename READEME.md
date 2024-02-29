メモ

- モジュール
https://pkg.go.dev/std


- intは基本的にint64を使う
- メモリを気にする場合は、キャパシティを使う
- try～catch～finallyの例外処理はない
- ポインタを使う場合は、newを使う、使わない場合は、makeを使う
- コンポジ型とは、構造体の中に構造体を入れること
- 小文字の場合は、パケージ内でしか参照できない
- interfaceは、メソッドの集まりで、必ず実装する必要があるメソッドがある場合に使う

- tyeps
https://go.dev/ref/spec



コマンド
- go run [fileName]      ファイルを実行
- gofmt -w [fileName]   フォーマットを整える
- go doc [packageName]   ドキュメントを見る