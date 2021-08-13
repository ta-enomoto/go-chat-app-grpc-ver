//あるユーザーの個別管理ページにアクセスがあった時のハンドラ
package routers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"web-server/query"
	"web-server/sessions"
)

func AdminUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		url := r.URL.Path
		requestedUserId := strings.TrimPrefix(url, "/admin/users/")

		dbUsr, err := sql.Open("mysql", query.ConStrUsr)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer dbUsr.Close()

		user := query.SelectUserById(requestedUserId, dbUsr)
		ResponseUserId := user.UserId

		t := template.Must(template.ParseFiles("./templates/admin/adminuser/adminuser.html"))
		t.ExecuteTemplate(w, "adminuser.html", ResponseUserId)

	case "POST":
		//当該ユーザーに対して、ユーザーIDの変更・パスワードの変更・ユーザーの削除処理があった場合の処理

		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		//以下では、POSTされたフォームデータによって、if文で処理を分岐する
		//ユーザーIDを変更…change-userId: 変更
		//パスワードを変更…change-password: 変更
		//ユーザーを削除…delete-user: このユーザーを削除する
		if r.FormValue("change-userId") == "変更" {

			newUserId := r.FormValue("userId")
			if newUserId == "" {
				fmt.Fprintf(w, "ユーザーIDが入力されていません")
				return
			}

			escapeStrings := regexp.MustCompile(`\?|\$|\&|\=|\-|\>|\<|\+|\;|\:|\*|\||\'`)
			if escapeStrings.MatchString(newUserId) {
				fmt.Fprintf(w, "使用できない文字が含まれています。")
				return
			}

			url := r.URL.Path
			currentUserId := strings.TrimPrefix(url, "/admin/users/")

			dbUsr, err := sql.Open("mysql", query.ConStrUsr)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbUsr.Close()

			userIdChangedOrNot := query.ChangeUserId(newUserId, currentUserId, dbUsr)
			if userIdChangedOrNot {
				t := template.Must(template.ParseFiles("./templates/admin/adminuser/adminuseridchanged.html"))
				t.ExecuteTemplate(w, "adminuseridchanged.html", nil)
				return
			}
			return
		} else if r.FormValue("change-password") == "変更" {

			newPsw_string := r.FormValue("password")
			if newPsw_string == "" {
				fmt.Fprintf(w, "パスワードが入力されていません")
				return
			}

			escapeStrings := regexp.MustCompile(`\?|\$|\&|\=|\-|\>|\<|\+|\;|\:|\*|\||\'`)
			if escapeStrings.MatchString(newPsw_string) {
				fmt.Fprintf(w, "使用できない文字が含まれています。")
				return
			}

			newPsw_byte := []byte(newPsw_string)

			url := r.URL.Path
			requestedUserId := strings.TrimPrefix(url, "/admin/users/")

			dbUsr, err := sql.Open("mysql", query.ConStrUsr)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbUsr.Close()

			user := query.SelectUserById(requestedUserId, dbUsr)
			currentPassword := user.Password

			//パスワードのハッシュ化はChangePassword関数内で行う
			passwordChangedOrNot := query.ChangePassword(newPsw_byte, currentPassword, dbUsr)
			if passwordChangedOrNot {
				t := template.Must(template.ParseFiles("./templates/admin/adminuser/adminpasswordchanged.html"))
				t.ExecuteTemplate(w, "adminpasswordchanged.html", nil)
				return
			}
			return
		} else if r.FormValue("delete-user") == "このユーザーを削除する" {

			url := r.URL.Path
			requestedUserId := strings.TrimPrefix(url, "/admin/users/")

			dbUsr, err := sql.Open("mysql", query.ConStrUsr)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbUsr.Close()

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
