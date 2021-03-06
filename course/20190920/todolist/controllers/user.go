package controllers

import (
	"net/http"
	"strings"
	"todolist/models"
	"todolist/session"
	"todolist/utils"
)

func UserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	q := strings.TrimSpace(r.FormValue("q"))

	utils.Render(w, []string{"views/layouts/base.html", "views/user/user.html"}, "base.html", struct {
		Query string
		Users []models.User
	}{q, models.GetUsers(q)})

}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.Render(w, []string{"views/user/login.html"}, "login.html", nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")

		user, err := models.GetUserByName(name)
		if err != nil || !user.ValidatePassword(password) {
			utils.Render(w, []string{"views/user/login.html"}, "login.html", struct {
				Name  string
				Error string
			}{name, "用户名或密码错误"})
		} else {
			sessionObj := session.DefaultManager.SessionStart(w, r)
			//登陆成功
			sessionObj.Set("user", user)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func LogoutAction(w http.ResponseWriter, r *http.Request) {
	session.DefaultManager.SessionDestory(w, r)
	http.Redirect(w, r, "/user/login/", http.StatusFound)
}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	var context interface{}
	if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		birthday := r.PostFormValue("birthday")
		tel := r.PostFormValue("tel")
		desc := r.PostFormValue("desc")
		addr := r.PostFormValue("addr")

		errors := models.ValidateCreateUser(name, password, birthday, tel, addr, desc)
		if len(errors) == 0 {
			models.CreateUser(name, password, birthday, tel, addr, desc)
			http.Redirect(w, r, "/users/", http.StatusFound)
			return
		}

		context = struct {
			Errors   map[string]string
			Name     string
			Password string
			Birthday string
			Tel      string
			Addr     string
			Desc     string
		}{errors, name, password, birthday, tel, addr, desc}
	}

	utils.Render(w, []string{"views/layouts/base.html", "views/user/create.html"}, "base.html", context)
}

func init() {
	http.HandleFunc("/users/", UserAction)
	http.HandleFunc("/users/create/", UserCreateAction)
	http.HandleFunc("/user/login/", LoginAction)
	http.HandleFunc("/user/logout/", LogoutAction)
}
