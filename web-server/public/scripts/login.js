const loginFunc = document.getElementById("login");

//ログインボタンクリック時の処理
loginFunc.addEventListener("click", function(){
  let userId = document.getElementById("inputedUserId").value;
  let password = document.getElementById("inputedPassword").value;

  //フォームが空欄の場合はsubmitせず、alertで通知
  if (userId == "" && password == "") {
    window.alert("ユーザーID・パスワードが入力されていません");
    return;
  } else if (userId == "") {
    window.alert("ユーザーIDが入力されていません");
    return;
  } else if (password == "") {
    window.alert("パスワードが入力されていません");
    return;
  } else {
    this.form.submit();
  };
});