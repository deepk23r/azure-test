# conftest policy to evaluate results for secret issues
package main

deny[msg] {

  severities := ["HIGH", "CRITICAL"]
  input.secrets.Results[i].Secrets[j].Severity == severities[_]

  msg = sprintf(
      "%s - Found %s secret: %s: %s",
      [ "secret.rego failed",
        input.secrets.Results[i].Secrets[j].Severity,
        input.secrets.Results[i].Target,
        input.secrets.Results[i].Secrets[j].Title
      ] )
}
