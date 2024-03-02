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


goroutine
- fan-out fan-in
  - fan-out: 複数のgoroutineを起動すること
  - fan-in: 複数のgoroutineからの結果を受け取ること
  - channelでパイプラインを持たせたいときに使う

コマンド
- go run [fileName]      ファイルを実行
- gofmt -w [fileName]   フォーマットを整える
- go doc [packageName]   ドキュメントを見る

batchリポジトリの構成
```
.
├── app                    # アプリケーションのコアロジックを含むディレクトリ
│   └── mail               # メール処理タスク専用のディレクトリ
│       ├── logic          # メール処理のビジネスロジックを格納
│       └── executor.go    # メール処理ロジックの実行を担当するファイル
├── config                 # アプリケーション設定ファイルを保存するディレクトリ。データベース接続やAPIキーなどを含む
│   └── XXXX_config.go
├── library                # 共通のコード、ライブラリ、ユーティリティを含むディレクトリ。ヘルパー関数やミドルウェアなどを含む
├── go.mod                 # プロジェクトの依存関係を定義するGoモジュールファイル
└── go.sum                 # 依存関係の整合性を保証するためのチェックサムを提供するファイル
```

[参考になりそうなディレクトリ構成](https://github.com/golang-standards/project-layout)