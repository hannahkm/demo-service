// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js

package http

import _ "unsafe" // for linkname

// RoundTrip should be an internal detail,
// but widely used packages access it using linkname.
// Notable members of the hall of shame include:
//   - github.com/erda-project/erda-infra
//
// Do not remove or change the type signature.
// See go.dev/issue/67401.
//
//go:linkname badRoundTrip net/http.(*Transport).RoundTrip
func badRoundTrip(*Transport, *Request) (*Response, error)

// RoundTrip implements the [RoundTripper] interface.
//
// For higher-level HTTP client support (such as handling of cookies
// and redirects), see [Get], [Post], and the [Client] type.
//
// Like the RoundTripper interface, the error types returned
// by RoundTrip are unspecified.
func (t *Transport) RoundTrip(req *Request) (_unnamedRetVal0 *Response, _unnamedRetVal1 error) {
	//line <generated>:1
	if hookContext2011018656, _ := OtelBeforeTrampoline_RoundTrip2011018656(&t, &req); false {
	} else {
		defer OtelAfterTrampoline_RoundTrip2011018656(hookContext2011018656, &_unnamedRetVal0, &_unnamedRetVal1)
	}
	//line roundtrip.go:30:2
	if t == nil {
		panic("transport is nil")
	}
	//line roundtrip.go:33:2
	return t.roundTrip(req)
}

//line <generated>:1
type HookContextImpl2011018656 struct {
	params      []interface{}
	returnVals  []interface{}
	skipCall    bool
	data        interface{}
	funcName    string
	packageName string
}

func (c *HookContextImpl2011018656) SetSkipCall(skip bool)    { c.skipCall = skip }
func (c *HookContextImpl2011018656) IsSkipCall() bool         { return c.skipCall }
func (c *HookContextImpl2011018656) SetData(data interface{}) { c.data = data }
func (c *HookContextImpl2011018656) GetData() interface{}     { return c.data }
func (c *HookContextImpl2011018656) GetKeyData(key string) interface{} {
	if c.data == nil {
		return nil
	}
	return c.data.(map[string]interface{})[key]
}

func (c *HookContextImpl2011018656) SetKeyData(key string, val interface{}) {
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	c.data.(map[string]interface{})[key] = val
}

func (c *HookContextImpl2011018656) HasKeyData(key string) bool {
	if c.data == nil {
		return false
	}
	_, ok := c.data.(map[string]interface{})[key]
	return ok
}

func (c *HookContextImpl2011018656) GetParam(idx int) interface{} {
	switch idx {
	case 0:
		return *(c.params[0].(**Transport))
	case 1:
		return *(c.params[1].(**Request))
	}
	return nil
}

func (c *HookContextImpl2011018656) SetParam(idx int, val interface{}) {
	if val == nil {
		c.params[idx] = nil
		return
	}
	switch idx {
	case 0:
		*(c.params[0].(**Transport)) = val.(*Transport)
	case 1:
		*(c.params[1].(**Request)) = val.(*Request)
	}
}

func (c *HookContextImpl2011018656) GetReturnVal(idx int) interface{} {
	switch idx {
	case 0:
		return *(c.returnVals[0].(**Response))
	case 1:
		return *(c.returnVals[1].(*error))
	}
	return nil
}

func (c *HookContextImpl2011018656) SetReturnVal(idx int, val interface{}) {
	if val == nil {
		c.returnVals[idx] = nil
		return
	}
	switch idx {
	case 0:
		*(c.returnVals[0].(**Response)) = val.(*Response)
	case 1:
		*(c.returnVals[1].(*error)) = val.(error)
	}
}
func (c *HookContextImpl2011018656) GetParamCount() int     { return len(c.params) }
func (c *HookContextImpl2011018656) GetReturnValCount() int { return len(c.returnVals) }
func (c *HookContextImpl2011018656) GetFuncName() string    { return c.funcName }
func (c *HookContextImpl2011018656) GetPackageName() string { return c.packageName }

// Trampoline Template
func OtelBeforeTrampoline_RoundTrip2011018656(recv0 **Transport, param0 **Request) (hookContext *HookContextImpl2011018656, skipCall bool) {
	defer func() {
		if err := recover(); err != nil {
			println("failed to exec Before hook", "BeforeRoundTrip")
			if e, ok := err.(error); ok {
				println(e.Error())
			}
			fetchStack, printStack := OtelGetStackImpl, OtelPrintStackImpl
			if fetchStack != nil && printStack != nil {
				printStack(fetchStack())
			}
		}
	}()
	hookContext = &HookContextImpl2011018656{}
	hookContext.params = []interface{}{recv0, param0}
	hookContext.funcName = "RoundTrip"
	hookContext.packageName = "http"
	if BeforeRoundTrip != nil {
		BeforeRoundTrip(hookContext, *recv0, *param0)
	}
	return hookContext, hookContext.skipCall
}

func OtelAfterTrampoline_RoundTrip2011018656(hookContext HookContext, arg0 **Response, arg1 *error) {
	defer func() {
		if err := recover(); err != nil {
			println("failed to exec After hook", "AfterRoundTrip")
			if e, ok := err.(error); ok {
				println(e.Error())
			}
			fetchStack, printStack := OtelGetStackImpl, OtelPrintStackImpl
			if fetchStack != nil && printStack != nil {
				printStack(fetchStack())
			}
		}
	}()
	hookContext.(*HookContextImpl2011018656).returnVals = []interface{}{arg0, arg1}
	if AfterRoundTrip != nil {
		AfterRoundTrip(hookContext, *arg0, *arg1)
	}
}

//go:linkname BeforeRoundTrip github.com/DataDog/dd-trace-go/contrib/net/http/v2/otelc.BeforeRoundTrip
func BeforeRoundTrip(hookContext HookContext, recv0 *Transport, param0 *Request)

//go:linkname AfterRoundTrip github.com/DataDog/dd-trace-go/contrib/net/http/v2/otelc.AfterRoundTrip
func AfterRoundTrip(hookContext HookContext, arg0 *Response, arg1 error)
