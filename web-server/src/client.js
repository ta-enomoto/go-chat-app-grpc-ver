const {GoaChat} = require('./modules/chatapi_pb');
const {ChatapiClient} = require('./modules/chatapi_grpc_web_pb');

let allchats ="";

//ウィンドウ表示時に、APIからのチャットの取得する
window.onload = async function() {
  
  // //APIで使用するURLからのルームID取得処理
  // let url = location.href;
  // let roomid = url.replace("http://172.26.0.2/admin/chatrooms/chatroom","");

  // //APIリクエスト(GET)先のURL
  // const urlForApiGet = "http://172.26.0.3:8000/chatroom/" + roomid;

  // //headersにAPIキー認証用のAuthorizationヘッダーを設定
  // const axiosConfig = {
  //   headers: {
  //     "Authorization": "apikey",
  //   }
  // };

  const client = new ChatapiClient('http://172.26.0.6:9000', null, null);

  const request = new GoaChat();
  let id = 1;
  request.setId(id);

  async function getChatFromApi() {
    await client.getchat(request, {}, (err, response) => {//{"Authorization": "apikey"}, (err, response) => {
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
          elUser.id ="user";
          elUser.style = "display: inline-block; _display: inline;";
          
          let elPostDt = document.createElement("div");
          elPostDt.appendChild(textPostDtNode);
          elPostDt.id ="postdt";
          elPostDt.style = "display: inline-block; _display: inline;";
          
          let elChat = document.createElement("div");
          elChat.appendChild(textChatNode);
          elChat.id ="chatText";
  
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

      return;
      };
    });
  };

  //APIのチャット取得処理が完了するまで、以降の処理を待つ
  await getChatFromApi();

  //全チャット表示後、ページ最下部にスクロールする
  let element = document.documentElement;
  let bottom = element.scrollHeight - element.clientHeight;
  window.scroll(0, bottom);
};

//チャットルーム削除ボタンがクリックされた時の処理
//削除実行前に、確認ウィンドウを表示する
function deleteChtrmFunc(){

	if(window.confirm("本当にこのルームを削除しますか？")){
    this.form.submit();
	} else {
    window.alert("ルーム削除をキャンセルしました");
    return false
  };
};