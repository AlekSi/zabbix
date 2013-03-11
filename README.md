zabbix [![Is maintained?](http://stillmaintained.com/AlekSi/zabbix.png)](http://stillmaintained.com/AlekSi/zabbix)
======

WARNING! API IS NOT STABLE YET

TODO description

Install it: `go get github.com/AlekSi/zabbix`

You *have* to run tests before using this package – Zabbix API doesn't match documentation in few details, which are changing in patch releases. Tests should not be destructive, but you are advised to run them against not-production instance or at least make a backup.

    export TEST_ZABBIX_URL=http://host/api_jsonrpc.php
    export TEST_ZABBIX_USER=User
    export TEST_ZABBIX_PASSWORD=Password
    go test -v

`TEST_ZABBIX_URL` may contain HTTP basic auth username and password: `http://username:password@host/api_jsonrpc.php`. Also, in some setups URL should be like `http://host/zabbix/api_jsonrpc.php`.

Documentation is available on [godoc.org](http://godoc.org/github.com/AlekSi/zabbix).

License: Simplified BSD License (see LICENSE).
