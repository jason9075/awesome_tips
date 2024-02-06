# pre-commit

要在專案中使用 pre-commit 來檢查 `.env` 和 `.env-example` 文件是否匹配，你可以透過撰寫一個自定義的 pre-commit hook 腳本來實現。這個腳本會在每次提交前運行，檢查 `.env` 和 `.env-example` 文件中的變數是否一致。下面是一個實現這個功能的基本步驟：

1. **安裝 Pre-commit：**
   確保你已經安裝了 pre-commit。如果還沒有安裝，可以透過 pip 安裝：

   ```sh
   pip install pre-commit
   ```

2. **建立 Pre-commit Hook 設定：**
   在專案根目錄下建立一個 `.pre-commit-config.yaml` 檔案，並加入自定義 hook 的設定。你將需要指定執行的腳本。

3. **撰寫檢查腳本：**
   建立一個檢查 `.env` 和 `.env-example` 文件的腳本。這個腳本會比較這兩個文件中定義的變數，並在發現不一致時給出警告或錯誤。

4. **安裝到 Git Hooks**
   在專案目錄下運行了 pre-commit install 命令。這個命令會將 pre-commit 安裝到你的 git hooks 中，確保每次 commit 時都會觸發 pre-commit 檢查。

- **.pre-commit-config.yaml 配置範例：**

  ```yaml
  repos:
    - repo: local
      hooks:
        - id: env-check
          name: Check .env and .env-example
          entry: bash scripts/check_env_example.sh
          language: script
          always_run: true
          files: '\.env$|\.env-example$'
  ```

- **check_env_example.sh 腳本範例：**

  ```bash
  #!/bin/bash

  if [ ! -f .env ]; then
    exit 0
  fi

  ENV_FILE_KEY=$(grep -v -e '^#' -e '^$' .env | cut -d= -f1 | sort)
  ENV_EXAMPLE_FILE_KEY=$(grep -v -e '^#' -e '^$' .env-example | cut -d= -f1 | sort)

  diff <(echo "$ENV_FILE_KEY") <(echo "$ENV_EXAMPLE_FILE_KEY") > /dev/null

  if [ $? -ne 0 ]; then
    echo "Error: .env and .env-example files do not match!"
    exit 1
  fi
  ```
