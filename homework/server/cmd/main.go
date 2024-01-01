package main

import (
	"flag"
	"homework/server/internal/adapters/foldersys"
	"homework/server/internal/app"
	"homework/server/internal/ports/grpcserver"
	"log"
	"net"

	"go.uber.org/zap"
)

// вообще изначально хотел сделать отдельные папки /client и /server и в них уже свои /internal и /cmd, но из клиента мы явно хотим использовать protoc + структуры файлов
// А они внутри internal , и получется я не могу импортировать из клиента , но добавать префиксы/суффиксы к internal и cmd тоже не хочется
// так что решил разделить сделав вот так общий main
// Я поискал альтернативу слову port, но для клиента - но так и не нашёл хорошего, так что starborad

// TODO add server gracefull shutdown

func RunServer(addr string, repoaddr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// вообще, очень хочется 2 логера, тк кажется, что это как почти 2 разных продукта
	// тут, конечно есть небольшая делема - нужно ли делать логгер внутри app при том условии
	// что сам app не генерирует ошибки, а просто прослойка
	serverLogger, err := zap.NewProduction(zap.WithCaller(true))
	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}
	defer serverLogger.Sync()

	repoLogger, err := zap.NewProduction(zap.WithCaller(true))
	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}
	defer repoLogger.Sync()

	s := grpcserver.NewServer(app.NewApp(foldersys.NewRepo(repoaddr, repoLogger)), serverLogger)
	serverLogger.Info("Server is listenning", zap.String("addr", addr))
	if err := s.Serve(lis); err != nil {
		serverLogger.Fatal("Failed to serve", zap.Error(err))
	}
}

func main() {
	var addr = flag.String("addr", ":8081", "Server listening address")
	var repoaddr = flag.String("repoaddr", "", "Server folder system address")
	flag.Parse()

	if addr == nil || repoaddr == nil {
		log.Fatal("Can't parse server address or repoaddress from flags")
	}

	RunServer(*addr, *repoaddr)
}
