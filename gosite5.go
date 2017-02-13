/*day 5 of week 5 assignment
* 13.02.2017
* began work 9:03
*/

package main

import(
"net/http"
"html/template"
)

var tpl template.Template

func init{
tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
http.("/", homepage)
http.ListenAndServer(":8080", nil)
}

func homepage(w http.ResponseWrite, r *httpRequest){
tpl.ExecuteTemplate(w, `home.gohtml`, nil)
}
