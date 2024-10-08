package permissions

type PermissionEnum string

const (
	PERMISSION_READ_ALL_ORGANIZATIONS  = "123e4567-e89b-12d3-a456-426614173000"
	PERMISSION_ADMIN_ALL_ORGANIZATIONS = "123e4567-e89b-12d3-a456-426614173001"

	// Application permissions
	PERMISSION_CREATE_APPLICATION = "123e4567-e89b-12d3-a456-426614174000"
	PERMISSION_UPDATE_APPLICATION = "123e4567-e89b-12d3-a456-426614174001"
	PERMISSION_DELETE_APPLICATION = "123e4567-e89b-12d3-a456-426614174002"
	PERMISSION_READ_APPLICATION   = "123e4567-e89b-12d3-a456-426614174003"

	// Organization permissions
	PERMISSION_ADD_ORGANIZATION_MEMBER         = "123e4567-e89b-12d3-a456-426614175000"
	PERMISSION_REMOVE_ORGANIZATION_TEAM_MEMBER = "123e4567-e89b-12d3-a456-426614175001"

	PERMISSION_CREATE_TEAM    = "123e4567-e89b-12d3-a456-426614175002"
	PERMISSION_UPDATE_TEAM    = "123e4567-e89b-12d3-a456-426614175003"
	PERMISSION_DELETE_TEAM    = "123e4567-e89b-12d3-a456-426614175004"
	PERMISSION_READ_ALL_TEAMS = "123e4567-e89b-12d3-a456-426614175005"

	PERMISSION_UPDATE_ORGANIZATION_BASIC_DETAILS = "123e4567-e89b-12d3-a456-426614175006"
	PERMISSION_CHAGE_ORGANIZATION_PLAN           = "123e4567-e89b-12d3-a456-426614175007"
	PERMISSION_UPDATE_ORGANIZATION_BILLING       = "123e4567-e89b-12d3-a456-426614175008"
	PERMISSION_CREATE_ORGANIZATION_INTEGRATION   = "123e4567-e89b-12d3-a456-426614175009"
	PERMISSION_UPDATE_ORGANIZATION_INTEGRATION   = "123e4567-e89b-12d3-a456-426614175010"
	PERMISSION_READ_ORGANIZATION_INTEGRATIONS    = "123e4567-e89b-12d3-a456-426614175011"

	PERMISSION_READ_ORGANIZATION_DETAILS  = "123e4567-e89b-12d3-a456-426614175012"
	PERMISSION_READ_ORGANIZATION_COSTS    = "123e4567-e89b-12d3-a456-426614175013"
	PERMISSION_READ_ORGANIZATION_PROJECTS = "123e4567-e89b-12d3-a456-426614175014"

	PERMISSION_CREATE_PROJECT    = "123e4567-e89b-12d3-a456-426614175015"
	PERMISSION_UPDATE_PROJECT    = "123e4567-e89b-12d3-a456-426614175016"
	PERMISSION_DELETE_PROJECT    = "123e4567-e89b-12d3-a456-426614175017"
	PERMISSION_READ_ALL_PROJECTS = "123e4567-e89b-12d3-a456-426614175018"
	PERMISSION_READ_PROJECT      = "123e4567-e89b-12d3-a456-426614175019"

	PERMISSION_CREATE_PROJECT_ENVIRONMENT    = "123e4567-e89b-12d3-a456-426614175020"
	PERMISSION_UPDATE_PROJECT_ENVIRONMENT    = "123e4567-e89b-12d3-a456-426614175021"
	PERMISSION_DELETE_PROJECT_ENVIRONMENT    = "123e4567-e89b-12d3-a456-426614175022"
	PERMISSION_READ_ALL_PROJECT_ENVIRONMENTS = "123e4567-e89b-12d3-a456-426614175023"
	PERMISSION_READ_PROJECT_ENVIRONMENT      = "123e4567-e89b-12d3-a456-426614175024"

	PERMISSION_CREATE_PROJECT_ENVIRONMENT_VARIABLE = "123e4567-e89b-12d3-a456-426614175025"
	PERMISSION_UPDATE_PROJECT_ENVIRONMENT_VARIABLE = "123e4567-e89b-12d3-a456-426614175026"
	PERMISSION_DELETE_PROJECT_ENVIRONMENT_VARIABLE = "123e4567-e89b-12d3-a456-426614175027"
	PERMISSION_READ_PROJECT_ENVIRONMENT_VARIABLE   = "123e4567-e89b-12d3-a456-426614175029"

	//TEAM PERMISSIONS
	PERMISSION_TEAM_ADMIN       = "123e4567-e89b-12d3-a456-426614176000"
	PERMISSION_TEAM_CONTRIBUTOR = "123e4567-e89b-12d3-a456-426614176000"
	PERMISSION_TEAM_READ_ONLY   = "123e4567-e89b-12d3-a456-426614176000"
	PERMISSION_TEAM_DEPLOY      = "123e4567-e89b-12d3-a456-426614176000"
)
