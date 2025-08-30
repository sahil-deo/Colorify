package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
)

type Message struct {
	Color string `json:"text"`
}

func cloneSlice(slice []string) []string {
	newSlice := make([]string, len(slice))
	copy(newSlice, slice)
	return newSlice
}
func HexToRGB(hex string) ([]int64, error) {

	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	r, err := strconv.ParseInt(hex[0:2], 16, 0)

	if err != nil {
		return []int64{0, 0, 0}, err
	}

	g, err := strconv.ParseInt(hex[2:4], 16, 0)

	if err != nil {
		return []int64{0, 0, 0}, err
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 0)

	if err != nil {
		return []int64{0, 0, 0}, err
	}

	return []int64{r, g, b}, nil

}

var protanopia = [3][3]float64{
	{0.567, 0.433, 0.000},
	{0.558, 0.442, 0.000},
	{0.000, 0.242, 0.758},
}

var deuteranopia = [3][3]float64{
	{0.625, 0.375, 0.000},
	{0.700, 0.300, 0.000},
	{0.000, 0.300, 0.700},
}

var tritanopia = [3][3]float64{
	{0.950, 0.050, 0.000},
	{0.000, 0.433, 0.567},
	{0.000, 0.475, 0.525},
}

func clamp(x float64) float64 {
	return math.Max(0, math.Min(1, x))
}

func simulateColorBlind(r, g, b float64, matrix [3][3]float64) []float64 {
	rr := r*matrix[0][0] + g*matrix[0][1] + b*matrix[0][2]
	gg := r*matrix[1][0] + g*matrix[1][1] + b*matrix[1][2]
	bb := r*matrix[2][0] + g*matrix[2][1] + b*matrix[2][2]
	return []float64{clamp(rr), clamp(gg), clamp(bb)}
}
func RGBToHex(c []float64) string {
	return fmt.Sprintf("#%02X%02X%02X", int(c[0]), int(c[1]), int(c[2]))
}
func HextoHex(hex string) (string, string, string) {

	var rgb []int64 = make([]int64, 3)

	rgb, _ = HexToRGB(hex)
	var normalized []float64 = make([]float64, 3)

	for i, c := range rgb {
		normalized[i] = (float64(c) / 255.0)
	}

	d1 := simulateColorBlind(normalized[0], normalized[1], normalized[2], protanopia)
	d2 := simulateColorBlind(normalized[0], normalized[1], normalized[2], deuteranopia)
	d3 := simulateColorBlind(normalized[0], normalized[1], normalized[2], tritanopia)

	for i, c := range d1 {
		d1[i] = c * 255
	}

	for i, c := range d2 {
		d2[i] = c * 255
	}

	for i, c := range d3 {
		d3[i] = c * 255
	}

	return RGBToHex(d1), RGBToHex(d2), RGBToHex(d3)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	colors := r.URL.Query()["color"]

	var d1 map[string]string = make(map[string]string, len(colors))
	var d2 map[string]string = make(map[string]string, len(colors))
	var d3 map[string]string = make(map[string]string, len(colors))

	for _, c := range colors {
		d1[c], d2[c], d3[c] = HextoHex(c)
	}

	result := map[string]map[string]string{"protanopia": d1, "deuteranopia": d2, "tritanopia": d3}
	r2 := map[string]map[string]map[string]string{"colors": result}
	json.NewEncoder(w).Encode(r2)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/api", handler)

	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		fmt.Println("H")
		port = "5555"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
