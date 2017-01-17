hephaestus
===================
[![works badge](https://cdn.rawgit.com/nikku/works-on-my-machine/v0.2.0/badge.svg)](https://github.com/nikku/works-on-my-machine)

>*In Greek mythology, the god of fire, and of the arts which need fire in the execution. He was said to be the son of Zeus and Hera, or (according to Hesiod) of the latter only.*

**Description**

Hephaestus is a CLI-based firewall ruleset generator built for UNIX/Linux systems running iptables written in Go 1.7.4.

**Usage**
```
> $ ./core -i "(permit|deny) (inbound|outbound) (ssh|icmp|http|https)) 
```

**Demonstration**
```
[1450:9545 - 0:1693] 2017-01-16T19:28:40CST [REDACTED_USER@REDACTED_HOST:o +2] ~/git/hephaestus/src 
$ ./core -i "deny inbound ssh"
[DEBUG] verb validity check: PASSED
[DEBUG] flow validity check: PASSED
[DEBUG] protocol validity check: PASSED
[DEBUG] isvalidstatement: PASSED
-------------------------------------
-A INPUT -p tcp -m --dport 22 -j DENY

[1450:9545 - 0:1694] 2017-01-16T19:28:51CST [REDACTED_USER@REDACTED_HOST:o +2] ~/git/hephaestus/src 
$ ./core -i "deny inbound icmp"
[DEBUG] verb validity check: PASSED
[DEBUG] flow validity check: PASSED
[DEBUG] protocol validity check: PASSED
[DEBUG] isvalidstatement: PASSED
-------------------------------------
-A INPUT -p icmp -m icmp --icmp-type 0 -j DENY
-A INPUT -p icmp -m icmp --icmp-type 3 -j DENY
-A INPUT -p icmp -m icmp --icmp-type 11 -j DENY

[1450:9545 - 0:1696] 2017-01-16T19:34:01CST [REDACTED_USER@REDACTED_HOST:o +2] ~/git/hephaestus/src 
$ ./core -i "permit outbound icmp"
[DEBUG] verb validity check: PASSED
[DEBUG] flow validity check: PASSED
[DEBUG] protocol validity check: PASSED
[DEBUG] isvalidstatement: PASSED
-------------------------------------
-A OUTPUT -p icmp -m icmp --icmp-type 0 -j PERMIT
-A OUTPUT -p icmp -m icmp --icmp-type 3 -j PERMIT
-A OUTPUT -p icmp -m icmp --icmp-type 11 -j PERMIT

[1450:9545 - 0:1697] 2017-01-16T19:34:14CST [REDACTED_USER@REDACTED_HOST:o +2] ~/git/hephaestus/src 
$ 
```
