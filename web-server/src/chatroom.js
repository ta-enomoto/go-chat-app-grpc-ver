const { GoaChat } = require('./modules/chatapi_pb');
const { ChatapiClient } = require('./modules/chatapi_grpc_web_pb');

let socket = null;
let data = "";
let wsuri = "ws://172.26.0.2/wsserver";
let allchats = "";

//ウィンドウ表示時に、APIからのチャットの取得と、WebSocketのハンドシェイク処理を行う
window.onload = function () {

  //   //API、WebSocket共通で使用するURLからのルームID取得処理
  //   let url = location.href;
  //   let roomid = url.replace("http://172.26.0.2/mypage/chatroom","");

  // //チャット読み込み処理
  //   //APIリクエスト(GET)先のURL
  //   const urlForApiGet = "http://172.26.0.3:8000/chatroom/" + roomid;

  //   //headersにAPIキー認証用のAuthorizationヘッダーを設定
  //   const axiosConfig = {
  //     headers: {
  //       "Authorization": "apikey",
  //     }
  //   };

  const client = new ChatapiClient('http://172.26.0.6:9000', null, null);

  const request = new GoaChat();
  let roomid = 2;
  request.setId(roomid);

  client.getchat(request, {}, (err, response) => {//{"Authorization": "apikey"}, (err, response) => {
    if (err) {
      console.log(`Unexpected error for getChat: code = ${err.code}` + `, message = "${err.message}"`);
    } else {
      let res = response.toObject();
      allchats = res['fieldList']
      //console.log(allchats['fieldList'][0]['userId']);

      for (const chat of allchats) {
        //各チャット毎にHTML要素を生成・順に追加していく
        let textUserNode = document.createTextNode(chat['userId']);
        let textPostDtNode = document.createTextNode(chat['postDt']);
        let textChatUnescaped = unescape(chat['chat']).replace(/\\/g, "");
        let textChatNode = document.createTextNode(textChatUnescaped);

        let elUser = document.createElement("div");
        elUser.appendChild(textUserNode);
        elUser.id = "user";
        elUser.style = "display: inline-block; _display: inline;";

        let elPostDt = document.createElement("div");
        elPostDt.appendChild(textPostDtNode);
        elPostDt.id = "postdt";
        elPostDt.style = "display: inline-block; _display: inline;";

        let elChat = document.createElement("div");
        elChat.appendChild(textChatNode);
        elChat.id = "chatText";

        let newLi = document.createElement("li");
        newLi.appendChild(elUser);
        newLi.appendChild(elPostDt);
        newLi.appendChild(elChat);
        let chatList = document.getElementById("chats");
        chatList.appendChild(newLi);
      };

      //ルーム名を一番目のルーム作成時投稿チャットから取得しHTML要素を生成・追加する
      let roomNameUnescaped = allchats[0]['roomName'].replace(/\\/g, "");
      let roomnameText = document.createTextNode(roomNameUnescaped);
      let newH2 = document.createElement("h2");
      newH2.appendChild(roomnameText);
      let roomname = document.getElementById("roomname-header");
      roomname.appendChild(newH2);

      //以降の処理でも使用するため、allchatsをreturnする
      return allchats;
    };
  });

  //全チャット表示後、ページ最下部にスクロールする
  var element = document.documentElement;
  var bottom = element.scrollHeight - element.clientHeight;
  window.scroll(0, bottom);

  //async awaitが使えないためsetTimeを使用
  setTimeout(
    function () {
      //WebSocketハンドシェイク
      socket = new WebSocket(wsuri);

      //WebSocket開通時の処理
      socket.onopen = function () {
        console.log("connected");

        //WebSocket用チャットルームの識別のため、WebSocket開通時に初期投稿を行っておく
        class Newchat {
          constructor(id, userid, roomname, member, chat, postdt) {
            this.id = id;
            this.userid = userid;
            this.roomname = roomname;
            this.member = member;
            this.chat = chat;
            this.postdt = postdt;
          }
        }
        let roomname = allchats[0]['roomName'];
        let userid = allchats[0]['userId'];
        let member = allchats[0]['member'];
        let postdt = Date.now();
        let chat = "first contact";
        const newchat = new Newchat(roomid, userid, roomname, member, chat, postdt);
        socket.send(JSON.stringify(newchat));
        console.log(JSON.stringify(newchat));

        //メッセージ受信時の処理
        socket.onmessage = function (e) {
          console.log("message recieved" + e.data);
          //受信したメッセージを元にHTML要素を生成・追加する
          let chatobj = JSON.parse(e.data);
          let textUserNode = document.createTextNode(chatobj.Chatroom.userId);
          let textPostDtNode = document.createTextNode(chatobj.PostDt);
          let textChatUnescaped = unescape(chatobj.Chat).replace(/\\/g, "");
          let textChatNode = document.createTextNode(textChatUnescaped);

          let elUser = document.createElement("div");
          elUser.appendChild(textUserNode);
          elUser.id = "user";
          elUser.style = "display: inline-block; _display: inline;";

          let elPostDt = document.createElement("div");
          elPostDt.appendChild(textPostDtNode);
          elPostDt.id = "postdt";
          elPostDt.style = "display: inline-block; _display: inline;";

          let elChat = document.createElement("div");
          elChat.appendChild(textChatNode);
          elChat.id = "chatText";

          let newLi = document.createElement("li");
          newLi.appendChild(elUser);
          newLi.appendChild(elPostDt);
          newLi.appendChild(elChat);
          let chatList = document.getElementById("chats");
          chatList.appendChild(newLi);

          var element = document.documentElement;
          var bottom = element.scrollHeight - element.clientHeight;
          window.scroll(0, bottom);
        };
        //WebSocket終了時の処理
        socket.onclose = function (e) {
          console.log("connection closed");
        };
      };
    },
    "1000"
  );
};

//チャット投稿時の処理
window.send = function send() {
  //投稿用のチャットクラスを定義
  class Newchat {
    constructor(id, userid, roomname, member, chat, postdt, cookie) {
      this.id = id;
      this.userid = userid;
      this.roomname = roomname;
      this.member = member;
      this.chat = chat;
      this.postdt = postdt;
      this.cookie = cookie;
    };
  };

  //URLから当該ルームIDを取得
  let url = location.href;
  let roomid = url.replace("http://172.26.0.2/mypage/chatroom", "");

  //チャット欄が空欄で投稿ボタンが押された時は処理中断・alertで通知する
  let chat = document.getElementById('chat').value;
  if (chat == "") {
    window.alert("チャットが入力されていません");
    return;
  };
  let chatEscaped = escape(chat);

  //チャット・cookie外の値を代入する(ルーム名・ユーザーID・メンバーはルーム作成時投稿チャットから取得)
  let roomname = allchats[0]['roomName'];
  let userid = allchats[0]['userId'];
  let member = allchats[0]['member'];
  let date = Date.now();
  let postdt = new Date(date);

  //ブラウザに保存されているcookieを取得する
  let cookieValue = document.cookie;
  let cookie = cookieValue.replace("cookieName=", "");

  //新規チャットインスタンスを生成する
  const newchat = new Newchat(roomid, userid, roomname, member, chatEscaped, postdt, cookie);

  //JSONにエンコードして投稿
  const newchatJSON = JSON.stringify(newchat);
  socket.send(newchatJSON);

  //APIリクエスト(POST)先のURLを設定
  const urlForApiPost = "http://172.26.0.3:8000/chatroom/chat";

  //headersにAPIキー認証用のAuthorizationヘッダーを設定
  // const axiosConfig = {
  //   headers: {
  //     "Authorization": "apikey",
  //   }
  // };

  //APIを叩く関数(POST)
  async function postChatToApi() {
    try {
      res = await axios.post(urlForApiPost, newchatJSON)//, axiosConfig);
      if (res.data == true) {
        console.log("投稿成功");
      } else {
        console.log("投稿失敗");
        return;
      };
    } catch (error) {
      const {
        status,
        statusText
      } = error.response;
      console.log(`Error! HTTP Status: ${status} ${statusText}`);
    };
  };
  postChatToApi();

  //投稿後、チャット投稿欄は空欄に戻す
  document.chatform.reset();
  console.log(JSON.stringify(newchat));
};

//チャットルーム削除ボタンがクリックされた時の処理
//削除実行前に、確認ウィンドウを表示する
window.deleteChtrmFunc = function deleteChtrmFunc() {

  if (window.confirm("本当にこのルームを削除しますか？")) {
    this.form.submit();
  } else {
    window.alert("ルーム削除をキャンセルしました");
    return false
  };
};