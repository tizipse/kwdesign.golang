package site

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"gorm.io/gorm"
	"kwd/app/constant"
	"kwd/app/model"
	"kwd/app/request/admin/site"
	res "kwd/app/response/admin/site"
	authService "kwd/app/service/site/manage"
	"kwd/kernel/app"
	"kwd/kernel/authorize"
	"kwd/kernel/response"
)

func DoPermissionByCreate(ctx *gin.Context) {

	var request site.DoPermissionByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var module model.SysModule

	fm := app.Database.Where("`is_enable`=?", constant.IsEnableYes).First(&module, request.Module)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "模块不存在")
		return
	} else if fm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("模块查找失败：%v", fm.Error))
		return
	}

	var parent1, parent2 int

	var parent model.SysPermission

	if request.Parent > 0 {

		if app.Database.Find(&parent, request.Parent); parent.Id <= 0 {
			response.Fail(ctx, "父级权限不存在")
			return
		} else if parent.ParentI2 > 0 {
			response.Fail(ctx, "该权限已是最低等级，无法继续添加")
			return
		} else if parent.ParentI1 > 0 {

			if request.Method == "" || request.Path == "" {
				response.Fail(ctx, "接口不能为空")
				return
			}

			parent2 = parent.Id
			parent1 = parent.ParentI1
		} else {
			parent1 = parent.Id
		}
	}

	var permission model.SysPermission

	if request.Method != "" && request.Path != "" {

		var count int64

		app.Database.Model(model.SysPermission{}).Where("`method`=? and `path`=?", request.Method, request.Path).Count(&count)

		if count > 0 {
			response.Fail(ctx, "权限已存在")
			return
		}
	}

	permission = model.SysPermission{
		ModuleId: module.Id,
		ParentI1: parent1,
		ParentI2: parent2,
		Name:     request.Name,
		Slug:     request.Slug,
		Method:   request.Method,
		Path:     request.Path,
	}

	if cp := app.Database.Create(&permission); cp.Error != nil {
		response.Fail(ctx, fmt.Sprintf("添加失败：%v", cp.Error))
		return
	}

	response.Success[any](ctx)
}

func DoPermissionByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var request site.DoPermissionByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var permission model.SysPermission

	fp := app.Database.First(&permission, id)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "权限不存在")
		return
	} else if fp.Error != nil {
		response.Fail(ctx, fmt.Sprintf("权限查找失败：%v", fp.Error))
		return
	}

	method := permission.Method
	path := permission.Path

	if permission.ModuleId != request.Module {

		var module model.SysModule

		fm := app.Database.Where("`is_enable`=?", constant.IsEnableYes).First(&module, request.Module)

		if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
			response.NotFound(ctx, "模块不存在")
			return
		} else if fm.Error != nil {
			response.Fail(ctx, fmt.Sprintf("模块查找失败：%v", fm.Error))
			return
		}
	}

	var parent1, parent2 int

	var parent model.SysPermission

	if request.Parent > 0 {

		if app.Database.Find(&parent, request.Parent); parent.Id <= 0 {
			response.NotFound(ctx, "父级权限不存在")
			return
		} else if parent.ParentI2 > 0 {
			response.Fail(ctx, "该权限已是最低等级，无法继续添加")
			return
		} else if parent.ParentI1 > 0 {

			if request.Method == "" || request.Path == "" {
				response.Fail(ctx, "接口不能为空")
				return
			}

			parent2 = parent.Id
			parent1 = parent.ParentI1
		} else {
			parent1 = parent.Id
		}
	}

	if request.Method != "" && request.Path != "" {

		var count int64

		app.Database.Model(model.SysPermission{}).Where("`id`<>? and `method`=? and `path`=?", id, request.Method, request.Path).Count(&count)

		if count > 0 {
			response.Fail(ctx, "权限已存在")
			return
		}
	}

	permission.ParentI1 = parent1
	permission.ParentI2 = parent2
	permission.Name = request.Name
	permission.Slug = request.Slug
	permission.Method = request.Method
	permission.Path = request.Path

	tx := app.Database.Begin()

	if up := app.Database.Save(&permission); up.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("修改失败：%v", up.Error))
		return
	}

	if method != request.Method || path != request.Path { //	变更权限
		if method != "" || path != "" {
			if _, err := app.Casbin.DeletePermission(method, path); err != nil {
				tx.Rollback()
				response.Fail(ctx, fmt.Sprintf("修改失败：%v", err))
				return
			}
		}

		if request.Method != "" || request.Path != "" {

			var bindings []model.SysRoleBindPermission

			if tx.Find(&bindings, "`permission_id`=?", permission.Id); len(bindings) > 0 {
				for _, item := range bindings {
					if _, err := app.Casbin.AddPermissionForUser(authorize.NameByRole(item.RoleId), permission.Method, permission.Path); err != nil {
						tx.Rollback()
						response.Fail(ctx, fmt.Sprintf("修改失败：%v", err != nil))
						return
					}
				}
			}
		}
	}

	tx.Commit()

	response.Success[any](ctx)
}

func DoPermissionByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var permission model.SysPermission

	fp := app.Database.First(&permission, id)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "权限不存在")
		return
	} else if fp.Error != nil {
		response.Fail(ctx, fmt.Sprintf("权限查找失败：%v", fp.Error))
		return
	}

	tx := app.Database.Begin()

	if dp := tx.Delete(&permission); dp.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("删除失败：%v", dp.Error))
		return
	}

	if permission.Method != "" && permission.Path != "" {
		if _, err := app.Casbin.DeletePermission(permission.Method, permission.Path); err != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("删除失败：%v", err))
			return
		}
	} else if permission.ParentI1 > 0 {

		var children []model.SysPermission

		if tx.Find(&children, "`parent_i2`=? and `method`<>? and `path`<>?", permission.Id, "", ""); len(children) > 0 {

			if dp2 := tx.Delete(&model.SysPermission{}, "`parent_i2`=?", permission.Id); dp2.Error != nil {
				tx.Rollback()
				response.Fail(ctx, fmt.Sprintf("删除失败：%v", dp2.Error))
				return
			}

			for _, item := range children {
				if _, err := app.Casbin.DeletePermission(item.Method, item.Path); err != nil {
					tx.Rollback()
					response.Fail(ctx, fmt.Sprintf("删除失败：%v", err))
					return
				}
			}
		}
	} else {

		var children []model.SysPermission

		if tx.Find(&children, "`parent_i1`=? and `method`<>? and `path`<>?", permission.Id, "", ""); len(children) > 0 {

			if dp1 := tx.Delete(&model.SysPermission{}, "`parent_i1`=?", permission.Id); dp1.Error != nil {
				tx.Rollback()
				response.Fail(ctx, fmt.Sprintf("删除失败：%v", dp1.Error))
				return
			}

			for _, item := range children {
				if _, err := app.Casbin.DeletePermission(item.Method, item.Path); err != nil {
					tx.Rollback()
					response.Fail(ctx, fmt.Sprintf("删除失败：%v", err))
					return
				}
			}
		}
	}

	tx.Commit()

	response.Success[any](ctx)
}

func ToPermissionByTree(ctx *gin.Context) {

	var request site.ToPermissionByTree

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	responses := authService.TreePermission(request.Module, false, false)

	response.Success[[]res.TreePermission](ctx, responses)
}

func ToPermissionByParents(ctx *gin.Context) {

	var request site.ToPermissionByTree

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	responses := authService.TreePermission(request.Module, true, true)

	response.Success[[]res.TreePermission](ctx, responses)
}

func ToPermissionBySelf(ctx *gin.Context) {

	responses := make([]any, 0)

	var results []res.TreePermission
	var modules []res.TreePermission

	var children, children1, children2 []model.SysPermission

	var permissions []model.SysPermission

	tx := app.Database.
		Preload("Module").
		Preload("Parent1").
		Preload("Parent2").
		Where("method<>? and path<>?", "", "")

	if !authorize.Root(authorize.Id(ctx)) {
		tx = tx.
			Where("exists (?)", app.Database.
				Select("1").
				Model(model.SysRoleBindPermission{}).
				Where(fmt.Sprintf("`%s`.`id`=`%s`.`permission_id`", model.TableSysPermission, model.TableSysRoleBindPermission)).
				Where("exists (?)", app.Database.
					Select("1").
					Model(model.SysAdminBindRole{}).
					Where(fmt.Sprintf("`%s`.`role_id`=`%s`.`role_id` and `%s`.`admin_id`=?", model.TableSysRoleBindPermission, model.TableSysAdminBindRole, model.TableSysAdminBindRole), authorize.Id(ctx)),
				),
			)
	}

	tx.Find(&permissions)

	if len(permissions) > 0 {
		for _, item := range permissions {
			mark := true
			for _, value := range modules {
				if item.Module.Id == value.Id {
					mark = false
				}
			}
			if mark {
				modules = append(modules, res.TreePermission{
					Id:   item.Module.Id,
					Name: item.Module.Name,
				})
			}
		}

		child := map[int]model.SysPermission{}
		child1 := map[int]model.SysPermission{}
		child2 := map[int]model.SysPermission{}

		for _, item := range permissions {
			if item.ParentI2 > 0 {
				if item.Parent1 != nil {
					item.Parent1.Module = item.Module
					child[item.ParentI1] = *item.Parent1
				}
				if item.Parent2 != nil {
					item.Parent2.Module = item.Module
					child1[item.ParentI2] = *item.Parent2
				}
				child2[item.Id] = item
			} else if item.ParentI1 > 0 {
				if item.Parent1 != nil {
					item.Parent1.Module = item.Module
					child[item.ParentI1] = *item.Parent1
				}
				child1[item.Id] = item
			} else {
				child[item.Id] = item
			}
		}

		for _, item := range child {
			children = append(children, item)
		}
		for _, item := range child1 {
			children1 = append(children1, item)
		}
		for _, item := range child2 {
			children2 = append(children2, item)
		}
	}

	if len(modules) > 0 && (len(children2) > 0 || len(children1) > 0 || len(children) > 0) {

		for _, item := range modules {

			//	处理模块一层
			child := res.TreePermission{
				Id:   item.Id,
				Name: item.Name,
			}

			for _, value := range children {
				if child.Id == value.Module.Id {

					//	处理第二层
					child1 := res.TreePermission{
						Id:   value.Id,
						Name: value.Name,
					}

					for _, val := range children1 {
						if child1.Id == val.ParentI1 {

							//	处理第三层
							child2 := res.TreePermission{
								Id:   val.Id,
								Name: val.Name,
							}

							for _, v := range children2 {
								if child2.Id == v.ParentI2 {

									//	处理第四层
									child3 := res.TreePermission{
										Id:   v.Id,
										Name: v.Name,
									}

									if child3.Children != nil && len(child3.Children) > 0 || v.Method != "" && v.Path != "" {
										child2.Children = append(child2.Children, child3)
									}
								}
							}

							if child2.Children != nil && len(child2.Children) > 0 || val.Method != "" && val.Path != "" {
								child1.Children = append(child1.Children, child2)
							}
						}
					}

					if child1.Children != nil && len(child1.Children) > 0 || value.Method != "" && value.Path != "" {
						child.Children = append(child.Children, child1)
					}
				}
			}

			if child.Children != nil && len(child.Children) > 0 {
				results = append(results, child)
			}
		}

		for _, item := range results {
			responses = append(responses, item)
		}
	}

	response.Success(ctx, responses)
}
