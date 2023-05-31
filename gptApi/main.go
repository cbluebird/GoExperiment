package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

// Chatgpt的api参数
type ChatgptMessage struct {
	//数值在0-1.0之间
	Temperature float64   `json:"temperature"`
	Messages    []message `json:"messages"`
}

type message struct {
	//role有三种值，“system”代表预设，“user”代表用户，“assistant”代表gpt的回复消息
	Role string `json:"role"`
	//content为具体的聊天内容
	Content string `json:"content"`
}

type ChatCompletion struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Message      struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
	} `json:"choices"`
	Created int64  `json:"created"`
	Id      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type ChatResponse struct {
	Content string `json:"content"`
}

func main() {
	var data1 ChatgptMessage
	var model = "gpt-3.5-turbo"
	data1.Temperature = 0.2
	var messages []message
	messages = append(messages, message{"user", "你好用英语怎么说"})
	data1.Messages = messages
	data1.Temperature = 0.5
	//proxyUrl, err := url.Parse("http://192.168.5.120:7890")
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS11,
	}
	//client := &http.Client{}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: cfg,
		}}

	url := "https://api.openai.com.5435486.xyz/v1/chat/completions"
	method := "POST"
	payloadData := map[string]interface{}{
		"model":       model,
		"messages":    data1.Messages,
		"temperature": data1.Temperature,
	}
	payloadBuffer := new(bytes.Buffer)
	json.NewEncoder(payloadBuffer).Encode(payloadData) //使用 json.NewEncoder编码,编码结果暂存到 buffer
	payload := bytes.NewReader(payloadBuffer.Bytes())
	req, err := http.NewRequest(method, url, payload)
	InitConfig()
	key := Config.GetString("chatgpt-key")
	req.Header.Add("Authorization", key)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.openai.com")
	req.Header.Add("Connection", "keep-alive")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)
	req.Close = true
	var data ChatCompletion
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		return
	}
	response := ChatResponse{}
	response.Content = data.Choices[0].Message.Content
	log.Println(response.Content)
}

var Config = viper.New()

func InitConfig() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")
	Config.WatchConfig() // 自动将配置读入Config变量
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal("Config not find", err)
	}
}
