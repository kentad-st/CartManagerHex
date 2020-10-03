package postgresDao

import (
	model "CartManagerHex/domain/model"

	repository "CartManagerHex/domain/repository"
	"database/sql"
	"fmt"

    _ "github.com/lib/pq"
)

type cartPersistence struct{}

func NewCartPersistence() repository.CartRepository {
	return &cartPersistence{}
}

func (cp cartPersistence) GetItems(user string) ([]model.Cart, error) {

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
    defer db.Close()

    if err != nil {
        fmt.Println(err)
	}
	script := " SELECT"
	script = script + "      cart_id"
	script = script + "      , user_nm"
	script = script + "      , item"
	script = script + "      , price"
	script = script + "      , quantity "
	script = script + "  FROM"
	script = script + "      public.cart "
	script = script + "  WHERE"
	script = script + "      user_nm = $1"
	script = script + "  ORDER BY"
	script = script + "      cart_id"
	rows, err := db.Query(script, user)
	var res []model.Cart
	
	for rows.Next() {
		var e model.Cart
		rows.Scan(&e.Id, &e.User, &e.Item, &e.Price, &e.Quantity)
		res = append(res, e)
	}
	
	return res, nil
}

func (cp cartPersistence) Post(cart *model.Cart) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
    defer db.Close()

    if err != nil {
        fmt.Println(err)
    }

	// INSERT
	script := "  INSERT  INTO public.cart(user_nm, item, price, quantity) "
	script = script + "  VALUES ($1, $2, $3, $4)"
	db.QueryRow(script, cart.User, cart.Item, cart.Price, cart.Quantity)
}

func (cp cartPersistence) DeleteItem(cart_id string) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
    defer db.Close()

    if err != nil {
        fmt.Println(err)
    }

	// DELETE
	script := "  DELETE FROM public.cart WHERE cart_id = $1 "
	db.QueryRow(script, cart_id)
}