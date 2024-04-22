# msfgo
Go library for interacting with Metasploit's RPC API

Thank you to [hupe1980](https://github.com/hupe1980) and his [gomsf](https://github.com/hupe1980/gomsf) library for inspiration. This library is based off of his amazing work.

## Rapid7 Documentation
- [RPC Api](https://docs.rapid7.com/metasploit/rpc-api)
- [Msf::RPC](https://docs.metasploit.com/api/Msf/RPC.html)

# Adding this Libarary to your Project
```
go get github.com/deranged0tter/msfgo
```

```
import (
    "github.com/deranged0tter/msfgo"
)
```

# Starting the RPC Server for Metasploit
From msfconsole:
```
msf > load msgrpc [Pass=password] [User=username]
```

From Command Line:
```
msfrpcd -P password -U username
```

# Connecting to Metasploit's RPC Server
```
// create the client
client, err := msfgo.NewClient("localhost:55552")
if err != nil {
    log.Fatal(err)
}

// connect to the rpc server
err = client.Login("username", "password")
if err != nil {
    log.Fatal(err)
}
defer client.Logout()
```

# Find a Bug?
Submit an [issue](https://github.com/Deranged0tter/msfgo/issues)

# Want to add a feature or missing module?
Submit a pull request

# Supported Methods
## Auth
- [X] auth.login
- [X] auth.logout
- [X] auth.token_list
- [X] auth.token_add
- [X] auth.token_generate
- [X] auth.token_remove

## Console
- [X] console.create
- [X] console.destroy
- [X] console.list
- [X] console.read
- [X] console.write
- [X] console.tab
- [X] console.session_kill
- [X] console.session_detach

## Core
- [X] core.version
- [X] core.stop
- [X] core.setg
- [X] core.unsetg
- [X] core.save
- [X] core.reload_modules
- [X] core.module_stats
- [X] core.add_module_path
- [X] core.thread_list
- [X] core.thread_kill

## DB
- [ ] db.hosts
- [ ] db.services
- [ ] db.vulns
- [ ] db.workspaces
- [ ] db.current_workspace
- [ ] db.get_workspace
- [ ] db.set_workspace
- [ ] db.del_workspace
- [ ] db.add_workspace
- [ ] db.get_host
- [ ] db.report_host
- [ ] db.report_service
- [ ] db.get_service
- [ ] db.get_note
- [ ] db.get_client
- [ ] db.report_client
- [ ] db.report_note
- [ ] db.notes
- [ ] db.get_ref
- [ ] db.del_vul
- [ ] db.del_note
- [ ] db.del_service
- [ ] db.del_host
- [ ] db.report_vuln
- [ ] db.events
- [ ] db.report_event
- [ ] db.report_loot
- [ ] db.loots
- [ ] db.report_cred
- [ ] db.creds
- [ ] db.import_data
- [ ] db.get_vuln
- [ ] db.clients
- [ ] db.del_client
- [ ] db.driver
- [ ] db.connect
- [ ] db.status
- [ ] db.disconnect

## Job
- [X] job.list
- [X] job.stop
- [X] job.info

## Module
- [X] module.exploits
- [X] module.evasion
- [X] module.auxiliary
- [X] module.payloads
- [X] module.encoders
- [X] module.nops
- [X] module.platforms
- [X] module.post
- [X] module.info
- [X] module.compatible_payloads
- [X] module.compatible_sessions
- [X] module.options
- [X] module.execute
- [X] module.search
- [ ] module.check
- [ ] module.results

## Plugin
- [ ] plugin.load
- [ ] plugin.unload
- [ ] plugin.loaded

## Session
- [ ] session.list
- [ ] session.stop
- [ ] session.shell_read
- [ ] session.shell_write
- [ ] session.shell_upgrade
- [ ] session.meterpreter_read
- [ ] session.ring_read
- [ ] session.ring_put
- [ ] session.ring_last
- [ ] session.ring_clear
- [ ] session.meterpreter_write
- [ ] session.meterpreter_session_detach
- [ ] session.meterpreter_session_kill
- [ ] session.meterpreter_tabs
- [ ] session.meterpreter_run_single
- [ ] session.meterpreter_script
- [ ] session.meterpreter_directory_separator
- [ ] session.compatible_modules

# Tested Functions
## Auth
- [ ] auth.login
- [ ] auth.logout
- [ ] auth.token_list
- [ ] auth.token_add
- [ ] auth.token_generate
- [ ] auth.token_remove

## Console
- [ ] console.create
- [ ] console.destroy
- [ ] console.list
- [ ] console.read
- [ ] console.write
- [ ] console.tab
- [ ] console.session_kill
- [ ] console.session_detach

## Core
- [ ] core.version
- [ ] core.stop
- [ ] core.setg
- [ ] core.unsetg
- [ ] core.save
- [ ] core.reload_modules
- [ ] core.module_stats
- [ ] core.add_module_path
- [ ] core.thread_list
- [ ] core.thread_kill

## DB
- [ ] db.hosts
- [ ] db.services
- [ ] db.vulns
- [ ] db.workspaces
- [ ] db.current_workspace
- [ ] db.get_workspace
- [ ] db.set_workspace
- [ ] db.del_workspace
- [ ] db.add_workspace
- [ ] db.get_host
- [ ] db.report_host
- [ ] db.report_service
- [ ] db.get_service
- [ ] db.get_note
- [ ] db.get_client
- [ ] db.report_client
- [ ] db.report_note
- [ ] db.notes
- [ ] db.get_ref
- [ ] db.del_vul
- [ ] db.del_note
- [ ] db.del_service
- [ ] db.del_host
- [ ] db.report_vuln
- [ ] db.events
- [ ] db.report_event
- [ ] db.report_loot
- [ ] db.loots
- [ ] db.report_cred
- [ ] db.creds
- [ ] db.import_data
- [ ] db.get_vuln
- [ ] db.clients
- [ ] db.del_client
- [ ] db.driver
- [ ] db.connect
- [ ] db.status
- [ ] db.disconnect

## Job
- [ ] job.list
- [ ] job.stop
- [ ] job.info

## Module
- [ ] module.exploits
- [ ] module.evasion
- [ ] module.auxiliary
- [ ] module.payloads
- [ ] module.encoders
- [ ] module.nops
- [ ] module.platforms
- [ ] module.post
- [ ] module.info
- [ ] module.compatible_payloads
- [ ] module.compatible_sessions
- [ ] module.options
- [ ] module.execute
- [ ] module.search
- [ ] module.check
- [ ] module.results

## Plugin
- [ ] plugin.load
- [ ] plugin.unload
- [ ] plugin.loaded

## Session
- [ ] session.list
- [ ] session.stop
- [ ] session.shell_read
- [ ] session.shell_write
- [ ] session.shell_upgrade
- [ ] session.meterpreter_read
- [ ] session.ring_read
- [ ] session.ring_put
- [ ] session.ring_last
- [ ] session.ring_clear
- [ ] session.meterpreter_write
- [ ] session.meterpreter_session_detach
- [ ] session.meterpreter_session_kill
- [ ] session.meterpreter_tabs
- [ ] session.meterpreter_run_single
- [ ] session.meterpreter_script
- [ ] session.meterpreter_directory_separator
- [ ] session.compatible_modules

# License
[Apache 2.0](https://github.com/Deranged0tter/msfgo/tree/main?tab=Apache-2.0-1-ov-file)
