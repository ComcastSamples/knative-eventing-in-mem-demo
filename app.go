/**
* Copyright 2018 Comcast Cable Communications Management, LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
* SPDX-License-Identifier: Apache-2.0
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handlePost(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	fmt.Fprintf(rw, "%s", body)
	log.Println("############ EVENTING Demo using In-Memory CHANNEL ################")
	log.Printf("%s", body)
}

func main() {
	log.Print("Starting server on port 8080...")
	http.HandleFunc("/", handlePost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
