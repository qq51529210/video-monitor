package weekplans

import (
	"net/http"

	"recordplan/db"
	"recordplan/webapi/internal"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

type patchReq struct {
	// 名称
	Name *string `json:"name" binding:"omitempty,max=32"`
	// 是否禁用
	// 0: 禁用
	// 1: 启用
	Enable *int8 `json:"enable" binding:"omitempty,oneof=0 1"`
	// 保存的天数
	SaveDays *uint32 `json:"saveDays" binding:"omitempty,min=1"`
	// 是一个 RecordTime 的 JSON 数组
	Peroids [][]*db.TimePeroid `json:"peroids" binding:"omitempty,min=1,max=7,dive,timePeroid"`
}

// @Summary	修改
// @Tags		周计划
// @Param		id		path	string		true	"WeekPlan.ID"
// @Param		data	body	patchReq	true	"数据"
// @Accept		json
// @Produce	json
// @Success	201	{object}	internal.RowResult
// @Failure	400	{object}	internal.Error
// @Failure	500	{object}	internal.Error
// @Router		/week_plans/{id} [patch]
func patch(ctx *gin.Context) {
	// 参数
	var id internal.IDPath[string]
	err := ctx.ShouldBindJSON(&id.ID)
	if err != nil {
		internal.Handle400(ctx, err)
		return
	}
	var req patchReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		internal.Handle400(ctx, err)
		return
	}
	// 参数是空的
	if util.IsNilOrEmpty(&req) {
		// 返回
		ctx.JSON(http.StatusCreated, &internal.RowResult{
			Row: 0,
		})
		return
	}
	// 数据库
	var model db.WeekPlan
	util.CopyStruct(&model, &req)
	model.ID = id.ID
	if len(req.Peroids) > 0 {
		model.Peroids = jsonTimePeroid(req.Peroids)
	}
	rows, err := db.UpdateWeekPlan(&model)
	if err != nil {
		internal.HandleDB500(ctx, err)
		return
	}
	// 返回
	ctx.JSON(http.StatusCreated, &internal.RowResult{
		Row: rows,
	})
}
