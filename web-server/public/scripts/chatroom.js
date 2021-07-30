let socket = null;
let data = "";
let wsuri = "ws://172.25.0.2/wsserver";
let allchats =""
window.onload = async function() {
  
  let url = location.href;
  let roomid = url.replace("http://172.25.0.2/mypage/chatroom","");

  const urlForApiGet = "http://172.25.0.3:8000/chatroom/" + roomid
  
  //チャット読み込み
  async function getChatFromApi() {
    try {
       res = await axios.get(urlForApiGet, { headers: { "Authorization": "apikey" } });
       allchats = res.data;
      for (const chat of allchats) {
        let textUser = document.createTextNode(chat.UserId);
        let textPostDt = document.createTextNode(chat.PostDt);
        let textChat = document.createTextNode(chat.Chat);

        let elUser = document.createElement("div");
        elUser.appendChild(textUser);
        elUser.id ="user";
        elUser.style = "display: inline-block; _display: inline;";
        
        let elPostDt = document.createElement("div");
        elPostDt.appendChild(textPostDt);
        elPostDt.id ="postdt";
        elPostDt.style = "display: inline-block; _display: inline;";
        
        let elChat = document.createElement("div");
        elChat.appendChild(textChat);
        elChat.id ="chatText";

        let newLi = document.createElement("li");
        newLi.appendChild(elUser);
        newLi.appendChild(elPostDt);
        newLi.appendChild(elChat);
        let chatList = document.getElementById("chats");
        chatList.appendChild(newLi);
      }
      
      let roomnameText = document.createTextNode(allchats[0].RoomName);
      let newH2 = document.createElement("h2");
      newH2.appendChild(roomnameText);
      let roomname = document.getElementById("roomname-header");
      roomname.appendChild(newH2);
      
      return allchats;
    } catch(error){
      const {
        status,
        statusText
      } = error.response;
      console.log(`Error! HTTP Status: ${status} ${statusText}`);
    }
  }
  await getChatFromApi()
  
  var element = document.documentElement;
  var bottom = element.scrollHeight - element.clientHeight;
  window.scroll(0, bottom);
  
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
    let chatobj = JSON.parse(e.data);
    let textUser = document.createTextNode(chatobj.Chatroom.userId);
    let textPostDt = document.createTextNode(chatobj.PostDt);
    let textChat = document.createTextNode(chatobj.Chat);

    let elUser = document.createElement("div");
    elUser.appendChild(textUser);
    elUser.id ="user";
    elUser.style = "display: inline-block; _display: inline;";
    
    let elPostDt = document.createElement("div");
    elPostDt.appendChild(textPostDt);
    elPostDt.id ="postdt";
    elPostDt.style = "display: inline-block; _display: inline;";
    
    let elChat = document.createElement("div");
    elChat.appendChild(textChat);
    elChat.id ="chatText";

    let newLi = document.createElement("li");
    newLi.appendChild(elUser);
    newLi.appendChild(elPostDt);
    newLi.appendChild(elChat);
    let chatList = document.getElementById("chats");
    chatList.appendChild(newLi);
    
    var element = document.documentElement;
    var bottom = element.scrollHeight - element.clientHeight;
    window.scroll(0, bottom);
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
  let date = Date.now();
  let postdt = new Date(date);

  let cookieValue = document.cookie;
  let cookie = cookieValue.replace("cookieName=","");
  const newchat = new Newchat(roomid, userid, roomname, member, chat, postdt, cookie);
  const newchatJSON = JSON.stringify(newchat);
  socket.send(newchatJSON);

  const urlForApiPost = "http://172.25.0.3:8000/chatroom/chat";
  const axiosConfig = {
    headers: {
      "Authorization": "apikey",
    }
  };

  axios.post(urlForApiPost, newchatJSON, axiosConfig).then(response => {
    console.log('body:', response.data); 
});
  
  document.chatform.reset();
  console.log(JSON.stringify(newchat));
};