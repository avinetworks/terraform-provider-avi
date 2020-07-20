############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_alertconfig"
sidebar_current: "docs-avi-datasource-alertconfig"
description: |-
  Get information of Avi AlertConfig.
---

# avi_alertconfig

This data source is used to to get avi_alertconfig objects.

## Example Usage

```hcl
data "avi_alertconfig" "foo_alertconfig" {
    uuid = "alertconfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AlertConfig by name.
* `uuid` - (Optional) Search AlertConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `action_group_ref` - The alert config will trigger the selected alert action, which can send notifications and execute a controlscript.
* `alert_rule` - List of filters matching on events or client logs used for triggering alerts.
* `autoscale_alert` - This alert config applies to auto scale alerts.
* `category` - Determines whether an alert is raised immediately when event occurs (realtime) or after specified number of events occurs within rolling time window.
* `description` - A custom description field.
* `enabled` - Enable or disable this alert config from generating new alerts.
* `expiry_time` - An alert is expired and deleted after the expiry time has elapsed.
* `name` - Name of the alert configuration.
* `obj_uuid` - Uuid of the resource for which alert was raised.
* `object_type` - The object type to which the alert config is associated with.
* `recommendation` - Placeholder for description of property recommendation of obj type alertconfig field type string  type str.
* `rolling_window` - Only if the number of events is reached or exceeded within the time window will an alert be generated.
* `source` - Signifies system events or the type of client logsused in this alert configuration.
* `summary` - Summary of reason why alert is generated.
* `tenant_ref` - It is a reference to an object of type tenant.
* `threshold` - An alert is created only when the number of events meets or exceeds this number within the chosen time frame.
* `throttle` - Alerts are suppressed (throttled) for this duration of time since the last alert was raised for this alert config.
* `uuid` - Unique object identifier of the object.

