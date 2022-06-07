# Go API DEMO
api server made in golang

# Try
下記いずれかのパターンを行えば、ブラウザで `localhost:8080/healthcheck` で動作確認できます。

## pattern 1
* go build main.go
* go run ./main

## pattern 2
* VS Codeで本プロジェクトを開く
* 左メニューにある[実行とデバッグ]を押す
* Launch すると `.vscode/launch.json` の内容に則って実行されます。

## pattern 3
* ルートで `make serve` コマンドを実行