import secrets

URL_Reset = secrets.token_urlsafe(32)
print(f"URL Token: http://example.com.{URL_Reset}")