// generated by stringer -type=LeapIndicator,Mode,VersionNumber -output=protocol_string.go protocol.go; DO NOT EDIT

package ntp

import "fmt"

const _LeapIndicator_name = "LEAP_NONELEAP_ADDLEAP_SUBLEAP_NOSYNC"

var _LeapIndicator_index = [...]uint8{0, 9, 17, 25, 36}

func (i LeapIndicator) String() string {
	if i+1 >= LeapIndicator(len(_LeapIndicator_index)) {
		return fmt.Sprintf("LeapIndicator(%d)", i)
	}
	return _LeapIndicator_name[_LeapIndicator_index[i]:_LeapIndicator_index[i+1]]
}

const _Mode_name = "MODE_RESERVEDMODE_SYMMETRIC_ACTIVEMODE_SYMMETRIC_PASSIVEMODE_CLIENTMODE_SERVERMODE_BROADCASTMODE_CONTROLMODE_PRIVATE"

var _Mode_index = [...]uint8{0, 13, 34, 56, 67, 78, 92, 104, 116}

func (i Mode) String() string {
	if i+1 >= Mode(len(_Mode_index)) {
		return fmt.Sprintf("Mode(%d)", i)
	}
	return _Mode_name[_Mode_index[i]:_Mode_index[i+1]]
}

const _VersionNumber_name = "NTPv4"

var _VersionNumber_index = [...]uint8{0, 5}

func (i VersionNumber) String() string {
	i -= 4
	if i+1 >= VersionNumber(len(_VersionNumber_index)) {
		return fmt.Sprintf("VersionNumber(%d)", i+4)
	}
	return _VersionNumber_name[_VersionNumber_index[i]:_VersionNumber_index[i+1]]
}
