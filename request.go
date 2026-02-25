package gripp

import (
	"encoding/json"
	"log"
)

type BaseRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	ID     int         `json:"id"`
}

type RequestType []BaseRequest

type RequestFilter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type RequestPaging struct {
	FirstResult int `json:"firstresult"`
	MaxResults  int `json:"maxresults"`
}

type RequestOrdering struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type Response struct {
	Id     int             `json:"id"`
	Thread string          `json:"thread"`
	Result json.RawMessage `json:"result"`
	Error  string          `json:"error,omitempty"`
}

type GetResult[T any] struct {
	Rows                  []T  `json:"rows"`
	Count                 int  `json:"count"`
	Start                 int  `json:"start"`
	Limit                 int  `json:"limit"`
	NextStart             int  `json:"next_start"`
	MoreItemsInCollection bool `json:"more_items_in_collection"`
}

type RequestBuilder[T any] struct {
	client   *Client
	base     string
	filters  []RequestFilter
	paging   *RequestPaging
	ordering *RequestOrdering
}

const maxPageSize = 250

func (rb *RequestBuilder[T]) Filter(input ...interface{}) *RequestBuilder[T] {
	var field, operator string
	var value interface{}
	if len(input) == 3 {
		field = input[0].(string)
		operator = input[1].(string)
		value = input[2]
	} else if len(input) == 2 {
		field = input[0].(string)
		operator = "equals"
		value = input[1]
	} else {
		return rb
	}
	rb.filters = append(rb.filters, RequestFilter{
		Field:    rb.base + "." + field, // example: "project.name"
		Value:    value,
		Operator: operator,
	})
	return rb
}

func (rb *RequestBuilder[T]) Page(firstResult, maxResults int) *RequestBuilder[T] {
	if maxResults <= 0 || maxResults > maxPageSize {
		maxResults = maxPageSize
	}
	rb.paging = &RequestPaging{
		FirstResult: firstResult,
		MaxResults:  maxResults,
	}
	return rb
}

func (rb *RequestBuilder[T]) OrderBy(field, direction string) *RequestBuilder[T] {
	rb.ordering = &RequestOrdering{
		Field:     rb.base + "." + field, // example: "project.name"
		Direction: direction,
	}
	return rb
}

func (rb *RequestBuilder[T]) Delete(itemId int) error {
	// params example
	//"params":[
	//   {{itemid}}
	//],
	// build the request body
	request := BaseRequest{
		Method: rb.base + ".delete", // example: "project.delete"
		Params: []interface{}{
			itemId,
		},
		ID: 1,
	}

	// send the request and return the response
	var requests RequestType
	requests = append(requests, request)

	_, err := rb.client.makeRequest(requests)
	if err != nil {
		log.Println("Error making request:")
		return err
	}

	return nil
}

func (rb *RequestBuilder[T]) Get() ([]T, error) {
	// params example
	//"params":[
	//	[
	//		{
	//			"field":"project.id",
	//			"operator":"greaterequals",
	//			"value":1
	//		}
	//	],
	//	{
	//		"paging":{
	//			"firstresult":0,
	//			"maxresults":10
	//		}
	//		"orderings":[
	//			{
	//				"field":"project.id",
	//				"direction":"asc"
	//			}
	//		]
	//	}
	//],
	// build the request body
	buildRequest := func(paging *RequestPaging) BaseRequest {
		return BaseRequest{
			Method: rb.base + ".get", // example: "project.get"
			Params: []interface{}{
				rb.filters,
				struct {
					Paging    *RequestPaging    `json:"paging,omitempty"`
					Orderings []RequestOrdering `json:"orderings,omitempty"`
				}{
					Paging: paging,
					Orderings: func() []RequestOrdering {
						if rb.ordering != nil {
							return []RequestOrdering{*rb.ordering}
						}
						return nil
					}(),
				},
			},
			ID: 1,
		}
	}

	resolvePaging := func(paging *RequestPaging) *RequestPaging {
		if paging == nil {
			return nil
		}
		resolved := *paging
		if resolved.MaxResults <= 0 || resolved.MaxResults > maxPageSize {
			resolved.MaxResults = maxPageSize
		}
		return &resolved
	}

	if rb.paging != nil {
		request := buildRequest(resolvePaging(rb.paging))
		responses, err := rb.client.makeRequest(RequestType{request})
		if err != nil {
			log.Println("Error making request:")
			return nil, err
		}

		for _, response := range responses {
			var result GetResult[T]
			err = json.Unmarshal(response.Result, &result)
			if err != nil {
				log.Printf("Error unmarshalling response: %v\nResponse: %s\n", err, string(response.Result))
				return nil, err
			}
			return result.Rows, nil
		}

		return nil, nil
	}

	start := 0
	maxResults := maxPageSize
	var allRows []T
	for {
		paging := &RequestPaging{FirstResult: start, MaxResults: maxResults}
		request := buildRequest(paging)
		responses, err := rb.client.makeRequest(RequestType{request})
		if err != nil {
			log.Println("Error making request:")
			return nil, err
		}

		var hadResponse bool
		for _, response := range responses {
			hadResponse = true
			var result GetResult[T]
			err = json.Unmarshal(response.Result, &result)
			if err != nil {
				log.Printf("Error unmarshalling response: %v\nResponse: %s\n", err, string(response.Result))
				return nil, err
			}
			allRows = append(allRows, result.Rows...)
			if !result.MoreItemsInCollection {
				return allRows, nil
			}
			start = result.NextStart
		}

		if !hadResponse {
			return allRows, nil
		}
	}
}
