package routers

import (
	"database/sql"
	"fmt"
	"goserver/query"
	"goserver/sessions"
	"html/template"
	"net/http"
	"net/url"
	"strings"
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

		t := template.Must(template.ParseFiles("./templates/admin/users/user.html"))
		t.ExecuteTemplate(w, "user.html", ResponseUserId)

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
				fmt.Println("成功")
				return
			}
			return
			//t := template.Must(template.ParseFiles("./templates/admin/users/useridchanged.html"))
			//t.ExecuteTemplate(w, "useridchanged.html", nil)
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
				fmt.Println("成功")
				return
			}
			return

			//t := template.Must(template.ParseFiles("./templates/admin/users/passwordchanged.html"))
			//t.ExecuteTemplate(w, "passwordchanged.html", nil)
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
				fmt.Println("削除成功")
				return
			}
			return
		}
	}
}
