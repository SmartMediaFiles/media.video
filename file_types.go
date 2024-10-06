package media_video

import (
	"github.com/smartmediafiles/media/media/maps"
	"github.com/smartmediafiles/media/media/types"
)

// List of supported media.Video file types.
const (
	VideoAvi types.FileType = "avi"
	VideoMkv types.FileType = "mkv" // Matroska Multimedia Container
	VideoM4v types.FileType = "m4v" // Apple iTunes MPEG-4 Container, optionally with DRM copy protection
	VideoMov types.FileType = "mov" // Apple QuickTime Movie
	VideoMp4 types.FileType = "mp4" // MPEG-4 Container based on QuickTime, can contain AVC, HEVC,...
	VideoMpg types.FileType = "mpg" // Moving Picture Experts Group (MPEG)
	VideoWmv types.FileType = "wmv" // Windows Media Video
)

// VideoFileTypesExtensions is a map of media.Video file types to file extensions.
var VideoFileTypesExtensions = maps.MapFileTypeExtensions{
	VideoAvi: {".avi"},
	VideoMkv: {".mkv"},
	VideoM4v: {".m4v"},
	VideoMov: {".mov", ".qt"},
	VideoMp4: {".mp4", "mp"},
	VideoMpg: {".mpg", ".mpeg"},
	VideoWmv: {".wmv"},
}
