package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter()
	routes := map[string]string{
		"base":     "/{width:[0-9]+}x{height:[0-9]+}",
		"extended": "/{width:[0-9]+}x{height:[0-9]+}/{background:(?:[0-9]|[a-f])+}",
		"full":     "/{width:[0-9]+}x{height:[0-9]+}/{background:(?:[0-9]|[a-f])+}/{color:(?:[0-9]|[a-f])+}",
	}

	// dimensions routes
	router.HandleFunc(routes["base"], handleRequest).Queries("label", "{label}")
	router.HandleFunc(routes["base"], handleRequest)

	// dimensions, background color routes
	router.HandleFunc(routes["extended"], handleRequest).Queries("label", "{label}")
	router.HandleFunc(routes["extended"], handleRequest)

	// dimensions, background color, label color routes
	router.HandleFunc(routes["full"], handleRequest).Queries("label", "{label}")
	router.HandleFunc(routes["full"], handleRequest)

	// Index handler
	router.HandleFunc("/", index)

	// 404 handler
	router.NotFoundHandler = http.HandlerFunc(render404)

	// serve
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handles the index route
func index(writer http.ResponseWriter, request *http.Request) {
	vars := map[string]string{
		"label": "Please provide dimensions /{width}x{height}",
	}
	vars = prepareVars(vars)
	_ = generateAndWriteImage(&vars, writer)
	log.Println("index::")
}

// Handles the 404 Error-Page
func render404(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content-type", "text/html")
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprint(writer, getErrorMessage())
}

// Handles the 400 Error-Page
func render400(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content-type", "text/html")
	writer.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(writer, getErrorMessage())
}

// Default function from which to handle all correct requests
func handleRequest(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	vars = prepareVars(vars)
	err := generateAndWriteImage(&vars, writer)
	log.Println("image::", vars, err)
	if err != nil {
		render400(writer, request)
	}
}

// Builds the default error message
func getErrorMessage() string {
	link := fmt.Sprintf("<a href='%s'>GitHub</a>", "https://github.com/jschaefer-io/fizz-image")
	message := fmt.Sprintln("No Image could be generated.<br />")
	message += fmt.Sprintln("Please check the documentation at " + link)
	return message
}

// Merges the request var string map with a
// default-value set
func prepareVars(vars map[string]string) map[string]string {
	defaults := map[string]string{}
	defaults["width"] = "400"
	defaults["height"] = "300"
	defaults["background"] = "a5a5a5"
	defaults["color"] = "fff"
	defaults["label"] = "" // if label stays empty, it will be replaced to {width}x{height}

	for key, value := range vars {
		if defaults[key] != "" || key == "label" {
			defaults[key] = value
		}
	}
	return defaults
}

// Generate the image from the given vars
// and writes it to the given io.Writer
func generateAndWriteImage(vars *map[string]string, writer io.Writer) error {
	width, height, _ := readSizes(vars)

	if width*height > 4000*4000 {
		return errors.New("image dimensions to big")
	}

	backgroundColor, _ := readColor(vars, "background")
	labelColor, _ := readColor(vars, "color")
	label, _ := readLabel(vars)
	img := buildImage(backgroundColor, width, height)

	if label == "" {
		label = fmt.Sprintf("%dx%d", width, height)
	}
	addLabel(label, labelColor, &img)

	options := jpeg.Options{
		Quality: 100,
	}
	_ = jpeg.Encode(writer, &img, &options)
	return nil
}

// Tries to resolve the image label from
// the given map of request vars
func readLabel(vars *map[string]string) (string, error) {
	label := (*vars)["label"]
	if len(label) == 0 {
		return "", errors.New("given label is empty")
	}
	return label, nil
}

// Tries to resolve the image dimensions from the
// given map of request vars
func readSizes(vars *map[string]string) (int, int, error) {
	width, errWidth := strconv.Atoi((*vars)["width"])
	height, errHeight := strconv.Atoi((*vars)["height"])

	if errWidth != nil || errHeight != nil {
		return 0, 0, errors.New("unable to read image dimensions")
	}

	return width, height, nil
}

// Tries to resolve the background image from the
// given map of request vars
func readColor(vars *map[string]string, field string) (color.RGBA, error) {
	bg := (*vars)[field]

	// Handle shorthand syntax
	if len(bg) == 3 {
		bg = bg + bg
	}

	// basic validation
	if len(bg) != 6 {
		return color.RGBA{}, errors.New("unable to read background color")
	}

	// Calculate Colors
	red, e1 := hexStringToUInt8(string(bg[0]) + string(bg[1]))
	green, e2 := hexStringToUInt8(string(bg[2]) + string(bg[3]))
	blue, e3 := hexStringToUInt8(string(bg[4]) + string(bg[5]))

	// Validated hex2uint8
	if e1 != nil || e2 != nil || e3 != nil {
		return color.RGBA{}, errors.New("unable to read background color")
	}

	// Build and return
	backgroundColor := color.RGBA{
		R: red,
		G: green,
		B: blue,
		A: 255,
	}
	return backgroundColor, nil
}

// Converts the given string to an unsigned 8 bit int
// assuming the input string contains base 16 digits
func hexStringToUInt8(str string) (uint8, error) {
	result, err := strconv.ParseInt(str, 16, 9)
	return uint8(result), err
}

// Generates an image with the given background color and dimensions
func buildImage(backgroundColor color.RGBA, width int, height int) image.RGBA {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{C: backgroundColor}, image.Point{}, draw.Src)
	return *img
}

// renders the given string-label on the given image
func addLabel(label string, labelColor color.RGBA, img *image.RGBA) {
	size := img.Bounds().Size()
	x := size.X / 2
	y := size.Y / 2

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(labelColor),
		Face: basicfont.Face7x13,
	}

	// String offset to center the string
	offset := d.MeasureBytes([]byte(label)) / 2

	point := fixed.Point26_6{
		X: fixed.Int26_6(x*64) - offset,
		Y: fixed.Int26_6(y * 64),
	}
	d.Dot = point
	d.DrawString(label)
}
