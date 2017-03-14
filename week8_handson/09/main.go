package main

import(

)

var tpl *template.Template

func init(){
tpl = template.Must(template.ParseGlob("templates/index.gohtml"))
}

func main(){
http.HandleFunc("/", index)
http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
}

func index(w http.ResponseWriter, _ *http.Request){
tpl.ExecuteTemplate(w, "index.gohtml", nil)
}