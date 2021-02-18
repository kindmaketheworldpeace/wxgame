package main
import (
    "fmt"
    "net/http"
)
func getRank(w http.ResponseWrite,r *http.Request ) {
    _,_=w.Write([]byte("test"))

}
func addRank(w http.ResponseWrite,r *http.Request ) {

}
func main() {

    defer system.MyLogger.File.Close()
    http.HandleFunc("/getRank", getRank)
    http.ListenAndServe(":"+strconv.Itoa(conf.MyConfig.System.Port),nil)
}