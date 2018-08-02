// /*
// Copyright 2017 The Kubernetes Authors.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

package main

// import (
// 	"flag"
// 	"os"
// 	"runtime"

// 	_ "k8s.io/apimachinery/pkg/apimachinery/announced"
// 	_ "k8s.io/apimachinery/pkg/apimachinery/registered"
// 	"k8s.io/apimachinery/pkg/util/wait"
// 	_ "k8s.io/apiserver/pkg/endpoints/discovery"
// 	_ "k8s.io/apiserver/pkg/server"
// 	"k8s.io/apiserver/pkg/util/logs"
// 	// Following deps are not saved in vendor/ by godep save, so import it explicitly

// 	// "github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_adaptor/server"
// 	"github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_adaptor/server"
// )

// func main() {
// 	logs.InitLogs()
// 	defer logs.FlushLogs()

// 	if len(os.Getenv("GOMAXPROCS")) == 0 {
// 		runtime.GOMAXPROCS(runtime.NumCPU())
// 	}

// 	cmd := server.NewCommandStartSampleAdapterServer(os.Stdout, os.Stderr, wait.NeverStop)
// 	cmd.Flags().AddGoFlagSet(flag.CommandLine)
// 	if err := cmd.Execute(); err != nil {
// 		panic(err)
// 	}
// }
