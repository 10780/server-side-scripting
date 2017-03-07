/*day 6 of week 8 assignment
*started work 12:20
*/

package main

import(
"net/http"
"html/template"
)

var tpl *template.Template

func init(){
tpl = template.Must(template.ParseGlob("templates/*.gohtml")
}

func main(){
http.HandleFunc("/", index)
http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("assets"))))
http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
http.ExecuteTemplate(w, "index6.gohtml")
}