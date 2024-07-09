package main

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"github.com/pzierahn/chatbot_services/account"
	"github.com/pzierahn/chatbot_services/auth"
	"github.com/pzierahn/chatbot_services/chat"
	"github.com/pzierahn/chatbot_services/collections"
	"github.com/pzierahn/chatbot_services/datastore"
	"github.com/pzierahn/chatbot_services/documents"
	"github.com/pzierahn/chatbot_services/llm"
	"github.com/pzierahn/chatbot_services/llm/anthropic"
	"github.com/pzierahn/chatbot_services/llm/openai"
	"github.com/pzierahn/chatbot_services/llm/vertex"
	"github.com/pzierahn/chatbot_services/notion"
	pb "github.com/pzierahn/chatbot_services/proto"
	"github.com/pzierahn/chatbot_services/search"
	"github.com/pzierahn/chatbot_services/search/qdrant"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const credentialsFile = "service_account.json"

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func initDatastore(ctx context.Context) *datastore.Service {
	db, err := datastore.New(ctx)
	if err != nil {
		log.Fatalf("failed to create datastore service: %v", err)
	}

	return db
}

func initFirebase(ctx context.Context) *firebase.App {
	var opts []option.ClientOption
	if _, err := os.Stat(credentialsFile); err == nil {
		serviceAccount := option.WithCredentialsFile(credentialsFile)
		opts = append(opts, serviceAccount)
	}

	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		log.Fatalf("failed to create firebase app: %v", err)
	}

	return app
}

func initBucket(ctx context.Context, app *firebase.App) *storage.BucketHandle {
	firebaseStorage, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("failed to create firebase storage client: %v", err)
	}

	bucket, err := firebaseStorage.Bucket("brainboost-399710.appspot.com")
	if err != nil {
		log.Fatalf("did not get bucket: %v", err)
	}

	return bucket
}

func initModels(ctx context.Context) []llm.Chat {
	openaiClient, err := openai.New()
	if err != nil {
		log.Fatalf("failed to create openai client: %v", err)
	}

	vertexClient, err := vertex.New(ctx)
	if err != nil {
		log.Fatalf("failed to create vertex client: %v", err)
	}

	claude, err := anthropic.New()
	if err != nil {
		log.Fatalf("failed to create anthropic client: %v", err)
	}

	models := []llm.Chat{
		openaiClient,
		vertexClient,
		claude,
	}

	return models
}

func initSearch(engine llm.Embedding) search.DB {
	search, err := qdrant.New(engine)
	if err != nil {
		log.Fatalf("failed to create qdrant search: %v", err)
	}

	return search
}

func initAuth(ctx context.Context, app *firebase.App) auth.Service {
	service, err := auth.WithFirebase(ctx, app)
	if err != nil {
		log.Fatalf("failed to create auth service: %v", err)
	}

	return service
}

func setTestEnv() {
	_ = os.Setenv("CHATBOT_MONGODB_URI", "mongodb://localhost:27017")
	_ = os.Setenv("CHATBOT_QDRANT_KEY", "")
	_ = os.Setenv("CHATBOT_QDRANT_URL", "localhost:6334")
	_ = os.Setenv("CHATBOT_QDRANT_INSECURE", "true")
}

func main() {
	ctx := context.Background()

	setTestEnv()

	app := initFirebase(ctx)

	database := initDatastore(ctx)
	models := initModels(ctx)

	engine := models[0].(llm.Embedding)
	search := initSearch(engine)
	bucket := initBucket(ctx, app)
	//authService := initAuth(ctx, app)

	userService := &account.InsecureVerifier{
		//Database: database,
		//Auth:     authService,
	}

	chatService := &chat.Service{
		Models:   models,
		Auth:     userService,
		Database: database,
		Search:   search,
	}

	documentsService := &documents.Service{
		Auth:        userService,
		Database:    database,
		Storage:     bucket,
		SearchIndex: search,
	}

	collectionService := &collections.Service{
		Auth:     userService,
		Database: database,
		Storage:  bucket,
		Search:   search,
	}

	notionService := &notion.Client{
		Chat:      chatService,
		Documents: documentsService,
		Database:  database,
		Auth:      userService,
		Cache:     make(map[string]string),
	}

	grpcServer := grpc.NewServer()
	//pb.RegisterAccountServiceServer(grpcServer, userService)
	pb.RegisterChatServiceServer(grpcServer, chatService)
	pb.RegisterDocumentServiceServer(grpcServer, documentsService)
	pb.RegisterCollectionServiceServer(grpcServer, collectionService)
	pb.RegisterNotionServer(grpcServer, notionService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9055"
		log.Printf("defaulting to port %s", port)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("starting server on %v", listener.Addr().String())
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
