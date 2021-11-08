# tf-cloud
CLI tool help with managing workpaces and other stuff in terraform cloud.  This is a very rough work in progress.

# Build
Eventually I'll get around to doing packages and stuff, but for now just do `go build`
```bash
go build .
```

# Setup
This tool expects to have a config file and by default uses `$HOME/.tf-cloud.yaml`. Can also specify the config file with `--config ~/.tf-cloud-bw.yaml`.

Example config file
```bash
token: <TF-CLOUD-USER-API-TOKEN>
organization: <YOUR-ORGANIZATION> # This will be the organization you want use when working with workspaces
```

Can also use ENV VARs prefixed with `TFCLOUD_` for config settings. This will check for the ENV VAR first and if they are not set, it will fallback to the config file.

Example using ENV VARs
```bash
export TFCLOUD_TOKEN=fakeToken
export TFCLOUD_ORGANIZATION=fakeOrg
```

# Listing organizations
To view list organizations your account is associated with
```bash
$ ./tf-cloud org list --config ~/.tf-cloud-bw.yaml
```

# To list workspaces
This will list the first 20 workspaces under your organization defined in the `$HOME/.tf-cloud.yaml`.

```bash
$ ./tf-cloud workspace list
Using config file: /home/flynshue/.tf-cloud.yaml
/api/v2/organizations/kot-labs/workspaces

Workspaces
ID, Name, Created, Locked, Execution Mode
ws-8hwbdW2pzZBqByut, kops-utils-test, 2021-10-22T17:11:46.914Z, false, local
ws-kT8YYQf4ArtQJZ3d, kops-dns-sbx, 2021-11-05T21:46:47.798Z, false, local
ws-Xea9oUG8qr8pRW3U, foo-bar-lab, 2021-11-05T20:29:51.207Z, false, local
ws-rB4MvZUqs4KvD5oB, kops-dns-lab, 2021-11-05T20:02:46.422Z, false, local
```

# To create workspace
This will create workspace using local execution-mode. Currently, local execution-mode is hardcoded in the code because I don't want to do remote operations.
```bash
tf-cloud workspace create <WORKSPACE-NAME>
```
Example
```bash
$ ./tf-cloud workspace create foo-bar-stg
Using config file: /home/flynshue/.tf-cloud.yaml

Successfully created workspace
id: ws-Q1mEQK8AkURZT8Xp, name: foo-bar-stg, created: 2021-11-08T15:36:11.174Z, locked: false, execution-mode: local
```

# To search for workspace
```bash
$ ./tf-cloud workspace search foo-bar
Using config file: /home/flynshue/.tf-cloud.yaml
/api/v2/organizations/kot-labs/workspaces

Workspaces
ID, Name, Created, Locked, Execution Mode
ws-icZgjSmkA3WreyVb, foo-bar-prd, 2021-11-08T15:36:31.032Z, false, local
ws-nZaAK5Ewp1kxxxph, foo-bar-sbx, 2021-11-08T15:36:27.108Z, false, local
ws-Q1mEQK8AkURZT8Xp, foo-bar-stg, 2021-11-08T15:36:11.174Z, false, local
ws-Xea9oUG8qr8pRW3U, foo-bar-lab, 2021-11-05T20:29:51.207Z, false, local
```

# To delete workspace
```bash
$ ./tf-cloud workspace delete foo-bar-prd
Using config file: /home/flynshue/.tf-cloud.yaml

200 Successfully deleted /api/v2/organizations/kot-labs/workspaces/foo-bar-prd
```

You can now see workspace has been removed
```bash
$ ./tf-cloud workspace search foo-bar-prd
Using config file: /home/flynshue/.tf-cloud.yaml
/api/v2/organizations/kot-labs/workspaces

Workspaces
ID, Name, Created, Locked, Execution Mode
```