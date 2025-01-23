package ui

import (
	"fmt"
	"image/color"
	"strings"
	"time"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"

	"github.com/fanchann/gowor/utils"
)

func (w *WordDictionary) repeatingKeyPressed(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func (w *WordDictionary) Update() error {
	// Deteksi perubahan input
	if w.Text != w.LastText {
		w.Suggestions = utils.GetSuggestions(w.Trie, w.Text)
		w.LastUpdate = time.Now()
		w.LastText = w.Text
	}

	// Keluar dari aplikasi jika tombol F9 ditekan
	if ebiten.IsKeyPressed(ebiten.KeyF9) {
		return ebiten.Termination
	}

	// Keluar dari popup jika tombol ESC ditekan dan popup sedang ditampilkan
	if w.ShowPopup && ebiten.IsKeyPressed(ebiten.KeyEscape) {
		w.ShowPopup = false
		w.Meaning = ""
	}

	w.Runes = ebiten.AppendInputChars(w.Runes[:0])
	for _, r := range w.Runes {
		if unicode.IsPrint(r) {
			w.Text += string(r)
		}
	}

	ss := strings.Split(w.Text, "\n")
	if len(ss) > 10 {
		w.Text = strings.Join(ss[len(ss)-10:], "\n")
	}

	if w.repeatingKeyPressed(ebiten.KeyEnter) || w.repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		w.Text += "\n"
	}

	// jika tombol backspace ditekan, hapus karakter yang terakhir
	if w.repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(w.Text) >= 1 {
			w.Text = w.Text[:len(w.Text)-1]
		}
	}

	// update saran berdasarkan kata terakhir yang ditekan
	words := strings.Fields(w.Text)
	if len(words) > 0 {
		lastWord := words[len(words)-1]
		w.Suggestions = utils.GetSuggestions(w.Trie, lastWord)

		if w.repeatingKeyPressed(ebiten.KeySpace) || w.repeatingKeyPressed(ebiten.KeyEnter) {
			w.Meaning = utils.GetMeaning(w.Trie, lastWord)
			w.ShowPopup = true // Show popup
		}
	} else {
		w.Suggestions = nil
		w.Meaning = ""
	}

	w.Counter++
	return nil
}

func (w *WordDictionary) Draw(screen *ebiten.Image) {
	// Blink the cursor.
	t := w.Text
	if w.Counter%60 < 30 {
		t += "_"
	}
	ebitenutil.DebugPrint(screen, t) // Menampilkan teks yang diketik beserta kursor

	// Suggestions
	if len(w.Suggestions) > 0 {
		if time.Since(w.LastUpdate) >= suggestionTime {
			suggestionsText := "Suggestions: " + strings.Join(w.Suggestions, ", ")
			ebitenutil.DebugPrintAt(screen, suggestionsText, 10, ScreenHeight-60)
		}
	}

	if w.ShowPopup && w.Meaning != "" {
		// Ambil kata terakhir yang dicari
		words := strings.Fields(w.Text)
		var lastWord string
		if len(words) > 0 {
			lastWord = words[len(words)-1]
		}

		meaningText := fmt.Sprintf("%s Meaning: %s", lastWord, w.Meaning)

		// Hitung ukuran popup berdasarkan panjang teks
		font := basicfont.Face7x13
		charWidth := 7        // Lebar setiap karakter dalam font monospace
		maxCharsPerLine := 50 // Jumlah maksimum karakter per baris
		lineHeight := 16      // Tinggi setiap baris teks

		// Memecah teks menjadi beberapa baris
		lines := utils.WrapText(meaningText, maxCharsPerLine)

		// Hitung ukuran popup berdasarkan jumlah baris
		popupWidth := maxCharsPerLine*charWidth + 40 // Tambahkan padding
		popupHeight := len(lines)*lineHeight + 40

		popupX := (ScreenWidth - popupWidth) / 2 // Posisi tengah layar
		popupY := (ScreenHeight - popupHeight) / 2

		// Box popup
		ebitenutil.DrawRect(screen, float64(popupX), float64(popupY), float64(popupWidth), float64(popupHeight), color.RGBA{0, 0, 0, 200}) // Latar belakang semi-transparan
		ebitenutil.DrawRect(screen, float64(popupX), float64(popupY), float64(popupWidth), 2, color.White)                                 // Border atas
		ebitenutil.DrawRect(screen, float64(popupX), float64(popupY+popupHeight-2), float64(popupWidth), 2, color.White)                   // Border bawah
		ebitenutil.DrawRect(screen, float64(popupX), float64(popupY), 2, float64(popupHeight), color.White)                                // Border kiri
		ebitenutil.DrawRect(screen, float64(popupX+popupWidth-2), float64(popupY), 2, float64(popupHeight), color.White)                   // Border kanan

		// Tampilkan teks arti kata di dalam popup (setiap baris)
		textX := popupX + 20
		textY := popupY + 20
		for _, line := range lines {
			text.Draw(screen, line, font, textX, textY, color.White)
			textY += lineHeight
		}

		exitPopupText := "Press ESC to close popup"
		text.Draw(screen, exitPopupText, font, textX, textY+30, color.White)
	}

	exitAppText := "Press F9 to exit application"
	ebitenutil.DebugPrintAt(screen, exitAppText, 10, ScreenHeight-20)
}

func (w *WordDictionary) Layout(outSideWidth int, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight

}
