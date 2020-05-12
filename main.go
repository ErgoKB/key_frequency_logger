package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("can not get executable directory, err: %s", err)
	}

	filePath := flag.String("o", filepath.Join(currentDir, "output.csv"), "output file name")
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
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sig:
	case <-frequencyLogger.doneCh:
	}
	fmt.Println()
}
