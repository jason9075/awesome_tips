## netcat

使用 netcat（通常被簡稱為 nc）來監聽端口時，它會在處理完一次連接後自動結束。如果您想要讓 netcat 持續監聽，而不是在第一個連接被處理後就停止，您需要使用一個無限循環（如在shell腳本中），或者使用支持持續監聽模式的 netcat 變體。

使用 netcat 的 -k 選項:

某些 netcat 版本支持 -k 選項，這個選項使 netcat 在客戶端斷開連接後繼續監聽。您可以嘗試下面的命令：

```bash
sudo nc -lk 8080
```

不過，請注意並非所有版本的 netcat 都支持 -k 選項。

## socat

以下是一個基本的範例，展示了如何使用 socat 監聽8080端口，並對進來的請求回應一個簡單的HTTP頁面：

```bash
socat TCP-LISTEN:8080,reuseaddr,fork SYSTEM:'echo HTTP/1.1 200 OK; echo Content-Type\: text; echo; echo response msg.'
```

這行命令做了以下幾件事：

1. TCP-LISTEN:8080,reuseaddr,fork：這告訴 socat 監聽TCP 8080端口，reuseaddr 允許重用本地地址和端口，fork 則為每個連接創建一個新的進程。
2. SYSTEM:'...'：這部分執行一個shell命令，該命令生成HTTP響應。這裡的命令回應了一個簡單的HTML頁面。

這個命令會對所有發送到該端口的HTTP請求回應相同的頁面。這對於測試或模擬簡單的HTTP服務器非常有用。

## Request

```bash
curl -X POST http://192.168.94.5:8080/event \
     -H "Content-Type: multipart/form-data; boundary=MIME_boundary" \
     --data-binary @simple_body.txt
```
