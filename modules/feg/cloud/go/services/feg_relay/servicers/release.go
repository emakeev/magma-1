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

package servicers

import (
	"github.com/go-magma/magma/lib/go/protos"
	fegprotos "github.com/go-magma/magma/modules/feg/cloud/go/protos"

	"golang.org/x/net/context"
)

// ReleaseReq relays the ReleaseRequest sent from VLR->FeG->Access Gateway
func (srv *FegToGwRelayServer) ReleaseReq(
	ctx context.Context,
	req *fegprotos.ReleaseRequest,
) (*protos.Void, error) {
	if err := validateFegContext(ctx); err != nil {
		return nil, err
	}
	return srv.ReleaseRequestUnverified(ctx, req)
}

func (srv *FegToGwRelayServer) ReleaseRequestUnverified(
	ctx context.Context,
	req *fegprotos.ReleaseRequest,
) (*protos.Void, error) {
	conn, ctx, err := getGWSGSServiceConnCtx(ctx, req.Imsi)
	if err != nil {
		return &protos.Void{}, err
	}
	client := fegprotos.NewCSFBGatewayServiceClient(conn)
	return client.ReleaseReq(ctx, req)
}
