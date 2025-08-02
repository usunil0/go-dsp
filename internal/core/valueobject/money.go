package valueobject

type CPMMicros int64

func (c CPMMicros) USD() float64        { return float64(c) / 1e6 }
func (c CPMMicros) PerImpMicros() int64 { return (int64(c) + 999) / 1000 } // ceil(cpm/1000)
