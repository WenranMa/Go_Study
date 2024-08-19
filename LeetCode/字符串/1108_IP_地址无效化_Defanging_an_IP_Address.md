# 1108_IP_地址无效化_Defanging_an_IP_Address

Given a valid (IPv4) IP address, return a defanged version of that IP address.
A defanged IP address replaces every period "." with "[.]".

Example:
Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0"

### 解：

最后的n表示前替换前n个，-1就全替换。

```go
func defangIPaddr(address string) string {
    return strings.Replace(address, ".", "[.]", -1)
}
```

