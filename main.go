package main

import (
	handler "CartManagerHex/adapter/http"
	postgres "CartManagerHex/adapter/postgres"
	app2app "CartManagerHex/adapter/app2app"
	service "CartManagerHex/service"
)

func main(){
	cartPersistence := postgres.NewCartPersistence()
	cartApp2AppHooks := app2app.NewOrderHooks()
	cartService := service.NewCartManagementService(cartPersistence, cartApp2AppHooks)
	cartHandler := handler.NewCartHandler(cartService)

	cartHandler.MainHttpd()
}