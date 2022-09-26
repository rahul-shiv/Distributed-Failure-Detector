// Package main implements a server for worker service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	pb "mp1/protos"
	"net"
	"os"
	"os/exec"
	"strings"

	"google.golang.org/grpc"
)

var (
	port              = flag.Int("port", 50052, "The server port")
	GREP_COMMAND_NAME = "grep" // Name of grep command on linux
	VM_LOG_DIR        = "/mp1/logs/"
	LOCAL_LOG_DIR     = "/Users/rshiv/Documents/UIUC FA22/Distributed Systems - 425/cs425-mp1/logs/"
)

// variable to switch log dir location when using local machinevs vms
var LOG_DIR = VM_LOG_DIR

type server struct {
	pb.UnimplementedGrepInterfaceServer
}

/*this rpc is called using the test_client to create known logs in all of the vms, with each index in the logs array mapping to each vm
 */
func (s *server) CallWorkerCreateLogs(ctx context.Context, in *pb.WorkerCreateLogsRequest) (*pb.WorkerCreateLogsResponse, error) {
	log.Printf("Received Create Logs Request: %v", in.GetVmIndex())
	os.Chdir(LOG_DIR)
	err := os.Mkdir("tests", 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	var logs []string
	logs = append(logs, "Lorem ipsum dolor sit amet. Et excepturi ratione At obcaecati non harum distributed distributed distributed voluptatem id veniam amet aut dolor sequi est optio ullam non temporibus dolorem. Qui eius sunt sed dolore voluptatem est accusantium exercitationem. Et sequi quaerat qui atque distributed distributed distributed alias non debitis nostrum sed nihil eius et nihil facilis non omnis veniam vel accusamus quos. Vel suscipit corporis in nesciunt amet et nihil doloremque non dolorem soluta in quaerat porro ad harum voluptatem distributed distributed distributed est autem aliquam. Ut rerum asperiores id rerum enim non asperiores delectus a illo totam ut omnis veniam quo ducimus reprehenderit. Est necessitatibus impedit id possimus distributed distributed distributed fugit quo laborum beatae et excepturi itaque qui ipsa nobis. Aut animi galisum ex rerum enim rem consequuntur voluptatum ut quam iusto.")
	logs = append(logs, "Aut internos cupiditate qui fuga mollitia qui fugit internos. Rem distributed distributed distributed rerum molestiae non eius itaque est nobis fuga est aliquid veniam ex esse accusantium. Et similique suscipit eum aspernatur facilis qui quis quas quo consequatur animi sed magni enim. Sed enim quia sit explicabo consequatur vel impedit assumenda eum distributed distributed distributed autem distinctio est voluptatem minima aut nulla Quis et quibusdam atque. Vel sequi vero aut recusandae repellat et numquam quas aut soluta dignissimos hic dolor autem. 33 voluptatem repellat 33 molestiae consequatur a quaerat commodi et natus eaque eos ipsam omnis est omnis ducimus. 33 dolor deserunt aut distributed distributed distributed nisi quia vel quae architecto quo error saepe non dolores dignissimos. Ut aspernatur aliquam et modi deserunt distributed distributed distributed est nemo ipsum. Eos eveniet enim et veniam blanditiis qui blanditiis recusandae et voluptatum molestiae non tempore dolor At sint molestiae id atque voluptates!")
	logs = append(logs, "Ea consequatur vero est voluptatum dolor sit reprehenderit consectetur aut distributed distributed distributed consequatur ipsum in porro velit qui enim nihil. Quo quod nostrum est molestias ullam 33 iure consectetur. Totam nemo aut voluptas provident et nemo facilis. Et quasi iusto sit iste odit id odit velit aut dolores veritatis aut quam illo? Eos assumenda vitae et inventore laborum aut iste omnis. Aut Quis ducimus in inventore officia 33 dolorem aperiam qui voluptas facere id nemo dolorem aut reprehenderit enim nam commodi reprehenderit. Vel veniam deserunt sit dolor vero sit dicta officia.")
	logs = append(logs, "Sed harum ratione distributed distributed distributed sit quia quibusdam et illum nulla. Non dolor distributed distributed distributed adipisci non placeat sequi distributed distributed distributed temporibus sint. A exercitationem accusamus qui amet quisquam ut distributed distributed distributed tempore quidem. Quo nulla pariatur aut doloribus rerum aut eligendi molestiae distributed distributed distributed sed iure eveniet. At obcaecati reprehenderit aut asperiores voluptatem et eligendi illum ut saepe inventore distributed distributed distributed et ducimus saepe et Quis fugiat illum ut saepe inventore et ducimus saepe et Quis fugiat illum ut saepe inventore et ducimus saepe et Quis fugiat illum ut saepe inventore et ducimus saepe et Quis fugiat")
	logs = append(logs, "Sed culpa velit distributed distributed distributed et dolor voluptatum et ratione voluptas quo temporibus nemo et numquam nobis quo dignissimos distributed distributed distributed mollitia ut illum quos! Eum vero maiores aut nemo inventore et odio distributed distributed distributed architecto aut assumenda delectus est quia itaque ut quidem totam distributed distributed distributed sit pariatur iste. Et quas quas aut suscipit iure sed sequi alias id delectus voluptatem? Qui iusto quia ut culpa galisum id eligendi architecto et excepturi dicta ex nulla quibusdam et distributed distributed distributed suscipit aperiam. Aut incidunt laborum in quasi distributed distributed distributed facere et nobis architecto?")
	logs = append(logs, "Ea accusantium rerum distributed distributed distributed id voluptatem nobis qui nobis ullam ut aliquam minima ex voluptate natus et rerum distributed distributed distributed recusandae ab quia neque. Et sequi tempora aut optio velit est reiciendis iste ea nulla sequi hic dolore distributed distributed distributed delectus id voluptatum quos qui odit nisi. Ab accusantium repellendus non repudiandae ipsum eos aliquam consequatur distributed distributed distributed est aliquid quidem et sequi aspernatur cum enim debitis nam dolore modi. Sit itaque neque sit distributed distributed distributed odio sint ea aliquam perferendis id voluptas laborum. A esse impedit et distributed distributed distributed rerum molestiae ea officia nihil At labore rerum qui totam laudantium aut voluptas rerum!")
	logs = append(logs, "Sit consequatur sunt et nostrum reiciendis et illo voluptatem. Ad expedita expedita est quia consequatur non molestiae maxime. Et natus iste et officiis enim et consequatur quisquam? Qui similique sint sed perferendis saepe ab optio deleniti in esse tenetur est autem totam. Ut voluptas pariatur et quam necessitatibus id inventore doloribus sed voluptatem quod in voluptates nihil eum assumenda quidem et pariatur voluptatum. Sit libero accusantium qui animi nostrum distributed distributed distributed et ratione dolore et expedita nemo qui sint reprehenderit. Quo temporibus alias vel ipsam quis distributed distributed distributed id fuga doloremque id quia aspernatur distributed distributed distributed sit mollitia alias.")
	logs = append(logs, "Sed voluptatem molestiae id quisquam ducimus repellat assumenda qui repellendus voluptates? Hic voluptates dignissimos sit galisum vero et nemo dolorum At recusandae minima et libero maiores. Qui obcaecati aliquam aut aspernatur asperiores ut laborum deleniti qui maxime dignissimos. Sit error iste 33 necessitatibus non ducimus totam ab officia porro quo optio laboriosam rem illo temporibus qui voluptas dolorum. Sit officiis delectus est nulla necessitatibus aut officiis vero et quaerat optio non distributed distributed distributed quidem nulla id voluptas quam sed dicta veniam? distributed distributed distributed Sit dolores odit ad voluptate voluptatem cum quam quisquam et Quis internos. Non error quia et nemo debitis ex molestiae nihil. Ut nobis possimus est mollitia sunt vel beatae dolorum est doloribus quidem non consequatur inventore vel earum alias qui quasi ducimus. Vel eius Quis ut delectus quae ab vitae incidunt quo porro dolores At fugit deleniti.")
	logs = append(logs, "Quo eveniet harum id reiciendis unde vel quasi tenetur ab dolor voluptates. Non expedita rerum ea laboriosam temporibus aut beatae quos eum tempora ullam. In molestiae quos in sunt eaque est voluptatem ducimus. Et dolores laborum eos omnis dignissimos id odit ullam aut voluptatem inventore qui accusamus dolores non repudiandae incidunt? A impedit quia non architecto placeat sit iusto dolor non minus eaque qui commodi labore. Ad nulla natus et tenetur architecto ut atque commodi distributed distributed distributed id numquam necessitatibus. Ut pariatur quos ad voluptas iusto et possimus galisum eum iste repudiandae ut ipsa eligendi est esse nostrum distributed distributed distributed id velit maxime.")
	logs = append(logs, "Qui consequuntur fugiat At eaque galisum et internos optio. Ut fugiat eius ut sint unde ex expedita quibusdam et accusantium quidem. Est blanditiis eius ab provident veniam aut iure corrupti sit sequi perferendis qui architecto officiis. Vel rerum iure iure ab sint quisquam nam ipsum dolor non error tempora ut repudiandae corporis eos sapiente recusandae. Ut tempore fuga id omnis quos quo voluptates aliquam sed totam voluptatem. distributed distributed distributed Ex consequatur necessitatibus non repellat sapiente et ipsum error et voluptates distributed distributed distributed minima. Et officiis consectetur quo veniam iure aut fugit voluptates ut enim doloremque qui deserunt laudantium ut accusantium nulla aut adipisci totam.")
	f, err := os.Create(fmt.Sprintf("tests/test%d.log", in.GetVmIndex()))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	for i := 0; i < 135700; i++ {
		for _, line := range strings.Split(logs[in.GetVmIndex()-1], ". ") {
			f.WriteString(fmt.Sprintf("%s\n", line))
		}
	}
	return &pb.WorkerCreateLogsResponse{}, nil
}

/*this rpc is called using the test_client to delete previously created known logs in all of the vms
 */
func (s *server) CallWorkerDeleteLogs(ctx context.Context, in *pb.WorkerDeleteLogsRequest) (*pb.WorkerDeleteLogsResponse, error) {
	log.Printf("Received Delete Logs Request")
	os.Chdir(LOG_DIR)
	err := os.RemoveAll("tests")
	if err != nil {
		log.Printf(err.Error())
	}
	return &pb.WorkerDeleteLogsResponse{}, nil
}

/*this rpc runs the grep program on the log files
 */
func (s *server) Grep(ctx context.Context, in *pb.GrepRequest) (*pb.GrepResponse, error) {
	log.Printf("Received: %v", in.GetSearchString())
	os.Chdir(LOG_DIR)
	command := fmt.Sprintf("%s %s '%s' %s", GREP_COMMAND_NAME, in.GetSearchOptions(), in.GetSearchString(), in.GetSearchFile())
	cmd := exec.Command("/bin/sh", "-c", command)
	fmt.Printf(cmd.String() + "\n")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf(err.Error())
	}
	return &pb.GrepResponse{GrepOutput: string(output)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGrepInterfaceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
