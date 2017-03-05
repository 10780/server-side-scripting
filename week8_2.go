/*day 3 of week 8 assignment
*work started 6:16, 03.04.2017
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
http.Handle("/public/", http.StripPrefix("/public," http.FileServer(http.Dir("assets"))))
http.ListenAndServe(":8080", nil)
}

func index(w http.ReponseWriter, r *http.Request){
http.ExecuteTemplate(w, "index3.gohtml")
}