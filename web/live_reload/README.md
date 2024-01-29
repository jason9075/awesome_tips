### 使用 Live Reload 工具

有許多工具和擴展可以幫助實現這一功能，比如 `LiveReload` 或 `BrowserSync`。這些工具會監視文件變化並自動刷新瀏覽器。

#### 安裝 LiveReload

這些工具可以透過 npm 安裝，並且可以很容易地集成到你的開發工作流中。例如，使用 `BrowserSync`：

```bash
npm install -g browser-sync
```

然後運行：

```bash
browser-sync start --proxy "localhost:8080" --files "."
```

這將會開啟一個代理伺服器，監聽目錄下的文件變化，並在變化時刷新連接到該伺服器的所有瀏覽器。
