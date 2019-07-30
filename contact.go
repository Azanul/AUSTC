package AUSTC_website

import(
    "net/http"
    "os"
    "encoding/json"
)

type Data struct {
Name string
Email string
Dept string
Occupation string
message string
}

func save(w http.ResponseWriter, r *http.Request) {
name := r.FormValue("first_name")
email := r.FormValue("email")
dept := r.FormValue("dept")
occup := r.FormValue("occupation")
msg := r.FormValue("msg")

data := &Data{name, email, dept, occup, msg}

b, err := json.Marshal(data)
if err != nil {
http.Error(w, err.Error(), 500)
return
}

f, err := os.Open("somefile.json")
if err != nil {
http.Error(w, err.Error(), 500)
return
}

f.Write(b)
f.Close()
}

func init() {
http.HandleFunc("/contact", save)
}