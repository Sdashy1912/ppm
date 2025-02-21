package dao

import (
	"ppm/libnvth/internal/vultemplate/bean"
)

// VulTemplateDAO interface
type VulTemplateDAO interface {
	List(param VulTemplateListDAOParam) ([]bean.VulTemplateBean, error)
	Get(ID string) (bean.VulTemplateBean, error)
	Insert(param VulTemplateDAOParam) error
	Update(param VulTemplateDAOParam) error
	Remove(ID string) error
	RemoveAll() error
}
