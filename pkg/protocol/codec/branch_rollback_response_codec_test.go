/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package codec

import (
	"testing"

	serror "github.com/seata/seata-go/pkg/util/errors"

	model2 "github.com/seata/seata-go/pkg/protocol/branch"

	"github.com/seata/seata-go/pkg/protocol/message"
	"github.com/stretchr/testify/assert"
)

func TestBranchRollbackResponseCodec(t *testing.T) {
	msg := message.BranchRollbackResponse{
		AbstractBranchEndResponse: message.AbstractBranchEndResponse{
			Xid:          "123344",
			BranchId:     56678,
			BranchStatus: model2.BranchStatusPhaseoneFailed,
			AbstractTransactionResponse: message.AbstractTransactionResponse{
				TransactionErrorCode: serror.ErrorCodeBeginFailed,
				AbstractResultMessage: message.AbstractResultMessage{
					ResultCode: message.ResultCodeFailed,
					Msg:        "FAILED",
				},
			},
		},
	}

	codec := BranchRollbackResponseCodec{}
	bytes := codec.Encode(msg)
	msg2 := codec.Decode(bytes)

	assert.Equal(t, msg, msg2)
}
