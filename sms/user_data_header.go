package sms

type UserDataHeader struct {
	TotalNumber int
	Sequence    int
	Tag         int
}

func (udh *UserDataHeader) ReadFrom(octets []byte) error {
	octetsLng := len(octets)
	headerLng := int(octets[0]) + 1
	if (octetsLng-headerLng) <= 0 || headerLng <= 5 {
		return ErrIncorrectUserDataHeaderLength
	}

	h := octets[:headerLng]
	if int(h[1]) == 0 && headerLng == 6 {
		udh.Sequence = int(h[5])
		udh.TotalNumber = int(h[4])
		udh.Tag = int(h[3])
	} else if int(h[1]) == 8 && headerLng == 7 {
		udh.Sequence = int(h[6])
		udh.TotalNumber = int(h[5])
		udh.Tag = int(h[3])*256 + int(h[4])
	}

	return nil
}
