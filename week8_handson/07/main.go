package main

import(
"net/http"
"html/template"
"log"
)

var tpl *template.Template

func init(){
tpl = template.Must(template.ParseGlob("templates/index.gohtml"))
}

func main(){
http.HandleFunc("/", dogs)
http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public")
http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, r *http.Request){
err := tpl.Execute(w, nil)
if err != nil{
log.Fatalln("template did not execute: ", err)
}
}