package constants

type Status string

const (
	StatusOutOfUpRange  Status = "out_of_up_range"
	StatusOutOfLowRange Status = "out_of_low_range"
	StatusInRange       Status = "in_range"
	StatusClose         Status = "close"
)

func (s Status) String() string {
	return string(s)
}
