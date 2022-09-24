package site

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"gorm.io/gorm"
	"kwd/app/model"
	"kwd/app/request/admin/site"
	res "kwd/app/response/admin/site"
	"kwd/kernel/app"
	"kwd/kernel/authorize"
	"kwd/kernel/response"
	"strconv"
)

func DoRoleByCreate(ctx *gin.Context) {

	var request site.DoRoleByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var permissionsIds []int

	var modules, children1, children2, children3 []int

	for _, item := range request.Permissions {
		if len(item) >= 4 {
			children3 = append(children3, item[3])
		} else if len(item) >= 3 {
			children2 = append(children2, item[2])
		} else if len(item) >= 2 {
			children1 = append(children1, item[1])
		} else if len(item) >= 1 {
			modules = append(modules, item[0])
		}
	}

	if len(modules) > 0 {

		var permissions []model.SysPermission

		app.Database.Find(&permissions, "`module_id` in (?) and `method`<>? and `path`<>?", modules, "", "")

		for _, item := range permissions {
			permissionsIds = append(permissionsIds, item.Id)
		}
	}
	if len(children3) > 0 {

		var permissions []model.SysPermission

		app.Database.Find(&permissions, "`id` in (?) and `method`<>? and `path`<>?", children3, "", "")

		for _, item := range permissions {
			permissionsIds = append(permissionsIds, item.Id)
		}
	}
	if len(children2) > 0 {
		var permissions []model.SysPermission

		if app.Database.Find(&permissions, "`parent_i2` in (?) and `method`<>? and `path`<>?", children2, "", ""); len(permissions) > 0 {
			for _, item := range permissions {
				permissionsIds = append(permissionsIds, item.Id)
			}
		} else if app.Database.Find(&permissions, "`id` in (?) and `method`<>? and `path`<>?", children2, "", ""); len(permissions) > 0 {
			for _, item := range permissions {
				permissionsIds = append(permissionsIds, item.Id)
			}
		}
	}
	if len(children1) > 0 {

		var permissions []model.SysPermission

		if app.Database.Find(&permissions, "`parent_i1` in (?) and `method`<>? and `path`<>?", children1, "", ""); len(permissions) > 0 {
			for _, item := range permissions {
				permissionsIds = append(permissionsIds, item.Id)
			}
		} else if app.Database.Find(&permissions, "`id` in (?) and `method`<>? and `path`<>?", children1, "", ""); len(permissions) > 0 {
			for _, item := range permissions {
				permissionsIds = append(permissionsIds, item.Id)
			}
		}
	}

	var temp = make(map[int]int, len(permissionsIds))

	for _, item := range permissionsIds {
		temp[item] = item
	}

	var bind []int

	for _, item := range temp {
		bind = append(bind, item)
	}

	if len(bind) <= 0 {
		response.Fail(ctx, "可用权限不能为空")
		return
	}

	tx := app.Database.Begin()

	role := model.SysRole{
		Name:    request.Name,
		Summary: request.Summary,
	}

	if cr := tx.Create(&role); cr.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("添加失败：%v", cr.Error))
		return
	}

	var binds []model.SysRoleBindPermission

	for _, item := range bind {
		binds = append(binds, model.SysRoleBindPermission{
			RoleId:       role.Id,
			PermissionId: item,
		})
	}

	if cb := tx.CreateInBatches(binds, 100); cb.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("添加失败：%v", tx.Error))
		return
	}

	var permissions []model.SysRoleBindPermission

	tx.
		Preload("Permission", app.Database.Where("`method`<>? and `path`<>?", "", "")).
		Where("`role_id`=?", role.Id).
		Find(&permissions)

	if len(permissions) > 0 {

		var items = make([][]string, len(permissions))

		for idx, item := range permissions {
			items[idx] = []string{item.Permission.Method, item.Permission.Path}
		}

		if _, err := app.Casbin.AddPermissionsForUser(authorize.NameByRole(role.Id), items...); err != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("添加失败：%v", err))
			return
		}
	}

	tx.Commit()

	response.Success[any](ctx)
}

func DoRoleByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	if id == strconv.Itoa(authorize.ROOT) {
		response.Fail(ctx, "开发组权限，无法修改")
		return
	}

	var request site.DoRoleByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var role model.SysRole

	fr := app.Database.First(&role, id)

	if errors.Is(fr.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "角色不存在")
		return
	} else if fr.Error != nil {
		response.Fail(ctx, fmt.Sprintf("角色查找失败：%v", fr.Error))
		return
	}

	var permissionsIds []int

	var modules, children1, children2, children3 []int

	for _, item := range request.Permissions {
		if len(item) >= 4 {
			children3 = append(children3, item[3])
		} else if len(item) >= 3 {
			children2 = append(children2, item[2])
		} else if len(item) >= 2 {
			children1 = append(children1, item[1])
		} else if len(item) >= 1 {
			modules = append(modules, item[0])
		}
	}

	if len(modules) > 0 {

		var permissions []model.SysPermission

		app.Database.Find(&permissions, "`module_id` in (?) and `method`<>? and `path`<>?", modules, "", "")

		for _, item := range permissions {
			permissionsIds = append(permissionsIds, item.Id)
		}
	}
	if len(children3) > 0 {

		var permissions []model.SysPermission

		app.Database.Find(&permissions, "`id` in (?) and `method`<>? and `path`<>?", children3, "", "")

		for _, item := range permissions {
			permissionsIds = append(permissionsIds, item.Id)
		}
	}
	if len(children2) > 0 {
		var permissions []model.SysPermission

		if app.Database.Find(&permissions, "`parent_i2` in (?) and `method`<>? and `path`<>?", children2, "", ""); len(permissions) > 0 {
			for _, item := range permissions {
				permissionsIds = append(permissionsIds, item.Id)
			}
		} else if app.Database.Find(&permissions, "`id` in (?) and `method`<>? and `path`<>?", children2, "", ""); len(permissions) > 0 {
			for _, item := range permissions {
				permissionsIds = append(permissionsIds, item.Id)
			}
		}
	}
	if len(children1) > 0 {

		var permissions []model.SysPermission

		if app.Database.Find(&permissions, "`parent_i1` in (?) and `method`<>? and `path`<>?", children1, "", ""); len(permissions) > 0 {
			for _, item := range permissions {
				permissionsIds = append(permissionsIds, item.Id)
			}
		} else if app.Database.Find(&permissions, "`id` in (?) and `method`<>? and `path`<>?", children1, "", ""); len(permissions) > 0 {
			for _, item := range permissions {
				permissionsIds = append(permissionsIds, item.Id)
			}
		}
	}

	var temp = make(map[int]int, len(permissionsIds))

	for _, item := range permissionsIds {
		temp[item] = item
	}

	var bind []int

	for _, item := range temp {
		bind = append(bind, item)
	}

	if len(bind) <= 0 {
		response.Fail(ctx, "可用权限不能为空")
		return
	}

	role.Name = request.Name
	role.Summary = request.Summary

	app.Database.Find(&role.BindPermissions, "`role_id`=?", role.Id)

	var creates []model.SysRoleBindPermission
	var deletes []int

	if len(bind) > 0 {

		for _, item := range bind {

			mark := true

			for _, value := range role.BindPermissions {
				if item == value.PermissionId {
					mark = false
					break
				}
			}

			if mark {
				creates = append(creates, model.SysRoleBindPermission{
					RoleId:       role.Id,
					PermissionId: item,
				})
			}
		}

		for _, item := range role.BindPermissions {

			mark := true

			for _, value := range bind {
				if item.PermissionId == value {
					mark = false
					break
				}
			}

			if mark {
				deletes = append(deletes, item.Id)
			}
		}
	}

	tx := app.Database.Begin()

	if ur := tx.Save(&role); ur.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("修改失败：%v", ur.Error != nil))
		return
	}

	if len(creates) > 0 {

		if cb := tx.CreateInBatches(creates, 100); cb.Error != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("修改失败：%v", cb.Error))
			return
		}

		var ids []int

		for _, item := range creates {
			ids = append(ids, item.PermissionId)
		}

		var permissions []model.SysPermission

		tx.Where("method<>? and path<>?", "", "").Find(&permissions, ids)

		if len(permissions) > 0 {
			var items [][]string
			for _, item := range permissions {
				items = append(items, []string{item.Method, item.Path})
			}
			if _, err := app.Casbin.AddPermissionsForUser(authorize.NameByRole(role.Id), items...); err != nil {
				tx.Rollback()
				response.Fail(ctx, fmt.Sprintf("修改失败：%v", err))
				return
			}
		}
	}

	if len(deletes) > 0 {

		if db := tx.Where("`role_id`=?", role.Id).Delete(&model.SysRoleBindPermission{}, deletes); db.Error != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("修改失败：%v", db.Error))
			return
		}
	}

	if len(deletes) > 0 {
		if _, err := app.Casbin.DeletePermissionsForUser(authorize.NameByRole(role.Id)); err != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("修改失败：%v", err))
			return
		}
	}

	if len(creates) > 0 || len(deletes) > 0 {

		var permissions []model.SysRoleBindPermission

		tx.
			Preload("Permission", app.Database.Where("method <> ? and path <> ?", "", "")).
			Find(&permissions, "`role_id`=?", role.Id)

		if len(permissions) > 0 {

			var items = make([][]string, len(permissions))

			for idx, item := range permissions {
				items[idx] = []string{item.Permission.Method, item.Permission.Path}
			}

			if _, err := app.Casbin.AddPermissionsForUser(authorize.NameByRole(role.Id), items...); err != nil {
				tx.Rollback()
				response.Fail(ctx, fmt.Sprintf("修改失败：%v", err))
				return
			}
		}
	}

	tx.Commit()

	response.Success[any](ctx)
}

func DoRoleByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	if id == strconv.Itoa(authorize.ROOT) {
		response.Fail(ctx, "开发组权限，无法修改")
		return
	}

	var role model.SysRole

	fr := app.Database.First(&role, id)

	if errors.Is(fr.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "角色不存在")
		return
	} else if fr.Error != nil {
		response.Fail(ctx, fmt.Sprintf("角色查找失败：%v", fr.Error))
		return
	}

	tx := app.Database.Begin()

	if dr := app.Database.Delete(&role); dr.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("删除失败：%v", dr.Error))
		return
	}

	//bind := model.SysRoleBindPermission{RoleId: role.Id}

	if db := tx.Delete(&model.SysRoleBindPermission{}, "`role_id`=?", role.Id); db.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("删除失败：%v", db.Error))
		return
	}

	if _, err := app.Casbin.DeleteRole(authorize.NameByRole(role.Id)); err != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("删除失败：%v", err))
		return
	}

	tx.Commit()

	response.Success[any](ctx)
}

func ToRoleByPaginate(ctx *gin.Context) {

	var request site.ToRoleByPaginate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	tx := app.Database.Where("`id`<>?", authorize.ROOT)

	responses := response.Paginate[res.ToRoleByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx.Model(model.SysRole{}).Count(&responses.Total)

	if responses.Total > 0 {

		var roles []model.SysRole

		tx.
			Preload("BindPermissions.Permission").
			Offset(request.GetOffset()).
			Limit(request.GetLimit()).
			Order("`id` desc").
			Find(&roles)

		responses.Data = make([]res.ToRoleByPaginate, len(roles))

		for index, item := range roles {

			responses.Data[index] = res.ToRoleByPaginate{
				Id:        item.Id,
				Name:      item.Name,
				Summary:   item.Summary,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}

			for _, value := range item.BindPermissions {
				var permissions []int
				if value.Permission.ModuleId > 0 {
					permissions = append(permissions, value.Permission.ModuleId)
				}
				if value.Permission.ParentI1 > 0 {
					permissions = append(permissions, value.Permission.ParentI1)
				}
				if value.Permission.ParentI2 > 0 {
					permissions = append(permissions, value.Permission.ParentI2)
				}
				if value.PermissionId > 0 {
					permissions = append(permissions, value.PermissionId)
				}
				responses.Data[index].Permissions = append(responses.Data[index].Permissions, permissions)
			}
		}
	}

	response.Success[any](ctx, responses)
}

func ToRoleByEnable(ctx *gin.Context) {

	var roles []model.SysRole

	tx := app.Database

	if !authorize.Root(authorize.Id(ctx)) {
		tx.Where("`role_id`<>?", authorize.ROOT)
	}

	tx.Find(&roles)

	responses := make([]res.ToRoleByOnline, len(roles))

	for index, item := range roles {
		responses[index] = res.ToRoleByOnline{
			Id:   item.Id,
			Name: item.Name,
		}
	}

	response.Success(ctx, responses)
}
