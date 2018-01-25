package telerik

type HtmlInputHidden struct{
  /*
  The name of the control, which is submitted with the form data.
  @see typeNamesForAutocomplete.go
  Ex.: const NAMES_FOR_AUTOCOMPLETE_NAME
  */
  Name                        String

  /*
  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
  checkbox.
  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
  changed before the reload.
  */
  Value                       String

  /*
  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
  just as descendants of their form elements. An input can only be associated with one form.
  */
  Form                        String

  /*
  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
  the autocomplete attribute to control this feature.
  */
  Disabled                    Boolean

  /*
  This attribute indicates whether the value of the control can be automatically completed by the browser.
  Possible values are:
  off: The user must explicitly enter a value into this field for every use, or the document provides its own
  auto-completion method. The browser does not automatically complete the entry.
  on: The browser is allowed to automatically complete the value based on values that the user has entered during
  previous uses, however on does not provide any further information about what kind of data the user might be expected
  to enter.
  @see typeNamesForAutocomplete.go
  */
  AutoComplete                Boolean

  Global                      HtmlGlobalAttributes
}
func(el *HtmlInputHidden)String() string {
  return `<input ` + el.Global.String() + ` type="hidden" ` + el.Name.ToAttr("name") + el.Value.ToAttr("value") + el.AutoComplete.ToAttrOnOff("autocomplete") + `>`
}