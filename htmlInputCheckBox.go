package telerik

// <input> elements of type checkbox are rendered by default as square boxes that are checked (ticked) when activated,
// like you might see in an official government paper form. They allow you to select single values for submission in a
// form (or not).
//
// <input id="checkBox" type="checkbox">
type HtmlInputCheckBox struct{
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
  When the value of the type attribute is radio or checkbox, the presence of this Boolean attribute indicates that the
  control is selected by default, otherwise it is ignored.
  Unlike other browsers, Firefox will by default persist the dynamic checked state of an <input> across page loads. Use
  the autocomplete attribute to control this feature.
  */
  Checked                     Boolean

  Global                      HtmlGlobalAttributes
}
func(el *HtmlInputCheckBox)String() string {
  return `<input ` + el.Global.String() + ` type="checkbox" ` + el.Name.ToAttr("name") + el.Value.ToAttr("value") + el.Checked.ToAttrSet("checked") + `>`
}