/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type S7MessageObjectResponse struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	Parent        *S7DataAlarmMessage
}

// The corresponding interface
type IS7MessageObjectResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7MessageObjectResponse) CpuFunctionType() uint8 {
	return 0x08
}

func (m *S7MessageObjectResponse) InitializeParent(parent *S7DataAlarmMessage) {
}

func NewS7MessageObjectResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7DataAlarmMessage {
	child := &S7MessageObjectResponse{
		ReturnCode:    returnCode,
		TransportSize: transportSize,
		Parent:        NewS7DataAlarmMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7MessageObjectResponse(structType interface{}) *S7MessageObjectResponse {
	castFunc := func(typ interface{}) *S7MessageObjectResponse {
		if casted, ok := typ.(S7MessageObjectResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*S7MessageObjectResponse); ok {
			return casted
		}
		if casted, ok := typ.(S7DataAlarmMessage); ok {
			return CastS7MessageObjectResponse(casted.Child)
		}
		if casted, ok := typ.(*S7DataAlarmMessage); ok {
			return CastS7MessageObjectResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7MessageObjectResponse) GetTypeName() string {
	return "S7MessageObjectResponse"
}

func (m *S7MessageObjectResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7MessageObjectResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7MessageObjectResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7MessageObjectResponseParse(readBuffer utils.ReadBuffer) (*S7DataAlarmMessage, error) {
	if pullErr := readBuffer.PullContext("S7MessageObjectResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, pullErr
	}
	returnCode, _returnCodeErr := DataTransportErrorCodeParse(readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field")
	}
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, pullErr
	}
	transportSize, _transportSizeErr := DataTransportSizeParse(readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field")
	}
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("S7MessageObjectResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7MessageObjectResponse{
		ReturnCode:    returnCode,
		TransportSize: transportSize,
		Parent:        &S7DataAlarmMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7MessageObjectResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageObjectResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (returnCode)
		if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
			return pushErr
		}
		_returnCodeErr := m.ReturnCode.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
			return popErr
		}
		if _returnCodeErr != nil {
			return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
		}

		// Simple Field (transportSize)
		if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
			return pushErr
		}
		_transportSizeErr := m.TransportSize.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
			return popErr
		}
		if _transportSizeErr != nil {
			return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("S7MessageObjectResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *S7MessageObjectResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}