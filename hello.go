package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	//obtenerExtension()
	//almacenarNombres()
	//imprimirHost()
	//convertir64()
	htmlPage()
}

func obtenerExtension() {
	archivos, err := ioutil.ReadDir("src/assets/img")

	if err != nil {
		log.Fatal(err)
	}

	for _, archivo := range archivos {
		if archivo.Name()[len(archivo.Name())-3:len(archivo.Name())] == "jpg" || archivo.Name()[len(archivo.Name())-4:len(archivo.Name())] == "jpeg" || archivo.Name()[len(archivo.Name())-3:len(archivo.Name())] == "png" {
			fmt.Println(archivo.Name())
		}
	}
}

func almacenarNombres() {
	archivos, err := ioutil.ReadDir("src/assets/img")
	imagenesArray := []string{}

	if err != nil {
		log.Fatal(err)
	}

	for _, archivo := range archivos {
		if archivo.Name()[len(archivo.Name())-3:len(archivo.Name())] == "jpg" || archivo.Name()[len(archivo.Name())-4:len(archivo.Name())] == "jpeg" || archivo.Name()[len(archivo.Name())-3:len(archivo.Name())] == "png" {
			imagenesArray = append(imagenesArray, archivo.Name())
		}

	}
	fmt.Println(len(imagenesArray))
}

func imprimirHost() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s", hostname)
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func convertir64() string {
	archivos, err := ioutil.ReadDir("src/assets/img")
	//contador
	imagenes := []string{}
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {

		if archivo.Name()[len(archivo.Name())-3:len(archivo.Name())] == "jpg" || archivo.Name()[len(archivo.Name())-4:len(archivo.Name())] == "jpeg" || archivo.Name()[len(archivo.Name())-3:len(archivo.Name())] == "png" {
			//fmt.Println(archivo.Name())
			imagenes = append(imagenes, archivo.Name())
		}

	}

	fmt.Println("cantidad de imagenes del array: ", len(imagenes))

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(imagenes) - 1
	aux := rand.Intn(max-min+1) + min

	fmt.Println("La imagen aleatoria elegida es:  ", imagenes[aux])

	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile("src/assets/img/" + imagenes[aux])
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpg":
		base64Encoding += "data:image/jpg;base64,"
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)

	return base64Encoding
}

func htmlPage() {
	fs := http.FileServer(http.Dir("./src/app"))
	http.Handle("/", fs)

	log.Print("Abierto el puerto 4 lukitas")

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
