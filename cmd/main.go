pa kage main
                          
import (
	"fmt"
	"log"                                      
	"net"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/config"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/di"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/util"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if util.HasError(err) {
		log.Fatalln("failed to load configs", err)
	}

	service, err := di.InitializeService(config)
	if util.HasError(err) {
		log.Fatalln("failed to initialize deps", err)
	}

	listener, err := net.Listen("tcp", config.Port)

	if util.HasError(err) {
		log.Fatalln("failed to create listener ", err)
	}

	server := grpc.NewServer()
	pb.RegisterInventoryServiceServer(server, service)

	fmt.Println("service raising up ...")
	fmt.Println("serving on port:", config.Port)

	log.Fatalln(server.Serve(listener))
}
