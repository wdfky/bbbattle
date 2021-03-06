package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type AutoGenerated struct {
	Code     int        `json:"code"`
	Msg      string     `json:"msg"`
	Newslist []Newslist `json:"newslist"`
}
type Newslist struct {
	Naming  string `json:"naming"`
	Wordnum int    `json:"wordnum"`
	Sex     int    `json:"sex"`
}
type Msg struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

// func str2bytes(str string) []byte {
// 	x := (*[2]uintptr)(unsafe.Pointer(&str))
// 	h := [3]uintptr{x[0], x[1], x[1]}
// 	return *(*[]byte)(unsafe.Pointer(&h))
// }
// func str2uint64(str string) (num uint64) {
// 	num = 0
// 	level := uint64(10)
// 	bytes := str2bytes(str)
// 	for i := len(bytes) - 1; i >= 0; i-- {
// 		num += uint64(bytes[i]-'0') * level
// 		level *= 10
// 	}
// 	return
// }
func defaultHandle(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	answer := `{"status":"OK","message":"请先登录"}`
	rw.Write([]byte(answer))
}
func postLoginHandle(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//fmt.Printf("r.PostForm.Get(\"password\"): %v\n", r.PostForm.Get("password"))
	fmt.Printf("r.RemoteAddr: %v\n", r.RemoteAddr)
	// if strings.Split(r.RemoteAddr, ":")[0] != "127.0.0.1" {
	// 	return
	// }
	r.ParseForm()
	//fmt.Println("到这里")
	//fmt.Println("body", r.Body)
	//fmt.Println(r.Header)
	userid := r.PostForm.Get("account")
	password := r.PostForm.Get("password")
	//fmt.Println(userid)
	//fmt.Println(password)
	if ok := redismgr.CheckId(userid, password); !ok {
		rw.Write([]byte(`{"message":"请检查账号密码", "name":""}`))
	} else {
		fmt.Println("密码正确")
		name, err := rdb.HGet(ctx, "name", userid).Result()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(name, 123)
		//rw.Write([]byte(`{"message":"登录成功"}`))
		msg := &Msg{
			Message: "1",
			Name:    name,
		}
		data, err := json.Marshal(msg)
		if err != nil {
			log.Println(err)
			rw.Write([]byte(`{"message":"1", "name":""}`))
			return
		}
		rw.Write(data)
	}
	//fmt.Println(userid, password)
	//rw.Write([]byte(userid + password))
}
func postRegHandle(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	userid := r.PostForm.Get("account")
	password := r.PostForm.Get("password")
	if ok := redismgr.RegId(userid, password); !ok {
		rw.Write([]byte(`{"message":"注册失败,请尝试别的用户名"}`))
	} else {
		rw.Write([]byte(`{"message":"注册成功"}`))
		redismgr.SetName(userid, GetRandName())
	}
}
func GetRandName() string {
	resp, err := http.Get("http://api.tianapi.com/txapi/cname/index?key=5192177fc77c82beb272933ad00f293e")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return "111"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return ""
	}
	data := &AutoGenerated{}
	err = json.Unmarshal(body, data)
	if err != nil {
		log.Fatal(err)
	}
	return data.Newslist[0].Naming
}

// func tcpDial(roomid string, token string, addr string) {
// 	conn, err := net.Dial("tcp", ":1235")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	defer conn.Close()
// 	room := Room{
// 		RoomId: roomid,
// 		Token:  token,
// 		Addr:   addr,
// 	}
// 	b, err := json.Marshal(room)
// 	if err != nil {
// 		log.Println("marshal failed, error:", err)
// 		return
// 	}
// 	msg, err := Encode(string(b))
// 	if err != nil {
// 		log.Println("encode error:", err)
// 		return
// 	}
// 	_, err = conn.Write(msg)
// 	if err != nil {
// 		log.Println("conn write failed error:", err)
// 		return
// 	}
// }
func postSetNameHandle(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	//之后改成cookie之类的
	userid := r.PostForm.Get("userid")
	name := r.PostForm.Get("name")
	ok := redismgr.SetName(userid, name)
	if !ok {
		rw.Write([]byte(`{"message":"false"}`))
	} else {
		rw.Write([]byte(`{"message":"true"}`))
	}
}
func postRoomHandle(rw http.ResponseWriter, r *http.Request) {
	//fmt.Println("daozhele")
	defer r.Body.Close()
	r.ParseForm()
	userid := r.PostForm.Get("userid")
	name := r.PostForm.Get("name")
	//fmt.Println(userid, name)
	roomid, token, addr, err := RpcClient(userid, name)
	if err != nil {
		log.Println(err.Error())
		return
	}
	ip := strings.Split(addr, ":")

	room := &Room{
		RoomId: roomid,
		Token:  token,
		Host:   ip[0],
		Port:   ip[1],
	}
	data, err := json.Marshal(room)
	if err != nil {
		log.Println(err)
	}
	rw.Write(data)
	//fmt.Println(roomid, token, addr)
	//go tcpDial(roomid, token, addr)
}
func main() {
	redismgr.Init()
	//fmt.Println("到这里")
	http.HandleFunc("/", defaultHandle)
	http.HandleFunc("/login", postLoginHandle)
	http.HandleFunc("/register", postRegHandle)
	http.HandleFunc("/startmatch", postRoomHandle)
	http.HandleFunc("/setname", postSetNameHandle)
	go http.ListenAndServe(":8080", nil)
	select {}
}
