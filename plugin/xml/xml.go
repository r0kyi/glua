package xml

import (
	"github.com/clbanning/mxj/v2"
	"github.com/r0kyi/glua/core"
)

type Xml struct {
	raw string
	xml map[string]interface{}
}

func (x *Xml) encode() error {
	xml := mxj.Map(x.xml)
	raw, err := xml.Xml()
	if err != nil {
		return err
	}
	x.raw = core.B2S(raw)

	return nil
}

func (x *Xml) decode() error {
	xml, err := mxj.NewMapXml(core.S2B(x.raw))
	if err != nil {
		return err
	}
	x.xml = xml

	return nil
}
