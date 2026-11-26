//go:build darwin && !mygo_appstore

package darwin

// The private selectors (SPI) of AppKit and WebKit the darwin backend
// calls, every one looked up with respondsToSelector: first. A build with
// the mygo_appstore tag, for the Mac App Store, whose review rejects apps
// that use non-public API (guideline 2.5.1), takes spi_appstore.go
// instead: there each is "", which respondsTo answers false, so that the
// backend takes the path it takes when AppKit or WebKit lacks one, and
// the binary holds none of their names. No other file of the package
// names a private selector (TestSPIOnlyInSPIFiles).

// appStore reports a build for the Mac App Store.
const appStore = false

// WebKit's.
const (
	spiSetAllowsPictureInPictureMediaPlayback = "_setAllowsPictureInPictureMediaPlayback:"
	spiAllowsPictureInPictureMediaPlayback    = "_allowsPictureInPictureMediaPlayback"
	spiSetWindowOcclusionDetectionEnabled     = "_setWindowOcclusionDetectionEnabled:"
	spiWindowOcclusionDetectionEnabled        = "_windowOcclusionDetectionEnabled"

	spiHiddenPageDOMTimerThrottlingEnabled             = "_hiddenPageDOMTimerThrottlingEnabled"
	spiSetHiddenPageDOMTimerThrottlingEnabled          = "_setHiddenPageDOMTimerThrottlingEnabled:"
	spiHiddenPageCSSAnimationSuspensionEnabled         = "_hiddenPageCSSAnimationSuspensionEnabled"
	spiSetHiddenPageCSSAnimationSuspensionEnabled      = "_setHiddenPageCSSAnimationSuspensionEnabled:"
	spiPageVisibilityBasedProcessSuppressionEnabled    = "_pageVisibilityBasedProcessSuppressionEnabled"
	spiSetPageVisibilityBasedProcessSuppressionEnabled = "_setPageVisibilityBasedProcessSuppressionEnabled:"

	spiIsPlayingAudio         = "_isPlayingAudio"
	spiMediaMutedState        = "_mediaMutedState"
	spiSetPageMuted           = "_setPageMuted:"
	spiDisplayCaptureState    = "_displayCaptureState"
	spiSetDisplayCaptureState = "_setDisplayCaptureState:completionHandler:"

	spiCanTogglePictureInPicture = "_canTogglePictureInPicture"
	spiTogglePictureInPicture    = "_togglePictureInPicture"
	spiIsPictureInPictureActive  = "_isPictureInPictureActive"

	spiUserAgent            = "_userAgent"
	spiWebProcessIdentifier = "_webProcessIdentifier"
	spiGPUProcessIdentifier = "_gpuProcessIdentifier"
	spiSetUsesPageCache     = "_setUsesPageCache:"
	spiScrollPosition       = "_scrollPosition"
	spiInspector            = "_inspector"
	spiIsUserInitiated      = "_isUserInitiated"

	// The GUI tests' (testhooks.go).
	spiSetMockCaptureDevicesEnabled       = "_setMockCaptureDevicesEnabled:"
	spiSetMediaDevicesEnabled             = "_setMediaDevicesEnabled:"
	spiSetMockCaptureDevicesPromptEnabled = "_setMockCaptureDevicesPromptEnabled:"
	spiSetGetUserMediaRequiresFocus       = "_setGetUserMediaRequiresFocus:"
)

// AppKit's.
const (
	spiCornerRadius           = "_cornerRadius"
	spiSetForceMainAppearance = "_setForceMainAppearance:"
	spiHasActiveAppearance    = "_hasActiveAppearance"

	spiWaitCursor                           = "_waitCursor"
	spiHelpCursor                           = "_helpCursor"
	spiZoomInCursor                         = "_zoomInCursor"
	spiZoomOutCursor                        = "_zoomOutCursor"
	spiCellCursor                           = "_cellCursor"
	spiWindowResizeNorthWestSouthEastCursor = "_windowResizeNorthWestSouthEastCursor"
	spiWindowResizeNorthEastSouthWestCursor = "_windowResizeNorthEastSouthWestCursor"
)
