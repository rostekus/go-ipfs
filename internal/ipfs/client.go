package ipfs

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

type Client struct {
	Url string
}

func (c *Client) Add(content string) (*http.Response, error) {
	var contentBuffer bytes.Buffer
	contentBuffer.WriteString(content)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileField, err := writer.CreateFormField("file")
	if err != nil {
		return &http.Response{}, nil
	}

	_, err = io.Copy(fileField, &contentBuffer)
	if err != nil {
		return &http.Response{}, nil
	}

	writer.Close()

	req, err := http.NewRequest("POST", c.Url, body)
	if err != nil {
		return &http.Response{}, nil
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp, nil
}
