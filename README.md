# ReviewPolice👮‍♀️

### 概要
AdventCalendar 2020で作ったやつです。
https://qiita.com/yuta-ron/items/9288e2e31ffd91a17e18

### 使い方
1. `GOOGLE_APPLICATION_CREDENTIALS` のjsonを発行します。
2. GitHubのSecrets に下記を設定します。
- キー名: `GOOGLE_APPLICATION_CREDENTIALS_BODY`
- 値: 1. で生成したJSONをbase64エンコードしたもの。
