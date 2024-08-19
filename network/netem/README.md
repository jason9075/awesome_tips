# 使用netem模擬網絡問題

在Linux系統中，可以使用netem工具來模擬網絡延遲、丟包等情況。以下是一些基本的命令：

## 模擬延遲

```bash
sudo tc qdisc add dev wlan0 root netem delay 100ms
```

## 模擬丟包

```bash
sudo tc qdisc add dev wlan0 root netem loss 10%
```

## 模擬延遲和丟包

```bash

sudo tc qdisc add dev wlan0 root netem delay 100ms loss 10%
```

這些方法可以幫助你模擬WiFi訊號很差的狀況，從而測試你的設備或應用在這種環境下的表現。

## Install Steps

1. 安裝netem工具
   netem 是內建在 Linux 的 tc（Traffic Control）工具中的一部分，因此只需要安裝 tc 即可。

執行以下命令來安裝 tc 工具：

```bash
sudo apt-get update
sudo apt-get install iproute2
```

iproute2 套件中包含了 tc 工具，安裝完成後即可使用 netem 來模擬網絡問題。

2. 檢查netem是否正常運行
   安裝完成後，你可以使用以下命令來檢查 netem 是否可以正常運行：

```bash
sudo tc qdisc show dev wlan0
```

這應該會顯示出當前網絡設備（如 wlan0）的隊列紀律（qdisc），確認已安裝的 netem 可用。

3. 使用netem模擬網絡狀況
   安裝好 tc 工具後，你可以按照之前提供的命令來模擬網絡延遲、丟包等狀況。

例如，模擬 100ms 的網絡延遲：

```bash
sudo tc qdisc add dev wlan0 root netem delay 100ms
```

這樣，你的樹莓派將會在網絡通信中模擬出 100 毫秒的延遲，這對測試 WiFi 訊號差的狀況非常有幫助。

4. 恢復正常網絡狀態
   當你不再需要模擬網絡問題時，可以使用以下命令來移除 netem 設置，恢復正常的網絡狀態：

```bash
sudo tc qdisc del dev wlan0 root netem
```

這樣你的網絡設備就會回到正常狀態，不再模擬任何網絡問題。
