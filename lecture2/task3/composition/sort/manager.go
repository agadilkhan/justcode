package sort

type AlorithmType string

const (
	Bubble AlorithmType = "bubble"
	Quick  AlorithmType = "quick"
)

type SortingAlgo interface {
	IsMatch(algorithmType AlorithmType) bool
	Sort()
}

type Manager struct {
	algorithms []SortingAlgo
}

func NewManager(algorithms ...SortingAlgo) *Manager {
	return &Manager{algorithms}
}

func (m *Manager) Execute(algorithmType AlorithmType) {
	for _, algo := range m.algorithms {
		if !algo.IsMatch(algorithmType) {
			continue
		}
		algo.Sort()
	}
}
