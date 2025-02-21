package stats

// Interactor acts as services
type Interactor struct {
	dao DAO
}

// Stats stats
func (interactor Interactor) Stats() (map[string]interface{}, error) {
	return interactor.dao.Stats()
}

// NewInteractor return a new Interactor instances
func NewInteractor(dao DAO) Interactor {
	return Interactor{dao: dao}
}
