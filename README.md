# OverlayTool
This repository will contain some kind of overlay iteration script written in Golang.
Currently available:
- Point inside Polygon => to find any ODP inside Clusters.
```bash
(_bin)$ ./point-inside-polygon -input-point Point-ODP_R4_20190428.csv -input-polygon Polygon-Cluster-TREG-4.csv
2019/04/28 23:57:00 Parse Polygon File...
2019/04/28 23:57:00 skip at row 252: the coordinates (B) is unabled to be parsed: wrong coordinate format: 'odp-ken-fbd'
2019/04/28 23:57:00 Parse Point File...
2019/04/28 23:57:00 Start Iteration Process...
 111788 / 111788 [========================================================================================================================================================================================================================================================] 100.00% 8s
Iteration Complete.
2019/04/28 23:57:09 Store Result...
All Done. Check Your Result: output.csv
```

## Status
Experimental

## How to Start Development
```bash
JUST CLONE IT
```
## License
Not Yet Decided