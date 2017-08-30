// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package capability

import (
	"github.com/pingcap/tipb/go-mysqlx/Connection"
	"github.com/pingcap/tipb/go-mysqlx/Datatypes"
)

type HandleTLS struct {

}

func (h *HandleTLS) IsSupport() bool {
	return false
}

func (h *HandleTLS) GetName() string {
	return "tls"
}

func (h *HandleTLS) Get() Mysqlx_Connection.Capability {
	return Mysqlx_Connection.Capability{}
}

func (h *HandleTLS) Set(any *Mysqlx_Datatypes.Any) bool {
	return false
}
