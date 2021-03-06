// Copyright 2015 Matthew Collins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

type mode int

const (
	mScaled mode = iota
	mUnscaled
)

// AttachPoint is a side of the screen that an element
// can be attached to
type AttachPoint int

// Attachment points
//     VAlign, HAlign
const (
	Top, Left AttachPoint = iota, iota
	Middle, Center
	Bottom, Right
)
