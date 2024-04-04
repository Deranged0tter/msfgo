package methods

const (
	SessionList                          string = "session.list"             // priority
	SessionStop                          string = "session.stop"             // priority
	SessionShellRead                     string = "session.shell_read"       // priority
	SessionShellWrite                    string = "session.shell_write"      // priority
	SessionShellUpgrade                  string = "session.shell_upgrade"    // priority
	SessionMeterpreterRead               string = "session.meterpreter_read" // priority
	SessionRingRead                      string = "session.ring_read"
	SessionRingPut                       string = "session.ring_put"
	SessionRingLast                      string = "session.ring_last"
	SessionRingClear                     string = "session.ring_clear"
	SessionMeterpreterWrite              string = "session.meterpreter_write"          // priority
	SessionMeterpreterSessionDetach      string = "session.meterpreter_session_detach" // priority
	SessionMeterpreterSessionKill        string = "session.meterpreter_session_kill"   // priority
	SessionMeterpreterTabs               string = "session.meterpreter_tabs"
	SessionMeterpreterRunSingle          string = "session.meterpreter_run_single" // priority
	SessionMeterpreterScript             string = "session.meterpreter_script"
	SessionMeterpreterDirectorySeparator string = "session.meterpreter_directory_separator"
	SessionCompatibleModules             string = "session.compatible_modules"
)
