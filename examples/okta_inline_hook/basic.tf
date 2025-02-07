resource "okta_inline_hook" "test" {
  name    = "testAcc_replace_with_uuid"
  version = "1.0.1"
  type    = "com.okta.oauth2.tokens.transform"

  channel = {
    type    = "HTTP"
    version = "1.0.0"
    uri     = "https://example.com/test"
    method  = "POST"
  }

  auth = {
    key   = "Authorization"
    type  = "HEADER"
    value = "123"
  }
}

resource "okta_inline_hook" "twilio" {
  name = "twillio"
  version = "1.0.0"
  type = "com.okta.telephony.provider"

  channel = {
    version = "1.0.0"
    uri = "https://example.com/test"
    method = "POST"
  }

  auth = {
    key = "Authorization"
    type = "HEADER"
    value = "secret"
  }
}
