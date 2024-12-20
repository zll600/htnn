// Copyright The HTNN Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helper

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/yaml"
)

func MustReadInput(fn string, out *[]map[string]interface{}) {
	input, err := os.ReadFile(fn)
	Expect(err).NotTo(HaveOccurred())
	Expect(yaml.UnmarshalStrict(input, out, yaml.DisallowUnknownFields)).To(Succeed())
	// shuffle the input to detect bugs relative to the order
	res := *out
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
}

func WaitServiceUp(port string, service string) {
	Eventually(func() bool {
		c, err := net.DialTimeout("tcp", port, 10*time.Millisecond)
		if err != nil {
			return false
		}
		c.Close()
		return true
	}, 10*time.Second, 50*time.Millisecond,
		fmt.Sprintf("%s is unavailable. Please run `make start-service` in ./controller to make it up.", service))
}
