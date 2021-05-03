package interservice_communication

type Http struct {
	dev bool
}

func (h *Http) Get() *Http {

	h.dev = false
	return h
}
