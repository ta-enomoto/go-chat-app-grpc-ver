let socket = null;
let data = "";
let wsuri = "ws://172.25.0.2/wsserver";
let allchats =""
window.onload = async function() {
  
  let url = location.href;
  let roomid = url.replace("http://172.25.0.2/mypage/chatroom","");

  const urlForApi = "http://172.25.0.4:8000/chatroom/" + roomid
  
  //チャット読み込み
  async function getChatFromApi() {
    try {
       res = await axios.get(urlForApi);
       allchats = res.data;
      for (const chat of allchats) {
        let text = document.createTextNode(chat.Chat + '　(投稿者：' + chat.UserId + '　投稿日：' + chat.PostDt + ')');
        let newLi = document.createElement("li");
        newLi.appendChild(text);
        let chatList = document.getElementById("chats")
        chatList.appendChild(newLi);
      }
      return allchats
    } catch(error){
      const {
        status,
        statusText
      } = error.response;
      console.log(`Error! HTTP Status: ${status} ${statusText}`);
    }
  }
  await getChatFromApi()
  
    //WebSocket
    socket = new WebSocket(wsuri);

    socket.onopen = function() {
      console.log("connected");
      class　Newchat {
        constructor(id, userid, roomname, member, chat, postdt){
          this.id = id;
          this.userid = userid;
          this.roomname = roomname;
          this.member = member;
          this.chat = chat;
          this.postdt = postdt;
        }
      }
    let roomname = allchats[0].RoomName;
    let userid = allchats[0].UserId;
    let member = allchats[0].Member;
    let postdt = Date.now();
    let chat = "first contact";
    const newchat = new Newchat(roomid, userid, roomname, member, chat, postdt);
    socket.send(JSON.stringify(newchat));
    console.log(JSON.stringify(newchat));
    }
    socket.onmessage = function(e) {
      console.log("message recieved" + e.data);
      let chatobj = JSON.parse(e.data)
      let text = document.createTextNode(chatobj.chat + '　(投稿者：' + chatobj.userid + '　投稿日：' + chatobj.postdt + ')');
      let newLi = document.createElement("li");
      newLi.appendChild(text);
      let chatList = document.getElementById("chats")
      chatList.appendChild(newLi);
    }
    socket.onclose = function(e) {
      console.log("connection closed");
    }
};
  
  function send() {
      class　Newchat {
        constructor(id, userid, roomname, member, chat, postdt, cookie){
          this.id = id;
          this.userid = userid;
          this.roomname = roomname;
          this.member = member;
          this.chat = chat;
          this.postdt = postdt;
          this.cookie = cookie;
        }
      }

      let url = location.href;
      let roomid = url.replace("http://172.25.0.2/mypage/chatroom","");
      
      let chat = document.getElementById('chat').value;
      if (chat == "") {
        window.alert("チャットが入力されていません");
        return;
      };
      
      let roomname = allchats[0].RoomName;
      let userid = allchats[0].UserId;
      let member = allchats[0].Member;
      var date = Date.now();
      let postdt = new Date(date);
                
      let cookieValue = document.cookie;
      let cookie = cookieValue.replace("cookieName=","");
      const newchat = new Newchat(roomid, userid, roomname, member, chat, postdt, cookie);
      socket.send(JSON.stringify(newchat));
      document.chatform.reset();
      console.log(JSON.stringify(newchat));
    };