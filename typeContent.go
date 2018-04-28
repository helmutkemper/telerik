package telerik

import (
  //"fmt"
  "fmt"
  "bytes"
  "sort"
  "reflect"
)

type Content []interface{}

func(el Content) ToHtmlSupport() []byte {
  var buffer bytes.Buffer

  keys := make([]int, 0)
  for k := range el {
    keys = append(keys, k)
  }
  sort.Ints(keys)
  for _, k := range keys {
    switch outConverted := el[k].(type) {
    case KendoUiMultiSelect:
      buffer.Write( outConverted.ToHtml() )

      if !reflect.DeepEqual( outConverted.Dialog, KendoUiDialog{} ) {
        buffer.Write( outConverted.Dialog.ToHtml() )
      }
    }
  }

  return buffer.Bytes()
}

func(el Content) ToHtml() []byte {
  var buffer bytes.Buffer

  keys := make([]int, 0)
  for k := range el {
    keys = append(keys, k)
  }
  sort.Ints(keys)
  for _, k := range keys {
    switch outConverted := el[k].(type) {
    case string:
      buffer.WriteString( outConverted )
    case HtmlElementDiv:
      buffer.Write( outConverted.ToHtml() )
    case HtmlInputGeneric:
      buffer.Write( outConverted.ToHtml() )
    case HtmlInputNumber:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementSpan:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormLabel:
      buffer.Write( outConverted.ToHtml() )
    case HtmlInputText:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormLegend:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormFieldSet:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementForm:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormSelect:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormTextArea:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormMeter:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormButton:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormDataList:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiAutoComplete:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiMobileSwitch:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiNumericTextBox:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiComboBox:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiDialog:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementScript:
      buffer.Write( outConverted.ToHtml() )
    case KendoUiMultiSelect:
      buffer.Write( outConverted.ToHtml() )

      if !reflect.DeepEqual( outConverted.Dialog, KendoUiDialog{} ) {
        buffer.Write( outConverted.Dialog.ToHtml() )
      }

    case HtmlScriptType:
      buffer.WriteString( outConverted.String() )
    case KendoDataSource:
      //do noting

    case HtmlInputCheckBox:
      buffer.Write( outConverted.ToHtml() )

    /*case HtmlElementFormOutput:
      buffer.Write( outConverted.ToHtml() )
    case HtmlElementFormProgress:
      buffer.Write( outConverted.ToHtml() )*/

    default:
      fmt.Printf("\n\n\nToHtml() - typeContent.go: error: type %T not found in typeContent\n\n\n", outConverted)
    }
  }

  return buffer.Bytes()
}
func(el *Content) ToJavaScript() []byte {
  var buffer bytes.Buffer
      //buffer.WriteString( "\n" )

  keys := make([]int, 0)
  for k := range *el {
    keys = append(keys, k)
  }
  sort.Ints(keys)
  for _, k := range keys {
    switch outConverted := (*el)[k].(type) {
    case string:
      buffer.WriteString( outConverted )
      //buffer.WriteString( "\n" )
    case KendoUiDialog:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDateTimePicker:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDatePicker:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDateInput:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiContextMenu:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiConfirm:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiComboBox:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiColorPicker:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiColorPalette:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiCalendar:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiButton:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiMultiSelect:
      buffer.Write( outConverted.ToJavaScript() )

      if !reflect.DeepEqual( outConverted.Dialog, KendoUiDialog{} ) {
        buffer.Write( outConverted.Dialog.ToJavaScript() )
      }

      //buffer.WriteString( "\n" )
    case KendoUiAutoComplete:
      buffer.Write( outConverted.ToJavaScript() )
    case KendoUiMobileSwitch:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDraggable:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDropDownList:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDropTarget:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiDropTargetArea:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementDiv:
      buffer.Write( outConverted.ToJavaScript() )
    //  //buffer.WriteString( "\n" )
    case HtmlElementFormLabel:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoUiNumericTextBox:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementFormSelect:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementScript:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementFormButton:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case HtmlElementSpan:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    case KendoDataSource:
      buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )
    //case KendoSchema:
      //buffer.Write( outConverted.ToJavaScript() )
      //buffer.WriteString( "\n" )

    case HtmlElementFormTextArea:
    case HtmlInputCheckBox:
    case HtmlInputColor:
    case HtmlInputDate:
    case HtmlInputDateTimeLocal:
    case HtmlInputEmail:
      buffer.Write( outConverted.ToJavaScript() )
    case HtmlInputFile:
    case HtmlInputGeneric:
      buffer.Write( outConverted.ToJavaScript() )
    case HtmlInputHidden:
    case HtmlInputImage:
    case HtmlInputMonth:
    case HtmlInputNumber:
    case HtmlInputPassword:
      buffer.Write( outConverted.ToJavaScript() )
    case HtmlInputRadio:
    case HtmlInputRange:
    case HtmlInputSearch:
      buffer.Write( outConverted.ToJavaScript() )
    case HtmlInputTel:
      buffer.Write( outConverted.ToJavaScript() )
    case HtmlInputText:
      buffer.Write( outConverted.ToJavaScript() )
    case HtmlInputTime:
    case HtmlInputUrl:
      buffer.Write( outConverted.ToJavaScript() )
    case HtmlElementForm:
      buffer.Write( outConverted.ToJavaScript() )
    case HtmlInputWeek:

    default:
      fmt.Printf("\n\n\nToJavaScript() - typeContent.go: error: type %T not found in typeContent\n\n\n", outConverted)
    }
  }

  return buffer.Bytes()
}
func(el *Content)FilterFormElements() []interface{} {

  var contentProcessedList = make( []interface{}, 0 )
  var contentUnprocessedList = make( []interface{}, 0 )
  var contentFoundList = make( []interface{}, 0 )

  el.processContent( &contentProcessedList, &contentUnprocessedList, &contentFoundList, *el)
  for {
    if len( contentUnprocessedList ) == 0 {
      break
    }

    popElement := contentUnprocessedList[0]
    contentUnprocessedList = contentUnprocessedList[1:]

    el.processContent(&contentProcessedList, &contentUnprocessedList, &contentFoundList, popElement.(Content))
  }

  /*for _, v := range contentFoundList{
    fmt.Printf( "%T\n", v )
  }*/

  return contentFoundList
}
func(el *Content)processContent( contentProcessedList, contentUnprocessedList, contentFoundList *[]interface{}, content Content ) {
  *contentProcessedList  =append( *contentProcessedList, content )

  for _, contentElement := range content {
    el.addToUnprocessedList(contentUnprocessedList, contentFoundList, contentElement)
  }
}
func(el *Content)addToUnprocessedList( contentUnprocessedList, contentFoundList *[]interface{}, content interface{} ) {
  switch converted := content.(type) {
  case *Content:
  case HtmlElementSpan:
    *contentUnprocessedList  =append( *contentUnprocessedList, converted.Content )
  case HtmlElementDiv:
    *contentUnprocessedList  =append( *contentUnprocessedList, converted.Content )

    // Elementos de formulário que necessitam de javascript - início

  case KendoUiNumericTextBox:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiComboBox:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiAutoComplete:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiButton:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiCalendar:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiColorPalette:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiColorPicker:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiDateInput:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiDatePicker:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiDateTimePicker:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiDropDownList:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiMultiSelect:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoUiMobileSwitch:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlElementFormSelect:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlElementFormTextArea:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputCheckBox:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputColor:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputDate:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputDateTimeLocal:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputEmail:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputFile:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputGeneric:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputHidden:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputImage:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputMonth:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputNumber:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputPassword:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputRadio:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputRange:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputSearch:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputTel:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputText:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputTime:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputUrl:
    *contentFoundList        =append( *contentFoundList, &converted )
  case HtmlInputWeek:
    *contentFoundList        =append( *contentFoundList, &converted )
  case KendoDataSource:
    *contentFoundList        =append( *contentFoundList, &converted )

    // Elementos de formulário que necessitam de javascript - fim

  case HtmlElementFormLabel:
  default:
    fmt.Printf( "HtmlElementForm.addToUnprocessedList().type: %T\n", converted )
  }
}
func (el *Content)MakeJsObject() []byte {
  var pass = false
  var buffer bytes.Buffer
  var key, jsCode []byte
  var formElements = el.FilterFormElements()
  // fixme: mfalta um getName() ou algo parecido
  // fixme: KendoUiCalendar e KendoUiColorPalette devem ter name como obrigatórios

  for _, v := range formElements {
    switch converted := v.(type) {
    case *KendoDataSource:
      buffer.Write( []byte( "var " ) )
      buffer.Write( []byte( converted.VarName ) )
      buffer.Write( []byte( ";\n" ) )
    default: continue
    }
  }

  buffer.Write( []byte( "function addNewItemToKendoDataSource( id ){\n" ) )
  buffer.Write( []byte( "  console.log('id: ', id);\n" ) )
  buffer.Write( []byte( "  switch( id ){\n" ) )
  for _, v := range formElements {
    switch converted := v.(type) {
    case *KendoUiMultiSelect:
      if reflect.DeepEqual( converted.Dialog, KendoUiDialog{} ) {
        continue
      }

      key = []byte( converted.Html.GetId() )
      jsCode = []byte( `$('#` + string( converted.Dialog.GetId() ) + `').data('kendoDialog').open()` )
      buffer.Write( []byte( "    case 'id:" ) )
      buffer.Write( key )
      buffer.Write( []byte( "': " ) )
      buffer.Write( jsCode )
      buffer.Write( []byte( ";\n" ) )

      switch convertedFromInterface := converted.Dialog.Content.(type) {
      case Content:
        buffer.Write( []byte( "      " ) )
        buffer.Write( convertedFromInterface.ToJavaScript() )
        buffer.Write( []byte( "      break;\n" ) )
      }


      //buffer.Write( []byte( ";\n" ) )

    default: continue
    }
  }
  buffer.Write( []byte( "  }\n" ) )
  buffer.Write( []byte( "}\n" ) )



  buffer.Write( []byte( "function getFormValue( id ){\n" ) )
  buffer.Write( []byte( "  switch( id ){\n" ) )
  for _, v := range formElements {
    pass = false
    switch converted := v.(type) {
    case *KendoUiAutoComplete:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoAutoComplete').value()` )
    case *KendoUiButton:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoButton').value()` )
    case *KendoUiCalendar:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoCalendar').value()` )
    case *KendoUiColorPalette:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoColorPalette').value()` )
    case *KendoUiColorPicker:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoColorPicker').value()` )
    case *KendoUiComboBox:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoComboBox').value()` )
    case *KendoUiNumericTextBox:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoNumericTextBox').value()` )
    case *KendoUiDateInput:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoDateInput').value()` )
    case *KendoUiDatePicker:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoDatePicker').value()` )
    case *KendoUiDateTimePicker:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoDateTimePicker').value()` )
    case *KendoUiDropDownList:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoDropDownList').value()` )
    case *KendoUiMultiSelect: // fixme: colocar o conteúdo da janela aqui também
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoMultiSelect').value()` )
    case *HtmlElementFormSelect:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlElementFormTextArea:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputCheckBox:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputColor:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputDate:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputDateTimeLocal:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputEmail:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputFile:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputGeneric:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputHidden:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputImage:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputMonth:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputNumber:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputPassword:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputRadio:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputRange:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputSearch:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputTel:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputText:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputTime:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputUrl:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *HtmlInputWeek:
      pass = true
      key = []byte( converted.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )
    case *KendoUiMobileSwitch:
      pass = true
      key = []byte( converted.Html.Name )
      jsCode = []byte( `$('#` + string( converted.GetId() ) + `').value()` )

    case *KendoDataSource:
    default:
      fmt.Printf( "\n\n\ngetFormValue.id.type.%T.not.found.\n\n\n", converted )
    }

    if pass == true {
      buffer.Write([]byte( "    case 'id:" ))
      buffer.Write(key)
      buffer.Write([]byte( "': return " ))
      buffer.Write(jsCode)
      buffer.Write([]byte( ";\n" ))
    }
  }
  buffer.Write( []byte( "  }\n" ) )
  buffer.Write( []byte( "}\n" ) )

  return buffer.Bytes()
}
func (el *Content)MakeJsScript() []byte {
  var buffer bytes.Buffer

  var elementList = el.FilterFormElements()
  for k, v := range elementList {
    switch converted := v.(type) {
    case *KendoUiMultiSelect:

      elementList[k].(*KendoUiMultiSelect).Html.Global.Id = string( converted.GetId() )

      switch tplConverted := converted.FooterTemplate.(type) {
      case HtmlElementScript:
        if tplConverted.Type == SCRIPT_TYPE_KENDO_TEMPLATE {
          buffer.Write( tplConverted.ToElementScriptTag() )
        }
      }
      switch tplConverted := converted.NoDataTemplate.(type) {
      case HtmlElementScript:
        if tplConverted.Type == SCRIPT_TYPE_KENDO_TEMPLATE {
          buffer.Write( tplConverted.ToElementScriptTag() )
        }
      }
      switch tplConverted := converted.HeaderTemplate.(type) {
      case HtmlElementScript:
        if tplConverted.Type == SCRIPT_TYPE_KENDO_TEMPLATE {
          buffer.Write( tplConverted.ToElementScriptTag() )
        }
      }
    }
  }

  return buffer.Bytes()
}