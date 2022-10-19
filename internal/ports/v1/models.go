package v1

type GreetingRequest struct {
	Name string `form:"name"`
}

type GreetingResponse struct {
	Text string `json:"text"`
}

type Error struct {
	Msg string `json:"msg"`
}
