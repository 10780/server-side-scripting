/*day 2 of week 5 assignment
* Friday 10.02.2017
* began work 10:12
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
http:ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, r *httpRequest){
tpl.ExecuteTemplate(w, `home.gohtml`, nil)
}
