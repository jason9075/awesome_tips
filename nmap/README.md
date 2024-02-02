使用 `nmap` 掃描特定的子網段（例如 `192.168.43.x`）並找出哪些裝置是活躍的，可以透過以下幾種不同的命令和掃描技術來完成。以下是一些基礎的命令示例，您可以根據您的需要進行調整：

### 1. 基本掃描

要掃描 `192.168.43.0/24` 子網段中的所有 IP 地址（即 `192.168.43.1` 到 `192.168.43.254`），您可以使用以下命令：

```sh
nmap 192.168.43.0/24
```

這個命令會對每個 IP 地址執行一個基本的掃描，檢查預設的端口是否開放。

### 2. 快速掃描

如果您想要更快地獲得哪些裝置是活躍的結果，可以使用 `-sn` 選項進行「ping 掃描」（不掃描端口）：

```sh
nmap -sn 192.168.43.0/24
```

這個命令將只檢查裝置是否在線，而不檢查任何端口。

### 3. 掃描特定端口

如果您只對某些端口感興趣，可以使用 `-p` 選項指定端口。例如，要掃描所有裝置的 80 和 443 端口：

```sh
nmap -p 80,443 192.168.43.0/24
```

### 4. 使用更進階的掃描技術

`nmap` 提供了多種進階掃描選項，例如版本偵測（`-sV`）、作業系統偵測（`-O`）等：

```sh
nmap -sV -O 192.168.43.0/24
```

這個命令會嘗試識別端口上運行的服務版本以及網絡上裝置的作業系統。

### 注意事項

- 在執行 `nmap` 掃描之前，確保您有權掃描目標網絡。未經授權的掃描可能會違反法律或政策。
- 根據您的網絡環境和安全設置，某些掃描技術可能無法獲得預期的結果。
- 使用 `nmap` 的進階功能時，可能需要更多的時間來完成掃描，並且對網絡造成的影響也會更大。

以上就是使用 `nmap` 掃描特定子網段中有反應的裝置的基本方法。您可以根據具體需求調整命令和選項。