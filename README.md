# ha-ctrl

ha-ctrl is the efficient, intelligent and user friendly CLI which check pacemaker cluster status


# Features:

## checks
- control needed process status
- control needed systemd service units status
- control cluster status with `crm_mon`

For a short-term roadmap check the issues.

## Auto-detecting optional components:

`ha-ctrl` can detect if a component is installed, e.g `sbd` and run in that case the needed checks.
Or for example is `stonith-enabled` or other options are enabled, will run the different checks.
