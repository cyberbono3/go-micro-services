package main 

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/ewanvalentine/shippy/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
    
)

const (
	address = "localhost:50051"
	defaultFilename = "consignment.json"
)



func parseFile(file string) (*pb.Consignment, error){
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err!= nil {
		return nil, err
	}
    //from json to interface type
	json.Unmarshal(data, &consignment)
	return consignment, err
}


func main()  {
	// set-upo connection to the server
	conn,err := grpc.Dial(address, grpc.WithInsecure())
	if err!= nil {
		log.Fatal("Didd connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)


	//contact the server and print out the response
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}


}

