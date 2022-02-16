package main

// import (
// 	f "fmt"
// 	"os"
// )

// func main() {

// 	path := "app.html"

// 	var _, err = os.Stat(path)

// 	if err != nil {
// 		var file, err = os.Create(path)
// 		if err != nil {
// 			return
// 		}
// 		defer file.Close()
// 		f.Println("file closed")
// 	} else {
// 		os.Remove(path)
// 		var file, err = os.Create(path)
// 		if err != nil {
// 			return
// 		}
// 		defer file.Close()
// 		f.Println("file closed")
// 	}
// 	fileToWrite, err := os.OpenFile(path, os.O_RDWR, 0644)
// 	if err != nil {
// 		return
// 	}
// 	defer fileToWrite.Close()

// 	_, err = fileToWrite.WriteString(`<!DOCTYPE html>
// 	<html lang="en">
// 	<head>
// 		<meta charset="UTF-8">
// 		<meta http-equiv="X-UA-Compatible" content="IE=edge">
// 		<meta name="viewport" content="width=device-width, initial-scale=1.0">
// 		<title>Document</title>
// 	</head>
// 	<body>
// 		<p>hi, number one</p>
// 		<p>hi, number one</p>
// 		<p>hi, number one</p>
// 		<p>hi, number one</p>
// 	</body>
// 	</html>`)

// 	err = fileToWrite.Sync()
// 	if err != nil {
// 		return
// 	}
// 	f.Println("done")

// 	fileC, err := os.OpenFile(path, os.O_RDWR, 0644)
// 	options := make([]byte, 1024)
// 	data, err := fileC.Read(options)

// 	if err != nil {
// 		return
// 	}

// 	f.Println(string(rune(data)))
// }
