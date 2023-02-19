package module

import (
	"context"
	"gf-admin/internal/co_interface"
	"gf-admin/internal/dao"
	"gf-admin/internal/logic/companyDept"
	"gf-admin/internal/logic/companyPost"
	model "gf-admin/internal/model"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/text/gstr"
)

type Modules struct {
	conf *model.Config

	dept co_interface.IDept
	post co_interface.IPost

	i18n *gi18n.Manager
	xDao *dao.XDao
}

func (m *Modules) GetConfig() *model.Config {
	return m.conf
}
func (m *Modules) Dao() *dao.XDao {
	return m.xDao
}

func (m *Modules) Dept() co_interface.IDept {
	return m.dept
}
func (m *Modules) Post() co_interface.IPost {
	return m.post
}

func (m *Modules) T(ctx context.Context, content string) string {
	return m.i18n.Translate(ctx, content)
}

// Tf is alias of TranslateFormat for convenience.
func (m *Modules) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.i18n.TranslateFormat(ctx, format, values...)
}

func (m *Modules) SetI18n(i18n *gi18n.Manager) error {
	if i18n == nil {
		i18n = gi18n.New()
		i18n.SetLanguage("zh-CN")
		err := i18n.SetPath("i18n/" + gstr.ToLower(m.conf.KeyIndex))
		if err != nil {
			return err
		}
	}

	m.i18n = i18n
	return nil
}

func NewModules(
	conf *model.Config,
	xDao *dao.XDao,
) *Modules {
	module := &Modules{
		conf: conf,
		xDao: xDao,
	}

	// 初始化默认多语言对象
	module.SetI18n(nil)

	module.dept = companyDept.NewDept(module)
	module.post = companyPost.NewPost(module)

	return module
}
