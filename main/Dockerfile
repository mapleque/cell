FROM debian:9.4-slim

COPY releases/cell-linux /opt/cell
COPY bin/templates /opt/templates
COPY bin/www /opt/www

###########################################################

# HTTP service

# listen address (with port) for TCP
ENV HTTP_LISTEN=0.0.0.0:80
# html client file path
ENV HTTP_STATIC_PATH=/opt/www/html

# SSL key and cert file path
# HTTPS open when both settings.
#
# ENV HTTPS_KEY_FILE=
# ENV HTTPS_CERT_FILE=

###########################################################

# database settings
# The database should support mysql driver
# ENV DB_DSN=root@tcp(localhost:3306)/cell?charset=utf8mb4&parseTime=true&loc=Local
ENV DB_MAX_IDLE=10
ENV DB_MAX_OPEN=10

###########################################################

# OIDC settings

# OIDC key pair, can be general with any rsa tool
# OIDC service can not work if empty.
#
# ENV OIDC_PUBLIC_KEY_FILE=
# ENV OIDC_PRIVATE_KEY_FILE=

###########################################################

# Webmaster Email account settings
#
# For captcha notify.

# account
ENV MAIL_USERNAME=
ENV MAIL_PASSWORD=

# mail service host
ENV MAIL_HOST=

# smtp address, with port
ENV MAIL_ADDRESS=

# sender email address, use <> for alias
ENV MAIL_FROM=

# mail tempalte path pattern
ENV MAIL_TEMPLATE_PATTERN=/opt/templates/mail_template/*

###########################################################

# crypt settings

# Kerberos TGS secret key
#
# set a random token with 32 bytes
# !!!Important: modify this or setting with env
ENV KERBEROS_TGS_SECRET_KEY=a_random_token_with_32_charactor

# Kerberos application secret key
#
# set a random token with 32 bytes
# !!!Important: modify this or setting with env
ENV KERBEROS_APP_SECRET_KEY=b_random_token_with_32_charactor

# user password hash salt
#
# !!!Important: modify this or setting with env
ENV REACT_APP_PASSWORD_SALT=

ENTRYPOINT /opt/cell
