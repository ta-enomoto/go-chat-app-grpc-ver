const { GetchatRequest } = require('./modules/chatapi_pb');
const { ChatapiClient } = require('./modules/chatapi_grpc_web_pb');

//ウィンドウ表示時に、APIからのチャットの取得する
window.onload = function () {

  //APIで使用するルームIDは、URLから取得
  let url = location.href;
  let roomid = url.replace("http://172.26.0.2/admin/chatrooms/chatroom","");

  const client = new ChatapiClient('http://172.26.0.6:9000', null, null);
  const request = new GetchatRequest();

  request.setId(roomid);

  client.getchat(request, {"Authorization": "apikey"}, (err, response) => {
    if (err) {
      console.log(`Unexpected error for getChat: code = ${err.code}` + `, message = "${err.message}"`);
    } else {
      let res = response.toObject();
      let allchats = res['fieldList']

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
    };
  });

  //gRPC-WebのPOST処理はasync awaitが機能しなかったためsetTimeを使用(要改良)
  setTimeout(
    function () {
      let element = document.documentElement;
      let bottom = element.scrollHeight - element.clientHeight;
      window.scroll(0, bottom);
    },
    "1000"
  );


};

window.deleteChtrmFunc = function deleteChtrmFunc() {

  if (window.confirm("本当にこのルームを削除しますか？")) {
    this.form.submit();
  } else {
    window.alert("ルーム削除をキャンセルしました");
    return false
  };
};