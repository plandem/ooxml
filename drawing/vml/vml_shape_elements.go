package vml

//shapeElements is direct mapping of EG_ShapeElements
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

	//reserved elements
	Skew          *Skew          `xml:"skew,omitempty"`
	Extrusion     *Extrusion     `xml:"extrusion,omitempty"`
	Callout       *Callout       `xml:"callout,omitempty"`
	Lock          *Lock          `xml:"lock,omitempty"`
	ClipPath      *ClipPath      `xml:"clippath,omitempty"`
	SignatureLine *SignatureLine `xml:"signatureline,omitempty"`
}
