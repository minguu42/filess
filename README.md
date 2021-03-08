# filess

## 概要

`filess`は、簡単にファイルを整理するためのCLIツールです。

`filess`は、ソースディレクトリから移動させるファイルを検索し、条件を満たすファイルをターゲットディレクトリに移動させます。
ソースディレクトリとターゲットディレクトリは、基本的に`$HOME/.config/filess/config.json`に記述します。

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

### 設定ファイルの初期化

```bash
$ filess init
2021/03/08 09:38:34 Initialize /Users/hoge/.config/filess/config.json
```

`filess init`コマンドは、初期化された設定ファイル`$HOME/.config/filess/config.json`が作成します。既に設定ファイルが存在する場合でも、上書きして作成されます。

初期化された設定ファイルの内容は以下のようになります。

```json:config.json
{
  "targets": [
    ""
  ],
  "sources": [
    ""
  ],
  "inspections": [
    ""
  ]
}
```

### ファイルを移動、削除する

```bash
$ filess
...
```

`filess`の最も基本的なコマンドになります。
`filess`コマンドは、ファイルの移動、削除の２つを行います。
ソースディレクトリに、ファイル名が`<ターゲットディレクトリ名>_`であるファイルがあるとき、ターゲットディレクトリに移動させます。
また、インスペクションディレクトリが指定されている場合は、そのディレクトリ内のファイルを削除するか確認し、削除します。

#### 動作例

```json: config.json
{
  "targets": [
    "/Users/hoge/Documents/hoge"
  ],
  "sources": [
    "/Users/hoge/Downloads"
  ],
  "inspections": [
    "/Users/hoge/Downloads"
  ]
}
```

```bash
$ filess
Move to /Users/hoge/Downloads/hoge_example.txt from /Users/hoge/Documents/hoge/hoge_example.txt
.DS_Store を削除しますか？（Y/n）: Y
Delete /Users/akira/Downloads/.DS_Store
.localized を削除しますか？（Y/n）: n
```

### 設定を追加する

```bash
$ filess [-t <filepath> || -s <filepath> || -i <filepath>] 

$ filess -t /Users/hoge/Documents/example
2021/03/08 10:22:09 Add /Users/hoge/Documents/example to sources
```

`filess`に`-t`、`-s`、`-i`オプションで、それぞれターゲットディレクトリ、ソースディレクトリ、インスペクションディレクトリのパスを設定ファイルに追加します。

### 設定の確認

```bash
$ filess -c
[targets]
/Users/hoge/Documents/hoge
/Users/hoge/Documents/example

[sources]
/Users/hoge/Downloads

[inspections]
/Users/hoge/Downloads
```

`filess -c`コマンドで、設定を確認できます。

### バージョンの確認

```bash
$ filess -v
version: 0.1.0
revision:
```

`filess -c`コマンドで、バージョンを確認できます。
`revision`はビルド時にコミットIDを埋め込んでいるので、`go install`でインストールした場合はなく、GitHubでバイナリをダウンロードした場合にのみ存在します。
