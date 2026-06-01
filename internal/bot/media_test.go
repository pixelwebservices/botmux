package bot

import (
	"testing"

	tgbotapi "github.com/OvyFlash/telegram-bot-api"
)

func TestExtractMediaLivePhoto(t *testing.T) {
	msg := &tgbotapi.Message{
		LivePhoto: &tgbotapi.LivePhoto{
			FileID: "live-file",
			Photo: []tgbotapi.PhotoSize{
				{FileID: "thumb"},
				{FileID: "preview"},
			},
		},
	}
	mediaType, fileID := ExtractMedia(msg)
	if mediaType != "live_photo" || fileID != "preview" {
		t.Fatalf("ExtractMedia() = (%q, %q), want (live_photo, preview)", mediaType, fileID)
	}
}

func TestLivePhotoFileID(t *testing.T) {
	id := LivePhotoFileID(map[string]any{
		"photo": []any{
			map[string]any{"file_id": "small"},
			map[string]any{"file_id": "large"},
		},
	})
	if id != "large" {
		t.Fatalf("LivePhotoFileID() = %q, want large", id)
	}
}
