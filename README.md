# go-etm, html hyperlink manipulation for email campaign

## Installation
    $ go get github.com/benibudiharto/go-etm

## How to use
```Go
package main

import (
  "github.com/benibudiharto/go-etm"
  "fmt"
)

func main() {
    params := map[string]string{"utm_campaign": "new-year-sale", "utm_source": "email"}
    html   := "<html>" +
                    "<head></head>" +
                    "<body>" +
                        "<a href='http://example.com' x-button='buy'>Buy</a>" +
                        "<br/>" +
                        "<div>" +
                            "<a href='http://example.com/register' x-button='register'>Register</a>" +
                        "</div>" +
                    "</body>" +
              	"</html>"
    etm.Convert(html, params)
    fmt.Println(etm.String())
}
```
