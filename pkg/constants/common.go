package constants

const (
	ENV_LOCAL = "LOCAL"
	PROJECT   = "project"
	PLAN      = "plan"
	FLIGHT    = "flight"
	ZONE      = "zone"

	PLANNER_TENANT              = "PLANNER"
	ORGANIZATION_ID             = "organizationId"
	ORGANIZATION_PERMISSION     = "organizationPermission"
	ORGANIZATION_PERMISSION_ALL = ":*"
	ORGANIZATION_PERMISSION_UT  = ":UT"
	USER_ID                     = "userId"
	USER_ROLE                   = "role"
	USER_ROLE_SA                = "SA"
	USER_ROLE_PUBLIC            = "PUBLIC"

	ORDER_DESC = "desc"
	ORDER_ASC  = "asc"

	TIMEOUT = 15

	ANNIVERSARY_KINDA_ANNIVERSARY = "ANNIVERSARY"

	BAG_FILE         = ".bag"
	MCAP_FILE        = ".mcap"
	TXT_FILE         = ".txt"
	MP4_FILE         = ".mp4"
	IMG_FILE         = ".jpg"
	REPORT_FILE      = ".pdf"
	ODOMETRY_FILE    = "odometry.csv"
	VIDEO_STAMP_FILE = "video_stamp.txt"

	POIS_SORT_BY_TIMESTAMP   = "TIMESTAMP"
	POIS_SORT_BY_CRITICALITY = "CRITICALITY"

	POI_TYPE_IMAGE = "IMAGE"
	POI_TYPE_UT    = "UT"

	DEFAULT_ZONE_TYPE = "NFZ"

	FILETYPE_UNKNOWN = "UNKNOWN"
	FILETYPE_PC      = "application/octet-stream"
	FILETYPE_TXT     = "text/plain"
	FILETYPE_MP4     = "video/mp4"
	FILETYPE_IMG     = "image/jpeg"

	FONT_NOTO = "notosans"

	S3_KINDA_DIR  = "DIR"
	S3_KINDA_ITEM = "ITEM"

	// tables
	PLANNER_TABLE = "PlannerTable"

	SUCCESS                = "success"
	CF_VIEWABLE_LIMIT_HOUR = 12
)
