//あるユーザーの個別管理ページにアクセスがあった時のハンドラ
package routers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"web-server/query"
	"web-server/sessions"
)

func AdminUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		//ユーザーのcookieからセッション変数(ユーザーID)を取得する
		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		//ユーザーIDがadmin以外のユーザーのアクセスを拒否する
		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		//リクエスト元のURLの文字列から、ユーザーIDの文字列を取得する
		url := r.URL.Path
		requestedUserId := strings.TrimPrefix(url, "/admin/users/")

		//ユーザーDBに接続する
		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		//URLから取得した文字列を元にデータベースからユーザーIDを取得する
		user := query.SelectUserById(requestedUserId, dbUsr)
		ResponseUserId := user.UserId

		t := template.Must(template.ParseFiles("./templates/admin/adminuser/adminuser.html"))
		t.ExecuteTemplate(w, "adminuser.html", ResponseUserId)

	case "POST":
		//当該ユーザーに対して、ユーザーIDの変更・パスワードの変更・ユーザーの削除処理があった場合の処理

		//ユーザーのcookieを元に有効なセッションが存在しているかチェックする
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		//ユーザーのcookieからセッション変数(ユーザーID)を取得する
		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		//ユーザーIDがadmin以外のユーザーのアクセスを拒否する
		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		//以下では、POSTされたフォームデータによって、if文で処理を分岐する
		//ユーザーIDを変更…change-userId: 変更
		//パスワードを変更…change-password: 変更
		//ユーザーを削除…delete-user: このユーザーを削除する
		if r.FormValue("change-userId") == "変更" {
			//ユーザーIDの変更処理

			//フォームに入力された値を取得する
			newUserId := r.FormValue("userId")
			if newUserId == "" {
				fmt.Fprintf(w, "ユーザーIDが入力されていません")
				return
			}

			//リクエスト元のURLの文字列から、ユーザーIDの文字列を取得する
			url := r.URL.Path
			currentUserId := strings.TrimPrefix(url, "/admin/users/")

			//ユーザーDBに接続する
			dbUsr, err := sql.Open("mysql", query.ConStrUsr)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbUsr.Close()

			//ユーザーIDの変更が成功したら、変更完了ページを表示する
			userIdChangedOrNot := query.ChangeUserId(newUserId, currentUserId, dbUsr)
			if userIdChangedOrNot {
				t := template.Must(template.ParseFiles("./templates/admin/adminuser/adminuseridchanged.html"))
				t.ExecuteTemplate(w, "adminuseridchanged.html", nil)
				return
			}
			return
		} else if r.FormValue("change-password") == "変更" {
			//パスワードの変更処理

			//フォームに入力された値を取得する
			newPsw_string := r.FormValue("password")
			if newPsw_string == "" {
				fmt.Fprintf(w, "パスワードが入力されていません")
				return
			}

			//ハッシュ化のため、パスワード文字列をbyte型に変換する(ハッシュ化はChangePassword関数内で行う)
			newPsw_byte := []byte(newPsw_string)

			//リクエスト元のURLの文字列から、ユーザーIDの文字列を取得する
			url := r.URL.Path
			requestedUserId := strings.TrimPrefix(url, "/admin/users/")

			//ユーザーDBに接続する
			dbUsr, err := sql.Open("mysql", query.ConStrUsr)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbUsr.Close()

			//URLから取得した文字列を元にデータベースからパスワードを取得する
			user := query.SelectUserById(requestedUserId, dbUsr)
			currentPassword := user.Password

			//パスワードのハッシュ化・変更が成功したら、変更完了ページを表示する
			passwordChangedOrNot := query.ChangePassword(newPsw_byte, currentPassword, dbUsr)
			if passwordChangedOrNot {
				t := template.Must(template.ParseFiles("./templates/admin/adminuser/adminpasswordchanged.html"))
				t.ExecuteTemplate(w, "adminpasswordchanged.html", nil)
				return
			}
			return
		} else if r.FormValue("delete-user") == "このユーザーを削除する" {
			//ユーザーの削除処理

			//リクエスト元のURLの文字列から、ユーザーIDの文字列を取得する
			url := r.URL.Path
			requestedUserId := strings.TrimPrefix(url, "/admin/users/")

			//ユーザーDBに接続する
			dbUsr, err := sql.Open("mysql", query.ConStrUsr)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbUsr.Close()

			//ユーザーの削除が成功したら、削除完了ページを表示する
			userDeletedFromDb := query.DeleteUserById(requestedUserId, dbUsr)
			if userDeletedFromDb {
				t := template.Must(template.ParseFiles("./templates/admin/adminuser/adminuserdeleted.html"))
				t.ExecuteTemplate(w, "adminuserdeleted.html", nil)
				return
			}
			return
		}
	}
}
