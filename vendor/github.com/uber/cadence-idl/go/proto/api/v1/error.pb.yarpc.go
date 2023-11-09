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
// source: uber/cadence/api/v1/error.proto

package apiv1

var yarpcFileDescriptorClosurec8f91786c9aff272 = [][]byte{
	// uber/cadence/api/v1/error.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x6f, 0xd3, 0x40,
		0x10, 0xc5, 0x95, 0x22, 0x22, 0x31, 0x85, 0xb6, 0x98, 0x50, 0x2a, 0x84, 0x5a, 0x6a, 0xfe, 0x34,
		0x97, 0xda, 0x8a, 0x38, 0xf6, 0xd4, 0x84, 0x44, 0x8a, 0x84, 0xaa, 0xd2, 0x08, 0x90, 0x7a, 0x89,
		0x36, 0xeb, 0x49, 0x18, 0x75, 0xbd, 0x6b, 0xf6, 0x8f, 0x89, 0x2f, 0x7c, 0x0f, 0xbe, 0x2d, 0xb2,
		0x77, 0x5d, 0x72, 0x08, 0x12, 0xc7, 0xfd, 0xcd, 0xdb, 0xe7, 0xb7, 0xe3, 0x19, 0x38, 0x71, 0x0b,
		0xd4, 0x29, 0x67, 0x19, 0x4a, 0x8e, 0x29, 0x2b, 0x28, 0x2d, 0x07, 0x29, 0x6a, 0xad, 0x74, 0x52,
		0x68, 0x65, 0x55, 0xf4, 0xac, 0x16, 0x24, 0x41, 0x90, 0xb0, 0x82, 0x92, 0x72, 0x10, 0xaf, 0xe0,
		0xed, 0x37, 0xa5, 0xef, 0x96, 0x42, 0xfd, 0x1c, 0xaf, 0x91, 0x3b, 0x4b, 0x4a, 0x5e, 0x0a, 0x8d,
		0x2c, 0xab, 0x66, 0x96, 0x69, 0x8b, 0xd9, 0xb8, 0xb6, 0x88, 0xfa, 0x70, 0x60, 0xea, 0xf3, 0x5c,
		0xe3, 0x0f, 0x87, 0xc6, 0xce, 0x29, 0x3b, 0xea, 0xbc, 0xee, 0xf4, 0x1f, 0xdd, 0xec, 0x35, 0xfc,
		0xc6, 0xe3, 0x69, 0x16, 0x3d, 0x87, 0xae, 0x76, 0xb2, 0xae, 0xef, 0x34, 0xf5, 0x87, 0xda, 0xc9,
		0x69, 0x16, 0x2f, 0xa1, 0x37, 0x96, 0x96, 0x6c, 0x75, 0xa5, 0xec, 0x78, 0x4d, 0xc6, 0x1a, 0x6f,
		0x7c, 0x06, 0xfb, 0xdc, 0x69, 0x8d, 0xd2, 0xce, 0xb9, 0x70, 0xc6, 0xa2, 0x6e, 0x7d, 0x03, 0x1e,
		0x79, 0x1a, 0xbd, 0x83, 0x3d, 0xc6, 0x2d, 0x95, 0x78, 0xaf, 0xf3, 0xfe, 0x4f, 0x3c, 0x0d, 0xb2,
		0xb8, 0x0f, 0xef, 0xff, 0xf5, 0xa0, 0x91, 0xca, 0x0b, 0x81, 0xed, 0x93, 0xe2, 0x5f, 0xd0, 0xfb,
		0xa8, 0x72, 0x46, 0xf2, 0x4a, 0xd9, 0xcb, 0xc6, 0xc3, 0x27, 0x3a, 0x84, 0x6e, 0xd6, 0xf0, 0x10,
		0x24, 0x9c, 0xb6, 0x25, 0xdd, 0xf9, 0xcf, 0xa4, 0x0f, 0xb6, 0x25, 0xfd, 0xdd, 0x81, 0xe3, 0x91,
		0x20, 0x94, 0xf6, 0x2b, 0x6a, 0x43, 0xaa, 0xce, 0x31, 0x73, 0x45, 0xa1, 0xfe, 0x76, 0xfd, 0x0c,
		0xf6, 0x97, 0xc8, 0xac, 0xd3, 0x38, 0x2f, 0xbd, 0xa6, 0x6d, 0x4e, 0xc0, 0xe1, 0x66, 0x74, 0x02,
		0xbb, 0xbc, 0xb1, 0x9a, 0x53, 0x5e, 0x88, 0x90, 0x0b, 0x3c, 0x9a, 0xe6, 0x85, 0x88, 0xce, 0x21,
		0x32, 0xad, 0x77, 0xeb, 0x65, 0x42, 0xae, 0xa7, 0xf7, 0x95, 0x60, 0x67, 0xe2, 0x0b, 0x38, 0x9c,
		0xf8, 0x2f, 0xd4, 0xbf, 0x4b, 0xb2, 0x85, 0x68, 0x23, 0x9d, 0xc2, 0xe3, 0x36, 0xd2, 0x52, 0xb0,
		0x55, 0xc8, 0xb3, 0x1b, 0xd8, 0x44, 0xb0, 0x55, 0xfc, 0x06, 0x4e, 0x47, 0x4c, 0x72, 0x14, 0x82,
		0x6d, 0x74, 0x3f, 0x4c, 0x48, 0xdb, 0xfd, 0x97, 0x70, 0xe4, 0xbb, 0x1f, 0xca, 0x1b, 0x33, 0x11,
		0xf7, 0x20, 0xfa, 0x44, 0x39, 0xd9, 0xf1, 0x9a, 0x23, 0x66, 0xed, 0x8d, 0x08, 0x0e, 0x3e, 0x3b,
		0xd4, 0xd5, 0x84, 0x91, 0xd8, 0x60, 0x33, 0xd4, 0x25, 0x71, 0x1c, 0x3a, 0x53, 0x79, 0x76, 0x0c,
		0xaf, 0x66, 0x96, 0xf8, 0x5d, 0x55, 0xcf, 0x01, 0xea, 0x2f, 0x92, 0x95, 0x8c, 0x44, 0xfd, 0x84,
		0xa6, 0x3e, 0xbc, 0x85, 0x17, 0x5c, 0xe5, 0xc9, 0x96, 0x6d, 0x18, 0x42, 0xa3, 0xb8, 0xae, 0xd7,
		0xe5, 0xba, 0x73, 0x3b, 0x58, 0x91, 0xfd, 0xee, 0x16, 0x09, 0x57, 0x79, 0xba, 0xb9, 0x5c, 0xe7,
		0x94, 0x89, 0x74, 0xa5, 0xd2, 0x66, 0xad, 0xc2, 0xa6, 0x5d, 0xb0, 0x82, 0xca, 0xc1, 0xa2, 0xdb,
		0xb0, 0x0f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xcb, 0xa2, 0x1d, 0x8d, 0x03, 0x00, 0x00,
	},
}
