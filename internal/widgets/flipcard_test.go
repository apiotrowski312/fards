package widgets

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestFlipCard(t *testing.T) {
	t.Parallel()

	fc := NewFlipCard("front", "back")
	_ = test.WidgetRenderer(fc)

	assert.Equal(t, fc.label.Text, "front")
	test.Tap(fc)
	assert.Equal(t, fc.label.Text, "back")
	test.Tap(fc)
	assert.Equal(t, fc.label.Text, "front")
}
