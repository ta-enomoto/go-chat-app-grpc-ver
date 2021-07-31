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
		//fmt.Println(requestedUserId)

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

		if r.FormValue("change-userId") == "変更" {

			newUserId := r.FormValue("userId")
			if newUserId == "" {
				fmt.Fprintf(w, "ユーザーIDが入力されていません")
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
