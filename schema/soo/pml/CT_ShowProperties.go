// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/dml"
)

type CT_ShowProperties struct {
	// Loop Slide Show
	LoopAttr *bool
	// Show Narration in Slide Show
	ShowNarrationAttr *bool
	// Show Animation in Slide Show
	ShowAnimationAttr *bool
	// Use Timings in Slide Show
	UseTimingsAttr *bool
	// Presenter Slide Show Mode
	Present *CT_Empty
	// Browse Slide Show Mode
	Browse *CT_ShowInfoBrowse
	// Kiosk Slide Show Mode
	Kiosk *CT_ShowInfoKiosk
	// All Slides
	SldAll *CT_Empty
	// Slide Range
	SldRg *CT_IndexRange
	// Custom Show
	CustShow *CT_CustomShowId
	// Pen Color for Slide Show
	PenClr *dml.CT_Color
	ExtLst *CT_ExtensionList
}

func NewCT_ShowProperties() *CT_ShowProperties {
	ret := &CT_ShowProperties{}
	return ret
}

func (m *CT_ShowProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LoopAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "loop"},
			Value: fmt.Sprintf("%d", b2i(*m.LoopAttr))})
	}
	if m.ShowNarrationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showNarration"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowNarrationAttr))})
	}
	if m.ShowAnimationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showAnimation"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowAnimationAttr))})
	}
	if m.UseTimingsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "useTimings"},
			Value: fmt.Sprintf("%d", b2i(*m.UseTimingsAttr))})
	}
	e.EncodeToken(start)
	if m.Present != nil {
		sepresent := xml.StartElement{Name: xml.Name{Local: "p:present"}}
		e.EncodeElement(m.Present, sepresent)
	}
	if m.Browse != nil {
		sebrowse := xml.StartElement{Name: xml.Name{Local: "p:browse"}}
		e.EncodeElement(m.Browse, sebrowse)
	}
	if m.Kiosk != nil {
		sekiosk := xml.StartElement{Name: xml.Name{Local: "p:kiosk"}}
		e.EncodeElement(m.Kiosk, sekiosk)
	}
	if m.SldAll != nil {
		sesldAll := xml.StartElement{Name: xml.Name{Local: "p:sldAll"}}
		e.EncodeElement(m.SldAll, sesldAll)
	}
	if m.SldRg != nil {
		sesldRg := xml.StartElement{Name: xml.Name{Local: "p:sldRg"}}
		e.EncodeElement(m.SldRg, sesldRg)
	}
	if m.CustShow != nil {
		secustShow := xml.StartElement{Name: xml.Name{Local: "p:custShow"}}
		e.EncodeElement(m.CustShow, secustShow)
	}
	if m.PenClr != nil {
		sepenClr := xml.StartElement{Name: xml.Name{Local: "p:penClr"}}
		e.EncodeElement(m.PenClr, sepenClr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ShowProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "loop" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LoopAttr = &parsed
		}
		if attr.Name.Local == "showNarration" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowNarrationAttr = &parsed
		}
		if attr.Name.Local == "showAnimation" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAnimationAttr = &parsed
		}
		if attr.Name.Local == "useTimings" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseTimingsAttr = &parsed
		}
	}
lCT_ShowProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "present"}:
				m.Present = NewCT_Empty()
				if err := d.DecodeElement(m.Present, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "browse"}:
				m.Browse = NewCT_ShowInfoBrowse()
				if err := d.DecodeElement(m.Browse, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "kiosk"}:
				m.Kiosk = NewCT_ShowInfoKiosk()
				if err := d.DecodeElement(m.Kiosk, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldAll"}:
				m.SldAll = NewCT_Empty()
				if err := d.DecodeElement(m.SldAll, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldRg"}:
				m.SldRg = NewCT_IndexRange()
				if err := d.DecodeElement(m.SldRg, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "custShow"}:
				m.CustShow = NewCT_CustomShowId()
				if err := d.DecodeElement(m.CustShow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "penClr"}:
				m.PenClr = dml.NewCT_Color()
				if err := d.DecodeElement(m.PenClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_ShowProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ShowProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ShowProperties and its children
func (m *CT_ShowProperties) Validate() error {
	return m.ValidateWithPath("CT_ShowProperties")
}

// ValidateWithPath validates the CT_ShowProperties and its children, prefixing error messages with path
func (m *CT_ShowProperties) ValidateWithPath(path string) error {
	if m.Present != nil {
		if err := m.Present.ValidateWithPath(path + "/Present"); err != nil {
			return err
		}
	}
	if m.Browse != nil {
		if err := m.Browse.ValidateWithPath(path + "/Browse"); err != nil {
			return err
		}
	}
	if m.Kiosk != nil {
		if err := m.Kiosk.ValidateWithPath(path + "/Kiosk"); err != nil {
			return err
		}
	}
	if m.SldAll != nil {
		if err := m.SldAll.ValidateWithPath(path + "/SldAll"); err != nil {
			return err
		}
	}
	if m.SldRg != nil {
		if err := m.SldRg.ValidateWithPath(path + "/SldRg"); err != nil {
			return err
		}
	}
	if m.CustShow != nil {
		if err := m.CustShow.ValidateWithPath(path + "/CustShow"); err != nil {
			return err
		}
	}
	if m.PenClr != nil {
		if err := m.PenClr.ValidateWithPath(path + "/PenClr"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}