package crmtemplate

const CRM_ISCSI = `<configuration>
    <resources>
      <primitive id="p_iscsi_{{.CRM_TARGET_NAME}}_ip" class="ocf" provider="heartbeat" type="IPaddr2">
        <instance_attributes id="p_iscsi_{{.CRM_TARGET_NAME}}_ip-instance_attributes">
          <nvpair name="ip" value="{{.CRM_SVC_IP}}" id="p_iscsi_{{.CRM_TARGET_NAME}}_ip-instance_attributes-ip"/>
          <nvpair name="cidr_netmask" value="24" id="p_iscsi_{{.CRM_TARGET_NAME}}_ip-instance_attributes-cidr_netmask"/>
        </instance_attributes>
        <operations>
          <op name="monitor" interval="15" timeout="40" id="p_iscsi_{{.CRM_TARGET_NAME}}_ip-monitor-15"/>
          <op name="start" timeout="40" interval="0" id="p_iscsi_{{.CRM_TARGET_NAME}}_ip-start-0"/>
          <op name="stop" timeout="40" interval="0" id="p_iscsi_{{.CRM_TARGET_NAME}}_ip-stop-0"/>
        </operations>
      </primitive>

      <primitive id="p_iscsi_{{.CRM_TARGET_NAME}}" class="ocf" provider="heartbeat" type="iSCSITarget">
        <instance_attributes id="p_iscsi_{{.CRM_TARGET_NAME}}-instance_attributes">
          <nvpair name="iqn" value="{{.TARGET_IQN}}" id="p_iscsi_{{.CRM_TARGET_NAME}}-instance_attributes-iqn"/>
          <nvpair name="incoming_username" value="{{.USERNAME}}" id="p_iscsi_{{.CRM_TARGET_NAME}}-instance_attributes-incoming_username"/>
          <nvpair name="incoming_password" value="{{.PASSWORD}}" id="p_iscsi_{{.CRM_TARGET_NAME}}-instance_attributes-incoming_password"/>
          <nvpair name="portals" value="{{.PORTALS}}" id="p_iscsi_{{.CRM_TARGET_NAME}}-instance_attributes-portals"/>
          <nvpair name="tid" value="{{.TID}}" id="p_iscsi_{{.CRM_TARGET_NAME}}-instance_attributes-tid"/>
        </instance_attributes>
        <operations>
          <op name="start" timeout="40" interval="0" id="p_iscsi_{{.CRM_TARGET_NAME}}-start-0"/>
          <op name="stop" timeout="40" interval="0" id="p_iscsi_{{.CRM_TARGET_NAME}}-stop-0"/>
          <op name="monitor" interval="15" timeout="40" id="p_iscsi_{{.CRM_TARGET_NAME}}-monitor-15"/>
        </operations>
        <meta_attributes id="p_iscsi_{{.CRM_TARGET_NAME}}-meta_attributes">
          <nvpair name="target-role" value="Started" id="p_iscsi_{{.CRM_TARGET_NAME}}-meta_attributes-target-role"/>
        </meta_attributes>
      </primitive>

      <primitive id="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}" class="ocf" provider="heartbeat" type="iSCSILogicalUnit">
        <instance_attributes id="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}-instance_attributes">
          <nvpair name="lun" value="{{.LUN}}" id="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}-instance_attributes-lun"/>
          <nvpair name="target_iqn" value="{{.TARGET_IQN}}" id="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}-instance_attributes-target_iqn"/>
          <nvpair name="path" value="{{.DEVICE}}" id="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}-instance_attributes-path"/>
        </instance_attributes>
        <operations>
          <op name="start" timeout="40" interval="0" id="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}-start-0"/>
          <op name="stop" timeout="40" interval="0" id="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}-stop-0"/>
          <op name="monitor" timeout="40" interval="15" id="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}-monitor-15"/>
        </operations>
      </primitive>
    </resources>

    <constraints>
      <rsc_location id="lo_iscsi_{{.CRM_TARGET_NAME}}" resource-discovery="never">
        <resource_set id="lo_iscsi_{{.CRM_TARGET_NAME}}-0">
          <resource_ref id="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}"/>
          <resource_ref id="p_iscsi_{{.CRM_TARGET_NAME}}"/>
        </resource_set>
        <rule score="-INFINITY" id="lo_iscsi_{{.CRM_TARGET_NAME}}-rule">
{{.TARGET_LOCATION_NODES}}
        </rule>
      </rsc_location>
      <rsc_colocation id="co_iscsi_{{.CRM_TARGET_NAME}}" score="INFINITY" rsc="p_iscsi_{{.CRM_TARGET_NAME}}" with-rsc="p_iscsi_{{.CRM_TARGET_NAME}}_ip"/>
      <rsc_order id="o_iscsi_{{.CRM_TARGET_NAME}}" score="INFINITY" first="p_iscsi_{{.CRM_TARGET_NAME}}_ip" then="p_iscsi_{{.CRM_TARGET_NAME}}"/>

      <rsc_location id="lo_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}" rsc="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}" resource-discovery="never">
        <rule score="0" id="lo_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}-rule">
{{.LU_LOCATION_NODES}}
        </rule>
      </rsc_location>
      <rsc_colocation id="co_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}" score="INFINITY" rsc="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}" with-rsc="p_iscsi_{{.CRM_TARGET_NAME}}"/>
      <rsc_order id="o_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}" score="INFINITY" first="p_iscsi_{{.CRM_TARGET_NAME}}" then="p_iscsi_{{.CRM_TARGET_NAME}}_{{.CRM_LU_NAME}}"/>
    </constraints>
</configuration>`
