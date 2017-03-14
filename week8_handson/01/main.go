pacakge main

import(
"net/http"
"html/template"
"fmt"
"io"
)

var tpl *template.Template

func init(){
tpl = template.Must(template.ParseGlob(*.gohtml))
}

func main(){
http.HandleFunc("/", foo)
http.HandleFunc("/dog/", dog)
http.HandleFunc("/dog.jpg", dogpic)
http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request){
io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request){
tpl, err := template.ParseFiles("dog.gohtml")
if err != nil{
log.Fatalln(err)
}
tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func dogpic(w http.ResponseWriter, r *http.Request){
http.ServeFile(w, r, "dog.jpg")
}