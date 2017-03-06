/*day 4 of week 8
*work started 7:06, 03.05.2017
*/
package main

import(
"net/http"
"html/template"
)

var tpl *template.Template

func init(){
tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
http.HandleFunc("/", index)
http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("assets"))))
http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
http.ExecuteTemplate(w, "index4.gohtml")
}