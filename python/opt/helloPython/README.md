# Hello Python!

- Cloud: Azure
- Azure Virtual Machine

---

## 1.VMの作成
1. Azureホーム画面左上のメニュを開き[VirtualMachines]を選択する。
2. 




## Nginxインストール
基本GCPと同じ。
以下のみ異なる

```s
sudo vi /etc/apt/sources.list

# 以下2行を末尾に追加
deb http://nginx.org/packages/ubuntu/ bionic nginx
deb-src http://nginx.org/packages/ubuntu/ bionic nginx
```

## Pythonインストール
Ubuntuならデフォルトで入っているが、一応

1. インストール
```
sudo apt update
sudo apt install -y python3
```

2. バージョン確認
```
which python
which python3
```

3. pipインストール
```
sudo apt install python-pip
```

4. SampleWebAppのインストール
```
git clone https://github.com/makocchan0509/simpleWebApps.git
```

5. 依存パッケージのインストール
```
pip install Flask
```

6.WebAppの起動
```
python simpleWebApps/python/opt/helloPython/webapp.py
```

---

### Nginxのリバースプロキシ設定

1. 設定ファイル作成
```
sudo vi /etc/nginx/conf.d/server.conf

#以下をペーストする
server{
    server_name    localhost;

    proxy_set_header    Host    $host;
    proxy_set_header    X-Real-IP    $remote_addr;
    proxy_set_header    X-Forwarded-Host       $host;
    proxy_set_header    X-Forwarded-Server    $host;
    proxy_set_header    X-Forwarded-For    $proxy_add_x_forwarded_for;

    location / {
        proxy_pass    http://localhost:8085/;
    }
}
```