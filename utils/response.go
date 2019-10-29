package utils

import "fmt"

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

type ArrayDataResponse struct {
	Head map[string]string `json:"head"`
	Body interface{}
}

var (
	SuccessResponse = Response{Head: map[string]string{"code": "1000", "msg": "success"}}
	ArrayResponse   = ArrayDataResponse{Head: map[string]string{"code": "1000", "msg": "success"}}
)

func BuildError(code string) Response {
	return Response{Head: map[string]string{"code": code}}
}

func (errorResponse Response) Error() string {
	return fmt.Sprintf("code:%s;msg:%s", errorResponse.Head["code"], errorResponse.Head["msg"])
}

func (arrayResponse *ArrayDataResponse) Init(data interface{}, page, count, per_page int) {
	total_page := count / per_page
	if (count % per_page) != 0 {
		total_page += 1
	}
	next_page := page + 1
	if next_page > total_page {
		next_page = total_page
	}
	previousPage := page - 1
	if previousPage < 1 {
		previousPage = 1
	}
	body := ArrayBodyStruct{}
	body.Data = data
	body.CurrentPage = page
	body.TotalPage = total_page
	body.PerPage = per_page
	body.NextPage = next_page
	body.PreviousPage = previousPage
	arrayResponse.Body = body
}
