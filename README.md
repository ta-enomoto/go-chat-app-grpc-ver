<<<<<<< HEAD
## チャットアプリケーション
### Go言語の標準・準標準ライブラリのみでの実装を目指したアプリケーションです。
### - 技術的特徴  
  - 会員登録機能
  - ログイン機能
  - Cookieによるセッション管理
  - XMLHttpRequestによる非同期通信(`/mypage`内)
  - WebSocketによる双方向通信(`/mypage/chatroom`内)
  - コンテナ内での作動(Docker)

### - 今後の課題
  - 投稿内容をWebAPIから取得
  - 日本語対応
  - フレームワークを用いた同様のサイト作成
  - コードの整理

### 【使用方法】
①「新規登録」ページへ移動し、フォームにユーザー名・パスワード(共に英数字、10文字以内)を入力し<br>　登録ボタンをクリックしてください。このアプリケーションは最低2名の登録が必要なため、2人ユーザー<br>　を登録してください。

②登録が完了すると、`/resistration`ページに遷移するので、「ログインページに戻る」をクリックして<br>　ください。

③登録したいずれかのユーザーでログインしてください。

#### 注意点  
　ログイン時にセッションがメモリ上に生成され、ブラウザにCookieがセットされます。ログイン中に<br>　コンテナを立ち上げ直すとメモリ上のセッションも破棄され、セッション切れとなります。その場合、<br>　開発者ツールよりCookieを一度削除してください。

④マイページが表示されます。新規ルーム名(英数字、20文字以内)と、もう一人の作成済ユーザーを相手<br>　ユーザー名に入力し、「作成する」ボタンを押してください。「ルーム一覧」に作成したルームが表示さ<br>　れます。

⑤作成したルームをクリックすると、チャットページへ遷移します。  「新規投稿」フォームにチャット内容<br>　(英数字、255文字以内)を入力し、「メッセージを投稿」ボタンを押すと投稿が行われます。投稿後は、自動<br>　でページが更新され、投稿内容が反映されます。

#### 注意点
　 ~~現在、投稿時にWebSocketの接続が途切れページの自動更新がされない不具合が発生することがあります。(改修予定)~~
　(2021/06/02改修済)



⑥「このルームを削除する」ボタンを押すと、そのチャットルームが削除されます。<br>　(クリック一回で削除されます。)

⑦マイページにて「退会はこちら」リンクをクリックすると、退会ページに遷移します。  ログイン中<br>　のユーザーのユーザーID・パスワードを入力して「送信」ボタンを押すと、ユーザーが削除されます。

=======
# go-chat-app-api-ver
>>>>>>> origin/main