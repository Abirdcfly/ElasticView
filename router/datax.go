package router

import (
	. "github.com/1340691923/ElasticView/controller"
	"github.com/1340691923/ElasticView/platform-basic-libs/api_config"
	. "github.com/gofiber/fiber/v2"
)

// ES基础操作 路由
func runDatax(app *App) {
	apiRouterConfig := api_config.NewApiRouterConfig()
	const AbsolutePath = "/api/datax"
	es := app.Group(AbsolutePath)
	{
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "数据抽取源数据库列表",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "LinkInfoList",
		}, es.(*Group), true, DataxController{}.LinkInfoList)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "新增数据抽取源数据库",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "InsertLink",
		}, es.(*Group), true, DataxController{}.InsertLink)
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "删除数据抽取源数据库",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "DelLinkById",
		}, es.(*Group), true, DataxController{}.DelLinkById)
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "测试连接数据抽取源数据库",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "TestLink",
		}, es.(*Group), true, DataxController{}.TestLink)
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "数据抽取源数据库下拉选",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "LinkSelectOpt",
		}, es.(*Group), true, DataxController{}.LinkSelectOpt)
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "获取数据抽取源数据表名",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "Tables",
		}, es.(*Group), true, DataxController{}.Tables)
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "获取数据抽取源数据表字段",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "GetTableColumns",
		}, es.(*Group), true, DataxController{}.GetTableColumns)
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "开始数据抽取任务",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "Transfer",
		}, es.(*Group), true, DataxController{}.Transfer)
	}
}
