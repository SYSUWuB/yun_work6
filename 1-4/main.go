package main
import (
"log"
"net/http"
"os"
"strconv"
)
func main(){
	http.HandleFunc("/abc", index)
	err := http.ListenAndServe(":5565", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	num:=os.Getenv("Num")
	if num==""{
		_,err:=w.Write([]byte("there is not a number\n"))
		if err!=nil{
		log.Println("err:"+err.Error()+" No\n")
		}
	}else{
		numInt,_:=strconv.Atoi(num)
		even := is_even(numInt)
		if even==1{
			_,err:=w.Write([]byte("there is an even number\n"))
		} else{
			_,err:=w.Write([]byte("there is an odd number\n"))
		}
		if err!=nil{
			log.Println("err:"+err.Error()+" Yes\n")
		}
	}
}
func is_even(n int)int{
	if n%2==0{
		return 1
	} else{
		return 2
	}
}