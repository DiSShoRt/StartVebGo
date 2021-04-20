package main

import (
	"fmt"
	"net/http"
	"html/template"
)
//структкра пользователь
type User struct {
	Name string
	Age uint32
	Mail string
	Phone string 
}
//метод обновления имени
func (u *User )  SetNewName(NewName string) {
u.Name = NewName
}
//метод обновления вораста 
func (u *User )  SetNewAge(NewAge uint32) {
	u.Age = NewAge
}

func Home(w http.ResponseWriter, r *http.Request) {
	Anton := User{"Anton",19,"anton@.ru","8-988-875-65-65"}
	// парсинг файла html
	t , _ := template.ParseFiles("home.html")
	t.Execute(w, Anton)
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"About me")
}
func HandelReqest() {
	// обработчика запросов 
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)
	// создание и прослушивание сервера 
	http.ListenAndServe(":8080",nil)
} 
func main() {
	HandelReqest()
}