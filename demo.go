package main

/*
  For all of the following, adapt `gerrit.wikimedia.org` to the Gerrit
  server you test against.

  To run this demo, first install go.
  For example on debian, run
     sudo apt-get install golang

  Then remove eventual pull/binaries that you might still be there
  from previous runs:

    rm -rf ~/go/src/gerrit.wikimedia.org/r/test/gerrit-ping
    rm -rf ~/go/bin/gerrit-ping

  Then fetch the sources from gerrit:

     go get gerrit.wikimedia.org/r/test/gerrit-ping

  (this will not print any confirmation)
  And finally run:

     go run gerrit.wikimedia.org/r/test/gerrit-ping

  This should output:

    demo works!

  To clean up again, run the above `rm` commands.
  Done!
*/

import "fmt"

func main() {
	fmt.Println("demo works!")
}
