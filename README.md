# ha-ctrl

ha-ctrl is the efficient, intelligent and user friendly CLI which check pacemaker cluster status


# Features:

## checks
- control needed process status
- control needed systemd service units status
- control cluster status with `crm_mon`

## Auto-healing:

If some of the check are failings, `ha-ctrl` depending on the failure will try to solve automatically the problem. ( note: it depends on the check failure)

## Auto-detecting optional components:

`ha-ctrl` can detect if a component is installed, e.g `sbd` and run in that case the needed checks.

