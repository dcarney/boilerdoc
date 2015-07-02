// The MIT License (MIT)
//
// Copyright (c) 2015 Dylan Carney
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

package main

const (

	// the template for Markdown output.  Unfortunately, there's no escaping a backtick(`) inside backticks in Go,
	// so we have to use concatenation to get backtick literals:
	// e.g. `some string` + "`" + `another string` + "`"
	markdownTemplate = `
### {{ .Method }} [{{ .URLTemplate }}]

{{ .Description }}

{{ if .Notes }}
**NOTE:** {{ .Notes }}
{{end}}

{{ if .URLParams }}
#### Parameters
  {{ range $param := .URLParams }}
  * {{ $param.Name }} ({{ if $param.Required }}required {{ end }}{{ if $param.Type }}{{ $param.Type }}{{ end }}) : {{ $param.Description }}
  {{ end }}
{{ end }}

{{ if .DataParams }}
#### Request Body Parameters
  {{ range $param := .DataParams }}
  * {{ $param.Name }} ({{ if $param.Required }}required {{ end }}{{ if $param.Type }}{{ $param.Type }}{{ end }}) : {{ $param.Description }}
  {{ end }}
{{ end }}

#### Example success response
` + "`" + `{{ .SuccessResponse.Code }}` + "`" + `: {{ statusText .SuccessResponse.Code }}

    {{ .SuccessResponse.Content }}

{{ if .ErrorResponses }}
#### Example error responses
  {{ range $resp := .ErrorResponses }}
  ` + "`" + `{{ $resp.Code }}` + "`" + `: {{ statusText $resp.Code }}

    {{ $resp.Content }}
  {{ end }}
{{ end }}

{{ if .Examples }}
#### Examples
  {{ range $call := .Examples }}
    {{ $call }}
  {{ end }}
{{ end }}
`
)
