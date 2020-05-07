package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	filePath := flag.String("o", "output.csv", "output file name")

	flag.Parse()

	asciiArt :=
		`

                                               &&&&&&&&@
                ,&&&&&&&     &&&&&&&        &&&&&&&&&&&&
               &&&@  @&&*  #&&&   &&&      &&&&&&&&&&&&              #&&&&,
              &&&@   @&&.  &&&   .&&@     %&&&&&.               *&&&&&&&&&&&&&
            &&&&&&&&&&&&&&&&&&&&&&&&&     ,&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
          .&&#                    &&&&&& &&&                   &&&&&,   &&&&&&@
          &&%                    @&&   &&&&                    &&@ &&&&&
         &&&                     &&%   &&&                    &&&     &&&&
        &&&       &&&&&&        &&&   &&&       &&&&&        @&&      %&&.
       @&&        &&&&&        &&&    &&/      @&&&&&        &&(      &&&
      *&&                     %&&.   &&&                    &&&      &&&
      &&(                     &&@   &&&                    &&&      @&&
     &&&&&@@@@%/,.         .&&&&   .&&&                   &&&       &&&
     &&%&&&&&&&&&&&&&&&&&&&&& &&&  &&&&&&&&&&&&&&&&&&&&&&&&&&&     &&&
    &&&                        &&&(&&/                      &&&@  %&&.
   (&&,                         &&&&&                         &&& &&&
   &&&.                          &&&.                          &&&&&
   &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&@



     @&&&&&&&&&&                                &&     &&&    &&&&&&&&&&
     @&&                                        &&  @&&@      &&@      &&,
     @&&&&&&&&&#  @&@&&&&  /&&&&&&%   &&&&&     &&&&&         &&&&&&&&&&%
     @&&          @&&     &&    &&/ &&@   @&&   && &&&(       &&@      &&@
     @&&          @&&     &&&   &&  &&@   &&&   &&   *&&&     &&@     *&&&
     @&&&&&&&&&&  @&&       %&&&&&    &&&&&     &&      &&&   &&&&&&&&&.
                          &&    &&(
                           *&&&&&
	`
	fmt.Println(asciiArt)
	fmt.Println("ErgoKB key frequency logger - log your key press to csv for further analysis")
	fmt.Println("for more information, please visit https://www.ergokb.tw")
	fmt.Println("To stop, use Ctrl-c to leave program")
	fmt.Printf("\n\n")
	defer log.Info("Thank you for using frequency logger, bye-bye!")

	log.Infof("The output is stored to file %s", *filePath)
	frequencyLogger, err := newLogger(*filePath)
	if err != nil {
		log.WithError(err).Fatal("Can not create frequency logger")
	}
	defer frequencyLogger.close()
	go frequencyLogger.run()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println()
}
