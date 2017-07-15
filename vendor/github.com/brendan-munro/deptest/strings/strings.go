package strings

type CoolString string

func (_ CoolString) String() string {
	return "Cool string"
}
