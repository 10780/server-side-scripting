/*day 6 of week 5 assignment
* 14.02.2107
* began work 8:48
*/ 

pacakge main

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
http.ListenAndServe(":8080", nil)
}

func homepage(w, httpResponseWriter, r *httpRequest){
tpl.ExecuteTemplate(w, `home.gohtml`, nil)
}
