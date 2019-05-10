package scriba_nginx

import(
	// "time"
	"log"
	"os"
	"os/exec"
	"fmt"
	"bytes"
	"net"
)

type NginxLine struct{
	SrcIP net.IP
	RspCode uint16
}
type NginxLog struct{
	LineList []*NginxLine
}

/**
 * @brief      Ge a "snap" of nginx access log
 *
 * @param      accessFilePath  The access file path
 *
 * @return
 */
func SnapNginxAccess(accessFilePath string){
	printBannerSnapAccess(accessFilePath)

	/**
	 * Open Nginx access log
	 */
	file, err := os.Open(accessFilePath)
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()

	stdout := execAwk()
	fmt.Printf("%s", stdout)

}

func execAwk() string{

	/**
	 * Build the awk arg string
	 */
	awkCommandArg := fmt.Sprintf("{print $%s\"|\"$%s}", os.Getenv("NGXACC_IP_POSITION"), os.Getenv("NGXACC_RSP_CODE"))
	

	/**
	 * Exec awk
	 */
	cmd := exec.Command("awk", awkCommandArg, os.Getenv("NGINX_ACCESS_LOG_PATH"))

	/**
	 * type io.ReaderCloser
	 */
	stdout, err := cmd.StdoutPipe()
	if err != nil {
	    log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
	    log.Fatal(err)
	}

	/**
	 * Create a buffer for the reader
	 */
	stdoutBuf := new(bytes.Buffer)
	/**
	 * Read from the reader
	 */
	stdoutBuf.ReadFrom(stdout)
	/**
	 * Copy the buffer as string
	 */
	stdoutString := stdoutBuf.String()

	/**
	 * Calling wait close the pipe
	 */
	if err := cmd.Wait(); err != nil {
	    log.Fatal(err)
	}

	return stdoutString
}

/**
 * @brief      Print snap access banner
 *
 * @param      accessFilePath  The access file path
 *
 * @return 
 */
func printBannerSnapAccess(accessFilePath string){
	fmt.Println()
	fmt.Println("Snapping Nginx access log.")
	fmt.Printf("Nginx accesso log path: %s", accessFilePath)
	fmt.Println()
}
