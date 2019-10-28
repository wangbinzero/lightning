package utils

type Response struct {
	Head map[string]string `json:"head"`
	Body interface{}       `json:"body"`
}

type ArrayBodyStruct struct {
	CurrentPage  int         `json:"current_page"`
	TotalPage    int         `json:"total_page"`
	PerPage      int         `json:"per_page"`
	NextPage     int         `json:"next_page"`
	PreviousPage int         `json:"previous_page"`
	Data         interface{} `json:"data"`
}
