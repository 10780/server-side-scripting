package main

import (
"net/http"
"fmt"
)

func main(){
http.HandleFunc("/", set)
http.HandleFunc("/dog/bowzer", dog_bowzer)
http.HandleFunc("/cat", cat)
http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request){
http.SetCookie(w, &http.Cookie {
Name: "site-cookie",
Value: "value", 
})
fmt.Fprintln(w, "COOKIE SET")
}

func dog_bowzer(w http.ResponseWriter, r *http.Request){
c, err := r.Cookie("site-cookie")
if err != nil {
http.Error(w, http.StatusText(400), http.StatusBadRequest)
return
}
fmt.Fprintln(w, "COOKIE: ", c)
}

func cat(w http.ResponseWriter, r *http.Request){
c, err := r.Cookie("site-cookie")
if err != nil {
http.Error(w, http.StatusText(400), http.StatusBadRequest)
return
}
fmt.Fprintln(w, "COOKIE: ", c)
}