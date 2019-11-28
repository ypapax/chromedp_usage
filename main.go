package main

// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("main")
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15 * time.Second)
	defer cancel()

	//u := `https://golang.org/pkg/time/`
	u := `https://www.whatismybrowser.com/detect/what-is-my-user-agent`
	selector := `#detected_value`
	log.Println("requesting", u)
	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(u),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(selector),
		// retrieve the value of the textarea
		chromedp.Text(selector, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
