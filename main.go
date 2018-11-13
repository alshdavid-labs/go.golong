package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main() {
	fmt.Println("lol")
	
	mySigningKey := []byte("secret")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	
	claims["admin"] = true
	claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	
	tokenString, _ := token.SignedString(mySigningKey)
	
	
	fmt.Println(tokenString)

}
