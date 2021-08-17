const { GetchatRequest } = require('./modules/chatapi_pb');
const { PostchatRequest } = require('./modules/chatapi_pb');
const { ChatapiClient } = require('./modules/chatapi_grpc_web_pb');

//WebSocketインスタンスと取得したチャットは各関数で使用するため、グローバルで定義しておく
let socket = null;
let allchats = "";

//ウィンドウ表示時に、APIからのチャットの取得と、WebSocketのハンドシェイク処理を行う
window.onload = function () {

  //API、WebSocket共通で使用するルームIDは、URLから取得
  let url = location.href;
  let roomid = url.replace("http://172.26.0.2/mypage/chatroom", "");

  const client = new ChatapiClient('http://172.26.0.6:9000', null, null);
  const request = new GetchatRequest();

  request.setId(roomid);

  client.getchat(request, { "Authorization": "apikey" }, (err, response) => {
    if (err) {
      console.log(`Unexpected error for getChat: code = ${err.code}` + `, message = "${err.message}"`);
    } else {
      let res = response.toObject();
      allchats = res['fieldList']

      for (const chat of allchats) {
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

      let roomNameUnescaped = allchats[0]['roomName'].replace(/\\/g, "");
      let roomnameText = document.createTextNode(roomNameUnescaped);
      let newH2 = document.createElement("h2");
      newH2.appendChild(roomnameText);
      let roomname = document.getElementById("roomname-header");
      roomname.appendChild(newH2);

      var element = document.documentElement;
      var bottom = element.scrollHeight - element.clientHeight;
      window.scroll(0, bottom);

      //以降の処理でも使用するため、allchatsをreturnする
      return allchats;
    };
  });

  //gRPC-WebのPOST処理はasync awaitが機能しなかったためsetTimeを使用(要改良)
  setTimeout(
    function () {
      let wsuri = "ws://172.26.0.2/wsserver";
      socket = new WebSocket(wsuri);

      socket.onopen = function () {
        console.log("connected");

        //どのWebSocket用チャットルームにアクセスがあったか識別するため、WebSocket開通時に初期投稿を行っておく
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

        socket.onmessage = function (e) {
          console.log("message recieved" + e.data);

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

        socket.onclose = function (e) {
          console.log("connection closed");
        };
      };
    },
    "1500"
  );
};

window.postChat = function postChat() {
  //投稿者を特定するため、cookieも一緒に送信する
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

  let url = location.href;
  let roomid = url.replace("http://172.26.0.2/mypage/chatroom", "");

  let chat = document.getElementById('chat').value;
  if (chat == "") {
    window.alert("チャットが入力されていません");
    return;
  };
  let chatEscaped = escape(chat);

  let roomname = allchats[0]['roomName'];
  let userid = allchats[0]['userId'];
  let member = allchats[0]['member'];
  let millisec = Date.now();
  let date = new Date(millisec);
  let postdt = date.toISOString();

  let cookieValue = document.cookie;
  let cookie = cookieValue.replace("cookieName=", "");

  const newchat = new Newchat(roomid, userid, roomname, member, chatEscaped, postdt, cookie);

  const newchatJSON = JSON.stringify(newchat);
  socket.send(newchatJSON);

  const client = new ChatapiClient('http://172.26.0.6:9000', null, null);
  const request = new PostchatRequest();

  request.setId(roomid)
  request.setUserId(userid)
  request.setRoomName(roomname)
  request.setMember(member)
  request.setChat(chatEscaped)
  request.setPostDt(postdt)
  request.setCookie(cookie)

  client.postchat(request, { "Authorization": "apikey" }, (err, response) => {
    if (err) {
      console.log(`Unexpected error for getChat: code = ${err.code}` + `, message = "${err.message}"`);
    }
  });

  document.chatform.reset();
  console.log(JSON.stringify(newchat));
};

window.deleteChtrmFunc = function deleteChtrmFunc() {

  if (window.confirm("本当にこのルームを削除しますか？")) {
    this.form.submit();
  } else {
    window.alert("ルーム削除をキャンセルしました");
    return false
  };
};