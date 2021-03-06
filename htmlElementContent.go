package telerik

import (
	"bytes"
)

type HtmlContent struct {
	/*
	  Content
	*/
	Content Content `htmlAttr:"-" jsonSchema_description:""`

	*ToJavaScriptConverter `htmlAttr:"-" jsonSchema_description:""`
}

func (el *HtmlContent) ToHtml() []byte {
	var buffer bytes.Buffer

	buffer.Write(el.Content.ToHtml())
	return buffer.Bytes()
}
func (el *HtmlContent) ToJavaScript() []byte {
	var buffer bytes.Buffer

	buffer.Write(el.Content.ToJavaScript())
	return buffer.Bytes()
}
