package pipeline

type ActivityState uint8

const (
	Booting ActivityState = iota
	Syncing
	Active
	Crashed
)

func (as ActivityState) String() string {
	switch as {
	case Booting:
		return "booting"

	case Syncing:
		return "syncing"

	case Active:
		return "active"

	case Crashed:
		return "crashed"
	}

	return "unknown"
}

const (
	couldNotCastErr = "could not cast component initializer function to %s constructor type"
	pIDNotFoundErr  = "could not find pipeline ID for %s"
	uuidNotFoundErr = "could not find matching UUID for pipeline entry"
)
