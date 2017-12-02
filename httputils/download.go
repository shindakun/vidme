package httputils

import (
	// "fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Download(file string, url string) (err error) {
	// Paths need to be fixed up to work cross platform.
	//ex, err := os.Executable()
	//if err != nil {
	//  panic(err)
	//}
	//exPath := filepath.Dir(ex)
	//fmt.Println(exPath)
	file1 := strings.Replace(file, "/", "\\", -1)
	//fmt.Println(file1)

	// Create the file
	out, err := os.Create(file1)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	//resp, err := http.Get(url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	//  defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	// fmt.Println(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
