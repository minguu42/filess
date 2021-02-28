# filess

## 概要

filessは、簡単にファイルを整理するためのCLIツールです。
ファイル名に`<ディレクトリ名>_`というプリフィックスをつけて、config.jsonファイルにソースとターゲットを指定しておくと`filess`というコマンドを打つだけでファイルを適切な場所に移動させることができます。

## インストール

### Goからインストールする

`go install github.com/minguu42/filess/cmd/filess`でインストールできます。

### バイナリを$PATHに読み込ませる

GitHubからバイナリをダウンロードして、PATHを読み込む場所に配置してください。

## 使い方

### filess init

`filess init`コマンドで`$HOME/.config/filess/config.json`が作成されます。
以下のようになっています。

```json:config.json
{
  "targets": [
    ""
  ],
  "sources": [
    ""
  ]
}
```

すでにこのファイルが存在している場合でも、上書きで作成されます。ご注意ください。

### filess

`filess`コマンドで実行できます。
第一引数に