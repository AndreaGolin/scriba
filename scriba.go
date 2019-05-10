package main

import(
	"log"
	"os"
	"fmt"
	"os/signal"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	scrbngx "scriba/scriba_nginx"
)

func main(){

	done := make(chan bool)

	setEnvVars()
	printBanner()

	go listenForSigs(done)

	go scrbngx.SnapNginxAccess(os.Getenv("NGINX_ACCESS_LOG_PATH"))

	for{
		select{
		case <-done:
			log.Println("Stopping, bye.")
			return
		}
	}
}


/**
 * @brief      Listen for system signals. Kill app wathever comes up :)
 *
 * @param      chan bool the kill channel
 */
func listenForSigs(done chan bool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		select {
		case s := <-c:
			fmt.Printf("Received os signal: %s", s.String())
			fmt.Println()
			done <- true
			return
		}
	}
}

/**
 * @brief      Set environmental variables
 *
 * @return
 */
func setEnvVars(){

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
}

/**
 * @brief      Print banner or MOTD
 *
 * @return     
 */
func printBanner(){
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Println()
	fmt.Printf("%s version %s\n\r", red(os.Getenv("PACKAGE_NAME")), blue(os.Getenv("VERSION")))
	fmt.Println()
}