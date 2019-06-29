package vml

import "github.com/plandem/ooxml/ml"

//N.B.:
// Microsoft Office extended VML for more elements, but in reality only limited subset is used,
// so most used elements are exposed for better usability and rest elements considered as reserved to capture if required

type shapeElements struct {
	Path       *Path
	Formulas   *Formulas
	Handles    *Handles
	Fill       *Fill
	Stroke     *Stroke
	Shadow     *Shadow
	TextBox    *TextBox
	TextPath   *TextPath
	ImageData  *ImageData
	ClientData *ClientData

	ml.ReservedElements
}
