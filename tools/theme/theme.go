package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Bios-Marcel/cordless/config"
	"github.com/Bios-Marcel/tview"
	"github.com/gdamore/tcell"
)

func main() {
	theme := &config.Theme{
		Theme: &tview.Theme{
			PrimitiveBackgroundColor:    tcell.NewRGBColor(22, 0, 40),       // Dark purple
			ContrastBackgroundColor:     tcell.NewRGBColor(226, 49, 82),     // Amaranth
			MoreContrastBackgroundColor: tcell.ColorWhite,                   // White
			BorderColor:                 tcell.NewRGBColor(240, 159, 127),   // Light salmon
			BorderFocusColor:            tcell.NewRGBColor(226, 49, 82),	 // Amaranth
			TitleColor:                  tcell.ColorWhite,					 // White
			GraphicsColor:               tcell.NewRGBColor(76,42,69),		 // Dark byzantium
			PrimaryTextColor:            tcell.NewRGBColor(240, 159, 127),   // Light salmon
			SecondaryTextColor:          tcell.ColorYellow, 				 // Yellow
			TertiaryTextColor:           tcell.NewRGBColor(136,89,142),      // Antique fuchsia
			InverseTextColor:            tcell.NewRGBColor(136,89,142),      // Antique fuchsia
			ContrastSecondaryTextColor:  tcell.ColorDarkCyan, 				 // Dark cyan
		},
		BlockedUserColor: tcell.NewRGBColor(145, 30, 51),
		InfoMessageColor: tcell.NewRGBColor(145, 30, 51),
		BotColor:         tcell.NewRGBColor(226, 49, 82),
		MessageTimeColor: tcell.NewRGBColor(145, 30, 51),
		LinkColor:        tcell.NewRGBColor(226, 49, 82),
		DefaultUserColor: tcell.NewRGBColor(226, 49, 82),
		AttentionColor:   tcell.NewRGBColor(226, 49, 82),
		ErrorColor:       tcell.ColorRed,
		RandomUserColors: []tcell.Color{
			tcell.NewRGBColor(0xd8, 0x50, 0x4e),
			tcell.NewRGBColor(0xd8, 0x7e, 0x4e),
			tcell.NewRGBColor(0xd8, 0xa5, 0x4e),
			tcell.NewRGBColor(0xd8, 0xc6, 0x4e),
			tcell.NewRGBColor(0xb8, 0xd8, 0x4e),
			tcell.NewRGBColor(0x91, 0xd8, 0x4e),
			tcell.NewRGBColor(0x67, 0xd8, 0x4e),
			tcell.NewRGBColor(0x4e, 0xd8, 0x7c),
			tcell.NewRGBColor(0x4e, 0xd8, 0xaa),
			tcell.NewRGBColor(0x4e, 0xd8, 0xcf),
			tcell.NewRGBColor(0x4e, 0xb6, 0xd8),
			tcell.NewRGBColor(0x4e, 0x57, 0xd8),
			tcell.NewRGBColor(0x75, 0x4e, 0xd8),
			tcell.NewRGBColor(0xa3, 0x4e, 0xd8),
			tcell.NewRGBColor(0xcf, 0x4e, 0xd8),
			tcell.NewRGBColor(0xd8, 0x4e, 0x9c),
		},
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(theme)
}

// Usage: fromHex("#FF0000")
func fromHex(hexString string) tcell.Color {
	trimmed := strings.TrimPrefix(strings.TrimSpace(hexString), "#")
	var r, g, b int32
	fmt.Sscanf(trimmed, "%02x%02x%02x", &r, &g, &b)
	return tcell.NewRGBColor(r, g, b)
}
