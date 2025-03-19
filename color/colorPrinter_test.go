package color

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/fatih/color"
	"github.com/spf13/viper"

	"github.com/bjoernkarma/gitctl/config"
)

const message = "test message"

func expectMessageIsPrinted(t *testing.T, buf bytes.Buffer, message string) {
	if !bytes.Contains(buf.Bytes(), []byte(message)) {
		t.Errorf("expected test message to be printed, got %v", buf.String())
	}
}

func expectNoMessageIsPrinted(t *testing.T, buf bytes.Buffer) {
	if buf.String() != "" {
		t.Errorf("expected no messages to be printed, got %v", buf.String())
	}
}

func TestPrintColored_InvalidOutput(t *testing.T) {
	var buf = errorWriter{}
	color.Output = &buf
	defer func() {
		color.Output = nil
	}()

	viper.Set(config.GitCtlQuiet, false)
	viper.Set(config.GitCtlColor, true)
	PrintError(message)
}

func TestPrintInfo_QuietMode(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	viper.Set(config.GitCtlQuiet, true)
	PrintInfo(message)

	expectNoMessageIsPrinted(t, buf)
}

func TestPrintInfo_ColoredOutput(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf
	defer func() {
		color.Output = nil
	}()

	viper.Set(config.GitCtlQuiet, false)
	viper.Set(config.GitCtlColor, true)
	PrintInfo(message)

	expectMessageIsPrinted(t, buf, message)
}

func TestPrintInfo_NonColoredOutput(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	viper.Set(config.GitCtlQuiet, false)
	viper.Set(config.GitCtlColor, false)
	PrintInfo(message)

	expectMessageIsPrinted(t, buf, message)
}

func TestPrintSubtleInfo_QuietMode(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	viper.Set(config.GitCtlQuiet, true)
	PrintSubtleInfo(message)

	expectNoMessageIsPrinted(t, buf)
}

func TestPrintSubtleInfo_ColoredOutput(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf
	defer func() {
		color.Output = nil
	}()

	viper.Set(config.GitCtlQuiet, false)
	viper.Set(config.GitCtlColor, true)
	PrintSubtleInfo(message)

	expectMessageIsPrinted(t, buf, message)
}

func TestPrintSubtleInfo_NonColoredOutput(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	viper.Set(config.GitCtlQuiet, false)
	viper.Set(config.GitCtlColor, false)
	PrintSubtleInfo(message)

	expectMessageIsPrinted(t, buf, message)
}

func TestPrintSuccess_ColoredOutput(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf
	defer func() {
		color.Output = nil
	}()

	viper.Set(config.GitCtlQuiet, false)
	viper.Set(config.GitCtlColor, true)
	PrintSuccess(message)

	expectMessageIsPrinted(t, buf, message)
}

func TestPrintSuccess_NonColoredOutput(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	viper.Set(config.GitCtlQuiet, false)
	viper.Set(config.GitCtlColor, false)
	PrintSuccess(message)

	expectMessageIsPrinted(t, buf, message)
}

func TestPrintError_ColoredOutput(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf
	defer func() {
		color.Output = nil
	}()

	viper.Set(config.GitCtlQuiet, false)
	viper.Set(config.GitCtlColor, true)
	PrintError(message)

	expectMessageIsPrinted(t, buf, message)
}

func TestPrintError_NonColoredOutput(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil)
	}()

	viper.Set(config.GitCtlQuiet, false)
	viper.Set(config.GitCtlColor, false)
	PrintError(message)

	expectMessageIsPrinted(t, buf, message)
}

// Simulating an error when writing to the output

type errorWriter struct{}

func (ew *errorWriter) Write(p []byte) (n int, err error) {
	log.Print("simulated write error for message: " + string(p))
	return 0, fmt.Errorf("simulated write error")
}
