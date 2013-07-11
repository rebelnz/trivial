package main

import (
	"net/http"
	"html/template"
	"io/ioutil"
)

func getWData(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent",`Trival/1.1 (http://enkoder.info; rebelbassnz@gmail.com)`)
	resp, _ := client.Do(req)		
	content,_ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()	
	return []byte(content)	
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w,nil)
}

func wDataHandler(w http.ResponseWriter, r *http.Request) {	
	url := "http://en.wikipedia.org/w/api.php?action=query&list=random&rnnamespace=0&rnlimit=5&format=json"
	c := getWData(url)
	w.Write(c)
}

func wDetailHandler(w http.ResponseWriter, r *http.Request) {	
	query := r.URL.RawQuery
	url := "http://en.wikipedia.org/w/api.php?action=query&prop=revisions&titles=" + query + "&rvprop=content&format=json"
	// url := "http://en.wikipedia.org/w/api.php?action=query&prop=revisions&titles=" + query + "&rvprop=content&rvparse=1&format=json"
	c := getWData(url)
	w.Write(c)
}

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/wdata",wDataHandler)
	http.HandleFunc("/wdetail",wDetailHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":9000",nil)
}
