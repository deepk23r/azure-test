# conftest policy to test docker image for permissions 
package main

# check if file has setgid bit set
deny[msg] {
  input.permissions[i].metadata.Type == "RegularFile"
  permission := format_int(input.permissions[i].metadata.Mode, 8)
  regex.match("20......", permission)

  msg = sprintf(
    "%s - 4.8 Level 1 benchmark - setgid file permission validation: (%s, %o)",
    [
      "permissions.rego failed", 
      input.permissions[i].path, 
      input.permissions[i].metadata.Mode
    ],
  )
}

# check if file has setuid bit set
deny[msg] {
  input.permissions[i].metadata.Type == "RegularFile"
  permission := format_int(input.permissions[i].metadata.Mode, 8)
  regex.match("40......", permission)

  msg = sprintf(
    "%s - 4.8 Level 1 benchmark - setuid file permission validation: (%s, %o)",
    [
      "permissions.rego failed", 
      input.permissions[i].path, 
      input.permissions[i].metadata.Mode
    ],
  )
}

# check if file has setuid and setgid bit set
deny[msg] {
  input.permissions[i].metadata.Type == "RegularFile"
  permission := format_int(input.permissions[i].metadata.Mode, 8)
  regex.match("60......", permission)

  msg = sprintf(
    "%s - 4.8 Level 1 benchmark - setuid and setgid file permission validation: (%s, %o)",
    [
      "permissions.rego failed", 
      input.permissions[i].path, 
      input.permissions[i].metadata.Mode
    ],
  )
}
