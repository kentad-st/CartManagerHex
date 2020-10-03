package app2app


import (
	model "CartManagerHex/domain/model"
	service "CartManagerHex/service"
    "errors"
    "encoding/json"
    "net/http"
	"bytes"
	"strconv"
)

type orderHooks struct{}

func NewOrderHooks() service.OrderHooksService {
	return &orderHooks{}
}

type OrderRequestBody struct {
	User string `json:"user"`
	Item string `json:"item"`
	Price string `json:"price"`
	Quantity string `json:"quantity"`
}

func (oh orderHooks) Order(items *[]model.Cart) error {
	requestItems := []OrderRequestBody{}
	for _, v := range *items {
		requestItems = append(requestItems, OrderRequestBody{
			User: *v.User,
			Item: *v.Item,
			Price: strconv.FormatInt(*v.Price, 10),
			Quantity: strconv.FormatInt(*v.Quantity, 10),
		})
	}
	
	b, err := json.Marshal(requestItems)
	if err != nil {
		return err
	}

    req, err := http.NewRequest(
        "POST",
        "http://localhost:8081/addorders",
        bytes.NewBuffer([]byte(b)),
    )
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
	}
	if resp == nil || resp.StatusCode != 204 {
		return errors.New("request failed")
	}
    defer resp.Body.Close()

    return err
}