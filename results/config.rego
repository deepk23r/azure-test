# conftest policy to evaluate results for misconfigurations issues
package main

deny[msg] {

  severities := ["HIGH", "CRITICAL"]
  input.configs.Results[i].Misconfigurations[j].Severity == severities[_]
  input.configs.Results[i].Misconfigurations[j].Status == "FAIL"

  msg = sprintf(
      "%s - Found %s issues in infrastructure as code: %s: %s",
      [ "config.rego failed",
        input.configs.Results[i].Misconfigurations[j].Severity,
        input.configs.Results[i].Target,
        input.configs.Results[i].Misconfigurations[j].Title
      ] )
}
