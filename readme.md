# zipをWebからみる


漫画ビューアの部分は下記の物をベースに改造
https://github.com/tomalatte001/comic-viewer

|key入力|説明|
|:--|:--|
|f|上下の枠を消す|
|↑|一ページ進む|
|←|二ページ進む|
|↓|一ページ戻る|
|→|二ページ戻る|

# 使い方

http://localhost:8080/view/1 にアクセスすると表示されます

# APIとか

$$ view
./id/page

|記号|説明|
|:--|:--|
|id|zipフォルダの指定番号|
|page|指定現在ページ(未実装)|

## ziplist
./ziplist
指定フォルダ内のファイルリストを表示

## ファイルデータ取得
対象ファイルからデータを取得します。
./zip?id=id&page=page
オプション

|記号|説明|
|:--|:--|
|id|zipフォルダの指定番号|
|page|指定現在データ番号|

