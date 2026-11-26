//go:build darwin && mygo_appstore

package darwin

// The private selectors of spi.go, none of them in a build for the Mac App
// Store: respondsTo answers false for "", as when AppKit or WebKit lacks
// one, and the binary holds none of their names.

// appStore reports a build for the Mac App Store.
const appStore = true

// WebKit's.
const (
	spiSetAllowsPictureInPictureMediaPlayback = ""
	spiAllowsPictureInPictureMediaPlayback    = ""
	spiSetWindowOcclusionDetectionEnabled     = ""
	spiWindowOcclusionDetectionEnabled        = ""

	spiHiddenPageDOMTimerThrottlingEnabled             = ""
	spiSetHiddenPageDOMTimerThrottlingEnabled          = ""
	spiHiddenPageCSSAnimationSuspensionEnabled         = ""
	spiSetHiddenPageCSSAnimationSuspensionEnabled      = ""
	spiPageVisibilityBasedProcessSuppressionEnabled    = ""
	spiSetPageVisibilityBasedProcessSuppressionEnabled = ""

	spiIsPlayingAudio         = ""
	spiMediaMutedState        = ""
	spiSetPageMuted           = ""
	spiDisplayCaptureState    = ""
	spiSetDisplayCaptureState = ""

	spiCanTogglePictureInPicture = ""
	spiTogglePictureInPicture    = ""
	spiIsPictureInPictureActive  = ""

	spiUserAgent            = ""
	spiWebProcessIdentifier = ""
	spiGPUProcessIdentifier = ""
	spiSetUsesPageCache     = ""
	spiScrollPosition       = ""
	spiInspector            = ""
	spiIsUserInitiated      = ""

	// The GUI tests' (testhooks.go).
	spiSetMockCaptureDevicesEnabled       = ""
	spiSetMediaDevicesEnabled             = ""
	spiSetMockCaptureDevicesPromptEnabled = ""
	spiSetGetUserMediaRequiresFocus       = ""
)

// AppKit's.
const (
	spiCornerRadius           = ""
	spiSetForceMainAppearance = ""
	spiHasActiveAppearance    = ""

	spiWaitCursor                           = ""
	spiHelpCursor                           = ""
	spiZoomInCursor                         = ""
	spiZoomOutCursor                        = ""
	spiCellCursor                           = ""
	spiWindowResizeNorthWestSouthEastCursor = ""
	spiWindowResizeNorthEastSouthWestCursor = ""
)
