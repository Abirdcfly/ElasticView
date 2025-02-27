package router

import (
	. "github.com/1340691923/ElasticView/api"
	api_config "github.com/1340691923/ElasticView/platform-basic-libs/api_config"
	. "github.com/gofiber/fiber/v2"
)

// ES备份 路由
func runEsBackUp(app *App) {
	apiRouterConfig := api_config.NewApiRouterConfig()
	const AbsolutePath = "/api/backUp"
	backUp := app.Group(AbsolutePath)
	{
		esBackUpController := &EsBackUpController{}
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "快照仓库列表",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "SnapshotRepositoryListAction",
		}, backUp.(*Group), true, esBackUpController.SnapshotRepositoryListAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "新建快照仓库",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "SnapshotCreateRepositoryAction",
		}, backUp.(*Group), true, esBackUpController.SnapshotCreateRepositoryAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "删除快照仓库",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "SnapshotDeleteRepositoryAction",
		}, backUp.(*Group), true, esBackUpController.SnapshotDeleteRepositoryAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "快照列表",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "SnapshotListAction",
		}, backUp.(*Group), true, esBackUpController.SnapshotListAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "清理快照仓库",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "CleanupeRepositoryAction",
		}, backUp.(*Group), true, esBackUpController.CleanupeRepositoryAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "创建快照",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "CreateSnapshotAction",
		}, backUp.(*Group), true, esBackUpController.CreateSnapshotAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "删除快照",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "SnapshotDeleteAction",
		}, backUp.(*Group), true, esBackUpController.SnapshotDeleteAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "快照详情",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "SnapshotDetailAction",
		}, backUp.(*Group), true, esBackUpController.SnapshotDetailAction)

		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "将索引恢复至快照时状态",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "SnapshotRestoreAction",
		}, backUp.(*Group), true, esBackUpController.SnapshotRestoreAction)
		apiRouterConfig.MountApi(api_config.MountApiBasePramas{
			Remark:       "得到快照状态",
			Method:       api_config.MethodPost,
			AbsolutePath: AbsolutePath,
			RelativePath: "SnapshotStatusAction",
		}, backUp.(*Group), true, esBackUpController.SnapshotStatusAction)
	}
}
