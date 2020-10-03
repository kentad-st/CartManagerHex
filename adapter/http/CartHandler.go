package handler

import (
	"net/http"
	
	service "CartManagerHex/service"

	model "CartManagerHex/domain/model"

	"github.com/gin-gonic/gin"
	
    "github.com/gin-contrib/cors"
	"strconv"
	"fmt"
)


type CartHandler interface {
	MainHttpd()
	ShowItems() gin.HandlerFunc
	AddToCart() gin.HandlerFunc
	DeleteFromCart() gin.HandlerFunc
}

type cartHandler struct {
	cartManagementService service.CartManagementService
}

func NewCartHandler(cs service.CartManagementService) CartHandler {
	return &cartHandler{
		cartManagementService :cs,
	}
}


func (ch cartHandler) MainHttpd(){
	r := gin.Default()

    // CORS 対応
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"}
    r.Use(cors.New(config))

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
	})
	
	r.POST("/showitems", ch.ShowItems())
	r.POST("/addtocart", ch.AddToCart())
	r.POST("/deletefromcart", ch.DeleteFromCart())
	r.POST("/orderfromcart", ch.OrderFromAllItems())

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

type GetItemsRequest struct {
	User string `json:"user"`
}

func (ch cartHandler) ShowItems() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := GetItemsRequest{}
        c.Bind(&requestBody)
		result, err := ch.cartManagementService.ShowItems(requestBody.User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, result)
		}
        c.JSON(http.StatusOK, result)
    }
}

type AddToCartRequest struct {
	User string `json:"user"`
	Item string `json:"item"`
	Price string `json:"price"`
	Quantity string `json:"quantity"`
}

func (ch cartHandler)AddToCart() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := AddToCartRequest{}
        c.Bind(&requestBody)
        p, err := strconv.ParseInt(requestBody.Price, 10, 64)    
        if err != nil {
            fmt.Println(err)
        }
        q, err := strconv.ParseInt(requestBody.Quantity, 10, 64)    
        if err != nil {
            fmt.Println(err)
        }

        item := model.Cart{
			User: &requestBody.User,
			Item: &requestBody.Item,
			Price: &p,
			Quantity: &q,
        }
        ch.cartManagementService.AddToCart(&item)

        c.Status(http.StatusNoContent)
    }
}

type DeleteFromCartRequest struct {
	CartId string `json:"cart_id"`
}

func (ch cartHandler)DeleteFromCart() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := DeleteFromCartRequest{}
        c.Bind(&requestBody)
        ch.cartManagementService.DeleteFromCart(requestBody.CartId)

        c.Status(http.StatusNoContent)
    }
}

type OrderFromCartRequest struct {
	User string `json:"user"`
}

func (ch cartHandler)OrderFromAllItems() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := OrderFromCartRequest{}
        c.Bind(&requestBody)
        err := ch.cartManagementService.OrderFromAllItems(requestBody.User)
        if err != nil {
            fmt.Println(err)
            c.Status(http.StatusInternalServerError)
        }else{
            c.Status(http.StatusNoContent)
        }
    }
}