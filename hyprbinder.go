package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// globals
var appDir string = "/.config/hyprbinder"
var appFile string = "/.config/hyprbinder/usr.json"
var homePath, _ = os.LookupEnv("HOME")

// initializing struct/path
var user = User{Path: homePath + "/.config/hypr/hyprland.conf"}

// for binds
var binds = make([]Bind, 1)

type User struct {
	Path string `json:"Path"`
}

type Bind struct {
	bind1 string
	bind2 string
	prog  string
	cmd   string
}

func main() {
	CheckingExist()
	LoadCfg()
	switch Menu() {
	case 0:
		Clear()
		log.Fatal("Oh no you closed the programm")
	case 1:
		Clear()
		Q_and_A()
	case 2:
		Clear()
		fmt.Println("developing weight")
	}
}

func CheckingExist() {
	var _, FileError = os.Stat(homePath + appFile)
	if os.IsNotExist(FileError) == true {
		fmt.Println("Whare is your bind config? Default is ~/.config/hypr/hyprland.conf")
		fmt.Scanln(&user.Path)
		_, err := os.Stat(user.Path)
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatal("File doesn't exist  \n")
			}
		}
		fmt.Printf("Exist \n")
	}
}

func LoadCfg() {
	jsn, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}
	_, error := os.Stat(homePath + appDir)

	_, FileError := os.Stat(homePath + appFile)

	if os.IsNotExist(error) == true {
		fmt.Printf("Creating directory for hyprbinder")
		os.Mkdir(homePath+appDir, 0777)
	}
	if os.IsNotExist(FileError) == true {
		os.Create(homePath + appFile)
		os.WriteFile(homePath+appFile, jsn, 0666)
	}
	cfg := ReadFile(homePath + appFile)
	json.Unmarshal(cfg, &user)
}

// ovo je menju, brate
func Menu() uint8 {
	var q uint8
	Clear()
	fmt.Println("***Menu***")
	fmt.Println("1:Bind")
	fmt.Println("2:Change bind")
	fmt.Println("0:Exit")
	fmt.Scanln(&q)
	return q
}

func Q_and_A() Bind {
	var udefbind Bind
	Clear()
	fmt.Println("What's the first button? ")
	fmt.Scanln(&udefbind.bind1)
	//
	Clear()
	fmt.Println("What's the second button? ")
	fmt.Scanln(&udefbind.bind2)
	//
	Clear()
	fmt.Println("command ")
	fmt.Scanln(&udefbind.cmd)
	//
	Clear()
	fmt.Println("programm ")
	fmt.Scanln(&udefbind.prog)
	return udefbind
}

func ReadFile(pth string) []byte {
	bts, err := os.ReadFile(pth)
	if err != nil {
		log.Fatal(err)
	}
	return bts
}

func Clear() {
	var cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Parse(str string, bind *Bind) error {

}
