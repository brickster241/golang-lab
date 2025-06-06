package learnlab

import (
	"fmt"
	"net/url"
)

func parseURL() {

	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]
	rawURL := "https://www.amazon.jobs/en/search?base_query=&loc_query=Cartagena+de+Indias%2C+Colombia&latitude=10.42513&longitude=-75.5383&loc_group_id=&invalid_location=false&country=COL&city=Cartagena&region=&county=Bolivar"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error Parsing URL :", err)
		return
	}

	fmt.Println("Scheme    :", parsedURL.Scheme)
	fmt.Println("Host      :", parsedURL.Host)
	fmt.Println("Port      :", parsedURL.Port())
	fmt.Println("Path      :", parsedURL.Path)
	fmt.Println("Raw Query :", parsedURL.RawQuery)
	fmt.Println("Fragment  :", parsedURL.Fragment)

	queryParams := parsedURL.Query()
	fmt.Println(queryParams)

	// Building URL
	baseURL := &url.URL{
		Scheme: "https",
		Host: "goatifi.com",
		Path: "/2021-abu-dhabi-gp",
	}
	query := baseURL.Query()
	query.Set("name", "Lewis Hamilton")
	query.Set("mood", "cooked")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL :", baseURL.String())
}