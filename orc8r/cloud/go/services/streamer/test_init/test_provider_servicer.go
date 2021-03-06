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

package test_init

import (
	"context"
	"testing"

	"github.com/go-magma/magma/lib/go/protos"
	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	streamer_protos "github.com/go-magma/magma/orc8r/cloud/go/services/streamer/protos"
	"github.com/go-magma/magma/orc8r/cloud/go/services/streamer/providers"
	"github.com/go-magma/magma/orc8r/cloud/go/test_utils"
)

type providerServicer struct {
	provider providers.StreamProvider
}

// StartNewTestProvider starts a new stream provider service which forwards
// calls to the passed provider.
func StartNewTestProvider(t *testing.T, provider providers.StreamProvider) {
	labels := map[string]string{
		orc8r.StreamProviderLabel: "true",
	}
	annotations := map[string]string{
		orc8r.StreamProviderStreamsAnnotation: provider.GetStreamName(),
	}
	srv, lis := test_utils.NewTestOrchestratorService(t, orc8r.ModuleName, provider.GetStreamName(), labels, annotations)
	servicer := &providerServicer{provider: provider}
	streamer_protos.RegisterStreamProviderServer(srv.GrpcServer, servicer)
	go srv.RunTest(lis)
}

func (p *providerServicer) GetUpdates(ctx context.Context, req *protos.StreamRequest) (*protos.DataUpdateBatch, error) {
	updates, err := p.provider.GetUpdates(req.GatewayId, req.ExtraArgs)
	res := &protos.DataUpdateBatch{Updates: updates}
	return res, err
}
