package main

import (
	"fmt"

	"github.com/helmimuzkr/golang-play/repository"
	"github.com/helmimuzkr/golang-play/service"
)

func main() {

	// register repo
	repo := repository.NewRepo()

	// register service
	service1 := service.NewService()
	service2 := service.NewService2(repo)

	fmt.Println(service1.Find())
	fmt.Println(service1.Delete(1))

	fmt.Println(service2.Find())
	fmt.Println(service2.Delete(1))
}
