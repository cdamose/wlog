//Package wlog creates simple to use UI structure.
//The UI is used to simply print to the screen.
//There a wrappers that will wrap each other to create a good looking UI.
//You can add color and prefixes as well as make it thread safe.
package wlog

import "io"

//New creates a BasicUI.
//This should be the first function you call.
//This is not thread safe and should only be used in serial applications.
func New(reader io.Reader, writer, errorWriter io.Writer) *BasicUI {
	return &BasicUI{
		Reader:      reader,
		Writer:      writer,
		ErrorWriter: errorWriter,
	}
}

// AddConcurrent will wrap a thread safe UI on top of ui.
// Safe to use inside of go routines.
func AddConcurrent(ui UI) *ConcurrentUI {
	return &ConcurrentUI{UI: ui}
}

//AddColor will wrap a colorful UI on top of ui.
//Use wlog's color variables for the color.
//All background colors are not changed by this function but you are able to change them manually.
//Just create this structure manually and change any of the background colors you want.
func AddColor(logColor, outputColor, successColor, infoColor, errorColor, warnColor, runningColor Color, ui UI) *ColorUI {
	return &ColorUI{
		LogFGColor:     logColor,
		LogBGColor:     None,
		OutputFGColor:  outputColor,
		OutputBGColor:  None,
		SuccessFGColor: successColor,
		SuccessBGColor: None,
		InfoFGColor:    infoColor,
		InfoBGColor:    None,
		ErrorFGColor:   errorColor,
		ErrorBGColor:   None,
		WarnFGColor:    warnColor,
		WarnBGColor:    None,
		RunningFGColor: runningColor,
		RunningBGColor: None,
		UI:             ui,
	}
}

//AddPrefix will wrap a UI that will prefix the message on top of ui.
//If a prefix is set to nothing ("") then there will be no prefix for that message type.
func AddPrefix(logPre, outputPre, successPre, infoPre, errorPre, warnPre, runningPre string, ui UI) *PrefixUI {
	return &PrefixUI{
		LogPrefix:     logPre,
		OutputPrefix:  outputPre,
		SuccessPrefix: successPre,
		InfoPrefix:    infoPre,
		ErrorPrefix:   errorPre,
		WarnPrefix:    warnPre,
		RunningPrefix: runningPre,
		UI:            ui,
	}
}
