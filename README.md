## 概要

日本郵便の提供する郵便番号データ(通称KEN_ALL.CSV)をパースし、使いやすい形に変換します。
主に以下のような処理を行います。

- 複数行に分割された行を結合します
- SHIFT_JIS を UTF-8に変換します
- 半角カナを全角カナに変換します
- 全角数字、記号を半角に変換します
- データとして使いにくいレコードを加工します(加工処理詳細参照)
- CSV形式をJSON形式、TSV形式に変換できます

## コマンドの使い方

住所データをjson形式に変換

```
$ ken-all convert -i address -o json -f /tmp/KEN_ALL.CSV

{"region_id":"01101","zip":"0600000","pref_kana":"ホッカイドウ","city_kana":"サッポロシチュウオウク","town_kana":"","pref":"北海道","city":"札幌市中央区","town":"","update_status":"0","update_reason":"0","pref_code":"01"}
 :
```

事業所データをjson形式に変換

```
$ ken-all convert -i office  -o json -f /tmp/JIGYOSYO.CSV

{"jis_code":"01101","kana":"(カブ) ニホンケイザイシンブンシヤ サツポロシシヤ","name":"株式会社 日本経済新聞社 札幌支社","pref":"北海道","city":"札幌市中央区","town":"北一条西","address":"6丁目1-2アーバンネット札幌ビル2F","zip7":"0608621","zip5":"060  ","post_office":"札幌中央","type":"0","is_multi":"0","update_status":"0","pref_code":"01"}
```

詳しくは helpを参照ください

```
ken-all convert -h
Convert csv file into other format.

Usage:
  ken-all convert [flags]

Flags:
  -f, --file string            input file
  -h, --help                   help for convert
  -i, --input_type string      input type (default "address")
  -o, --output_format string   output format (default "csv")
```


## 加工処理詳細

日本郵便の提供するCSVの特徴として、町域名が38文字を超える場合複数行に分かれています。
まずはこれを1行にした上で下記のような処理を行います。

### 町域名に特定の文字列が入る場合は空文字に変換します

**例**

- 以下に掲載がない場合
- ○○一円
- ○○の次に番地がくる場合

**加工前**

| 郵便番号 | 都道府県名 | 市区町村名    | 町域名                   |
|----------|------------|---------------|--------------------------|
| 0600000  | 北海道     | 札幌市中央区  | 以下に掲載がない場合     |
| 1000301  | 東京都     | 利島村        | 利島村一円               |
| 3060433  | 茨城県     | 猿島郡境町    | 境町の次に番地がくる場合 |

**加工後**

| 郵便番号 | 都道府県名 | 市区町村名    | 町域名                   |
|----------|------------|---------------|--------------------------|
| 0600000  | 北海道     | 札幌市中央区  |                          |
| 1000301  | 東京都     | 利島村        |                          |
| 3060433  | 茨城県     | 猿島郡境町    |                          |


### 町域名が X（A、B）の形式の場合、X、XA、XB の3行に展開します

ただし、括弧内の文字が地名以外と思われるものは削除しています

**例**

- 「その他」、「地階・階層不明」、「全域」、「成田国際空港内」など特定のワードを含むもの
- 「132〜156」、「367番地」 など番地と思われるもの


**加工前**

| 郵便番号 | 都道府県名 | 市区町村名    | 町域名                                            |
|----------|------------|---------------|---------------------------------------------------|
| 0893443  | 北海道     | 中川郡本別町  | 西美里別（１１３〜７９１番地、西活込、西上、西中）|


**加工後**

| 郵便番号 | 都道府県名 | 市区町村名    | 町域名         |
|----------|------------|---------------|----------------|
| 0893443  | 北海道     | 中川郡本別町  | 西美里別       |
| 0893443  | 北海道     | 中川郡本別町  | 西美里別西活込 |
| 0893443  | 北海道     | 中川郡本別町  | 西美里別西上   |
| 0893443  | 北海道     | 中川郡本別町  | 西美里別西中   |


### 岩手県の地割表記は削除します

例: 「越中畑64地割〜越中畑66地割」は「越中畑」に変換します

**加工前**

| 郵便番号 | 都道府県名 | 市区町村名     | 町域名                         |
|----------|------------|----------------|--------------------------------|
| 0295523  | 岩手県     | 和賀郡西和賀町 | 越中畑６４地割〜越中畑６６地割 |

**加工後**

| 郵便番号 | 都道府県名 | 市区町村名     | 町域名 |
|----------|------------|----------------|--------|
| 0295523  | 岩手県     | 和賀郡西和賀町 | 越中畑 |


## 参考

* [郵便番号データの説明](http://www.post.japanpost.jp/zipcode/dl/readme.html)

