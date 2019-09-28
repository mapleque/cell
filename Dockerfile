FROM debian:9.4-slim

COPY releases/cell-linux /opt/cell
COPY bin/templates /opt/templates
COPY bin/www /opt/www

ENV \
HTTP_LISTEN=0.0.0.0:80 \
HTTP_STATIC_PATH=/opt/www/html \
HTTPS_KEY_FILE= \
HTTPS_CERT_FILE= \
DB_DSN=root@tcp(localhost:3306)/cell?charset=utf8mb4&parseTime=true&loc=Local \
DB_MAX_IDLE=10 \
DB_MAX_OPEN=10 \
OIDC_PUBLIC_KEY_FILE= \
OIDC_PRIVATE_KEY_FILE= \
MAIL_USERNAME= \
MAIL_PASSWORD= \
MAIL_HOST= \
MAIL_ADDRESS= \
MAIL_FROM= \
MAIL_TEMPLATE_PATTERN=/opt/templates/mail_template/* \
KERBEROS_TGS_SECRET_KEY=a_random_token_with_32_charactor \
KERBEROS_APP_SECRET_KEY=b_random_token_with_32_charactor \
REACT_APP_PASSWORD_SALT=

ENTRYPOINT /opt/cell
