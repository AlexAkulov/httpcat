package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

var (
	noColor  = flag.Bool("no-color", false, "Disable colors")
	listen   = flag.String("listen", ":8888", "Listen address")
	responce = flag.Int("responce", 200, "Responce code")
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(color.Output, color.CyanString("--------------------------------------------------"))
	fmt.Fprintf(color.Output, "Method: %v  Location: %v\n", color.RedString(r.Method), color.BlueString(r.RequestURI))
	if len(r.Header) > 0 {
		fmt.Fprintln(color.Output, color.YellowString("Headers:"))
		for k, v := range r.Header {
			fmt.Printf("    \"%s\": \"%v\"\n", k, v)
		}
	} else {
		fmt.Fprintln(color.Output, color.YellowString("No Headers"))
	}
	fmt.Fprintln(color.Output, color.BlueString("Body:"))
	color.Set(color.FgGreen)
	io.Copy(color.Output, r.Body)
	color.Unset()
	fmt.Println()
	w.WriteHeader(*responce)
}

func main() {
	flag.Parse()
	color.NoColor = *noColor

	fmt.Fprintln(color.Output, color.YellowString("   /\\_/\\"))
	fmt.Fprintln(color.Output, color.YellowString(" =( °w° )="))
	fmt.Fprintf(color.Output, "%v     %v\n", color.YellowString("   ) U (  //"), color.GreenString("HTTP Server"))
	fmt.Fprintf(color.Output, "%v      %v\n", color.YellowString("  (..|..)//"), color.GreenString("Listen on %v", *listen))

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(*listen, nil); err != nil {
		fmt.Fprintln(color.Output, color.RedString(err.Error()))
	}
}
