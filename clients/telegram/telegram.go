package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"

	"yt-donwloader/lib/e"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
	sendAudioMethod   = "sendAudio"
)

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(ctx context.Context, offset int, limit int) (updates []Update, err error) {
	defer func() { err = e.WrapIfErr("can't get updates", err) }()

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(ctx, getUpdatesMethod, q, nil, nil)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendMessage(ctx context.Context, chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(ctx, sendMessageMethod, q, nil, nil)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil
}

func (c *Client) SendAudio(ctx context.Context, chatID int, audioFilePath string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("audio", audioFilePath)

	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		fmt.Println("Error opening audio file:", err)
		return err
	}
	defer audioFile.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("audio", audioFilePath)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return err
	}

	_, err = io.Copy(part, audioFile)
	if err != nil {
		fmt.Println("Error copying audio content:", err)
		return err
	}

	header := make(map[string]string)
	header["key"] = "key"
	header["value"] = writer.FormDataContentType()

	_, err = c.doRequest(ctx, sendAudioMethod, q, &requestBody, header)
	if err != nil {
		return e.Wrap("can't send audio", err)
	}

	return nil
}

func (c *Client) doRequest(ctx context.Context, method string, query url.Values, requestBody io.Reader, header map[string]string) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("can't do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), requestBody)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	if header != nil {
		req.Header.Set(header["key"], header["value"])
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	//var result map[string]interface{}
	//err = json.NewDecoder(resp.Body).Decode(&result)
	//if err != nil {
	//	log.Println("Error decoding response body:", err)
	//	return
	//}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
