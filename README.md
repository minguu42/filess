# filess

## 概要

`filess`は、簡単にファイルを整理するためのCLIツールです。

`filess`は、ソースディレクトリから移動させるファイルを検索し、条件を満たすファイルをターゲットディレクトリに移動させます。
ソースディレクトリとターゲットディレクトリは、基本的に`$HOME/.config/filess/config.json`に記述します。

ファイル名に`<ディレクトリ名>_`というプリフィックスをつけて、ファイルのあるべきディレクトリを指定します。config.jsonファイルにソースとターゲットを指定しておくと`filess`というコマンドを打つだけでファイルを適切な場所に移動させることができます。

## インストール

### Goからインストールする

```bash
$ go install github.com/minguu42/filess/cmd/filess
$ filess -v
version: 0.1.0
revision:
```

でインストールできます。

### バイナリを$PATHに読み込ませる

GitHubからバイナリをダウンロードして、`$PATH`に配置してください。

## コマンド一覧

### filess init

`filess init`コマンドで設定ファイル`$HOME/.config/filess/config.json`が作成されます。
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
