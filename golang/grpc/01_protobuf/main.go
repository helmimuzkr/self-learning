package main

import (
	"fmt"
	"log"

	"github.com/helmimuzkr/01_protobuf/pb"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	user := &pb.User{
		Id:          1,
		FirstName:   "helmi",
		LastName:    "muzakir",
		DateOfBirth: "1999-10-05",
	}
	pagination := &pb.Pagination{
		LimitData:   10,
		TotalData:   1,
		CurrentPage: 1,
		TotalPage:   1,
	}
	listResponse := &pb.ListResponse{
		Pagination: pagination,
		Data:       []*pb.User{user},
	}

	fmt.Println("======== Struct proto")
	fmt.Println(listResponse)
	fmt.Println()

	// Convert Struct proto to String
	fmt.Println("======== String")
	fmt.Println(listResponse.String())
	fmt.Println()

	// Convert Struct proto to JSON.
	buf, err := protojson.Marshal(listResponse)
	if err != nil {
		log.Println(err)
		return
	}
	strJSON := string(buf)
	fmt.Println("======== Buffer json string")
	fmt.Println(strJSON)
	fmt.Println()

	// Convert JSON to Struct proto
	bufJSON := buf
	listResponse2 := new(pb.ListResponse)
	err = protojson.Unmarshal(bufJSON, listResponse2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("======== Struct proto 2")
	fmt.Println(listResponse2)
	fmt.Println()
}
