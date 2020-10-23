/*

AUTHOR: MELVIN GEORGE
ASSIGNMENT: Cloudflare System Assignment

*/

package main

/*

------------ CLI USAGES -----------
===================================

--url <valid_url>
--profile <no_of_requests_to_the_URL>

--help to get all the usages in the app

*/

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"sort"
	"time"
)

// Structure for connection information
type connectionInfo struct {
	conn        *tls.Conn
	elapsedTime int64
}

// Structure for Single Request Profile
type requestsProfile struct {
	statusCode string
	details    *connectionInfo
}

func main() {

	// set flags for CLI
	host := flag.String("url", "https://linktree.melvingeorge10.workers.dev/links", "Provide the full URL") // --url
	profile := flag.Int("profile", 0, "Provide the number of requests to make to the defined URL")          // --profile
	flag.Parse()

	// Print details about request to console
	fmt.Printf("\nInformation\n")
	fmt.Printf("-----------\n\n")
	fmt.Printf("Using URL: %s\n", *host)
	if *profile != 0 {
		fmt.Printf("Requests Count: %d\n", *profile)
	} else {
		fmt.Printf("Requests Count: %d\n", (*profile)+1)
	}

	// Requests Profile Array
	// for storing details about
	// all the requests
	var RequestsProfileArray []requestsProfile

	// if the profile is greater than 0
	// then we will measure
	// the mean, median, slow, fast time of the URL
	if *profile > 0 {
		fmt.Printf("\nSending %d requests...🔥\n", *profile)
		for i := 0; i < *profile; i++ {
			details := makeRequest(host)
			profile := composeConnectionDetails(details)
			RequestsProfileArray = append(RequestsProfileArray, profile)
		}
		fmt.Printf("\nAnalysis Complete...⚡️\n")
		profileAndWriteToConsole(RequestsProfileArray)
		return
	}

	// else if the profile is 0
	// we will output the raw tcp responce
	details := makeRequest(host)
	writeResponseToConsole(details.conn)

}

// function to profile all the requests
// and write measurements to console
func profileAndWriteToConsole(requestsProfileArray []requestsProfile) {
	requestsProfileArrayLength := len(requestsProfileArray)

	if requestsProfileArrayLength > 0 {

		var slowestTime, fastestTime requestsProfile
		var sum, median float64
		var noOfSuccessRequests, noOfFailureRequests int
		var errorStatusCodeArray []string

		// Sort the Single Profile Array
		// Because, Need to calculate the median
		// and also the fastest and smallest time for request
		sort.Slice(requestsProfileArray, func(i, j int) bool {
			return requestsProfileArray[i].details.elapsedTime < requestsProfileArray[j].details.elapsedTime
		})

		// get the fastest and slowest time
		// easy since it is sorted already
		fastestTime = requestsProfileArray[0]
		slowestTime = requestsProfileArray[(requestsProfileArrayLength - 1)]

		// loop through all the requests
		for i := 0; i < requestsProfileArrayLength; i++ {
			// and check if the request is success or a failure
			// just a simple check
			// for more agreesive check we can use a switch statement
			// with all the 2xx Status codes
			if requestsProfileArray[i].statusCode == "200" {
				noOfSuccessRequests++
			} else {
				noOfFailureRequests++
				errorStatusCodeArray = append(errorStatusCodeArray, requestsProfileArray[i].statusCode)
			}

			// calculate total sum for mean time
			sum = sum + float64(requestsProfileArray[i].details.elapsedTime)
		}

		// calculate median time
		if requestsProfileArrayLength%2 == 0 {
			if requestsProfileArrayLength > 2 {
				// for even number of requests greater than 2
				median1 := requestsProfileArray[int(requestsProfileArrayLength/2)]
				median2 := requestsProfileArray[int((requestsProfileArrayLength/2)+1)]
				median = float64((median1.details.elapsedTime + median2.details.elapsedTime) / 2)
			} else {
				median = float64((requestsProfileArray[0].details.elapsedTime + requestsProfileArray[1].details.elapsedTime) / 2)
			}
		} else {
			// if array length is 1
			if requestsProfileArrayLength == 1 {
				median = float64(requestsProfileArray[0].details.elapsedTime)
			} else {
				// for odd number of requests greater than 1
				median = float64(requestsProfileArray[int((requestsProfileArrayLength+1)/2)].details.elapsedTime)
			}
		}

		// show time
		fmt.Printf("\nFastest Time: %d ms 🏃‍♀️", fastestTime.details.elapsedTime)
		fmt.Printf("\nSlowest Time: %d ms 🐌", slowestTime.details.elapsedTime)
		fmt.Printf("\nMean Time: %.2f ms 🚀", (sum / float64(requestsProfileArrayLength)))
		fmt.Printf("\nMedian Time: %.2f ms 🌟", median)
		fmt.Printf("\nSuccess rate: %d %% ✅", (noOfSuccessRequests/requestsProfileArrayLength)*100)
		fmt.Printf("\nFailure rate: %d %% ❌\n", (noOfFailureRequests/requestsProfileArrayLength)*100)

		// Show Error codes if any
		if len(errorStatusCodeArray) > 0 {
			fmt.Printf("\nError Code: %s 💀\n", errorStatusCodeArray[0])
		}

	}
}

// function to compose details about a single request
func composeConnectionDetails(details *connectionInfo) requestsProfile {

	// close connection at the end
	defer details.conn.Close()

	// get status code
	statusCode := getStatusCode(details.conn)

	req := requestsProfile{
		statusCode: statusCode,
		details:    details,
	}

	return req
}

// make request to a specific URL
func makeRequest(host *string) *connectionInfo {
	// parse URL
	u, err := url.Parse(*host)
	if err != nil {
		log.Fatal("--> Cannot Parse RAW URL\nURL should start with:  https://\n-->eg: https://google.com\n--> Provide absolute path", err)
	}

	// start the timer
	startTime := time.Now()

	// Dial to the website using the TLS protocol
	conn, err := tls.Dial("tcp", u.Host+":443", &tls.Config{})
	if err != nil {
		log.Fatal("\n--> URL should start with:  https://\n--> eg: https://google.com\n--> Provide absolute path")
	}

	// populate connection header
	connectionHeader := "GET " + u.Scheme + "://" + u.Host + u.Path + " HTTP/1.1\r\nHost: " + u.Host + ":443" + "\r\nConnection: close\r\nContent-type: text/html; charset=UTF-8\r\n\r\n"

	// Write to connection
	fmt.Fprintf(conn, connectionHeader)

	// end the timer
	endTime := time.Now()

	// substract endTime from startTime
	// and convert it to milliseconds
	totalTime := endTime.Sub(startTime).Milliseconds()

	// return the connection details
	return &connectionInfo{
		conn:        conn,
		elapsedTime: totalTime,
	}

}

// write raw tcp responce to console
func writeResponseToConsole(conn *tls.Conn) {

	fmt.Printf("\nResponse\n")
	fmt.Printf("-----------\n\n")

	// Intialise a byte buffer
	var buf bytes.Buffer

	// copy the connection raw repsonce to buffer

	io.Copy(&buf, conn)

	// print the buffer
	fmt.Println(buf.String())

}

// get the status code for a single request
func getStatusCode(conn *tls.Conn) string {
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	statusCode := string(status[9:12])

	return statusCode
}
