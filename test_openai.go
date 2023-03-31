package main  
  
import (  
	"bytes"  
	"encoding/json"  
	"fmt"  
	"io/ioutil"  
	"net/http"
)  
  
func main() {  
	// 设置API密钥和请求参数  
	apiKey := "85074218bf9a4219878b9c1cdd950ae5"  
	endpoint := "https://etranslate01.openai.azure.com/"  
	model := "gpt-35"
	prompt := "translate to zh-cn:home"
	modelType := "completions"
	apiversion := "2022-12-01"
  maxTokens := 400  
	// model := "embedding-ada"
	// input := "Sample Document goes here"
	// modelType := "embeddings"
  
	// 构建API请求  
	type RequestBody struct {  
		Prompt      string `json:"prompt"`  
		MaxTokens   int    `json:"max_tokens"` 
	}  
  
	requestBody, err := json.Marshal(RequestBody{  
		Prompt:     prompt,  
		MaxTokens:  maxTokens,
	})  
	if err != nil {  
		fmt.Println("Error creating request body:", err)  
		return  
	} 

	// type RequestBody struct {  
	// 	Input      string `json:"input"`
	// }

	// requestBody, err := json.Marshal(RequestBody{  
	// 	Input:     input,
	// })  
	// if err != nil {  
	// 	fmt.Println("Error creating request body:", err)  
	// 	return  
	// } 


	requestUrl := endpoint + "openai/deployments/" + model + "/" + modelType + "?api-version=" + apiversion  
  
	req, err := http.NewRequest("POST", requestUrl, bytes.NewReader(requestBody))  
	if err != nil {  
		fmt.Println("Error creating request:", err)  
		return  
	}  
	req.Header.Set("Content-Type", "application/json")  
	req.Header.Set("api-key", apiKey)  
  
	// 发送API请求  
	client := &http.Client{}  
	resp, err := client.Do(req)  
	if err != nil {  
		fmt.Println("Error sending request:", err)  
		return  
	}  
	defer resp.Body.Close()  
  
	// 解析API响应  
	body, err := ioutil.ReadAll(resp.Body)  
	if err != nil {  
		fmt.Println("Error reading response body:", err)  
		return  
	}  
	type ResponseChoices struct {
		Text string `json:"text"`
		Index int `json:"index"`
		FinishReason string `json:"finish_reason"`
		Logprobs interface{} `json:"logprobs"`
	}
	
	type ResponseUsage struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens int `json:"total_tokens"`
	}

	type Response struct {
		ID string `json:"id"`
		Object string `json:"object"`
		Created int `json:"created"`
		Model string `json:"model"`
		Choices []ResponseChoices `json:"choices"`
		Usage ResponseUsage `json:"usage"`
	}

	
	
	// type ResponseData struct {
	// 	Object string `json:"object"`
	// 	Index int `json:"index"`
	// 	Embedding []float64 `json:"embedding"`
	// }
	
	// type ResponseUsage struct {
	// 	PromptTokens int `json:"prompt_tokens"`
	// 	TotalTokens int `json:"total_tokens"`
	// }

	// type Response struct {
	// 	Object string `json:"object"`
	// 	Data []ResponseData `json:"data"`
	// 	Model string `json:"model"`
	// 	Usage ResponseUsage `json:"usage"`
	// }
  
	var response Response  
	err = json.Unmarshal(body, &response)  
	if err != nil {  
		fmt.Println("Error parsing response body:", err)  
		return  
	}  
  
	// 输出API响应  
	fmt.Println(response.Choices[0].Text)  
	// fmt.Println(response)
}  