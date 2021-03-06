//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package io

import (
	"encoding/json"
	"io"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	dto "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
)

// EventReader unmarshals a request body into an Event type
type EventReader interface {
	ReadAddEventRequest(reader io.Reader) ([]dto.AddEventRequest, errors.EdgeX)
	ReadUpdateEventPushedByIdRequest(reader io.Reader) ([]dto.UpdateEventPushedByIdRequest, errors.EdgeX)
}

// NewRequestReader returns a BodyReader capable of processing the request body
func NewEventRequestReader() EventReader {
	return NewJsonReader()
}

// jsonReader handles unmarshaling of a JSON request body payload
type jsonEventReader struct{}

// NewJsonReader creates a new instance of jsonReader.
func NewJsonReader() jsonEventReader {
	return jsonEventReader{}
}

// Read reads and converts the request's JSON event data into an Event struct
func (jsonEventReader) ReadAddEventRequest(reader io.Reader) ([]dto.AddEventRequest, errors.EdgeX) {
	var addEvents []dto.AddEventRequest
	err := json.NewDecoder(reader).Decode(&addEvents)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "event json decoding failed", err)
	}
	return addEvents, nil
}

// Read reads and converts the request's JSON event data into an Event struct
func (jsonEventReader) ReadUpdateEventPushedByIdRequest(reader io.Reader) ([]dto.UpdateEventPushedByIdRequest, errors.EdgeX) {
	var requests []dto.UpdateEventPushedByIdRequest
	err := json.NewDecoder(reader).Decode(&requests)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "event json decoding failed", err)
	}
	return requests, nil
}
