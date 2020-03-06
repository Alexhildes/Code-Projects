

package main

​

import (

	"bufio"

	"fmt"

	"net/http"

	"net/http/httputil"

	"os"

	"strings"

)

​

func main() {

	fmt.Println("Usage: cat {list} | " + os.Args[0])

	r := bufio.NewReader(os.Stdin)

​

	id, inErr := r.ReadString('\n')

	for inErr == nil {

		id = strings.TrimSpace(id)

		if id == "" {

			id, inErr = r.ReadString('\n')

			continue

		}

​

		url := "http://172.28.35.11:8182/cxf/rest/v1/bpm/bpmservice/deletetask/" + id

​

		req, err := http.NewRequest("DELETE", url, nil)

		must(err, "failed to construct request")

​

		fmt.Println("   DELETE", url)

		resp, err := http.DefaultClient.Do(req)

		must(err, "failed to execute request")

​

		b, err := httputil.DumpResponse(resp, true)

		must(err, "failed to dump response")

​

		fmt.Println("== Response ==")

		fmt.Println(string(b))

		fmt.Print("\n\n")

​

		id, inErr = r.ReadString('\n')

	}

​

	fmt.Println("--- Finished", inErr)

}

​

func must(err error, msg string) {

	if err == nil {

		return

	}

​

	fmt.Println(msg)

	fmt.Println("Error: ", err)

	os.Exit(1)

}

