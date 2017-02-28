/*day 1 of week 5 assignment
* Thursday 09.02.2017
* began work 12:23
*/

package main

import(
"net/http"
"html/template"
)

var tpl template.Template
/*there's a templates repository I made for all the .gohtml files for the assignment(s)-
https://github.com/circusex/server-side-scripting/tree/templates
*/

func init{
tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
http.HandleFunc("/", homepage)
http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, r *http.Request){
tpl.ExecuteTemplate(w, `home.gohtml`, nil)
}
