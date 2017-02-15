/*day 7 of week assignment
* Wednesday 15.02.2017
* work began 12:30
*/

package main

import(
"net/http"
"html/template"
)

var tpl template.Template

func init(){
tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
http.("/", homepage)
http.ListenAndServe(":8080", nil)
}

func homepage(w httpResponseWriter, r *httpRequest){
tpl.ExecuteTemplate(w, `home.gothml`, nil)
}
