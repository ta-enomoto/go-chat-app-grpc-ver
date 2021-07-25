let allchats =""
window.onload = async function() {
  
  let url = location.href;
  let roomid = url.replace("http://172.25.0.2/admin/chatrooms/chatroom","");

  const urlForApi = "http://172.25.0.4:8000/chatroom/" + roomid
  
  //チャット読み込み
  async function getChatFromApi() {
    try {
       res = await axios.get(urlForApi, { headers: { Authorization: "apikey" } });
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
  
  let element = document.documentElement;
  let bottom = element.scrollHeight - element.clientHeight;
  window.scroll(0, bottom);
};