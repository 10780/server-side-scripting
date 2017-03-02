//all the pages use the same picture
//copypasted first 2 HandleFuncs and made changes to make things work
package main

import(
"net/http"
"io"
)

func main(){
http.HandleFunc("/", pic)
http.HandleFunc("assets/dog.jpg", index)
http.HandleFunc("/pic2", pic2)
http.HandleFunc("assets/dog2.jpg", pic22)
http.HandleFunc("/pic3", pic3)
http.HandleFunc("assets/dog3.jpg", pic33)
http.HandleFunc("/pi4", pic4)
http.HandleFunc("assets/dog4.jpg", pic44)
http.HandleFunc("/pic5", pic5)
http.HandleFunc("assets/dog5.jpg", pic55)
http.Handle("/pictures/", http.StripPrefix("/pictures", http.FileServer(http.Dir("./assets"))))
http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "text/hmtl; charset=UTF-8")
io.WriteString(w, `<img src="dog.jpg">`)
}

func pic(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "assets/dog.jpg")
}

func pic22(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "text/hmtl; charset=UTF-8")
io.WriteString(w, `<img src="dog2.jpg">`)
}

func pic2(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "assets/dog2.jpg")
}

func pic33(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "text/hmtl; charset=UTF-8")
io.WriteString(w, `<img src="dog3.jpg">`)
}

func pic3(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "assets/dog3.jpg")
}

func pic44(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "text/hmtl; charset=UTF-8")
io.WriteString(w, `<img src="dog4.jpg">`)
}

func pic4(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "assets/dog4.jpg")
}

func pic55(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "text/hmtl; charset=UTF-8")
io.WriteString(w, `<img src="dog5.jpg">`)
}

func pic5(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w, r, "assets/dog5.jpg")
}
