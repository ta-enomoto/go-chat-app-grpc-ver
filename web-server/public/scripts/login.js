//ログインボタンクリック時の処理
function loginFunc(){
  let userId = document.getElementById("inputedUserId").value;
  let password = document.getElementById("inputedPassword").value;

  //フォームが空欄の場合はsubmitせず、alertで通知
  if (userId == "" && password == "") {
    window.alert("ユーザーID・パスワードが入力されていません");
    return false;
  } else if (userId == "") {
    window.alert("ユーザーIDが入力されていません");
    return false;
  } else if (password == "") {
    window.alert("パスワードが入力されていません");
    return false;
  } else {
    this.form.submit();
  };
};