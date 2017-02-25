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
    fmt.Println(etm.Convert(html, params))
}
```

Result
```
<html>
<head></head>
<body>
    <a href="http://example.com?utm_campaign=new-year-sale&amp;utm_source=email&amp;x-button=buy" x-button="buy">Buy</a>
    <br/>
    <div>
        <a href="http://example.com/register?utm_campaign=new-year-sale&amp;utm_source=email&amp;x-button=register" x-button="register">Register</a>
    </div>
</body>
</html>
```
