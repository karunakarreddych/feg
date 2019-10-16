/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package diameter

import (
	"sync"
)

// RequestTracker stores a mapping of keys to channels and is intended to be
// used to store unique identifiers of requests and channels to send answers to
// after they are received. The methods are thread safe and do not close the
// channel after use
type RequestTracker struct {
	mapMutex   sync.Mutex
	requestMap map[interface{}]chan interface{}
}

func NewRequestTracker() *RequestTracker {
	return &RequestTracker{requestMap: make(map[interface{}]chan interface{})}
}

// RegisterRequest tracks a request in the map and returns the channel stored
func (rt *RequestTracker) RegisterRequest(key interface{}, ch chan interface{}) chan interface{} {
	rt.mapMutex.Lock()
	defer rt.mapMutex.Unlock()
	rt.requestMap[key] = ch
	return ch
}

// DeregisterRequest finds the channel in the map, removes and returns it. If no
// channel is found, nil is returned
func (rt *RequestTracker) DeregisterRequest(key interface{}) chan interface{} {
	rt.mapMutex.Lock()
	defer rt.mapMutex.Unlock()
	channel, ok := rt.requestMap[key]
	if !ok {
		return nil
	}
	delete(rt.requestMap, key)
	return channel
}
