package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	"fmt"
	"LearnGo/protoc/protocolbuf"
)

func main() {

	out, err := WritePb()
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	fmt.Println(out)


	person := &protocolbuf.Person{}

	if err := ReadPb(out,person); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(person)
}

func WritePb() ([]byte, error){
	p := &protocolbuf.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*protocolbuf.Person_PhoneNumber{
			{Number: "555-4321", Type: protocolbuf.Person_HOME},
		},
	}
	return proto.Marshal(p)

}

func ReadPb(data []byte, person *protocolbuf.Person) (error) {
	return proto.Unmarshal(data, person);
}