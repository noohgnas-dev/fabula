package parser

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/mrizkimaulidan/fabula/file"
	"github.com/mrizkimaulidan/fabula/instagram"
)

const API_URL = "https://storiesig.info/api/ig/stories/%s"

type Parser struct {
	Instagram *instagram.Instagram
}

func NewParser(instagram *instagram.Instagram) *Parser {
	return &Parser{
		Instagram: instagram,
	}
}

// Call the third party API URL to fetching the story
// it will returning the response JSON and decode
// it to Response struct
func (p *Parser) Call() (*Response, error) {
	resp, err := http.Get(fmt.Sprintf(API_URL, p.Instagram.ProfileID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Parsing the Response struct and checking
// the story content is video or photo
// returning the slices of struct File
func (p *Parser) Parsing(response *Response) *[]file.File {
	var files []file.File
	for _, r := range response.Result {
		newFile := file.File{
			Filename: strconv.Itoa(rand.Intn(int(time.Now().UnixNano() / 1000000))),
		}

		if r.HasAudio {
			newFile.Extension = ".mp4"
			newFile.URL = r.VideoVersions[0].URL

			files = append(files, newFile)

		} else {
			newFile.Extension = ".jpg"
			newFile.URL = r.ImageVersions2.Candidates[0].URL

			files = append(files, newFile)
		}
	}

	return &files
}
