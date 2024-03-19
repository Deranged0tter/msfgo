# msfgo
Go library for interacting with Metasploit's RPC API

## Rapid7 Documentation
[RPC Api](https://docs.rapid7.com/metasploit/rpc-api)
[Msf::RPC](https://docs.metasploit.com/api/Msf/RPC.html)

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
client, err := msfgo.NewClient("localhost:55552")
if err != nil {
    log.Fatal(err)
}

err = client.Login("username", "password")
if err != nil {
    log.Fatal(err)
}
defer client.Logout()
```