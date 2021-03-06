// Copyright 2018 The Go Cloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mempubsub

import (
	"context"
	"errors"
	"testing"
	"time"

	"gocloud.dev/pubsub/driver"
	"gocloud.dev/pubsub/drivertest"
)

type harness struct{}

func newHarness(ctx context.Context, t *testing.T) (drivertest.Harness, error) {
	return &harness{}, nil
}

func (h *harness) MakeTopic(ctx context.Context) (driver.Topic, error) {
	return &topic{}, nil
}

func (h *harness) MakeNonexistentTopic(ctx context.Context) (driver.Topic, error) {
	// A nil *topic behaves like a nonexistent topic.
	return (*topic)(nil), nil
}

func (h *harness) MakeSubscription(ctx context.Context, dt driver.Topic, n int) (driver.Subscription, error) {
	if n < 0 || n >= 2 {
		return nil, errors.New("n must be 0 or 1")
	}
	return newSubscription(dt.(*topic), time.Second), nil
}

func (h *harness) MakeNonexistentSubscription(ctx context.Context) (driver.Subscription, error) {
	return newSubscription(nil, time.Second), nil
}

func (h *harness) Close() {}

func TestConformance(t *testing.T) {
	drivertest.RunConformanceTests(t, newHarness, nil)
}
