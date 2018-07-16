package main

import (
	"net/http"
	. "beego-session"
	"fmt"
)

func main() {
	redisConn := "172.18.254.12:6379"
	Session, _ := NewManager("redis", &ManagerConfig{CookieName:"minipcapisid",Gclifetime:30*24*3600,EnableSetCookie:true,ProviderConfig:redisConn, CookieLifeTime:0})
	defer Session.GC()

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		s, _ := Session.SessionStart(w, r)
		defer s.SessionRelease(w)
		s.Set("test_session_redis", 1)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		s, _ := Session.SessionStart(w, r)
		defer s.SessionRelease(w)
		fmt.Println(s.Get("test_session_redis"))
	})
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
