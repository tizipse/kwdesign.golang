package site

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kwd/app/constant"
	"kwd/app/model"
	"kwd/app/request/admin/site"
	res "kwd/app/response/admin/site"
	"kwd/kernel/app"
	"kwd/kernel/authorize"
	"kwd/kernel/response"
	"strconv"
)

func DoAdminByCreate(ctx *gin.Context) {

	var request site.DoAdminByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var count int64

	tc := app.Database.Model(model.SysRole{})

	if !authorize.Root(authorize.Id(ctx)) {
		tc = tc.Where("`id`<>?", authorize.ROOT)
	}

	tc.Where("`id` IN (?)", request.Roles).Count(&count)

	if len(request.Roles) != int(count) {
		response.NotFound(ctx, "部分角色不存在")
		return
	}

	app.Database.Model(model.SysAdmin{}).Where("`mobile`=?", request.Mobile).Count(&count)

	if count > 0 {
		response.Fail(ctx, "该手机号已被注册")
		return
	}

	app.Database.Model(model.SysAdmin{}).Where("username = ?", request.Username).Count(&count)

	if count > 0 {
		response.Fail(ctx, "该用户名已被注册")
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	tx := app.Database.Begin()

	admin := model.SysAdmin{
		Nickname: request.Nickname,
		Password: string(password),
		IsEnable: request.IsEnable,
	}

	if !strutil.IsEmpty(request.Username) {
		admin.Username = &request.Username
	}

	if !strutil.IsEmpty(request.Mobile) {
		admin.Mobile = &request.Mobile
	}

	if ca := tx.Create(&admin); ca.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("创建失败：%v", ca.Error))
		return
	}

	var binds []model.SysAdminBindRole

	for _, item := range request.Roles {
		binds = append(binds, model.SysAdminBindRole{
			AdminId: admin.Id,
			RoleId:  item,
		})
	}

	if cb := tx.CreateInBatches(binds, 100); cb.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("创建失败：%v", cb.Error))
		return
	}

	if len(binds) > 0 {

		var items = make([]string, len(binds))

		for idx, item := range binds {
			items[idx] = authorize.NameByRole(item.RoleId)
		}

		if _, err := app.Casbin.AddRolesForUser(authorize.NameByAdmin(admin.Id), items); err != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("创建失败：%v", err))
			return
		}
	}

	tx.Commit()

	response.Success[any](ctx)
}

func DoAdminByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	if strconv.Itoa(authorize.Id(ctx)) == id {
		response.Fail(ctx, "无法修改自身账号")
		return
	}

	var request site.DoAdminByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var count int64

	tc := app.Database.Model(model.SysRole{})

	if !authorize.Root(authorize.Id(ctx)) {
		tc = tc.Where("`id`<>?", authorize.ROOT)
	}

	tc.Where("`id` in (?)", request.Roles).Count(&count)

	if len(request.Roles) != int(count) {
		response.NotFound(ctx, "部分角色不存在")
		return
	}

	app.Database.Model(model.SysAdmin{}).Where("`id`<>? and `mobile`=?", id, request.Mobile).Count(&count)

	if count > 0 {
		response.Fail(ctx, "该手机号已被注册")
		return
	}

	var admin model.SysAdmin

	fa := app.Database.Preload("BindRoles").First(&admin, id)

	if errors.Is(fa.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "该账号不存在")
		return
	} else if fa.Error != nil {
		response.NotFound(ctx, fmt.Sprintf("账号查找失败：%v", fa.Error))
		return
	}

	admin.Nickname = request.Nickname
	admin.Mobile = nil
	admin.IsEnable = request.IsEnable

	if !strutil.IsEmpty(request.Mobile) {
		admin.Mobile = &request.Mobile
	}

	if request.Password != "" {

		password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

		admin.Password = string(password)
	}

	var creates []model.SysAdminBindRole
	var deletes []int
	var del []int

	for _, item := range request.Roles {

		mark := true

		for _, value := range admin.BindRoles {
			if item == value.RoleId {
				mark = false
				break
			}
		}

		if mark {
			creates = append(creates, model.SysAdminBindRole{
				AdminId: admin.Id,
				RoleId:  item,
			})
		}
	}

	for _, item := range admin.BindRoles {

		mark := true

		for _, value := range request.Roles {
			if item.RoleId == value {
				mark = false
				break
			}
		}

		if mark {
			del = append(del, item.RoleId)
			deletes = append(deletes, item.Id)
		}
	}

	tx := app.Database.Begin()

	if ua := tx.Save(&admin); ua.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("修改失败：%v", ua.Error))
		return
	}

	if request.IsEnable != constant.IsEnableYes { //	用户禁用，删除缓存角色
		if _, err := app.Casbin.DeleteRolesForUser(authorize.NameByAdmin(admin.Id)); err != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("修改失败：%v", err))
			return
		}
	} else {

		var items = make([]string, len(request.Roles))

		for idx, item := range request.Roles {
			items[idx] = authorize.NameByRole(item)
		}

		if _, err := app.Casbin.AddRolesForUser(authorize.NameByAdmin(admin.Id), items); err != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("修改失败：%v", err))
			return
		}
	}

	if len(deletes) > 0 {

		var bindings model.SysAdminBindRole

		if db := tx.Where("`admin_id`=?", admin.Id).Delete(&bindings, deletes); db.Error != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("修改失败：%v", db.Error))
			return
		}

		if len(del) > 0 && request.IsEnable == constant.IsEnableYes { //	用户启用，结算需要删除的角色
			for _, item := range del {
				if _, err := app.Casbin.DeleteRoleForUser(authorize.NameByAdmin(admin.Id), authorize.NameByRole(item)); err != nil {
					tx.Rollback()
					response.Fail(ctx, fmt.Sprintf("修改失败：%v", err))
					return
				}
			}
		}
	}

	if len(creates) > 0 {

		if ca := tx.CreateInBatches(creates, 100); ca.Error != nil {
			tx.Rollback()
			response.Fail(ctx, ca.Error.Error())
			return
		}

		if len(creates) > 0 && request.IsEnable == constant.IsEnableYes { //	用户启用，处理需要新加的角色

			var items = make([]string, len(creates))

			for idx, item := range creates {
				items[idx] = authorize.NameByRole(item.RoleId)
			}

			if _, err := app.Casbin.AddRolesForUser(authorize.NameByAdmin(admin.Id), items); err != nil {
				tx.Rollback()
				response.Fail(ctx, fmt.Sprintf("修改失败：%v", err))
				return
			}
		}
	}

	tx.Commit()

	response.Success[any](ctx)
}

func ToAdminByPaginate(ctx *gin.Context) {

	var request site.ToAdminByPaginate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	tx := app.Database

	if !authorize.Root(authorize.Id(ctx)) {
		tx = tx.Where("not exists (?)", app.Database.
			Select("1").
			Model(model.SysAdminBindRole{}).
			Where(fmt.Sprintf("%s.id=%s.admin_id", model.TableSysAdmin, model.TableSysAdminBindRole)).
			Where("`role_id` = ?", authorize.ROOT),
		)
	}

	responses := response.Paginate[res.ToAdminByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx.Model(model.SysAdmin{}).Count(&responses.Total)

	if responses.Total > 0 {

		var admins []model.SysAdmin

		tx.
			Preload("BindRoles.Role").
			Order("`id` desc").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Find(&admins)

		responses.Data = make([]res.ToAdminByPaginate, len(admins))

		for index, item := range admins {

			responses.Data[index] = res.ToAdminByPaginate{
				Id:        item.Id,
				Nickname:  item.Nickname,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}

			if item.Username != nil {
				responses.Data[index].Username = *item.Username
			}

			if item.Mobile != nil {
				responses.Data[index].Mobile = *item.Mobile
			}

			for _, value := range item.BindRoles {

				responses.Data[index].Roles = append(responses.Data[index].Roles, res.ToAdminByPaginateOfRoles{
					Id:   value.Role.Id,
					Name: value.Role.Name,
				})
			}
		}
	}

	response.Success(ctx, responses)
}

func DoAdminByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID不存在"))
		return
	}

	if strconv.Itoa(authorize.Id(ctx)) == id {
		response.Fail(ctx, "无法删除自身账号")
		return
	}

	var admin model.SysAdmin

	fa := app.Database.First(&admin, id)

	if errors.Is(fa.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "账号不存在")
		return
	} else if fa.Error != nil {
		response.NotFound(ctx, fmt.Sprintf("账号查找失败：%v", fa.Error))
		return
	}

	tx := app.Database.Begin()

	if da := app.Database.Delete(&admin); da.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("账号删除失败：%v", da.Error))
		return
	}

	bind := model.SysAdminBindRole{AdminId: admin.Id}

	if db := tx.Delete(&bind, "`admin_id`=?", admin.Id); db.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("账号删除失败：%v", db.Error))
		return
	}

	if _, err := app.Casbin.DeleteRolesForUser(authorize.NameByAdmin(admin.Id)); err != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("账号删除失败：%v", err))
		return
	}

	tx.Commit()

	response.Success[any](ctx)
}

func DoAdminByEnable(ctx *gin.Context) {

	var request site.DoAdminByEnable

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	if authorize.Id(ctx) == request.Id {
		response.Fail(ctx, "无法操作自身账号")
		return
	}

	var admin model.SysAdmin

	fa := app.Database.First(&admin, request.Id)

	if errors.Is(fa.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "账号不存在")
		return
	} else if fa.Error != nil {
		response.Fail(ctx, fmt.Sprintf("账号查找失败：%v", fa.Error))
		return
	}

	admin.IsEnable = request.IsEnable

	tx := app.Database.Begin()

	if ua := app.Database.Save(&admin); ua.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("启禁失败：%v", ua.Error))
		return
	}

	if request.IsEnable == constant.IsEnableNo {
		if _, err := app.Casbin.DeleteRolesForUser(authorize.NameByAdmin(admin.Id)); err != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("启禁失败：%v", err))
			return
		}
	} else if request.IsEnable == constant.IsEnableYes {

		tx.Find(&admin.BindRoles, "`admin_id`=?", admin.Id)

		if len(admin.BindRoles) > 0 {

			var items []string

			for _, item := range admin.BindRoles {
				items = append(items, authorize.NameByRole(item.RoleId))
			}

			if _, err := app.Casbin.AddRolesForUser(authorize.NameByAdmin(admin.Id), items); err != nil {
				tx.Rollback()
				response.Fail(ctx, fmt.Sprintf("启禁失败：%v", err))
				return
			}
		}
	}

	tx.Commit()

	response.Success[any](ctx)
}
