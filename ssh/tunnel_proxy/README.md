# 透SSH Tunnel 來遠端上網

情境：公司有一台樹莓派使用wlan0連在公司網段 192.168.2.12/24，樹莓派本身接著有線網路到電腦上，該電腦上樹莓派的網段上 192.168.3.13/24，而且有運作中的http host，我的筆電在公司網段 192.168.2.55，由於無法直接存取192.168.3.13，但是可以ssh到樹莓派上，想藉由ssh 建立 tunnel，在讓瀏覽器 proxy 使用此 tunnel 間接造訪 192.168.3.13。

## 1. 在樹莓派上建立ssh tunnel

```bash
ssh -D 8080 -N pi@192.168.2.12
```

## 2. 設定瀏覽器 proxy

### Firefox

到Network Settings開啟設定，選擇手動設定 proxy，設定 SOCKS Host 為localhost，Port 為 8080，並勾選 SOCKS v5。HTTP、HTTPS proxy 設定為空白。

### Chrome

```bash
google-chrome --proxy-server="socks5://localhost:8080"
```

## 3. 瀏覽器造訪

在瀏覽器造訪 http://192.168.3.13 即可看到網頁。
