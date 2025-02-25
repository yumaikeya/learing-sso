package utils

import (
	"angya-backend/pkg/constants"
	"context"
	"slices"
	"strings"
)

func GetOrganizationId(ctx context.Context) string {
	organizationId, ok := ctx.Value(constants.ORGANIZATION_ID).(string)
	if !ok {
		panic("Can't be got organizationId!")
	}
	return organizationId
}

func GetUserId(ctx context.Context) string {
	userId, ok := ctx.Value(constants.USER_ID).(string)
	if !ok {
		panic("Can't be got userId!")
	}
	return userId
}

func GetUserRole(ctx context.Context) string {
	userRole, ok := ctx.Value(constants.USER_ROLE).(string)
	if !ok {
		return constants.USER_ROLE_PUBLIC
	}
	return userRole
}

func ValidateOrganizationPermission(ctx context.Context) (permit bool) {
	permission, ok := ctx.Value(constants.ORGANIZATION_PERMISSION).(string)
	if !ok {
		return
	}

	// wild cardを含んでいる時
	if strings.Contains(permission, constants.ORGANIZATION_PERMISSION_ALL) {
		permit = true
	}
	// UT permissionが設定されている時
	if strings.Contains(permission, constants.ORGANIZATION_PERMISSION_UT) {
		permit = true
	}

	return
}

type PoiSortByStr string

func (s PoiSortByStr) IsContain() bool {
	return slices.Contains([]PoiSortByStr{constants.POIS_SORT_BY_TIMESTAMP, constants.POIS_SORT_BY_CRITICALITY}, s)
}

type TenantStr string

func (s TenantStr) IsContain() bool {
	return strings.Contains(string(s), constants.PLANNER_TENANT)
}
