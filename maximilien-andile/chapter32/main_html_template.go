package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Product struct {
	Name            string
	Price           string
	Description     string
	ShippingDate    time.Time
	Sale            bool
	SaleImagePath   []string
	MyFunc          func(string, string) string
	ShippingOptions []string
	Notes           [][]int
}

func (p Product) Foo() string {
	return "FOO"
}

func (p Product) Bar(test string) string {
	return fmt.Sprintf("Bar : %s", test)
}

func Bar(a string, b string) string {
	buf := bytes.NewBufferString(a)
	buf.WriteString(b)
	return buf.String()
}

func redTeaPotHandler(w http.ResponseWriter, r *http.Request) {
	/* var capitalizeFirstLetter = func(text string) string { return strings.ToUpper(text) }
	funcs := template.FuncMap{"capitalizeFirstLetter": capitalizeFirstLetter} */
	/* tmpl, err := template.New("product-dynamic.html").
	Funcs(funcs).ParseFiles("./views/product-dynamic.html") */

	var capitalizeFirstLetter = func(text string) string {
		return cases.Title(language.English).String(text)
	}
	funcs := template.FuncMap{"capitalizeFirstLetter": capitalizeFirstLetter}

	tmpl, err := template.New("product-dynamic.html").Funcs(funcs).ParseFiles(
		"./views/product.html",
		"./views/header.html",
		"./views/footer.html",
	)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	notes := [][]int{}
	note1 := []int{1, 2, 3}
	note2 := []int{4, 5, 6}
	notes = append(notes, note1)
	notes = append(notes, note2)

	teaPot := Product{
		Name:          "Red Tea Pot 250ml",
		Description:   "Test",
		Price:         "19.99",
		ShippingDate:  time.Now(),
		Sale:          true,
		SaleImagePath: []string{"path to image"},
		MyFunc:        Bar,
		ShippingOptions: []string{
			"Extra Priority",
			"Normal",
			"Low Priority",
		},
		Notes: notes,
	}

	err = tmpl.ExecuteTemplate(w, "product.html", teaPot)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/red-tea-pot", redTeaPotHandler)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
