package xml

import (
	"github.com/clbanning/mxj/v2"
	"github.com/r0kyi/glua/core"
)

type Xml struct {
}

func (x *Xml) encode(xml map[string]any) (string, error) {
	xml_ := mxj.Map(xml)
	raw, err := xml_.Xml()
	if err != nil {
		return "", err
	}

	return core.B2S(raw), nil
}

func (x *Xml) decode(raw string) (map[string]any, error) {
	xml, err := mxj.NewMapXml(core.S2B(raw))
	if err != nil {
		return nil, err
	}

	return xml, nil
}
