package data


type Data struct{
	Ip string `json:"ip"`
	User string `json:"user"`
	From string `json:"from"`
	Type string `json:"type"`
	Content string `json:"content"`
	Userlist []string `json:"userlist"`
}