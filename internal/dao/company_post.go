// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/SupenBysz/gf-admin-community/utility/daoctl/dao_interface"
	"gf-admin/internal/dao/internal"
)

type CompanyPost = dao_interface.TIDao[internal.CompanyPostColumns]

func NewCompanyPost(dao ...dao_interface.IDao) CompanyPost {
	return (CompanyPost)(internal.NewCompanyPostDao(dao...))
}