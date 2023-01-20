---
name: DOMAIN_ELSEWHERE_AUTO
parameters:
  - domain
  - registrar
  - dns provider
parameter_types:
  domain: string
  registrar: string
  dns provider: string
---

`DOMAIN_ELSEWHERE_AUTO()` is similar to `DOMAIN_ELSEWHERE()` but instead of
a hardcoded list of nameservers, a DnsProvider() is queried.

`DOMAIN_ELSEWHERE_AUTO` is useful when you control a domain's registrar but the
DNS zones are managed by another system. Luckily you have enough access to that
other system that you can query it to determine the zone's nameservers.

For example, suppose you own a domain but the DNS servers for it are in Azure.
Further suppose that something in Azure maintains the zones (automatic or
human). Azure picks the nameservers for the domains automatically, and that
list may change occasionally.  `DOMAIN_ELSEWHERE_AUTO` allows you to easily
query Azure to determine the domain's delegations so that you do not need to
hard-code them in your dnsconfig.js file.

For example these two statements are equivalent:

```javascript
DOMAIN_ELSEWHERE_AUTO("example.com", REG_NAMEDOTCOM, DSP_AZURE);

// ...is equivalent to...

D("example.com", REG_NAMEDOTCOM,
    NO_PURGE,
    DnsProvider(DSP_AZURE)
);
```

{% hint style="info" %}
**NOTE**: The `NO_PURGE` is used to prevent DNSControl from changing the records.
{% endhint %}