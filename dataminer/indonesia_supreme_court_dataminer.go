package dataminer

type IndonesiaSupremeCourtDataminer struct {
}

func NewIndonesiaSupremeCourtDataminer() Dataminer {
	return IndonesiaSupremeCourtDataminer{}
}

func (d IndonesiaSupremeCourtDataminer) MineMetadata() error {
	return nil
}
