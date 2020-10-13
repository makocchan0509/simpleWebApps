# Hello Golang!

- クラウド: GCP
- VM: Compute Engine (ubuntu-1604-xenial-v20200923)
- VMにはPublicIPが付与されている前提
- ファイアウォールではssh(22),http(80,8080)を許可しておくこと。

## 0.rootユーザのパスワード

以下のコマンドでrootユーザのパスワードを設定する。
```
su passwd
```

以下のコマンドでrootユーザへスイッチ
```
su -
```

---

## 1.Install Nginx(任意)

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
```

ここまででVMのファイアウォールで8080ポートを許可していればAPから応答があります。
nginxと連携させる場合は以下手順に進みます。

---

3. Nginxのリバースプロキシ設定

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
```