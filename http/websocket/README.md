# Websocket

```bash
go run main.go
```

這個範例包括了一個基本的 WebSocket 伺服器，它可以接受來自客戶端的連線，並將收到的訊息回傳給客戶端。

當這個網頁加載並與 WebSocket 服務器建立連線後，服務器會定期向客戶端發送訊息，這些訊息會被顯示在網頁上。
運行你的 Go WebSocket 服務器，然後在瀏覽器中訪問 http://localhost:8080。這樣，你的瀏覽器應該會自動加載並顯示 web/index.html 文件的內容。

## Deploy

當你使用 ls | entr -r go run main.go 來自動重啟 Web 伺服器時，同時想要讓網頁瀏覽器自動重新整理，你可以考慮使用一些前端的解決方案。

### 使用 Browser-sync 工具

```bash
sudo npm install -g browser-sync
```

然後運行：

```bash
browser-sync start --proxy "localhost:8080" --files "**"
```

這將會開啟一個代理伺服器，監聽目錄下的文件變化，並在變化時刷新連接到該伺服器的所有瀏覽器。
