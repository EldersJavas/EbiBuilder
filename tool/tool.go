package tool

import (
	"log"
	"os"
	"strings"

)

func IsEbitenGame() bool {
	var Versions string
	f,err:=os.ReadFile("./go.mod")
	if err!=nil{
		log.Println("No Ebitengine project in the dir.")
		return false
	}
	for _, v := range strings.Split(string(f),"\n") {
		if strings.Contains(v,"github.com/hajimehoshi/ebiten/v2"){
			Versions=v[strings.LastIndex(v,"v2")+1:]		
		}
	}
	log.Printf("str=%v,\n",Versions)
	return false
}
