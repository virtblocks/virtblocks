package devices

type MemballoonModel int

const (
	MEMBALLOON_MODEL_NONE                    = MemballoonModel(0)
	MEMBALLOON_MODEL_VIRTIO                  = MemballoonModel(1)
	MEMBALLOON_MODEL_VIRTIO_NON_TRANSITIONAL = MemballoonModel(2)
	MEMBALLOON_MODEL_VIRTIO_TRANSITIONAL     = MemballoonModel(3)
)

type Memballoon struct {
	model MemballoonModel
}

func (self MemballoonModel) String() string {
	var ret string

	switch self {
	case MEMBALLOON_MODEL_NONE:
		ret = ""
	case MEMBALLOON_MODEL_VIRTIO:
		ret = "virtio-memballoon"
	case MEMBALLOON_MODEL_VIRTIO_NON_TRANSITIONAL:
		ret = "virtio-memballoon-non-transitional"
	case MEMBALLOON_MODEL_VIRTIO_TRANSITIONAL:
		ret = "virtio-memballoon-transitional"
	}

	return ret
}

func NewMemballoon() *Memballoon {
	return &Memballoon{model: MEMBALLOON_MODEL_NONE}
}

func (self *Memballoon) SetModel(model MemballoonModel) {
	self.model = model
}

func (self Memballoon) GetModel() MemballoonModel {
	return self.model
}

func (self Memballoon) String() string {
	return self.model.String()
}
