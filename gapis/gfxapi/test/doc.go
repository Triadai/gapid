// Copyright (C) 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package test is the integration test suite for the api compiler and templates.
package test

// The following are the imports that generated source files pull in when present
// Having these here helps out tools that can't cope with missing dependancies
import (
	_ "github.com/google/gapid/core/data/id"
	_ "github.com/google/gapid/core/data/pod"
	_ "github.com/google/gapid/core/log"
	_ "github.com/google/gapid/core/math/u64"
	_ "github.com/google/gapid/framework/binary"
	_ "github.com/google/gapid/framework/binary/registry"
	_ "github.com/google/gapid/framework/binary/schema"
	_ "github.com/google/gapid/gapis/atom"
	_ "github.com/google/gapid/gapis/atom/atom_pb"
	_ "github.com/google/gapid/gapis/capture"
	_ "github.com/google/gapid/gapis/gfxapi"
	_ "github.com/google/gapid/gapis/gfxapi/test/gfxapi_test_import"
	_ "github.com/google/gapid/gapis/gfxapi/test/test_pb"
	_ "github.com/google/gapid/gapis/memory"
	_ "github.com/google/gapid/gapis/memory/memory_pb"
	_ "github.com/google/gapid/gapis/replay/builder"
	_ "github.com/google/gapid/gapis/replay/protocol"
	_ "github.com/google/gapid/gapis/replay/value"
	_ "github.com/google/gapid/gapis/service/path"
)