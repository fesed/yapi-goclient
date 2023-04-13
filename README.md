# yapi-goclient

### 功能
go实现的yapi客户端，用管理员账号模拟登录，调用yapi的接口

### 版本
- go: 1.18
- yapi: 1.11.0

### 用法
```golang
package main

import (
	"fmt"

	yapi "github.com/fesed/yapi-goclient"
	"github.com/fesed/yapi-goclient/model"

	"github.com/samber/lo"
)

func getClient() *yapi.Client {
	c, err := yapi.NewClient(yapi.Options{
		Url: "http://127.0.0.1:3000",
		// 管理员账号，直接修改数据库user表的role字段为admin即可改为管理员
		Username: "admin@test.com",
		Password: "12345678",
	})
	if err != nil {
		panic(err)
	}

	return c
}

func main() {
	var (
		// 测试的两个用户必须提前添加到数据库中
		ownerName  = "testuser1"
		memberName = "testuser2"

		mockName        = "test"
		mockProjectname = "testProject"
		mockDesc        = "for test"
		mockString      = "updated"
	)
	// init client
	c := getClient()

	memberId, err := c.GetUserIdByKey(memberName)
	if err != nil {
		panic(err)
	}

	// group
	groupId, err := c.AddGroupWithOwnernames(mockName, mockDesc, ownerName)
	if err != nil {
		panic(err)
	}
	defer c.DeleteGroup(groupId)

	groups, err := c.ListGroup()
	if err != nil {
		panic(err)
	}

	isGroupExist := false
	for _, g := range groups {
		if g.GroupName == mockName {
			isGroupExist = true
			break
		}
	}
	if !isGroupExist {
		panic("not found group")
	}

	// group member
	if err := c.AddGroupMember(groupId, []int{memberId}, model.RoleOwner); err != nil {
		panic(err)
	}

	members, err := c.GetGroupMemberList(groupId)
	if err != nil {
		panic(err)
	}
	isMemberExist := false
	for _, m := range members {
		if m.UID == memberId {
			isMemberExist = true
			break
		}
	}
	if !isMemberExist {
		panic("not found member")
	}

	if err := c.ChangeGroupMemberRole(groupId, memberId, model.RoleGuest); err != nil {
		if err != nil {
			panic(err)
		}
	}

	newMembers, err := c.GetGroupMemberList(groupId)
	if err != nil {
		panic(err)
	}
	for _, m := range newMembers {
		if m.UID == memberId {
			if m.Role != string(model.RoleGuest) {
				panic("wrong role")
			}
			break
		}
	}

	if err := c.DeleteGroupMember(groupId, memberId); err != nil {
		panic(err)
	}
	newMembers1, err := c.GetGroupMemberList(groupId)
	if err != nil {
		panic(err)
	}
	for _, m := range newMembers1 {
		if m.UID == memberId {
			panic("the member should be deleted")
		}
	}

	// project
	projectId, err := c.AddProject(mockProjectname, mockDesc, "", groupId)
	if err != nil {
		panic(err)
	}

	projects, err := c.ListProject(groupId)
	if err != nil {
		panic(err)
	}

	IsProjectExist := false
	for _, p := range projects {
		if p.Name == mockProjectname {
			IsProjectExist = true
			break
		}
	}
	if !IsProjectExist {
		panic("not found project")
	}

	projectInfo, err := c.GetProject(projectId)
	if err != nil {
		panic(err)
	}
	if projectInfo.Name != mockProjectname {
		panic("wrong name")
	}

	// interface
	catId, err := c.AddCat(mockName, mockDesc, projectId)
	if err != nil {
		panic(err)
	}
	if err := c.UpdateCat(catId, mockName, mockString); err != nil {
		panic(err)
	}
	menus, err := c.ListMenu(projectId)
	if err != nil {
		panic(err)
	}

	isCatExist := false
	for _, m := range menus {
		if m.ID == catId {
			if m.Desc != mockString {
				panic("update cat failed")
			}
			isCatExist = true
			break
		}
	}
	if !isCatExist {
		panic("not found cat")
	}

	interfaceId, err := c.AddInterface("POST", catId, projectId, mockName, "/api/v1/test1")
	if err != nil {
		panic(err)
	}

	interfac, err := c.GetInterface(interfaceId)
	if err != nil {
		panic(err)
	}
	if interfac.Method != "POST" {
		panic("wrong method")
	}

	if err := c.DeleteInterface(interfaceId); err != nil {
		panic(err)
	}
	if _, err := c.GetInterface(interfaceId); err != nil {
		if err.Error() != "不存在的" {
			panic("wrong error")
		}
	}

	// log
	groupLogs, count, err := c.GetLogs(model.LogTypeGroup, groupId, 1, 10)
	if err != nil {
		panic(err)
	}
	_ = count
	lo.ForEach(groupLogs, func(item *model.Log, _ int) {
		fmt.Println(item)
	})

	// delete
	if err := c.DeleteProject(projectId); err != nil {
		panic(err)
	}
}
```

