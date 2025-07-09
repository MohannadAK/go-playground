package main

import (
	"fmt"

	"github.com/MohannadAK/go-playground/pkg/htmlWordCounter"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
      <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`

func main() {
	words, pics := htmlWordCounter.Start(raw)
	fmt.Printf("Words: %d, Images: %d\n", words, pics)
}
