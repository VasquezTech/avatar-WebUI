package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go-avatar/internal/api"
	"go-avatar/internal/defaults"

	"github.com/disintegration/imaging"
	"github.com/gorilla/mux"

	"github.com/nfnt/resize"
)

func init() {
	defaults.Init()
}

// Trait represents a customizable trait of a person
type Trait struct {
	Name   string
	Values []string
}

// PersonTraits represents the traits of a person
type PersonTraits struct {
	Name      string
	Values    []string
	Jacket    string
	Mouth     string
	Neck      string
	Pants     string
	Shoes     string
	Socks     string
	Top       string
	Underwear string
	Wings     string
	Body      string
	Cape      string
	CapeBack  string
	Eyes      string
	Glasses   string
	Gloves    string
	HairLower string
	Hair      string
	Hat       string
}

// Available traits
var traits = []Trait{
	{"Body", getFiles("Body")},
	{"Wings", getFiles("Wings")},
	{"CapeBack", getFiles("CapeBack")},
	{"Mouth", getFiles("Mouth")},
	{"Neck", getFiles("Neck")},
	{"Shoes", getFiles("Shoes")},
	{"Socks", getFiles("Socks")},
	{"Underwear", getFiles("Underwear")},
	{"Cape", getFiles("Cape")},
	{"HairLower", getFiles("HairLower")},
	{"Hair", getFiles("Hair")},
	{"Hat", getFiles("Hat")},
	{"Pants", getFiles("Pants")},
	{"Jacket", getFiles("Jacket")},
	{"Gloves", getFiles("Gloves")},
	{"Eyes", getFiles("Eyes")},
	{"Glasses", getFiles("Glasses")},
	{"Top", getFiles("Top")},
}

// GenerateRandomPersonTraits generates random traits for a person
func GenerateRandomPersonTraits() []string {
	// Seed the random number generator to ensure consistent results across runs
	rand.Seed(time.Now().UnixNano())

	// Create an empty slice to store selected trait values
	selectedTraits := []string{}

	// Shuffle the traits to ensure consistent order across runs
	shuffledTraits := make([]Trait, len(traits))
	copy(shuffledTraits, traits)
	rand.Shuffle(len(shuffledTraits), func(i, j int) {
		shuffledTraits[i], shuffledTraits[j] = shuffledTraits[j], shuffledTraits[i]
	})

	// Iterate over the shuffled traits
	for _, trait := range shuffledTraits {
		if trait.Name == "Hat" || trait.Name == "Hair" || trait.Name == "Eyes" {
			continue
		}
		// Randomly decide whether to include the trait
		includeTrait := rand.Intn(4) == 0 // 50% chance to include trait

		if includeTrait {
			// Select a random value for the trait and add it to the selectedTraits slice
			selectedTrait := trait.Values[rand.Intn(len(trait.Values))]
			selectedTraits = append(selectedTraits, selectedTrait)
		}
	}
	for _, trait := range traits {
		if trait.Name == "Het" || trait.Name == "Hair" || trait.Name == "Eyes" {
			includeTrait := rand.Intn(4) == 0 // 50% chance to include trait

			if includeTrait {
				// Add traits that should always be last
				selectedTraits = append(selectedTraits, trait.Values[rand.Intn(len(trait.Values))])
			}
		}
	}
	return selectedTraits
}

// DrawTrait draws a trait on the image
func DrawTrait(baseImg *image.RGBA, traitsvalues []string) {
	// Create a new RGBA image as canvas for merging traits onto it
	mergedImg := image.NewRGBA(baseImg.Bounds())

	// Copy base image onto merged image
	draw.Draw(mergedImg, baseImg.Bounds(), baseImg, image.Point{}, draw.Over)

	for i := range traitsvalues {
		traitPath := traitsvalues[i]

		traitImg, err := loadImage(traitPath)
		if err != nil {
			log.Printf("Error loading trait image: %v", err)
			continue
		}

		// Resize trait image if necessary
		traitImg = resize.Resize(256, 256, traitImg, resize.Bicubic)
		sharpenedImage := imaging.Sharpen(traitImg, 5)
		// Draw trait image onto merged image
		draw.Draw(mergedImg, traitImg.Bounds().Add(image.Pt(0, 0)), sharpenedImage, image.Point{}, draw.Over)
	}

	// Copy merged image onto the base image
	draw.Draw(baseImg, baseImg.Bounds(), mergedImg, image.Point{}, draw.Over)
}

// loadImage loads an image file from the specified path
func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// getFiles retrieves a list of files in the specified directory
func getFiles(dir string) []string {
	var files []string
	filepath.Walk("output/files/"+dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files
}

// GenerateAvatar generates an avatar of a person with random traits
func GenerateAvatar() (*bytes.Buffer, error) {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Generate random person traits
	traits := GenerateRandomPersonTraits()

	// Create a 256x256 pixel image
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))

	// Draw traits based on traits
	DrawTrait(img, traits)

	// Encode the image as a PNG buffer
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func SaveImage(buffer *bytes.Buffer, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// IndexHandler handles the root route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// AvatarHandler generates and serves the avatar image
func AvatarHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("avatar handle")

	// Generate avatar image

	out, err := GenerateAvatar()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Image generated!")
	// Serve the image
	w.Header().Set("Content-Type", "image/png")
	w.Write(out.Bytes())
}

func main() {
	r := mux.NewRouter()

	// Define routes

	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/avatar", AvatarHandler).Methods("GET")

	//r.HandleFunc("/avatar", AvatarHandler).Methods("POST")
	r.HandleFunc("/api/v1/avatar", api.GetAvatar)
	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	log.Println("Server started on http://localhost:8051")
	log.Fatal(http.ListenAndServe("localhost:8051", r))
}
