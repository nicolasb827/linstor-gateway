package crmtemplate

const LU_LOCATION_NODES = `<expression attribute="#uname" operation="ne" value="{{.CRM_NODE_NAME}}" id="lo_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}-rule-expression-{{.NR}}"/>`
