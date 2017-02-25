package etm

import "testing"

func TestEtm_String (t *testing.T) {
	etm := Etm{
		"<html>" +
			"<head></head>" +
			"<body>" +
				"<a href='http://example.com' x-button='buy'>Buy</a>" +
				"<br/>" +
				"<div>" +
					"<a href='http://example.com/register' x-button='register'>Register</a>" +
				"</div>" +
			"</body>" +
		"</html>",
		map[string]string{"utm_campaign": "new-year-sale", "utm_source": "email"},
	}

	result := etm.String()
	expectedResult := "<html><head></head><body><a href=\"http://example.com?utm_campaign=new-year-sale&amp;utm_source=email&amp;x-button=buy\" x-button=\"buy\">Buy</a><br/><div><a href=\"http://example.com/register?utm_campaign=new-year-sale&amp;utm_source=email&amp;x-button=register\" x-button=\"register\">Register</a></div></body></html>"
	if result  != expectedResult {
		t.Error("Expeceted : " + expectedResult)
		t.Error("Got       : " + result)
	}
}
