package protocol

import (
	"github.com/jiderlesi/onionscan/config"
	"github.com/jiderlesi/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
