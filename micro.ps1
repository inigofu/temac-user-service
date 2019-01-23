$env:CORS_ALLOWED_HEADERS="authorization,content-type"
$env:CORS_ALLOWED_ORIGINS="*"
$env:CORS_ALLOWED_METHODS="POST"
$env:REGISTRY="consul"
micro --registry=consul api