package main

import (
"net/http"
"html/template"
"fmt"
"github.com/satori/go.uuid"
)

type user struct{
Id string
First string
Last string
Email string
Password string
}

var tpl *template.Template

var mUser = map[string]user{}
var mSession = map[string]string{}

func init(){
tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
http.HandleFunc("/", index)
http.HandleFunc("/other", other)
http.HandleFunc("/login", login)
http.HandleFunc("/logout", logout)
http.HandleFunc("/signup", signup)
http.Handle("/favicon.ico", http.NotFoundHandler)
http.ListenAndServe(":8080",  nil)
}

func index(w http.ResponseWriter, r *http.Request){
c, err := r.Cookie("session")
if err != nil {
u := uuid.NewV4()
c = &http.Cookie{
Name: "session",
Value: u.String(),
Path: "/",
HttpOnly: true,
}
}

uid := mSession[c.Value]
usr := mUser[uid]

fmt.Println("SESSION ID", c.Value)

http.SetCookie(w, c)
tpl.ExecuteTemplate(w, "index.gohtml", usr)
}

func other(w http.ResponseWriter, r *http.Request){
c, err := r.Cookie("session")
if err != nil {
u := uuid.NewV4()
c = &http.Cookie{
Name: "session",
Value: u.String(),
Path: "/",
HttpOnly: true,
}
}
uid := mSession[c.Value]
usr := mUser[uid]

fmt.Println("SESSION ID", c.Value)

http.SetCookie(w, c)
tpl.ExecuteTemplate(w, "other.gohtml", usr)
}

func login(w http.ResponseWriter, r *http.Request){
c, err := r.Cookie("session")
if err != nil {
http.Redirect(w, r, "/", http.StatusSeeOther)
return
}

if r.Method == "POST"{
em := r.FormValue("email")
pw := r.FormValue("password")
m := map[string]user{}
for _, v := range mUser {
m[v.Email] = v
}
usr, ok := m[em]
if !ok {
http.Redirect(w, r, "/login", http.StatusForbidden)
return
} 
if usr.Password != pw {
http.Redirect(w, r, "/login", http.StatusForbidden)
return
}
mSession[c.Value] = usr.Id
http.Redirect(w, r, "/", http.StatusSeeOther)
return
}

tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request){
c, err := r.Cookie("session")
if err != nil {
http.Redirect(w, r, "/", http.StatusSeeOther)
return
}
delete(mSession, c.Value)
c.MaxAge = -1
http.SetCookie(w, c)
http.Redirect(w, r, "/", http.StatusSeeOther)
}

func signup(w http.ResponseWriter, r *http.Request){
c, err := r.Cookie("session")
if err != nil {
http.Redirect(w, r, "/", http.StatusSeeOther)
return
}

uid := mSession[c.Value]
usr := mUser[uid]
if usr.Email != "" {
http.Redirect(w, r, "/", http.StatusSeeOther)
return
}

if r.Method == "POST" {
fn := r.FormValue("first")
ln := r.FormValue("last")
em := r.FormValue("email")
pw := r.FormValue("password")
u := uuid.NewV4()

usr = user {
Id: u.String(),
First: fn,
Last: ln,
Email: em,
Password: pw,
}
mUser[usr.Id] = usr
mSession[c.Value] = usr.Id
http.Redirect(w, r, "/", http.StatusSeeOther)
return
}

tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}