// Code generated by "stringer -type=Code"; DO NOT EDIT.

package irtt

import "strconv"

const (
	_Code_name_0 = "UnexpectedInitChannelCloseServerFillTooLongOpenTimeoutTooShortInvalidReceivedStatsStringInvalidReceivedStatsIntInvalidServerRestrictionOpenTimeoutServerClosedConnTokenZeroDurationNonPositiveIntervalNonPositiveNoSuchWaiterNoSuchTimeSourceNoSuchTimerNoSuchFillerNoSuchAveragerInvalidWaitDurationInvalidWaitFactorInvalidWaitStringInvalidSleepFactorUnexpectedSequenceNumberClockMismatchStampAtMismatchShortReplyExpectedReplyFlagTTLErrorDFErrorUnexpectedOpenFlagAllocateResultsPanicInvalidExpAvgAlphaInvalidWinAvgWindow"
	_Code_name_1 = "InvalidSyslogURISyslogNotSupportedAddressMismatchLargeRequestShortIntervalInvalidConnTokenNoSuitableAddressFoundUnexpectedReplyFlagInvalidGCModeStringUnspecifiedWithSpecifiedAddressesNoMatchingInterfacesUpNoMatchingInterfaces"
	_Code_name_2 = "ProtocolVersionMismatchInvalidParamValueParamOverflowShortParamBufferInvalidFlagBitsSetDFNotSupportedInconsistentClocksNonexclusiveMidpointTStampUnexpectedHMACBadHMACNoHMACBadMagicInvalidClockIntInvalidClockStringInvalidAllowStampStringInvalidStampAtIntInvalidStampAtStringFieldsCapacityTooLargeFieldsLengthTooLargeInvalidDFStringShortWrite"
	_Code_name_3 = "MultipleAddressesServerStartServerStopListenerStartListenerStopListenerErrorDropNewConnOpenCloseCloseConnNoDSCPSupportExceededDurationNoReceiveDstAddrSupportRemoveNoConnInvalidServerFill"
	_Code_name_4 = "ConnectingConnectedWaitForPacketsServerRestrictionNoTestConnectedClosed"
)

var (
	_Code_index_0 = [...]uint16{0, 26, 43, 62, 88, 111, 135, 146, 158, 171, 190, 209, 221, 237, 248, 260, 274, 293, 310, 327, 345, 369, 382, 397, 407, 424, 432, 439, 457, 477, 495, 514}
	_Code_index_1 = [...]uint8{0, 16, 34, 49, 61, 74, 90, 112, 131, 150, 183, 205, 225}
	_Code_index_2 = [...]uint16{0, 23, 40, 53, 69, 87, 101, 119, 145, 159, 166, 172, 180, 195, 213, 236, 253, 273, 295, 315, 330, 340}
	_Code_index_3 = [...]uint8{0, 17, 28, 38, 51, 63, 76, 80, 87, 96, 105, 118, 134, 157, 169, 186}
	_Code_index_4 = [...]uint8{0, 10, 19, 33, 50, 56, 71}
)

func (i Code) String() string {
	switch {
	case -2078 <= i && i <= -2048:
		i -= -2078
		return _Code_name_0[_Code_index_0[i]:_Code_index_0[i+1]]
	case -1035 <= i && i <= -1024:
		i -= -1035
		return _Code_name_1[_Code_index_1[i]:_Code_index_1[i+1]]
	case -21 <= i && i <= -1:
		i -= -21
		return _Code_name_2[_Code_index_2[i]:_Code_index_2[i+1]]
	case 1024 <= i && i <= 1038:
		i -= 1024
		return _Code_name_3[_Code_index_3[i]:_Code_index_3[i+1]]
	case 2048 <= i && i <= 2053:
		i -= 2048
		return _Code_name_4[_Code_index_4[i]:_Code_index_4[i+1]]
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
