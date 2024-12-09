# slack-reminder

`slack-reminder`は、Slackで使用できる`/remind`コマンドを簡単かつ対話的に生成することができるCLIツールです。  

## インストール方法

### インストール

以下のコマンドを実行することで、`$GOPATH/bin`に`slack-reminder`バイナリがインストールされます。

```bash
go install github.com/staticWagomU/go-slackreminder
```

インストール後、`slack-reminder`コマンドをPATHが通ったディレクトリから実行できるようにしてください。

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## 使い方の例

対話モードを起動するには、以下のように単純にコマンドを実行します。

```bash
slack-reminder
```

プロンプトが表示されたら、リマインド対象（ユーザーやチャンネル）、メッセージ、時刻・日付などを順に入力してください。最後に生成された`/remind`コマンドが標準出力されます。

## ライセンス

本プロジェクトは MIT License のもとで公開されています。詳細は[LICENSE](./LICENSE)ファイルを参照してください。
