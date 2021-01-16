# ReviewPolice👮‍♀️

### 概要
AdventCalendar 2020で作ったやつです。
https://github.com/yuta-ron/reviewpolice

### 使い方
1. `GOOGLE_APPLICATION_CREDENTIALS` のjsonを発行します。
2. GitHubのSecrets に下記を設定します。
- キー名: `GOOGLE_APPLICATION_CREDENTIALS_BODY`
- 値: 1. で生成したJSONをbase64エンコードしたもの。
