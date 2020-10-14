# Hello Golang!

- クラウド: GCP
- VM: Compute Engine (ubuntu-1604-xenial-v20200923)
- ファイアウォールではssh(22),http(80,8080)を許可する。
---

## VMの作成
GCPでVMインスタンスを作成します。
まずはGCPのコンソールへログインしてダッシュボードを表示します。

1. 画面左上の[ナビゲーションメニュ]をクリック
2. コンピューティングカテゴリの[ComputeEngine]→[VMインスタンス]をクリック
3. VMインスタンスメニュにて[インスタンスを作成]をクリック
4. インスタンスの作成画面にて以下の項目を入力(記載がないものはデフォルトのままでOK)
    - 名前：インスタンスの名称
    - リージョン：asia-northeast1(東日本)
    - マシンファミリーのマシンタイプ：e2-micro
    - ブートディスクの[変更]をクリック→オペレーティングシステム：Ubuntu →[選択]をクリック
    - ファイアウォールにて[HTTPトラフィックを許可する]をチェック
    - 最下部の[作成]をクリック
5. VMインスタンスメニュにて作成したVMインスタンスが表示される。（少し時間かかる）  
**本メニュ上の一覧に表示される[SSH]からログインして「rootユーザのパスワード」手順から進んでください**  
**自分の端末からリモートしたい場合は下の手順へ進んでください。**

---

## sshkey作成と登録
本手順はVMへのログインに必要なsshkeyを生成してメタデータへ登録することでプロジェクト共通で他のVMにもログインできるようになる手順です。

1. 画面右上の[>-]cloud shellの起動をクリックする。
2. 以下のコマンドを実行 []内は任意の名前をつけること。USERNAMEはVMログイン時のユーザになる。
  ```
  ssh-keygen -t rsa -f ~/.ssh/[KEY_FILENAME] -C [USERNAME]
  #パスワードを設定できるので任意のパスワードを入力する
  ```
3. 以下のコマンドで結果を確認する。
  ```
  ls ~/.ssh/
  #[KEY_FILENAME]と[KEY_FILENAME].pub というファイルが存在すること。
  chmod 400 ~/.ssh/[KEY_FILENAME]
  ＃書き込み禁止しておく
  ```
4. コンソール画面に戻り、画面左上の[ナビゲーションメニュ]→[ComputeEngine]→[VMインスタンス]をクリック
5. 画面左のメニュから[メタデータ]をクリック
6. [SSH認証鍵]タブをクリック
7. [編集]をクリック
8. [項目を追加]をクリック
9. 空欄に以下のコマンドを実行して表示される鍵情報をコピーしてペーストする
  ```
  cat ~/.ssh/test-key.pub
  #ssh-rsaから[USERNAME]までの情報をコピー
  ```
10. [保存]をクリック  

以降は個人端末からsshするために鍵をDLする手順です。

11. cloudshellコンソール上の右上[...]みたいなのをクリック
12. [ファイルをダウンロード]をクリック
13. ファイルパスへ[~/.ssh/[KEY＿FILENAME]]を入力して[ダウンロード]をクリック

詳しくはこの辺り参照
[GCPdoc-sshkey作成](https://cloud.google.com/compute/docs/instances/adding-removing-ssh-keys?hl=ja)

---
## VMへのssh
1. コンソール画面左上の[ナビゲーションメニュ]→[ComputeEngine]→[VMインスタンス]をクリック
2. 作成済みのインスタンスの[publicIP]を確認。

以下のコマンドでVMへリモート接続します。
#Windowsの方はTeratermにてpublicIPとダウンロードした鍵ファイルでログインしてください。
```
ssh -i {keyfile} {user}@{publicIP}
```

## rootユーザのパスワード

以下のコマンドでrootユーザのパスワードを設定する。
```
su passwd
```

以下のコマンドでrootユーザへスイッチ
```
su -
```

---

## 1.Install Nginx

以下のコマンドでkeyファイルをDL
```
wget https://nginx.org/keys/nginx_signing.key
 ***‘nginx_signing.key’ saved [1561/1561]
```

以下のコマンドでkeyファイルをリポジトリへ追加
```
apt-key add nginx_signing.key
OK
```

以下のコマンドでファイルを編集して保存する。
```
vi /etc/apt/sources.list

# 以下2行を末尾に追加
deb http://nginx.org/packages/ubuntu/ xenial nginx
deb-src http://nginx.org/packages/ubuntu/ xenial nginx
# esc + : + w + q + ! + enterでviを終了
```

リポジトリのUpdateを行う。
```
apt update
```

nginxをインストールする。
```
apt install nginx
```

nginxを起動する。
```
systemctl start nginx
```

起動確認
```
systemctl status nginx
```

画面の確認
```
#ブラウザで以下のURLへアクセス
http://{publicIP}:80
#or curlコマンド
curl {publicIP}:80
```

---

## 2.Install Golang

```
cd /usr/local/src
```

Golangのサイトより資材をDLする。
```
wget https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
```

以下のコマンドで圧縮ファイルを解凍する。
```
tar xf go1.15.2.linux-amd64.tar.gz -C /usr/local/
```

goバージョンの確認
```
/usr/local/go/bin/go version
```

GOROOT変数をprofileへ追記する。
```
echo 'export GOROOT=/usr/local/go' >> ~/.profile
```

PATHへGOROOTの追加
```
echo 'export PATH=$PATH:$GOROOT/bin' >> ~/.profile
```

profileの読み込み
```
source ~/.profile
```

envコマンドで環境変数の確認
```
env
#GOROOT、PATHに↑が反映されていることを確認する。
```

サンプルAPのDL
```
git clone https://github.com/makocchan0509/simpleWebApps.git
```

カレントディレクトリを変更する
```
cd ./simpleWebApps/golang/helloGolang/
```

Golangアプリをビルドする。
```
go build
```

結果確認
```
ls
helloGolang  main.go
```

Apの実行
```
./helloGolang
#停止はCTRL+C
```

結果確認
```
http://{publicIP:8080}

# 以下のJSONメッセージが返ってきます。
{"result":"Success","errMessage":"Hello Golang World!"}
```

ここまででVMのファイアウォールで8080ポートを許可していればAPから応答があります。
nginxと連携させる場合は以下手順に進みます。

---

## 3. Nginxのリバースプロキシ設定

以下のコマンドでserver.confを追加します。
```
vi /etc/nginx/conf.d/server.conf

#以下のテキストを貼り付けて保存して閉じます。

server{
    server_name    localhost;

    proxy_set_header    Host    $host;
    proxy_set_header    X-Real-IP    $remote_addr;
    proxy_set_header    X-Forwarded-Host       $host;
    proxy_set_header    X-Forwarded-Server    $host;
    proxy_set_header    X-Forwarded-For    $proxy_add_x_forwarded_for;

    location / {
        proxy_pass    http://localhost:8080/;
    }
}
# esc + : + w + q + ! + enterでviを終了

```

デフォルトで置かれているファイルが悪さするので除外します。
```
mv /etc/nginx/conf.d/default.conf /tmp/
#削除しても良いです。
```

Nginxを再起動
```
systemctl restart nginx
```

接続確認
```
# golangのAPが実行されていること。
./helloGolang

#ブラウザで以下URLへアクセス
#APのレスポンスが応答されるはず

http://{publicIP}:80

# 以下のJSONメッセージが返ってきます。
{"result":"Success","errMessage":"Hello Golang World!"}

```

---

## 4. クリーニング

1. コンソール画面左上の[ナビゲーションメニュ]→[ComputeEngine]→[VMインスタンス]をクリック
2. 作成したVMインスタンスを選択して画面中央上辺りのゴミ箱アイコンをクリック

以上
