package model
type Law struct {
	Id      int    `json: "id" `
	Content string `json: "content"`
	Object  string `json: "object"`
	Money   string `json: "money"`
	Extra   string `json: "extra"`
	Hold    string `json: "hold"`
	Source  string `json: "source"`
}
 type Respone struct{
	 Err string `json: "response_err"`
	 Success string `json: "respone_success"`
 }
 