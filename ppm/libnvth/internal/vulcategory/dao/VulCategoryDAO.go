package dao

import (
	"ppm/libnvth/internal/vulcategory/bean"
)

// VulCategoryDAO interface
type VulCategoryDAO interface {
	List() ([]bean.VulCategoryListBean, error)
	Get(ID string) (bean.VulCategoryGetBean, error)
	Insert(param VulCategoryDAOParam) error
	Update(param VulCategoryDAOParam) error
	Remove(ID string) error
	RemoveAll() error
}
