package api

import (
	"context"
	"errors"
	"github.com/1340691923/ElasticView/platform-basic-libs/escache"
	"github.com/1340691923/ElasticView/platform-basic-libs/response"
	es2 "github.com/1340691923/ElasticView/platform-basic-libs/service/es"
	"github.com/1340691923/ElasticView/platform-basic-libs/util"
	"github.com/cch123/elasticsql"
	. "github.com/gofiber/fiber/v2"
)

//Es 基本操作
type EsController struct {
	BaseController
}

// Ping
func (this EsController) PingAction(ctx *Ctx) error {
	esConnect := new(escache.EsConnect)
	err := ctx.BodyParser(esConnect)
	if err != nil {
		return this.Error(ctx, err)
	}

	if esConnect.Pwd != ""{
		pwd,decrptErr := escache.EsPwdESBDecrypt(esConnect.Pwd)
		if decrptErr == nil{
			esConnect.Pwd = pwd
		}
	}

	switch esConnect.Version {
	case 6:
		esClient, err := escache.NewEsClientV6(esConnect)
		if err != nil {
			return this.Error(ctx, err)
		}
		data, _, err := esClient.Ping(esConnect.Ip).Do(context.Background())
		if err != nil {
			return this.Error(ctx, err)
		}
		if data.Version.Number == "" {
			return this.Error(ctx, errors.New("ES地址OK，但是密码验证失败"))
		}
		return this.Success(ctx, response.OperateSuccess, data)
	case 7:
		esClient, err := escache.NewEsClientV7(esConnect)
		if err != nil {
			return this.Error(ctx, err)
		}
		data, _, err := esClient.Ping(esConnect.Ip).Do(context.Background())
		if err != nil {
			return this.Error(ctx, err)
		}
		if data.Version.Number == "" {
			return this.Error(ctx, errors.New("ES地址OK，但是密码验证失败"))
		}
		return this.Success(ctx, response.OperateSuccess, data)
	case 8:
		esClient, err := escache.NewEsClientV8(esConnect)
		if err != nil {
			return this.Error(ctx, err)
		}
		data, _, err := esClient.Ping(esConnect.Ip).Do(context.Background())
		if err != nil {
			return this.Error(ctx, err)
		}
		if data.Version.Number == "" {
			return this.Error(ctx, errors.New("ES地址OK，但是密码验证失败"))
		}
		return this.Success(ctx, response.OperateSuccess, data)
	default:

	}

	return this.Error(ctx, errors.New("版本暂时只支持6.x,7.x,8.x"))

}

// Es 的CAT API
func (this EsController) CatAction(ctx *Ctx) error {

	esCat := new(escache.EsCat)
	err := ctx.BodyParser(&esCat)
	if err != nil {
		return this.Error(ctx, err)
	}
	esConnect, err := escache.GetEsClientByID(esCat.EsConnect)
	if err != nil {
		return this.Error(ctx, err)
	}

	esService, err := es2.NewEsService(esConnect)
	if err != nil {
		return this.Error(ctx, err)
	}
	return esService.Cat(ctx, esCat)

}

func (this EsController) RunDslAction(ctx *Ctx) error {

	esRest := new(escache.EsRest)
	err := ctx.BodyParser(&esRest)
	if err != nil {
		return this.Error(ctx, err)
	}
	esConnect, err := escache.GetEsClientByID(esRest.EsConnect)

	if err != nil {
		return this.Error(ctx, err)
	}

	esService, err := es2.NewEsService(esConnect)
	if err != nil {
		return this.Error(ctx, err)
	}
	return esService.RunDsl(ctx, esRest)
}

// SQL 转换为 DSL
func (this EsController) SqlToDslAction(ctx *Ctx) error {
	sql := ctx.FormValue("sql")
	dsl, table, err := elasticsql.ConvertPretty(sql)
	if err != nil {
		return this.Error(ctx, err)
	}
	return this.Success(ctx, "转换成功!", util.Map{
		"dsl":       dsl,
		"tableName": table,
	})
}

// 一些索引的操作
func (this EsController) OptimizeAction(ctx *Ctx) error {
	esOptimize := new(escache.EsOptimize)
	err := ctx.BodyParser(&esOptimize)
	if err != nil {
		return this.Error(ctx, err)
	}
	esConnect, err := escache.GetEsClientByID(esOptimize.EsConnect)
	if err != nil {
		return this.Error(ctx, err)
	}

	esService, err := es2.NewEsService(esConnect)
	if err != nil {
		return this.Error(ctx, err)
	}
	return esService.Optimize(ctx, esOptimize)
}

// 将索引恢复为可写状态   由于不可抗力，ES禁止写后，默认不会自动恢复
func (this EsController) RecoverCanWrite(ctx *Ctx) error {
	esConnectID := new(escache.EsConnectID)
	err := ctx.BodyParser(&esConnectID)
	if err != nil {
		return this.Error(ctx, err)
	}
	esConnect, err := escache.GetEsClientByID(esConnectID.EsConnectID)
	if err != nil {
		return this.Error(ctx, err)
	}
	esService, err := es2.NewEsService(esConnect)
	if err != nil {
		return this.Error(ctx, err)
	}
	return esService.RecoverCanWrite(ctx)

}
