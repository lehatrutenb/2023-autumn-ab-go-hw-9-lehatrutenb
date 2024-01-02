package main

import (
	"flag"
	"homework/server/internal/adapters/foldersys"
	"homework/server/internal/app"
	"homework/server/internal/ports/grpcserver"
	"homework/server/internal/starboards/grpcclient"

	"log"

	"go.uber.org/zap"
)

// вообще изначально хотел сделать отдельные папки /client и /server и в них уже свои /internal и /cmd, но из клиента мы явно хотим использовать protoc + структуры файлов
// А они внутри internal , и получется я не могу импортировать из клиента , но добавать префиксы/суффиксы к internal и cmd тоже не хочется
// так что решил разделить сделав вот так общий main
// Я поискал альтернативу слову port, но для клиента - но так и не нашёл хорошего, так что starborad

// TODO add server gracefull shutdown

type ServerLoggers struct {
	ServerLogger *zap.Logger
	RepoLogger   *zap.Logger
}

func RunServer(addr string, repoaddr string, slg ServerLoggers, opts ...grpcserver.ServerOption) {
	// вообще, очень хочется 2 логера, тк кажется, что это как почти 2 разных продукта
	// тут, конечно есть небольшая делема - нужно ли делать логгер внутри app при том условии
	// что сам app не генерирует ошибки, а просто прослойка
	var err error

	if slg.ServerLogger == nil {
		slg.ServerLogger, err = zap.NewProduction(zap.WithCaller(true))
		if err != nil {
			log.Fatalf("Failed to init logger: %v", err)
		}
		defer slg.ServerLogger.Sync()
	}

	if slg.RepoLogger == nil {
		slg.RepoLogger, err = zap.NewProduction(zap.WithCaller(true))
		if err != nil {
			log.Fatalf("Failed to init logger: %v", err)
		}
		defer slg.RepoLogger.Sync()
	}

	lis, s := grpcserver.NewServer(app.NewApp(foldersys.NewRepo(repoaddr, slg.RepoLogger)), slg.ServerLogger, addr, opts...)

	slg.ServerLogger.Info("Server is listenning", zap.String("addr", addr))
	if err := s.Serve(lis); err != nil {
		slg.ServerLogger.Fatal("Failed to serve", zap.Error(err))
	}
}

// можно сказать, что это надо положить в client.go , но мне кажется нет, тк по смыслу всё-таки
// не так далеко от раннера для сервера
type ClientLogger struct {
	Logger *zap.Logger
}

func CreateClient(clg ClientLogger, opts ...grpcclient.ClientOption) *grpcclient.Client {
	var err error

	if clg.Logger == nil {
		clg.Logger, err = zap.NewProduction(zap.WithCaller(true))
		if err != nil {
			log.Fatalf("Failed to init logger: %v", err)
		}
	}

	return grpcclient.NewClient(clg.Logger, opts...)
}

func main() {
	var server = flag.Bool("server", false, "Set that flag to run sever, not client")
	var serverAddr = flag.String("serveraddr", ":8081", "Server listening address")
	var repoAddr = flag.String("repoaddr", "", "Server folder system address")
	flag.Parse()

	if serverAddr == nil || repoAddr == nil || server == nil {
		log.Fatal("Can't parse server address or repoaddress from flags")
	}

	if *server {
		RunServer(*serverAddr, *repoAddr, ServerLoggers{})
	} else {
		CreateClient(ClientLogger{})
	}
}
