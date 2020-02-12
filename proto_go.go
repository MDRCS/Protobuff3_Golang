package main

import (
	buffer_data "data"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	msg := GetData()
	//Read n write to a file usin proto3
	//WriteProto("message.bin",msg)
	ReadProto("message.bin",msg)
	fmt.Println(msg)

	//From proto to json
	getjson := ProtoJSON(msg)
	fmt.Println(getjson)

	// From json to proto
	JsonProto3(getjson,msg)
	fmt.Println(msg)
}

func JsonProto3(JSON string,pb proto.Message) proto.Message {

	err := jsonpb.UnmarshalString(JSON,pb)
	if err!= nil {
		log.Fatalln("Can't convert json to proto ...",err)
		return nil
	}

	return pb
}

func ProtoJSON(pb proto.Message) string {
	marshaller := jsonpb.Marshaler{}
	data,err := marshaller.MarshalToString(pb)

	if err!= nil {
		log.Fatalln("Can't convert proto to json ...",err)
		return "Nothing to display"
	}

	return data
}

//Write to file
func WriteProto(fname string, pb proto.Message) error{
	out,err := proto.Marshal(pb)
	if err!= nil {
		log.Fatalln("Can't serialise to bytes ",err)
		return err
	}

	if err := ioutil.WriteFile(fname,out,0644); err != nil {
		log.Fatalln("Can't write to file ",err)
		return err
	}

	fmt.Println("Data has been writing successfuly ..\n")
	return nil
}

//Read from a file
func ReadProto(fname string, pb proto.Message) error {
	in,err := ioutil.ReadFile(fname)
	if err!= nil {
		log.Fatalln("Can't serialise to bytes ",err)
		return err
	}

	if err := proto.Unmarshal(in,pb) ; err != nil {
		log.Fatalln("Couldn't read the file ",err)
		return err
	}

	fmt.Println("Data has been reading successfuly !!!\n")
	return nil
}

func GetData() *buffer_data.Message {
	data := buffer_data.Message{
		Sender:   "Med",
		Content:  "Hello world",
		Receiver: "Hamza",
		Ttl:      3300,
	}

	return &data
}

///go help gopath-get'
