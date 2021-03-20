package css

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyleRuleText(t *testing.T) {
	sr := CSSStyleRule{}
	sr.Selector = NewCSSValue(".box")
	sr.Styles = make([]*CSSStyleDeclaration, 2)
	sr.Styles[0] = NewCSSStyleDeclaration("width", "10px", false)
	sr.Styles[1] = NewCSSStyleDeclaration("height", "100px", false)

	assert.Equal(t, sr.Text(), ".box {\n    width: 10px;\n    height: 100px\n}")
}
