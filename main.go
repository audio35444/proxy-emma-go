package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"./loger"
	"./mwHtml"
	"github.com/karlseguin/ccache"
)

var cache = ccache.New(ccache.Configure()) // solo se esta probando la libreria ccache
func handleCache(url string) {
	item := cache.Get(url)
	if item != nil {
		// fmt.Println("-------------------- Ya estuviste en la pagina:", item.Value(), "--------------------")
		fmt.Println("--------- URL REPETIDA ---------")
	}
	cache.Set(url, url, time.Minute*10)
}
func handleHTTP(w http.ResponseWriter, req *http.Request) {
	timeReq := time.Now()
	url := req.URL.String()
	t := time.Now()
	go loger.InsertNewLog("./log-file", url)
	go handleCache(url)
	// fmt.Println("Runtime InsertNewLog:", time.Since(t))
	insertNewLogTime := time.Since(t)
	resp, err := http.Get(url) //http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	// t = time.Now()
	go copyHeader(w.Header(), resp.Header)
	// fmt.Println("Runtime copyHeader:", time.Since(t))
	data, error := ioutil.ReadAll(resp.Body)
	if error == nil {
		dataString := string(data)
		go loger.WriteLogFiles(url, dataString) //guarda los archivos de log, si no sequiere loguear los archivos se quita esto
		mwHtml.AddSpamInfo(&dataString)         //agrega basura al dataString
		w.Write([]byte(dataString))
	}
	fmt.Println("Page:", url, " | RunTimes >> Request:", time.Since(timeReq), " InsertLog:", insertNewLogTime)
}
func copyHeader(wHeader, respHeader http.Header) {
	for k, vv := range respHeader {
		for _, v := range vv {
			wHeader.Add(k, v)
		}
	}
}
func main() {
	var port string
	if len(os.Args) >= 2 {
		port = ":" + os.Args[1]
	} else {
		port = ":8000"
	}
	http.HandleFunc("/", handleHTTP)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
