/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	lteplugin "github.com/go-magma/magma/modules/lte/cloud/go/plugin"
	"github.com/go-magma/magma/orc8r/cloud/go/plugin"
)

func main() {}

func GetOrchestratorPlugin() plugin.OrchestratorPlugin {
	return &lteplugin.LteOrchestratorPlugin{}
}
