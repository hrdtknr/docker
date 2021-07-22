## nodeのinstarl(道具の調達)
## debianを取ってきたうえでnodeとかも設定する書き方
## alpineの書き方はやったことがある
FROM node:14

WORKDIR /user/src/app

#コピー
COPY package*.json ./
COPY tsconfig.json ./
COPY src/* ./src/

# packageのコピーではなく、最初にinistallさせる
RUN npm install


# 最初に再審でbuild
RUN npm run build

# COPY . .

#解放するポート番号
EXPOSE 3000

CMD ["node", "build/index.js"]


## 複数のデータセンターに同じものを置くときに便利
## dokcerだけintallしてあれば、同じ状況を再現できる
## 他のコンテナに影響が出ないので、好き勝手出来る
## 依存関係
## マイクロサービス
## 例えば、Amazonの検索と会計機能が独立してる
## 検索だけを止めて改修ができる
## 責任の分解にも 


## 差分バックアップ
## フルバック
## 差分とりつつ、定期的に全部バックアップする考え



## スクラム=>新規で来たものは受け付けない（つっぱねる）、が正解
## ただしそれができるとはかぎらない

