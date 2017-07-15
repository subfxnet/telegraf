# Unbound Input Plugin

This plugin gathers stats from [Unbound Caching DNS Server](https://unbound.net/)

Configure the unbound dns server to provide extended statistics as outlined [here](https://unbound.net/documentation/howto_statistics.html).

### Configuration:

```toml
 # A plugin to collect stats from Unbound DNS server
 [[inputs.unbound]]
   ## The default location of the unbound-control program can be overridden with:
   control = "/usr/local/sbin/unbound-control"
   ## Use sudo to execute unbound-control by providing a user to sudo to:
   sudo_user = "root"
```

### Measurements & Fields:

- unbound
    - total.num.queries
    - total.num.queries_ip_ratelimited
    - total.num.cachehits
    - total.num.cachemiss
    - total.num.prefetch
    - total.num.zero_ttl
    - total.num.recursivereplies
    - total.requestlist.avg
    - total.requestlist.max
    - total.requestlist.overwritten
    - total.requestlist.exceeded
    - total.requestlist.current.all
    - total.requestlist.current.user
    - total.recursion.time.avg
    - total.recursion.time.median
    - total.tcpusage
    - time.now
    - time.up
    - time.elapsed
    - mem.cache.rrset
    - mem.cache.message
    - mem.mod.iterator
    - mem.mod.validator
    - mem.mod.respip
    - num.query.type.A
    - num.query.type.AAAA
    - num.query.type.MX
    - num.query.type.NS
    - num.query.type.CNAME
    - num.query.type.PTR
    - num.query.type.SRV
    - num.query.type.TXT
    - num.query.type.Other
    - num.query.class.IN
    - num.query.opcode.QUERY
    - num.query.tcp
    - num.query.tcpout
    - num.query.ipv6
    - num.query.flags.QR
    - num.query.flags.AA
    - num.query.flags.TC
    - num.query.flags.RD
    - num.query.flags.RA
    - num.query.flags.Z
    - num.query.flags.AD
    - num.query.flags.CD
    - num.query.edns.present
    - num.query.edns.DO
    - num.answer.rcode.NOERROR
    - num.answer.rcode.FORMERR
    - num.answer.rcode.SERVFAIL
    - num.answer.rcode.NXDOMAIN
    - num.answer.rcode.NOTIMPL
    - num.answer.rcode.REFUSED
    - num.answer.secure
    - num.answer.bogus
    - num.rrset.bogus
    - unwanted.queries
    - unwanted.replies
    - msg.cache.count
    - rrset.cache.count
    - infra.cache.count
    - key.cache.count


### Example Output:

```
 telegraf --config etc/telegraf.conf --input-filter unbound --test
* Plugin: unbound, Collection 1
> unbound,host=subfx.net ... 1462765437090957980
```
