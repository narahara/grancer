# Grancer
## glancer.jp

## 方針
* 全ての時間軸の情報を一気に見れること(at a grance)
    * 短い時間軸は上位の長い時間軸に影響される
    * 過去のサポート、レジスタンス、テクニカル指標上のライン
* 機械的に意志決定ができること　→ 意思決定を阻害するような情報は表示しない
* 常に人間が見ていなくても取引してくれる（自動売買とSlackを使った取引アシスト）
    * 目標値になったらトレーリングストップする

* テクニカル指標が違和感なく動いた場合の価格を算定する


## anaconda
* https://www.anaconda.com/
* https://repo.anaconda.com/archive/Anaconda3-2022.10-MacOSX-arm64.pkg
* glancer環境はanaconda navigatorから作成


```
conda activate grancer
```

## mySQL

```
CREATE DATABASE glancer CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE USER 'app'@'%' identified by 'password';
GRANT ALL PRIVILEGES on glancer.* to 'app'@'%' ;
```


## OANDA
a768709c9055b58ec5ac9055ca406d35-d6440968f9154452ff16dda859788d48

http://developer.oanda.com/rest-live-v20/introduction/


