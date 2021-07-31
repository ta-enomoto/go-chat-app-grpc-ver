## チャットアプリケーション

### Go言語の標準・準標準ライブラリを中心に実装を目指したウェブアプリケーションです。
ウェブ関連技術について基礎から学ぶため、CRUDなど基礎的な機能から、双方向通信やAPIなど様々な技術が用いられている会員制のチャットアプリケーションを選択しました。またGo言語はフレームワークを用いずに実装する部分も多いと耳にし、ウェブアプリケーション構築を基礎から学ぶため、できる限りGo言語の標準・準標準ライブラリのみを使用しての実装を目指して作成しました。

### -アーキテクチャ図
<div  align="center">
<img  src="https://user-images.githubusercontent.com/63547862/127731206-94fad971-1510-4f1e-aed7-9152b30cc692.png"  width="700px">
</div>

### - 技術的特徴
  - **会員登録・ログイン機能**<br>
    チャットアプリケーションに不可欠な機能のため、会員登録・ログイン機能を実装しています。パスワードは平文ではなくハッシュ化して保存しています。
  - **Cookieによるセッション管理**<br>
    Cookieによるセッション管理を行い、セッション変数にはユーザー名を保存しています。セッションIDはオンメモリで保持、失効期限を設定し管理を行っています。また同時に、後述のAPIサーバーで投稿者を特定する際にセッション情報を利用するため、DB上でもセッション情報をオンメモリでの管理と連携して保持しています。
  - **双方向通信(WebSocket)によるリアルタイム更新(`/mypage/chatroom`内)**<br>
    チャットアプリケーションにとってリアルタイムのやり取りは不可欠な機能のため、各チャットルーム毎に独立してWebSocket接続が行われるよう実装を行っています。
  - **APIによるチャット内容の取得・投稿(`/mypage/chatroom`内)**<br>
    データ量が多いチャット内容の取得・頻繁に投稿されるチャットの処理もWEBサーバーで処理すると負荷が大きくなりがちなため、チャットの取得・投稿は別途API用サーバーを立てて行っています。(フレームワークにはgoaを使用)
  - **コンテナ内での作動(Docker)**<br>
    より実際の運用環境に近づけるため、webサーバー、DBサーバー、APIサーバーを別々のDockerコンテナで起動し、Docker-composeにより一つのシステムとして構築しています。(クラウドサーバー上でこのアプリ公開し続けることは費用・セキュリティ面などの問題で避けたく、ローカル環境で本アプリケーションの動作を確かめていただけるようにしたかったことも動機の一つです。)
  - **セキュリティ**<br>
    冒頭のパスワードのハッシュ化以外にも、メタ文字のエスケープによるXSS対策、プリペアドステートメントの使用によるSQLインジェクション対策、APIキーによる認証など基礎的なセキュリティ対策を行っています。また、あるチャットルームURLへ参加メンバーではないユーザーからアクセスがあっても、アクセスが拒否される設計になっています。
  - **管理ページ**<br>
    管理ページから、各ユーザーのID・パスワードの変更・ユーザーの削除、チャットルームの削除が行えます。管理ページもセッション管理によりアクセス制御を行っており、管理者以外のユーザーのアクセスも防止しています。


### - 今後の課題
  - ~~管理ページの実装~~ **(2021/7/25 管理ページの実装)**
  - ~~API認証の実装~~ **(2021/7/23 APIキーによる認証実装)**
  - ~~ルーム一覧とチャットスペースを同一ページで表示~~ **(2021/7/24 レイアウト変更)**
  - ~~ユーザー管理用とチャット管理用のデータベースの分離~~ **(2021/7/27 DBを別々のコンテナに分離)**
  - フレームワークを用いた同様のアプリケーションの作成
  - TypeScriptの適用
  - 新着メッセージの通知

**ユーザー登録**
①コンテナを立ち上げ後、ログインページ([http://172.25.0.2/login](http://172.25.0.2/login))へアクセスします。「新規登録はこちら」から`/resistration`ページへ移動し、フォームにユーザー名・パスワード(10文字以内、ユーザー名"admin"は不可、パスワードは英数字)を入力し登録ボタンをクリックしてください。チャット機能の試行には最低2名の登録が必要なため、2人分のユーザーを登録してください。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903125-5e54e460-2a72-45bc-a6ba-e08d068d62e9.png" width="700px">
</div>

<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903149-757def6a-7592-415c-a935-fb16be58b8c6.png" width="700px">
</div>

②登録が完了すると、登録完了ページに遷移するので、「ログインページに戻る」をクリックしてください。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903165-1a370c47-183f-4307-b2d3-771ae73b7dce.png" width="700px">
</div>

**ログイン・ルームの作成・チャットの投稿**
①登録した、いずれかのユーザーでログインしてください。

②マイページが表示されます。新規ルーム名(20文字以内)と、もう一人の作成済ユーザーを相手ユーザー名に入力し、「作成する」ボタンを押してください。「ルーム一覧」に作成したルームが表示されます。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903184-3d72e5c8-3e8e-4ec9-b8a9-71f4243c82ae.png" width="700px">
</div>

③作成したルームをクリックすると、右のスペースにチャットページが表示されます。「新規投稿」フォームにチャット内容(255文字以内)を入力し、「メッセージを投稿」ボタンを押すと投稿が行われます。投稿内容はページの更新を経ずに反映されます。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903203-1552f4f1-3482-43c4-9abd-c97e1610aa6d.png" width="700px">
</div>

<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903232-d31598c8-0f62-46f2-a1e7-6177ecb9625a.gif" width="700px">
</div>

④右上にある「このルームを削除する」ボタンを押すと、そのチャットルームが削除されます。(クリック一回で削除されます。)
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126905110-b1bdaf95-6874-4a30-822e-c79e9d3d779b.png" width="700px">
</div>

**ログアウト**
①「ログアウトをこちら」をクリックすると、ログアウトします。(Cookieも削除されます。)
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903536-9a08b64f-c0b9-49a2-a18e-3f51602c118b.png" width="700px">
</div>

**セッション切れ**
①セッションの有効期間は1時間になっています。操作しないまま1時間経つとセッション切れになり、ページ遷移時に以下のページが表示されます。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903695-2c248650-25ad-4277-bc21-524ad3b2c45d.png" width="700px">
</div>

**退会**
①マイページにて「退会はこちら」リンクをクリックすると、退会ページに遷移します。  ログイン中のユーザーのユーザーID・パスワードを入力して「送信」ボタンを押すと、ユーザーが削除されます。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903254-c2cdbc9a-ffe9-461a-bc5d-35ac14610da5.png" width="700px">
</div>

②ユーザーの削除に成功すると、以下の画面が表示されます。
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126903272-689fa1a5-094d-4371-aaae-e94a1a752b7d.png" width="700px">
</div>


### 【管理ページの方法】
**ログイン**
①管理ページログインページ([http://172.25.0.2/admin/login](http://172.25.0.2/admin/login))へアクセスします。(管理者ID：`admin`、パスワード：`pass`)
<div align="center">
<img src="https://user-images.githubusercontent.com/63547862/126904058-59a331cc-adee-478b-96ff-b1841b26ba02.png" width="700px">
</div>

②ログイン後、以下の管理ページへ遷移します。ログアウトは、右上のリンクから行えます。
<div align="center">
<kbd><img src="https://user-images.githubusercontent.com/63547862/126904134-75242cf4-7ced-45d4-9429-a39afa6f206a.png" width="700px"></kbd>
</div>

**ユーザーの管理**
①「登録ユーザー一覧」をクリックすると、以下のページに遷移します。このページでは、登録中のユーザーが全員表示されます。
<div align="center">
<kbd><img src="https://user-images.githubusercontent.com/63547862/126904211-0ce2cb6e-7848-4665-a857-8ce54ce24318.png" width="700px"></kbd>
</div>

②管理を行いたいユーザーをクリックすると、以下の個別ユーザー管理ページに遷移します。このページでは、(a)ユーザーIDの変更、(b)パスワードの変更、(c)ユーザーの削除が行えます。
<div align="center">
<kbd><img src="https://user-images.githubusercontent.com/63547862/126904324-4d4da3f3-b358-4a7e-8039-a8ac1217e281.png" width="700px"></kbd>
</div>

**チャットルームの管理**
①「チャットルーム一覧」をクリックすると、以下のページに遷移します。このページでは、全てのチャットルームが表示されます。
<div align="center">
<kbd><img src="https://user-images.githubusercontent.com/63547862/126904607-51b8645b-7cad-4b83-a2f9-ca623a3369a1.png" width="700px"></kbd>
</div>

②管理を行いたいユーザーをクリックすると、以下の個別チャットルーム管理ページに遷移します。このページではチャット内容が表示され、右上のボタンからチャットルームの削除が行えます。
<div align="center">
<kbd><img src="https://user-images.githubusercontent.com/63547862/126904775-2e0ee16f-3b53-49af-a2d5-fb23cdd61d44.png" width="700px"></kbd>
</div>