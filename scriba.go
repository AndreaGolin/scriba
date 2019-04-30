package main

import(
	"log"
	"os"
	"fmt"
	"os/signal"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"time"
)

func main(){
	log.Printf("first commit")

	done := make(chan bool)
	monitorStream := make(chan int64)

	setEnvVars()
	printBanner()

	go listenForSigs(done)

	go monitor(os.Getenv("NGINX_ACCESS_LOG_PATH"), monitorStream)

	for{
		select{
		case <-done:
			log.Println("Stopping, bye.")
			return
		case s := <-monitorStream:
			fmt.Printf("\n New monitor log size: %d\n", s)
		}
	}
}

/**
 * @brief      Dummy parse function
 *
 * @return
 */
func parse(){
	file := os.Getenv("NGINX_ACCESS_LOG_PATH")
	info, err := os.Stat(file)
	if err != nil{
		log.Printf("%s", err)
	}

	log.Printf("%s", info.Name())
	log.Printf("%d", info.Size())
	log.Printf("%s", info.Mode())
}

/**
 * @brief      Monitor file function
 *
 * @return     
 */
func monitor(filename string, monitorStream chan int64){
	
	info, err := os.Stat(filename)
	if err != nil{
		log.Printf("%s", err)
	}

	currentSize := info.Size()
	for{

		time.Sleep(3000 * time.Millisecond)
		infoTmp, tmpErrror := os.Stat(filename)
		if tmpErrror != nil{
			log.Printf("%s", tmpErrror)
		}

		if infoTmp.Size() > currentSize{
			monitorStream<-infoTmp.Size()
			currentSize = infoTmp.Size()
		}
		
	}

	// log.Printf("%s", info.Name())
	// log.Printf("%d", info.Size())
	// log.Printf("%s", info.Mode())

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