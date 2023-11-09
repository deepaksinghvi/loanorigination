// The MIT License (MIT)
// 
// Copyright (c) 2021 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/api/v1/tasklist.proto

package apiv1

var yarpcFileDescriptorClosure216fa006947e00a0 = [][]byte{
	// uber/cadence/api/v1/tasklist.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0xdb, 0x36,
		0x14, 0x9e, 0xe2, 0xb4, 0x4b, 0xe8, 0xd9, 0xd5, 0xb8, 0xb5, 0x8d, 0xdd, 0x75, 0xf3, 0x74, 0x51,
		0x04, 0xc5, 0x26, 0xc1, 0x19, 0x76, 0xb5, 0x8b, 0xc1, 0xb1, 0x83, 0x55, 0xb0, 0xe3, 0x1a, 0x92,
		0x1a, 0x20, 0x03, 0x06, 0x8e, 0x12, 0x59, 0x9b, 0xd0, 0x0f, 0x05, 0x92, 0x72, 0xe2, 0x17, 0xd9,
		0xc3, 0xec, 0x89, 0xf6, 0x18, 0x03, 0x29, 0xd9, 0xf3, 0x12, 0x6f, 0x77, 0xe4, 0xf9, 0xce, 0x77,
		0x7e, 0x3e, 0x9e, 0x43, 0xe0, 0x54, 0x31, 0x15, 0x5e, 0x82, 0x09, 0x2d, 0x12, 0xea, 0xe1, 0x92,
		0x79, 0xeb, 0xa1, 0xa7, 0xb0, 0x4c, 0x33, 0x26, 0x95, 0x5b, 0x0a, 0xae, 0x38, 0xfc, 0x42, 0xfb,
		0xb8, 0x8d, 0x8f, 0x8b, 0x4b, 0xe6, 0xae, 0x87, 0xfd, 0xaf, 0x97, 0x9c, 0x2f, 0x33, 0xea, 0x19,
		0x97, 0xb8, 0xfa, 0xe8, 0x91, 0x4a, 0x60, 0xc5, 0x78, 0x51, 0x93, 0xfa, 0xdf, 0x3c, 0xc4, 0x15,
		0xcb, 0xa9, 0x54, 0x38, 0x2f, 0x1b, 0x87, 0x47, 0x01, 0xee, 0x04, 0x2e, 0x4b, 0x2a, 0x64, 0x8d,
		0x3b, 0x1f, 0xc0, 0x49, 0x84, 0x65, 0x3a, 0x63, 0x52, 0x41, 0x08, 0x8e, 0x0b, 0x9c, 0xd3, 0x33,
		0x6b, 0x60, 0x9d, 0x9f, 0x06, 0xe6, 0x0c, 0x7f, 0x04, 0xc7, 0x29, 0x2b, 0xc8, 0xd9, 0xd1, 0xc0,
		0x3a, 0xef, 0x5e, 0x7c, 0xeb, 0x1e, 0x28, 0xd2, 0xdd, 0x06, 0x98, 0xb2, 0x82, 0x04, 0xc6, 0xdd,
		0xc1, 0xc0, 0xde, 0x5a, 0xaf, 0xa9, 0xc2, 0x04, 0x2b, 0x0c, 0xaf, 0xc1, 0x97, 0x39, 0xbe, 0x47,
		0xba, 0x6d, 0x89, 0x4a, 0x2a, 0x90, 0xa4, 0x09, 0x2f, 0x88, 0x49, 0xd7, 0xbe, 0xf8, 0xca, 0xad,
		0x2b, 0x75, 0xb7, 0x95, 0xba, 0x13, 0x5e, 0xc5, 0x19, 0xbd, 0xc1, 0x59, 0x45, 0x83, 0xcf, 0x73,
		0x7c, 0xaf, 0x03, 0xca, 0x05, 0x15, 0xa1, 0xa1, 0x39, 0x1f, 0x40, 0x6f, 0x9b, 0x62, 0x81, 0x85,
		0x62, 0x5a, 0x95, 0x5d, 0x2e, 0x1b, 0xb4, 0x52, 0xba, 0x69, 0x3a, 0xd1, 0x47, 0xf8, 0x06, 0x3c,
		0xe3, 0x77, 0x05, 0x15, 0x68, 0xc5, 0xa5, 0x42, 0xa6, 0xcf, 0x23, 0x83, 0x76, 0x8c, 0xf9, 0x1d,
		0x97, 0x6a, 0x8e, 0x73, 0xea, 0xfc, 0x65, 0x81, 0xee, 0x36, 0x6e, 0xa8, 0xb0, 0xaa, 0x24, 0xfc,
		0x0e, 0xc0, 0x18, 0x27, 0x69, 0xc6, 0x97, 0x28, 0xe1, 0x55, 0xa1, 0xd0, 0x8a, 0x15, 0xca, 0xc4,
		0x6e, 0x05, 0x76, 0x83, 0x8c, 0x35, 0xf0, 0x8e, 0x15, 0x0a, 0xbe, 0x06, 0x40, 0x50, 0x4c, 0x50,
		0x46, 0xd7, 0x34, 0x33, 0x39, 0x5a, 0xc1, 0xa9, 0xb6, 0xcc, 0xb4, 0x01, 0xbe, 0x02, 0xa7, 0x38,
		0x49, 0x1b, 0xb4, 0x65, 0xd0, 0x13, 0x9c, 0xa4, 0x35, 0xf8, 0x06, 0x3c, 0x13, 0x58, 0xd1, 0x7d,
		0x75, 0x8e, 0x07, 0xd6, 0xb9, 0x15, 0x74, 0xb4, 0x79, 0xd7, 0x3b, 0x9c, 0x80, 0x8e, 0x96, 0x11,
		0x31, 0x82, 0xe2, 0x8c, 0x27, 0xe9, 0xd9, 0x13, 0xa3, 0xe1, 0xe0, 0x3f, 0x9f, 0xc7, 0x9f, 0x5c,
		0x6a, 0xbf, 0xa0, 0xad, 0x69, 0x3e, 0x31, 0x17, 0xe7, 0x67, 0xd0, 0xde, 0xc3, 0x60, 0x0f, 0x9c,
		0x48, 0x85, 0x85, 0x42, 0x8c, 0x34, 0xcd, 0x7d, 0x6a, 0xee, 0x3e, 0x81, 0xcf, 0xc1, 0x53, 0x5a,
		0x10, 0x0d, 0xd4, 0xfd, 0x3c, 0xa1, 0x05, 0xf1, 0x89, 0xf3, 0x87, 0x05, 0xc0, 0x82, 0x67, 0x19,
		0x15, 0x7e, 0xf1, 0x91, 0xc3, 0x09, 0xb0, 0x33, 0x2c, 0x15, 0xc2, 0x49, 0x42, 0xa5, 0x44, 0x7a,
		0x14, 0x9b, 0xc7, 0xed, 0x3f, 0x7a, 0xdc, 0x68, 0x3b, 0xa7, 0x41, 0x57, 0x73, 0x46, 0x86, 0xa2,
		0x8d, 0xb0, 0x0f, 0x4e, 0x18, 0xa1, 0x85, 0x62, 0x6a, 0xd3, 0xbc, 0xd0, 0xee, 0x7e, 0x48, 0x9f,
		0xd6, 0x01, 0x7d, 0x9c, 0x3f, 0x2d, 0xd0, 0x0b, 0x15, 0x4b, 0xd2, 0xcd, 0xd5, 0x3d, 0x4d, 0x2a,
		0x3d, 0x1a, 0x23, 0xa5, 0x04, 0x8b, 0x2b, 0x45, 0x25, 0xfc, 0x05, 0xd8, 0x77, 0x5c, 0xa4, 0x54,
		0x98, 0x59, 0x44, 0x7a, 0x07, 0x9b, 0x3a, 0x5f, 0xff, 0xef, 0x7c, 0x07, 0xdd, 0x9a, 0xb6, 0x5b,
		0x98, 0x08, 0xf4, 0x64, 0xb2, 0xa2, 0xa4, 0xca, 0x28, 0x52, 0x1c, 0xd5, 0xea, 0xe9, 0xb6, 0x79,
		0xa5, 0x4c, 0xed, 0xed, 0x8b, 0xde, 0xe3, 0xb1, 0x6e, 0x36, 0x38, 0x78, 0xb1, 0xe5, 0x46, 0x3c,
		0xd4, 0xcc, 0xa8, 0x26, 0xbe, 0xfd, 0x1d, 0x7c, 0xb6, 0xbf, 0x51, 0xb0, 0x0f, 0x5e, 0x44, 0xa3,
		0x70, 0x8a, 0x66, 0x7e, 0x18, 0xa1, 0xa9, 0x3f, 0x9f, 0x20, 0x7f, 0x7e, 0x33, 0x9a, 0xf9, 0x13,
		0xfb, 0x13, 0xd8, 0x03, 0xcf, 0x1f, 0x60, 0xf3, 0xf7, 0xc1, 0xf5, 0x68, 0x66, 0x5b, 0x07, 0xa0,
		0x30, 0xf2, 0xc7, 0xd3, 0x5b, 0xfb, 0xe8, 0x2d, 0xf9, 0x27, 0x43, 0xb4, 0x29, 0xe9, 0xbf, 0x33,
		0x44, 0xb7, 0x8b, 0xab, 0xbd, 0x0c, 0xaf, 0xc0, 0xcb, 0x07, 0xd8, 0xe4, 0x6a, 0xec, 0x87, 0xfe,
		0xfb, 0xb9, 0x6d, 0x1d, 0x00, 0x47, 0xe3, 0xc8, 0xbf, 0xf1, 0xa3, 0x5b, 0xfb, 0xe8, 0xf2, 0x37,
		0xf0, 0x32, 0xe1, 0xf9, 0x21, 0x45, 0x2f, 0x3b, 0xbb, 0xcd, 0xd5, 0xaa, 0x2c, 0xac, 0x5f, 0x87,
		0x4b, 0xa6, 0x56, 0x55, 0xec, 0x26, 0x3c, 0xf7, 0xf6, 0xff, 0xca, 0xef, 0x19, 0xc9, 0xbc, 0x25,
		0xaf, 0xbf, 0xaf, 0xe6, 0xe3, 0xfc, 0x09, 0x97, 0x6c, 0x3d, 0x8c, 0x9f, 0x1a, 0xdb, 0x0f, 0x7f,
		0x07, 0x00, 0x00, 0xff, 0xff, 0x41, 0xb6, 0x75, 0xa3, 0x5c, 0x05, 0x00, 0x00,
	},
	// google/protobuf/duration.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
		0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
		0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
		0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
		0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xd4, 0xcc, 0xc8, 0x25, 0x9c,
		0x9c, 0x9f, 0xab, 0x87, 0x66, 0xa6, 0x13, 0x2f, 0xcc, 0xc4, 0x00, 0x90, 0x48, 0x00, 0x63, 0x94,
		0x21, 0x54, 0x45, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc2, 0x81, 0x25,
		0x95, 0x05, 0xa9, 0xc5, 0xfa, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x70, 0xc7, 0x16, 0x24, 0xfd, 0x60,
		0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0xa2, 0x39, 0x00, 0xaa,
		0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4, 0x35, 0x89, 0x0d, 0x6c, 0x94,
		0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xb4, 0xc3, 0xfb, 0x00, 0x00, 0x00,
	},
	// google/protobuf/timestamp.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d,
		0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x03, 0x0b, 0x09, 0xf1, 0x43, 0x14, 0xe8, 0xc1, 0x14, 0x28,
		0x59, 0x73, 0x71, 0x86, 0xc0, 0xd4, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5,
		0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x79, 0x89,
		0x79, 0xf9, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x53, 0x2b, 0x23, 0x97,
		0x70, 0x72, 0x7e, 0xae, 0x1e, 0x9a, 0xa1, 0x4e, 0x7c, 0x70, 0x23, 0x03, 0x40, 0x42, 0x01, 0x8c,
		0x51, 0x46, 0x50, 0x25, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0x48, 0x6e,
		0xac, 0x2c, 0x48, 0x2d, 0xd6, 0xcf, 0xce, 0xcb, 0x2f, 0xcf, 0x43, 0xb8, 0xb7, 0x20, 0xe9, 0x07,
		0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x00, 0xa7, 0x55, 0x4c, 0x72, 0xee, 0x10, 0xdd, 0x01, 0x50,
		0x2d, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x20, 0x0d, 0x21, 0x20, 0xbd, 0x49, 0x6c, 0x60, 0xb3,
		0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x65, 0xce, 0x7d, 0xff, 0x00, 0x00, 0x00,
	},
	// google/protobuf/wrappers.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2f, 0x4a, 0x2c,
		0x28, 0x48, 0x2d, 0x2a, 0xd6, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0xca,
		0x5c, 0xdc, 0x2e, 0xf9, 0xa5, 0x49, 0x39, 0xa9, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x42, 0x22, 0x5c,
		0xac, 0x65, 0x20, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x84, 0xa3, 0xa4, 0xc4, 0xc5,
		0xe5, 0x96, 0x93, 0x9f, 0x58, 0x82, 0x45, 0x0d, 0x13, 0x92, 0x1a, 0xcf, 0xbc, 0x12, 0x33, 0x13,
		0x2c, 0x6a, 0x98, 0x61, 0x6a, 0x94, 0xb9, 0xb8, 0x43, 0x71, 0x29, 0x62, 0x41, 0x35, 0xc8, 0xd8,
		0x08, 0x8b, 0x1a, 0x56, 0x34, 0x83, 0xb0, 0x2a, 0xe2, 0x85, 0x29, 0x52, 0xe4, 0xe2, 0x74, 0xca,
		0xcf, 0xcf, 0xc1, 0xa2, 0x84, 0x03, 0xc9, 0x9c, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x2c, 0x8a,
		0x38, 0x91, 0x1c, 0xe4, 0x54, 0x59, 0x92, 0x5a, 0x8c, 0x45, 0x0d, 0x0f, 0x54, 0x8d, 0x53, 0x33,
		0x23, 0x97, 0x70, 0x72, 0x7e, 0xae, 0x1e, 0x5a, 0xf0, 0x3a, 0xf1, 0x86, 0x43, 0xc3, 0x3f, 0x00,
		0x24, 0x12, 0xc0, 0x18, 0x65, 0x08, 0x55, 0x91, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f,
		0x94, 0x8e, 0x88, 0xab, 0x92, 0xca, 0x82, 0xd4, 0x62, 0xfd, 0xec, 0xbc, 0xfc, 0xf2, 0x3c, 0x78,
		0xbc, 0x15, 0x24, 0xfd, 0x60, 0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce,
		0x1d, 0xa2, 0x39, 0x00, 0xaa, 0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4,
		0x35, 0x89, 0x0d, 0x6c, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x92, 0x48, 0x30, 0x06,
		0x02, 0x00, 0x00,
	},
}