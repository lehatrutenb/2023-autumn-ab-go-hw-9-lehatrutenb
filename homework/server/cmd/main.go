package main

import (
	"flag"
	"fmt"
	"homework/server/internal/adapters/foldersys"
	"homework/server/internal/app"
	"homework/server/internal/ports/grpcserver"
	"homework/server/internal/starboards/grpcclient"
	"log"
	"net"
)

// вообще изначально хотел сделать отдельные папки /client и /server и в них уже свои /internal и /cmd, но из клиента мы явно хотим использовать protoc + структуры файлов
// А они внутри internal , и получется я не могу импортировать из клиента , но добавать префиксы/суффиксы к internal и cmd тоже не хочется
// так что решил разделить сделав вот так общий main и файлы уже с приписками
// Я поискал альтернативу слову port, но для клиента - но так и не нашёл хорошего, так что starborad

func RunServer() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpcserver.NewServer(app.NewApp(foldersys.NewRepo("")))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RunClient() {
	fmt.Println(grpcclient.GetFileInfo(":8081", "123.txt"))
	fmt.Println(grpcclient.GetFileNames(":8081"))
	fmt.Println(grpcclient.GetFileData(":8081", "123.txt"))
}

func main() {
	var server = flag.Bool("server", false, "Set that flag to run sever, not client")
	flag.Parse()

	if server == nil {
		// todo very bad
		log.Panic("how it can be?")
	}
	if *server {
		RunServer()
	} else {
		RunClient()
	}
}
