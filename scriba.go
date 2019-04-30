package main

import(
	"log"
	"os"
	"fmt"
	"os/signal"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main(){
	log.Printf("first commit")

	done := make(chan bool)

	setEnvVars()
	printBanner()

	go listenForSigs(done)
	go parse()

	for{
		select{
		case <-done:
			log.Println("Stopping, bye.")
			return
		}
	}
}

/**
 * @brief      Dummy parse function
 *
 * @return
 */
func parse(){
	file := "./dummy.txt"
	info, err := os.Stat(file)
	if err != nil{
		log.Printf("%s", err)
	}

	log.Printf("%s", info.Name())
	log.Printf("%d", info.Size())
	log.Printf("%s", info.Mode())
}

/**
 * @brief      { function_description }
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