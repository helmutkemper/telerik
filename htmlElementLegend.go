package telerik

import (
  "bytes"
  "reflect"
)

// The HTML <legend> element represents a caption for the content of its parent <fieldset>.
type HtmlElementFormLegend struct{
  /*
  The name of the control, which is submitted with the form data.
  */
  Name                        string                      `htmlAttr:"name"`

  /*
  Content inside html tag
  */
  Content                     Content                      `htmlAttr:"-"`

  /*
  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
  just as descendants of their form elements. An input can only be associated with one form.
  */
  Form                        string                      `htmlAttr:"form"`

  /*
  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
  the autocomplete attribute to control this feature.
  */
  Disabled                    Boolean                     `htmlAttrSet:"disabled"`

  Global                      HtmlGlobalAttributes        `htmlAttr:"-"`

  *ToJavaScriptConverter                                  `htmlAttr:"-"`
}
func(el *HtmlElementFormLegend)ToHtml() []byte {
  var buffer bytes.Buffer

  element := reflect.ValueOf(el).Elem()
  data := el.ToJavaScriptConverter.ToTelerikHtml(element)

  buffer.Write( []byte( `<legend` ) )
  buffer.Write( el.Global.ToHtml() )
  buffer.Write( data )
  buffer.Write( []byte( `>` ) )
  buffer.Write( el.Content.ToHtml() )
  buffer.Write( []byte( `</legend>` ) )

  return buffer.Bytes()
}