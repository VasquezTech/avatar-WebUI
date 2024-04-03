package avatar

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"go-avatar/internal/images"
	"image"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

// Trait represents a customizable trait of a person
type Trait struct {
	Name   string
	Values []string
}

func hash(s string) int64 {
	var h uint32
	for _, c := range s {
		h = 32*h + uint32(c)
	}
	return int64(h)
}
func GenerateRandomPersonTraits(seed string) []string {
	// Available traits
	var traits = []Trait{
		{"Body", getFiles("/Body")},
		{"Wings", getFiles("/Wings")},
		{"CapeBack", getFiles("/CapeBack")},
		{"Mouth", getFiles("/Mouth")},
		{"Neck", getFiles("/Neck")},
		{"Shoes", getFiles("/Shoes")},
		{"Socks", getFiles("/Socks")},
		{"Underwear", getFiles("/Underwear")},
		{"Cape", getFiles("/Cape")},
		{"HairLower", getFiles("/HairLower")},
		{"Hair", getFiles("/Hair")},
		{"Hat", getFiles("/Hat")},
		{"Pants", getFiles("/Pants")},
		{"Jacket", getFiles("/Jacket")},
		{"Gloves", getFiles("/Gloves")},
		{"Eyes", getFiles("/Eyes")},
		{"Glasses", getFiles("/Glasses")},
		{"Top", getFiles("/Top")},
	}
	rand.Seed(hash(seed))
	selectedTraits := []string{}

	shuffledTraits := make([]Trait, len(traits))

	copy(shuffledTraits, traits)
	rand.Shuffle(len(shuffledTraits), func(i, j int) {
		shuffledTraits[i], shuffledTraits[j] = shuffledTraits[j], shuffledTraits[i]
	})

	var selectedTrait string
	for i, trait := range traits {

		includeTrait := false
		if trait.Name == "Body" {
			count := 0
			for {
				includeTrait = rand.Intn(10) == 0
				if includeTrait {
					count++
					rnd := rand.Intn(len(traits[i].Values))
					selectedTrait = traits[i].Values[rnd]
					selectedTraits = append(selectedTraits, selectedTrait)

				}
				if count == 1 {
					break
				}
			}
		}
	}
	for i, trait := range shuffledTraits {
		if len(traits[i].Values) == 0 {
			fmt.Println("traits[i].Values is zero!\n", traits[i])

			images.Init()
			return []string{}
		}
		includeTrait := rand.Intn(2) == 0 // 33% chance to include trait
		if trait.Name == "Hat" ||
			trait.Name == "Hair" ||
			trait.Name == "CapeBack" ||
			trait.Name == "Wings" ||
			trait.Name == "Cape" ||
			trait.Name == "Body" ||
			trait.Name == "Eyes" {
			continue
		} else if includeTrait {
			rnd := rand.Intn(len(traits[i].Values))
			selectedTrait = traits[i].Values[rnd]
			selectedTraits = append(selectedTraits, selectedTrait)
		}
	}

	for i, trait := range traits {

		includeTrait := false
		if trait.Name == "Hat" ||
			trait.Name == "CapeBack" ||
			trait.Name == "Wings" ||
			trait.Name == "Cape" ||
			trait.Name == "Hair" {

			includeTrait = rand.Intn(10) == 0 // 50% chance to include trait

			if includeTrait {
				selectedTraits = append(selectedTraits, trait.Values[rand.Intn(len(trait.Values))])
			}
		} else if trait.Name == "Eyes" {
			count := 0
			for {
				includeTrait = rand.Intn(10) == 0
				if includeTrait {
					count++
					rnd := rand.Intn(len(traits[i].Values))
					selectedTrait = traits[i].Values[rnd]
					selectedTraits = append(selectedTraits, selectedTrait)

				}
				if count == 1 {
					break
				}
			}
		}
	}
	return selectedTraits
}
func DrawTrait(baseImg *image.RGBA, traitsvalues []string) {
	mergedImg := image.NewRGBA(baseImg.Bounds())

	draw.Draw(mergedImg, baseImg.Bounds(), baseImg, image.Point{}, draw.Over)

	for i := range traitsvalues {
		traitPath := traitsvalues[i]

		traitImg, err := loadImage(traitPath)
		if err != nil {
			log.Printf("Error loading trait image: %v", err)
			continue
		}

		traitImg = resize.Resize(256, 256, traitImg, resize.Bicubic)
		AdjustBrightness := imaging.AdjustBrightness(traitImg, 0)

		draw.Draw(mergedImg, traitImg.Bounds().Add(image.Pt(0, 0)), AdjustBrightness, image.Point{}, draw.Over)
	}

	draw.Draw(baseImg, baseImg.Bounds(), mergedImg, image.Point{}, draw.Over)
}

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
	fmt.Println("Getting traits..")
	filepath.Walk("./files"+dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsNotExist(err) {
				err := os.MkdirAll("./files/"+dir, 0755)
				if err != nil {
					fmt.Println(err)
					getFiles(dir)
					fmt.Println(err)
					return err
				}
				fmt.Println(err)
				getFiles(dir)
				return err

			}
		}
		if info.IsDir() {
			f, _ := os.ReadDir(filepath.Join(dir, info.Name()))
			for _, file := range f {
				files = append(files, filepath.Join(dir, file.Name()))
			}
		} else {
			files = append(files, path)

		}
		return nil
	})

	return files
}
func random() string {

	length := 20

	// Generate random bytes
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// Encode random bytes to a string
	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	// Trim padding characters if any
	randomString = randomString[:length]
	return randomString

}

// GenerateAvatar generates an avatar of a person with random traits
func GenerateAvatar(api_var string) (*bytes.Buffer, error) {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Generating avatar..")
	// Generate random person traits
	if api_var == "random" {
		api_var = random()
	}
	new_traits := GenerateRandomPersonTraits(api_var)
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	DrawTrait(img, new_traits)

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
